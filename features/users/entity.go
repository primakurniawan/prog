package users

type Core struct {
	ID       int
	Email    string
	Password string
	Fullname string
	Image    string
}

type Business interface {
	RegisterUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(id int) (Core, error)
	GetUserFollowingById(userId int) ([]Core, error)
	GetUserFollowersById(userId int) ([]Core, error)
	// UpdateUser(id int, data Core) error
	// DeleteUser(id int, data Core) error
}

type Data interface {
	CreateUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(userId int) (Core, error)
	GetUserFollowingById(userId int) ([]Core, error)
	GetUserFollowersById(userId int) ([]Core, error)
	// UpdateUserById(id int, data Core) error
	// DeleteUserById(id int, data Core) error
}
