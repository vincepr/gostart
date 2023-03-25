package main

import (
	"context"
	"log"
)

type jwt string
const auth jwt = "JWT"

func main(){
	ctx := context.WithValue(context.Background(), auth, "Some Data we want to pass down")

	bearer := ctx.Value(auth)
	str, ok := bearer.(string)
	if !ok{
		log.Fatal("not a string")
	}
	log.Println("value:", str)
}