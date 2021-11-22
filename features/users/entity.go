package users

type Core struct {
	ID       int
	Email    string
	Password string
	Fullname string
	Image    string
}

type Business interface {
	RegisterUser(data Core) error
	GetUsersByFullname(fullname string) ([]Core, error)
	GetUserById(id int) (Core, error)
	// UpdateUser(id int, data Core) error
	// DeleteUser(id int, data Core) error
}

type Data interface {
	CreateUser(data Core) error
	GetUsersByFullname(fullname string) ([]Core, error)
	GetUserById(userId int) (Core, error)
	// GetUserFollowing(userId int) ([]Core, error)
	// GetUserFollowers(userId int) ([]Core, error)
	// UpdateUserById(id int, data Core) error
	// DeleteUserById(id int, data Core) error
}
