package auth

import "prog/features/users"

type Core struct {
	Token string
}

type Business interface {
	Login(data users.Core) (accessTokenCore Core, refreshTokenCore Core, err error)
	ReLogin(data Core, userId int) (accessTokenCore Core, err error)
	Logout(data Core) error
}

type Data interface {
	AddRefreshToken(data Core) error
	VerifyRefreshToken(data Core) error
	DeleteRefreshToken(data Core) (err error)
	VerifyUserCredential(data users.Core) (users.Core, error)
}
