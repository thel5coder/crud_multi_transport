package main

import (
	"crud_multi_transport/grpc_server/client"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main(){
	err := godotenv.Load("../../../.env")
	if err != nil{
		log.Fatal(err)
	}
	clt := client.NewUserGrRPCClient(os.Getenv("APP_GRPC_PORT"))
}
