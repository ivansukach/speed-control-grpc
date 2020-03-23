package main

import (
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/ivansukach/speed-control-grpc/repository"
	"github.com/ivansukach/speed-control-grpc/server"
	"github.com/ivansukach/speed-control-grpc/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format("02.01.2006 15:04:05") // Аналогично: YYYY.MM.DD-hh.mm.ss
}
func formatTimeOnlyDate(t time.Time) string {
	return t.Format("02.01.2006") // Аналогично: YYYY.MM.DD-hh.mm.ss
}
func main() {
	date := formatTimeOnlyDate(time.Now())
	log.Println(date)
	log.Println(formatTime(time.Now()))
	rps := repository.New(&date)
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
