# Sending Data (ex files here) over Network async

## first without a stream, just sending a big batch of data:
```go

type FileServer struct {}

func (fs *FileServer) start(){
	listener, err := net.Listen("tcp", ":3000")
	if err != nil{
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readingLoop(connection)
	}
}
func (fs *FileServer) readingLoop(conn net.Conn){
	buf := make([]byte, 2048)	// our buffer we write into
	for {
		numberBytes, err := conn.Read(buf)
		if err !=nil{
			log.Fatal(err)
		}
		file := buf[:numberBytes]
		fmt.Println(file)
		fmt.Printf("received %d bytes \n", numberBytes)
	}
}

func main(){
	go func(){
		time.Sleep(4*time.Second)
		sendFile(1000)
	}()
	server := &FileServer{}
	server.start()

}

// sending some random file over for testing
func sendFile(size int) error{
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err !=nil{
		return err
	}
	con, err := net.Dial("tcp", ":3000")
	if err !=nil{
		return err
	}
	numberBytes, err := con.Write(file)
	if err !=nil{
		return err
	}
	fmt.Printf("sending %d bytes over \n", numberBytes)
	return nil
}
```