package pingpong

import (
	"configservice/models"
	"configservice/models/dto"
	"configservice/repository/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// add ping in database
func (s *Service) CreatePing(text string) (*dto.PingResponse, error) {
	ping := &models.PingPong{Text: text}
	err := s.repo.SavePingPong(ping)
	if err != nil {
		return nil, err
	}
	return &dto.PingResponse{Message: "data inserted successfully"}, nil
}

// GetAllPings retrieves all ping records.
func (s *Service) GetAllPings(ctx context.Context) (*dto.GetPingResponse, error) {
	pings, err := s.repo.GetAllPings(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.GetPingResponse{Pings: pings}, nil
}

// Get ping by id
func (s *Service) GetPingByID(ctx context.Context, id string) (*models.PingPong, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// PingPongResponse := objectID
	return s.repo.GetPingByID(ctx, objectID)
}

// UpdatePingByID updates a ping record by its ID.
func (s *Service) UpdatePingByID(ctx context.Context, id string, update *models.PingPong) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.repo.UpdatePingByID(ctx, objectID, update)
}

// DeletePingByID deletes a ping record by its ID.
func (s *Service) DeletePingByID(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.repo.DeletePingByID(ctx, objectID)
}
