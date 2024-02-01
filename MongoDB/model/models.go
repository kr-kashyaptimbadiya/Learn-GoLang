package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	//primitive.ObjectID is a type provided by the
	//go.mongodb.org/mongo-driver/bson/primitive package for storing MongoDB document IDs.
	Movie   string `json:"movie,omitempty"`
	Watched bool   `json:"watched,omitempty"`
}

//The json:"watched,omitempty" tag specifies how the field should be encoded and decoded in JSON format.
