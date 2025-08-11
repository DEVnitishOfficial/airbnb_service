package db

// here storage is a struct, which takes the responsibility of creating objects out of the userService

type Storage struct { // facilitates dependency injection for user repository
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}
