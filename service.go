package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCollegeDetail(context.Context) (*CollegeDetail, error)

}

type CollegeDetailService struct {
	url string
}

func NewCollegeDetailService(url string) Service {
	return &CollegeDetailService{
		url: url,
	}
}

func (s *CollegeDetailService) GetCollegeDetail(ctx context.Context) (*CollegeDetail,error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	detail := &CollegeDetail{}
	if err := json.NewDecoder(resp.Body).Decode(detail); err != nil {
		return nil, err
	}

	return detail, nil
}

