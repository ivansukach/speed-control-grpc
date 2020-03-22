package server

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/book-service/repositories"
	"github.com/ivansukach/book-service/service"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	bs *service.BookService
}

func New(bs *service.BookService) *Server {
	return &Server{bs: bs}
}

func (s *Server) Add(ctx context.Context, req *protocol.AddRequest) (*protocol.EmptyResponse, error) {
	book := &repositories.Book{
		Id:            req.Book.Id,
		Title:         req.Book.Title,
		Author:        req.Book.Author,
		Genre:         req.Book.Genre,
		Edition:       req.Book.Edition,
		NumberOfPages: req.Book.NumberOfPages,
		Year:          req.Book.Year,
		Amount:        req.Book.Amount,
		IsPopular:     req.Book.IsPopular,
		InStock:       req.Book.InStock,
	}
	err := s.bs.Create(book)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Delete(ctx context.Context, req *protocol.DeleteRequest) (*protocol.EmptyResponse, error) {
	err := s.bs.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Update(ctx context.Context, req *protocol.UpdateRequest) (*protocol.EmptyResponse, error) {
	book := &repositories.Book{
		Id:            req.Book.Id,
		Title:         req.Book.Title,
		Author:        req.Book.Author,
		Genre:         req.Book.Genre,
		Edition:       req.Book.Edition,
		NumberOfPages: req.Book.NumberOfPages,
		Year:          req.Book.Year,
		Amount:        req.Book.Amount,
		IsPopular:     req.Book.IsPopular,
		InStock:       req.Book.InStock,
	}
	err := s.bs.Update(book)
	if err != nil {
		log.Error(err)
		return &protocol.EmptyResponse{}, err
	}
	return &protocol.EmptyResponse{}, nil
}
func (s *Server) Get(ctx context.Context, req *protocol.GetRequest) (*protocol.GetResponse, error) {
	book, err := s.bs.Read(req.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	response := &protocol.Book{
		Id:            book.Id,
		Title:         book.Title,
		Author:        book.Author,
		Genre:         book.Genre,
		Edition:       book.Edition,
		NumberOfPages: book.NumberOfPages,
		Year:          book.Year,
		Amount:        book.Amount,
		IsPopular:     book.IsPopular,
		InStock:       book.InStock,
	}
	return &protocol.GetResponse{Book: response}, nil
}
