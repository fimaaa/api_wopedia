package entity

import (
	"fmt"
	"time"
)

const (
	CollectionEventTestimony string = "eventTestimony"
)

type EventTestimonyDataID string

func NewEventTestimonyDataID(RandomID string) (EventTestimonyDataID, error) {

	var obj = EventTestimonyDataID(fmt.Sprintf("EventTestimony-%s", RandomID))

	return obj, nil
}

func (r EventTestimonyDataID) String() string {
	return string(r)
}

type EventTestimonyData struct {
	ID        EventTestimonyDataID `json:"id" bson:"_id" form:"id"`
	Name      string               `json:"name" bson:"name" form:"name"`
	IDGuest   string               `json:"id_guest" bson:"id_guest" form:"id_guest"`
	Signature string               `json:"signature" bson:"signature" form:"signature"`
	Message   string               `json:"message" bson:"message" form:"message"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at" form:"updated_at"`
}

type EditEventTestimonyData struct {
	ID        EventTestimonyDataID `json:"id" bson:"_id" form:"id"`
	Name      *string              `json:"name" bson:"name" form:"name"`
	IDGuest   *string              `json:"id_guest" bson:"id_guest" form:"id_guest"`
	Signature *string              `json:"signature" bson:"signature" form:"signature"`
	Message   *string              `json:"message" bson:"message" form:"message"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at" form:"updated_at"`
}

type FindEventTestimonyData struct {
	Name          *string    `json:"name" bson:"name" form:"name"`
	IDGuest       *string    `json:"id_guest" bson:"id_guest" form:"id_guest"`
	Signature     *string    `json:"signature" bson:"signature" form:"signature"`
	Message       *string    `json:"message" bson:"message" form:"message"`
	CreatedAtFrom *time.Time `json:"created_at_from" form:"created_at_from" filter:"skip"`
	CreatedAtTo   *time.Time `json:"created_at_to" form:"created_at_to" filter:"skip"`
	UpdatedAtFrom *time.Time `json:"updated_at_from" form:"updated_at_from" filter:"skip"`
	UpdatedAtTo   *time.Time `json:"updated_at_to" form:"updated_at_to" filter:"skip"`
}
