package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	// Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Age      int    `json:"age,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
}
