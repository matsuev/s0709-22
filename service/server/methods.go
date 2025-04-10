package main

import (
	"context"
	"s0709-22/internal/demoapi"
)

// DemoService ...
type DemoService struct {
	demoapi.UnimplementedDemoServiceServer
}

// NewDemoService ...
func NewDemoService() *DemoService {
	return &DemoService{}
}

// Echo ...
func (s *DemoService) Echo(ctx context.Context, request *demoapi.EchoRequest) (*demoapi.EchoRequest, error) {
	return request, nil
}
