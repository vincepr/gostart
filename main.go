package main

import (
	"apiServer/api"
	"apiServer/storage"
	"flag"
	"fmt"
	"log"
)

func main(){
	listenAddr := flag.String("listenaddr", ":5050", "0.0.0.0")
	flag.Parse()

	store := storage.NewMemordyStorage()					//fake in memory data to test with
	
	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}