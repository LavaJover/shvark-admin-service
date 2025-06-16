package grpcclients

import (
	"context"
	"time"

	authzpb "github.com/LavaJover/shvark-authz-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthzClient struct {
	conn *grpc.ClientConn
	service authzpb.AuthzServiceClient
}

func NewAuthzClient(addr string) (*AuthzClient, error) {
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

	return &AuthzClient{
		conn: conn,
		service: authzpb.NewAuthzServiceClient(conn),
	}, nil
}

func (c *AuthzClient) AssignRole(request *authzpb.AssignRoleRequest) (*authzpb.AssignRoleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.AssignRole(ctx, request)
}

func (c *AuthzClient) RevokeRole(request *authzpb.RevokeRoleRequest) (*authzpb.RevokeRoleResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.RevokeRole(ctx, request)
}

func (c *AuthzClient) AddPolicy(request *authzpb.AddPolicyRequest) (*authzpb.AddPolicyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.AddPolicy(ctx, request)
}

func (c *AuthzClient) DeletePolicy(request *authzpb.DeletePolicyRequest) (*authzpb.DeletePolicyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.DeletePolicy(ctx, request)
}

func (c *AuthzClient) CheckPermission(request *authzpb.CheckPermissionRequest) (*authzpb.CheckPermissionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.service.CheckPermission(ctx, request)
}