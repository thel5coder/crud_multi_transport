package client

import (
	"google.golang.org/grpc"
	"log"
)

type Client interface {
	OpenConn(host string) *grpc.ClientConn
}

type GRPCCLient struct {
	host string
}

func (c *GRPCCLient) OpenConn() *grpc.ClientConn{
	cc ,err := grpc.Dial(c.host,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	return cc
}

func NewGRPCClient(host string) *GRPCCLient {
	return &GRPCCLient{host: host}
}