package service

import (
	"context"
	"encoding/json"
	"s0709-22/internal/proxyproto"
	"s0709-22/services/connect-service/internal/userdb"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Service ...
type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
	conn    *pgxpool.Pool
	storage *userdb.Queries
}

// New ...
func New(uri string) (*Service, error) {
	connCfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), connCfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		conn:    conn,
		storage: userdb.New(conn),
	}, nil
}

// AuthData ...
type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Connect ...
func (s *Service) Connect(ctx context.Context, request *proxyproto.ConnectRequest) (*proxyproto.ConnectResponse, error) {
	auth := &AuthData{}

	if err := json.Unmarshal(request.Data, auth); err != nil {
		return RespondError(107, "bad request")
	}

	account, err := s.storage.GetUserByUsermame(ctx, auth.Username)
	if err != nil {
		return RespondError(101, "unauthorized")
	}

	if auth.Password != account.Password {
		return RespondError(101, "unauthorized")
	}

	return &proxyproto.ConnectResponse{
		Result: &proxyproto.ConnectResult{
			User: strconv.FormatInt(account.ID, 10),
		},
	}, nil
}
