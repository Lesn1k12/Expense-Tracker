package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
}

type Expense struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	Description string             `json:"description" bson:"description"`
	OwnerID     string             `json:"owner_id" bson:"owner_id"`
}
