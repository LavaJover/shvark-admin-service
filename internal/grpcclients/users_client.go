package grpcclients

import (
	"context"
	"time"

	userpb "github.com/LavaJover/shvark-user-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	conn *grpc.ClientConn
	service userpb.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
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

	return &UserClient{
		conn: conn,
		service: userpb.NewUserServiceClient(conn),
	}, nil
}

func (c *UserClient) CreateUser(request *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.CreateUser(ctx, request)
}

func (c *UserClient) GetUsers(request *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.GetUsers(ctx, request)
}