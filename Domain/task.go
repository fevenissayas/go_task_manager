package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
}

type TaskRepository interface {
    Add(task *Task) (primitive.ObjectID, error)
    GetAll() ([]*Task, error)
    GetByID(id primitive.ObjectID) (*Task, error)
    DeleteByID(id primitive.ObjectID) error
    UpdateByID(id primitive.ObjectID, task *Task) error
}