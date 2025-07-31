package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID
    Username string
    Email    string
    Password string
    Role     string
}

type UserRepository interface {
	Create(user *User) (primitive.ObjectID, error)
	Authenticate(usernameOrEmail, password string) (*User, error)
	GetByID(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	PromoteUser(userID primitive.ObjectID, newRole string) error
	DeleteByID(id primitive.ObjectID) error
}