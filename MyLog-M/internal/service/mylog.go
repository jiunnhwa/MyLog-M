package service

import (
	"MyLog-M/errors"
	"MyLog-M/internal/domain"
)

//go:generate mockgen -source=./mylog.go -destination=./mock/mock.go -package=mock
type repository interface {
	Tail(limit int64) (*[]domain.Data, error)
}

// Service/Usecase/Controller
type Service struct {
	repo repository
}

func New(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Tail(limit int64) (*[]domain.Data, error) {
	res, err := s.repo.Tail(limit)
	if err != nil {
		return &[]domain.Data{}, errors.Error{
			ErrorCode:     "DATABASE_ERROR",
			InternalError: err,
			Message:       "error occurred in database",
		}
	}
	return res, nil
}
