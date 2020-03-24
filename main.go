package main

import (
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/ivansukach/speed-control-grpc/repository"
	"github.com/ivansukach/speed-control-grpc/server"
	"github.com/ivansukach/speed-control-grpc/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	rps := repository.New()
	sls := service.New(rps)
	srv := server.New(sls)
	s := grpc.NewServer()
	protocol.RegisterSpeedControlServiceServer(s, srv)
	listener, err := net.Listen("tcp", ":1323")
	if err != nil {
		log.Error(err)
		return
	}
	s.Serve(listener)
}
