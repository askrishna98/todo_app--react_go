package models

// defines structure used in webapp

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToDoList represents a task in the to-do list and stores data in MongoDB

type ToDoList struct {
	ID     primitive.ObjectID // ID is uniq identified for each task, represents as MongoDB ObjectID
	Task   string
	Status bool
}
