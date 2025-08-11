package db

// here storage is a struct, which takes the responsibility of creating objects out of the userService

type Storage struct {
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}
