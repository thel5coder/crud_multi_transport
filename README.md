## Crud REST API, CLI, and Protobuf With Golang

You need install
Go 1.13+ 
Postgresql 9.5+
Redis

create database with name crud_multiple_transport, then execute sql file in db directory

to run http rest go to http_server directory then run command go run main.go
to run protobuf go to grpc_server then run command go run main.go
import postman collection in the postman directory

to run cli mode go to cli directory and then execute command make
run it with 

./bin/rli browse for get list user

./bin/rli read for get single user   

./bin/rli edit for edit existing user

./bin/rli add for add new user
./bin/rli delete for delete existing users

./bin/rli browse/read/edit/add/delete --help to see description usage

