package pingpong

import (
	"configservice/models/dto"
	"configservice/repository/mongo"
	"context"
)

type Service struct {
	repo *mongo.Repository
}

func NewService(repo *mongo.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Ping(ctx context.Context) (*dto.PingResponse, error) {
	err := s.repo.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.PingResponse{Message: "DB okay"}, nil
}
