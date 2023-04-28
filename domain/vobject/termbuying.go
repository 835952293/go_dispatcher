package vobject

import (
	"github.com/google/uuid"
)

type TermBuyingDetail struct {
	Amount int       `json:"amount" bson:"amount"`
	From   uuid.UUID `json:"from" bson:"from"`
	To     uuid.UUID `json:"to" bson:"to"`
}
