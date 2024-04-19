package usecase

type Repository interface {
	CreateUser(username, password string) error
	CheckUser(username string) (bool, error)
	GetUserPassword(username string) (string, error)
}

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
