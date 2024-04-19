package auth

type UseCase interface {
	Register(username, password string) error
	Login(username, password string) error
}

type Server struct {
	UnimplementedAuthServiceServer
	useCase UseCase
}

func NewServer(useCase UseCase) *Server {
	return &Server{
		UnimplementedAuthServiceServer: UnimplementedAuthServiceServer{},
		useCase:                        useCase,
	}
}
