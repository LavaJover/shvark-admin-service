package grpcclients

import (
	"context"
	"time"

	ssopb "github.com/LavaJover/shvark-sso-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SSOClient struct {
	conn *grpc.ClientConn
	service ssopb.SSOServiceClient
}

func NewSSOClient(addr string) (*SSOClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		return nil, err
	}

	return &SSOClient{
		conn: conn,
		service: ssopb.NewSSOServiceClient(conn),
	}, nil
}

func (c *SSOClient) Register(request *ssopb.RegisterRequest) (*ssopb.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.Register(ctx, request)
}

func (c *SSOClient) Login(request *ssopb.LoginRequest) (*ssopb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.Login(ctx, request)
}

func (c *SSOClient) ValidateToken(request *ssopb.ValidateTokenRequest) (*ssopb.ValidateTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.ValidateToken(ctx, request)
}