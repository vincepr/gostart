package main

import (
	"context"
	"fmt"
	"log"
)



func main(){
	service := NewTrainingService("https://catfact.ninja/fact")
	service = NewLoggingService(service)

	task, err := service.GetTrainingTask(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(task)
	fmt.Printf("%+v\n", task)



}