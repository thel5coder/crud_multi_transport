package client

import (
	"context"
	"crud_multi_transport/grpc_server/protobuf"
	"log"
)

type UserClient interface {
	Browse(search,sort,order string,page, limit int) *protobuf.Response
	Read(ID string) *protobuf.Response
	Edit(fullName,email,password,mobilePhone,ID string) *protobuf.Response
	Add(fullName,email,password,mobilePhone string) *protobuf.Response
	Delete(ID string) *protobuf.Response
}

type UserGRPCClient struct {
	Client *GRPCCLient
}

func (c *UserGRPCClient) Browse(search,sort,order string,page,limit int) *protobuf.Response{
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := protobuf.NewUserClient(conn)
	req := protobuf.BrowseRequest{
		Search:               search,
		Sort:                 sort,
		Order:                order,
		Page:                 int32(page),
		Limit:                int32(limit),
	}
	res,err := client.Browse(context.Background(),&req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func (c *UserGRPCClient) Read(ID string) *protobuf.Response{
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := protobuf.NewUserClient(conn)
	res,err := client.Read(context.Background(),&protobuf.ReadRequest{ID: ID})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func (c *UserGRPCClient) Edit(fullName,email,password,mobilePhone,ID string) *protobuf.Response{
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := protobuf.NewUserClient(conn)
	res,err := client.Edit(context.Background(),&protobuf.EditRequest{
		ID:                   ID,
		UserRequest:          &protobuf.UserRequest{
			FullName:             fullName,
			Email:                email,
			Password:             password,
			MobilePhone:          mobilePhone,
		},
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func (c *UserGRPCClient) Add(fullName,email,password,mobilePhone string) *protobuf.Response{
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := protobuf.NewUserClient(conn)
	res,err := client.Add(context.Background(),&protobuf.UserRequest{
		FullName:             fullName,
		Email:                email,
		Password:             password,
		MobilePhone:          mobilePhone,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func (c *UserGRPCClient) Delete(ID string) *protobuf.Response{
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := protobuf.NewUserClient(conn)
	res,err := client.Delete(context.Background(),&protobuf.DeleteRequest{ID: ID})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func NewUserGrRPCClient(host string) *UserGRPCClient{
	return &UserGRPCClient{Client: NewGRPCClient(host)}
}

