package server

import (
	"context"
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/ivansukach/speed-control-grpc/repository"
	"github.com/ivansukach/speed-control-grpc/service"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	sls *service.SpeedLimitService
}

func New(sls *service.SpeedLimitService) *Server {
	return &Server{sls: sls}
}

func (s *Server) Add(ctx context.Context, req *protocol.AddRequest) (*protocol.EmptyResponse, error) {
	record := &repository.Record{
		Date:   req.Record.Date,
		Number: req.Record.Number,
		Speed:  req.Record.Speed,
	}
	err := s.sls.Create(record)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Listing(ctx context.Context, req *protocol.ListingRequest) (*protocol.ListingResponse, error) {
	records, err := s.sls.Listing(req.Date, req.SpeedLimit)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result := make([]*protocol.Record, 0)
	for _, value := range *records {
		rec := protocol.Record{Date: value.Date, Number: value.Number, Speed: value.Speed}
		result = append(result, &rec)
	}
	return &protocol.ListingResponse{Record: result}, nil

}
func (s *Server) GetMinMax(ctx context.Context, req *protocol.GetMinMaxRequest) (*protocol.GetMinMaxResponse, error) {
	min, max, err := s.sls.GetMinMax(req.Date)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	minSpeedRecord := protocol.Record{Date: min.Date, Number: min.Number, Speed: min.Speed}
	maxSpeedRecord := protocol.Record{Date: max.Date, Number: max.Number, Speed: max.Speed}
	return &protocol.GetMinMaxResponse{Min: &minSpeedRecord, Max: &maxSpeedRecord}, nil
}
