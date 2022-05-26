package services

import (
	"context"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/db"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/models"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
	pb.UnimplementedDocumentServiceServer
}

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	var document models.Document

	document.Title = req.Title
	document.Content = req.Content

	if result := s.H.DB.Create(&document); result.Error != nil {
		return &pb.AddReply{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.AddReply{
		Status: http.StatusCreated,
		Id:     document.Id,
	}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	var document models.Document

	if result := s.H.DB.First(&document, req.Id); result.Error != nil {
		return &pb.UpdateReply{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	document.Id = req.Id
	document.Title = req.Title
	document.Content = req.Content

	s.H.DB.First(&document, req.Id).Save(&document)

	return &pb.UpdateReply{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	var document models.Document
	if result := s.H.DB.First(&document, req.Id); result.Error != nil {
		return &pb.GetReply{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.GetReply{
		Status:  http.StatusOK,
		Title:   document.Title,
		Content: document.Content,
	}, nil
}

func (s *Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.RemoveReply, error) {
	var document models.Document

	if result := s.H.DB.First(&document, req.Id); result.Error != nil {
		return &pb.RemoveReply{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.RemoveReply{
		Status: http.StatusOK,
	}, nil
}
