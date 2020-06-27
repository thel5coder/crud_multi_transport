package main

import (
	"context"
	"crud_multi_transport/db"
	"crud_multi_transport/grpc_server/endpoint"
	"crud_multi_transport/grpc_server/protobuf"
	"crud_multi_transport/grpc_server/service"
	"crud_multi_transport/grpc_server/transport"
	"crud_multi_transport/helpers/jwe"
	"crud_multi_transport/helpers/jwt"
	"crud_multi_transport/helpers/str"
	"crud_multi_transport/usecase"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main(){
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	//jwe

	jweCredential := jwe.Credential{
		KeyLocation: os.Getenv("PRIVATE_KEY"),
		Passphrase:  os.Getenv("PASSPHRASE"),
	}

	//setup redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	//setup db connection
	dbInfo := db.Connection{
		Host:     os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	database, err := dbInfo.DbConnect()
	if err != nil {
		panic(err)
	}

	//jwtconfig
	jwtConfig := middleware.JWTConfig{
		Claims:     &jwt.CustomClaims{},
		SigningKey: []byte(os.Getenv("SECRET")),
	}

	//jwt credential
	jwtCred := jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}
	useCaseContract := usecase.UcContract{
		DB:        database,
		Redis:     redisClient,
		Jwe:       jweCredential,
		JwtConfig: jwtConfig,
		JwtCred:   jwtCred,
	}
	ctx := context.Background()
	userService := &service.UserService{UcContract:&useCaseContract}
	errors := make(chan error)

	go func() {
		listener,err := net.Listen("tcp",os.Getenv("APP_GRPC_PORT"))
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := grpc.NewServer()
		protobuf.RegisterUserServer(gRPCServer,transport.NewGRPCServer(ctx,endpoint.Endpoints{
			BrowseEndpoint: endpoint.MakeBrowseEndPoint(userService),
			ReadEndPoint:   endpoint.MakeReadEndPoint(userService),
			EditEndPoint:   endpoint.MakeEditEndPoint(userService),
			AddEndPoint:    endpoint.MakeAddEndPoint(userService),
			DeleteEndPoint: endpoint.MakeDeleteEndPoint(userService),
		}))

		fmt.Println("gRPC listen on "+os.Getenv("APP_GRPC_PORT"))
		errors <- gRPCServer.Serve(listener)
	}()

	fmt.Println(<-errors)
}
