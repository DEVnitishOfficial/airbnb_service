package db

type UserRepository interface {
	Create() error
	// Other methods like Get, Update, Delete can be added here
}

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) Create() error {
	return nil
}
