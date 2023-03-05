package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service{
	return &LoggingService{ next: next}
}

//middleware -ish logging
func (s *LoggingService) GetTrainingTask(ctx context.Context)(task *TrainingTask, err error){
	// defer trick to get endTime-startTime
	// since this defer func() runs after the s.next.GetTrainingTask
	// we can access that ones returns:(*TrainingTask, error)
	// named-return style.
	defer func(start time.Time){
		fmt.Printf("task:%v | err:%s | took:%v", task, err, time.Since(start))
	}(time.Now())

	return s.next.GetTrainingTask(ctx)
}