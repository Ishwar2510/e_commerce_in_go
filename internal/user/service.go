package user

// this is just a im memory psql
type Service struct {
	users  map[int]User
	nextId int
}

// Initializing the service
func NewService() *Service {
	return &Service{
		users:  make(map[int]User),
		nextId: 1,
	}
}

func (s *Service) Create(name, email string) User {
	user := User{
		Id:    s.nextId,
		Name:  name,
		Email: email,
	}
	s.users[user.Id] = user
	s.nextId++

	return user
}
func (s *Service) Get(id int) (User, bool) {
	u, ok := s.users[id]
	return u, ok
}
