package auth

import (
	"context"
	"fmt"
)

var (
	emptyResponse = &Empty{}

	registerErrorTemplate = "cannot register user: %w"
	loginErrorTemplate    = "cannot login user: %w"
)

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*Empty, error) {
	err := s.useCase.Register(request.Username, request.Password)
	if err != nil {
		err = fmt.Errorf(registerErrorTemplate, err)
	}

	return emptyResponse, err
}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*Empty, error) {
	err := s.useCase.Login(request.Username, request.Password)
	if err != nil {
		err = fmt.Errorf(loginErrorTemplate, err)
	}

	return emptyResponse, err
}

func (s *Server) mustEmbedUnimplementedAuthServiceServer() {}
