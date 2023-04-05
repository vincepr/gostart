package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"syscall"
	"time"
)

func main(){
	path := flag.String("path", "", "path/folder to walk")
	flag.Parse()
	if *path == ""{
		flag.Usage()
		return
	}
	doFolder(*path)
}

// running the main program
func doFolder(path string){
	root, err := Walk(path)
	if err != nil{
		log.Fatalln("Walk failed", err)
	}
	if root == nil {
		log.Fatalln("Could not parse root")
	}
	root.Print()
}

// tree structure we parse our file system into
type FileNode struct {
	Name		string
	Mode		os.FileMode		// permissions
	Size		int64			// "real" size the data takes
	SizeOnDisk	int64			// real-size*block_size = size the os allocates for the file and is actually used
	IsDir		bool
	LastChange	time.Time
	Children	[]*FileNode
	Parent		*FileNode
}

/*
*	Walking/Creating the Tree-structure logic:
*/

// walks the folder and returns a FileNode tree of all folders and files inside while calculating sizes
func Walk(path string) (*FileNode, error){
	file, err := os.Lstat(path)				//syscall to get os.FileInfo
	if err != nil{
		log.Printf("Walk() ERROR: %s: %v", path, err)
		return nil, nil
	}
	st, ok := file.Sys().(*syscall.Stat_t)	//type assertion of the syscall
	if !ok {
		return nil, fmt.Errorf("could not cast %T to syscall.Stat_t", file.Sys())
	}
	root := &FileNode{
		Name:		file.Name(),
		Mode:		file.Mode(),
		Size:		st.Size,
		SizeOnDisk:	st.Blocks*512,			// calculated by using Stat_t.Blocks -> the number of 512 byte blocks the file uses
		IsDir:		file.IsDir(),
		LastChange:	file.ModTime(),
	}
	if root.IsDir{
		names, err := readDirNames(path)
		if err != nil{
			log.Printf("Walk() ERROR: %s: %v", path, err)
			return root, nil
		}
		for _,name := range names{
			child, err := Walk(path+"/"+name)
			if err != nil{
				return nil, err
			}
			child.Parent = root
			// sum up the filesizes
			root.Size += child.Size
			root.SizeOnDisk += child.SizeOnDisk
			root.Children = append(root.Children, child)
		}
	}
	return root, nil
}

// reads (and sorts) all filenames in a folder
func readDirNames(path string) ([]string, error){
	file, err := os.Open(path)
	if err != nil{
		return nil, err
	}
	names, err := file.Readdirnames(-1)
	file.Close()
	if err != nil{
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

/*
*	Printing out the Filetree struct logic:
*/

func (f *FileNode) Print(){
	f.printWithDepth(0)
}

func (f *FileNode) printWithDepth(depth int){
	fmt.Printf("%s%s [%s]	[%s] \n", 
		strings.Repeat("  ", (depth*2)) ,
		f.Name,
		f.Mode.String(),
		formatSize(f.SizeOnDisk))
	for _, child := range f.Children{
		child.printWithDepth(depth+1)
	}
}

/*
*	helpers
*/

// 2048 B -> 4.00 KB, xxx B -> 3590MB -> 3,50 GB
func formatSize(bytes int64) string{
	const (
	KB = 1024
	MB = 1024*KB
	GB = 1024*MB
	TB = 1024*GB
	)
	if bytes >= TB {
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	} else 
	if bytes >= GB {
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	} else 
	if bytes >= MB {
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	} else 
	if bytes >= KB {
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	}
	return fmt.Sprintf("%v B", bytes)
}