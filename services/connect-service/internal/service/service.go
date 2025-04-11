package service

import (
	"context"
	"encoding/json"
	"s0709-22/internal/proxyproto"
)

// Service ...
type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
}

// New ...
func New() *Service {
	return &Service{}
}

// AuthData ...
type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Connect ...
func (s *Service) Connect(ctx context.Context, request *proxyproto.ConnectRequest) (*proxyproto.ConnectResponse, error) {
	auth := &AuthData{}

	response := &proxyproto.ConnectResponse{}

	if err := json.Unmarshal(request.Data, auth); err != nil {
		response.Error = &proxyproto.Error{
			Code:    107,
			Message: "bad request",
		}
		return response, nil
	}

	if auth.Username != "alex" || auth.Password != "qwerty" {
		response.Error = &proxyproto.Error{
			Code:    101,
			Message: "unauthorized",
		}
		return response, nil
	}

	response.Result = &proxyproto.ConnectResult{
		User: "alex",
	}

	return response, nil
}
