package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service interface {
	GetTrainingTask(context.Context) (*TrainingTask, error)
}

type TrainingService struct {
	url string
}

func NewTrainingService(url string) Service{
	return &TrainingService{ url: url}
}

func (s *TrainingService) GetTrainingTask(ctx context.Context) (*TrainingTask, error){
	response, err := http.Get(s.url)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println("BODY:", response.Body)

	task := &TrainingTask{}
	err = json.NewDecoder(response.Body).Decode(task)

	// if err := json.NewDecoder(response.Body).Decode(task); err != nil{
	// 	return nil, err
	// }
	fmt.Println("TASK: ",task.Question)
	return task, nil
}