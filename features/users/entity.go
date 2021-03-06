package users

type Core struct {
	ID       int
	Email    string
	Password string
	Fullname string
	Image    string
}

type Business interface {
	CreateUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(id int) (Core, error)
	UpdateUserById(userId int, data Core) error
	DeleteUserById(userId int) error
}

type Data interface {
	CreateUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(userId int) (Core, error)
	UpdateUserById(userId int, data Core) error
	DeleteUserById(userId int) error
}
