package users

const ServiceName = "UserService"

type UserService interface {
	Get() string
}

type userService struct {
	n string
}

func New(n string) UserService {
	return &userService{n: n}
}

func (s *userService) Get() string {
	return s.n
}
