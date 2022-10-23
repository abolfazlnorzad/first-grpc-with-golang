package main

import (
	"app/grpcserver"
	"app/user"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	op := flag.String("op", "c", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		runGrpcServer()
	case "c":
		runGrpcClient()
	}
}

func runGrpcServer() {
	grpclog.Infoln("start grpc server")
	l, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalln("Failed to listen ", err)
	}
	grpclog.Infoln("listening on 127.0.0.1:8282 ")

	var opts []grpc.ServerOption
	usersServer, err := grpcserver.NewGrpcServer()
	server := grpc.NewServer(opts...)
	if err != nil {
		log.Fatalln(err)
	}
	user.RegisterUserServiceServer(server, usersServer)
	e := server.Serve(l)
	if e != nil {
		log.Fatalln(e)
	}

}

func runGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)
	input := ""
	fmt.Println("All People? (y/n)")
	fmt.Scanln(&input)
	if strings.EqualFold(input, "y") {
		people, err := client.GetPeople(context.Background(), &user.Request{})
		if err != nil {
			log.Fatalln(err)
		}

		for {
			preson, err := people.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(preson)
		}
		return
	}

	//fmt.Println("name?")
	//fmt.Scanln(&input)
	//
	//person, err := client.GetPerson(context.Background(), &cmd.Request{Name: input})
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(*person)
}
