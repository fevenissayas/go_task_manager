package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}

type TaskRepository interface {
    Add(task *Task) (primitive.ObjectID, error)
    GetAll() ([]*Task, error)
    GetByID(id primitive.ObjectID) (*Task, error)
    DeleteByID(id primitive.ObjectID) error
    UpdateByID(id primitive.ObjectID, task *Task) error
}