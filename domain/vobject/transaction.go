package vobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Amount   int       `json:"amount" bson:"amount"`
	From     uuid.UUID `json:"from" bson:"from"`
	To       uuid.UUID `json:"to" bson:"to"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}
