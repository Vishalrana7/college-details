package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {

	return &LoggingService{
		next: next,
	}

}

func (s *LoggingService) GetCollegeDetail(ctx context.Context) (detail *CollegeDetail, err error) {

	defer func(staret time.Time) {
		fmt.Printf("detail=%v err=%s took=%v\n",detail.Detail,err,time.Since(staret))

	}(time.Now())


	return s.next.GetCollegeDetail(ctx)

}