package service

type User struct {
	Nama string
}

type UserService struct {
	db []*User
}

var dataUser UserService

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService() UserIface {
	return &UserService{}
}

func (u *UserService) Register(user *User) string {

	dataUser.db = append(dataUser.db, user)
	return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {

	return dataUser.db
}
