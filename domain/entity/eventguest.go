package entity

import (
	"fmt"
	"time"
)

const (
	CollectionEventGuest string = "eventGuest"
)

type EventGuestDataID string

func NewEventGuestDataID(RandomID string) (EventGuestDataID, error) {

	var obj = EventGuestDataID(fmt.Sprintf("EventGuest-%s", RandomID))

	return obj, nil
}

func (r EventGuestDataID) String() string {
	return string(r)
}

type EventGuestData struct {
	ID               EventGuestDataID `json:"id" bson:"_id" form:"id"`
	Name             string           `json:"name" bson:"name" form:"name"`
	From             string           `json:"from" bson:"from" form:"from"`
	Qty              int              `json:"qty" bson:"qty" form:"qty"`
	IsOverQtyAllowed bool             `json:"is_over_qty_allowed" bson:"is_over_qty_allowed" form:"is_over_qty_allowed"`
	IDEvent          string           `json:"id_event" bson:"id_event" form:"id_event"`
	CreatedAt        time.Time        `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at" bson:"updated_at" form:"updated_at"`
}

type EditEventGuestData struct {
	ID               EventGuestDataID `json:"id" bson:"_id" form:"id"`
	Name             *string          `json:"name" bson:"name" form:"name"`
	From             *string          `json:"from" bson:"from" form:"from"`
	Qty              *int             `json:"qty" bson:"qty" form:"qty"`
	IsOverQtyAllowed *bool            `json:"is_over_qty_allowed" bson:"is_over_qty_allowed" form:"is_over_qty_allowed"`
	UpdatedAt        time.Time        `json:"updated_at" bson:"updated_at" form:"updated_at"`
}

type FindEventGuestData struct {
	Name             *string    `json:"name" bson:"name" form:"name"`
	From             *string    `json:"from" bson:"from" form:"from"`
	MinQty           *int       `json:"min_qty" bson:"min_qty" form:"min_qty" filter:"lte"`
	MaxQty           *int       `json:"max_qty" bson:"max_qty" form:"max_qty" filter:"gte"`
	IsOverQtyAllowed *bool      `json:"is_over_qty_allowed" bson:"is_over_qty_allowed" form:"is_over_qty_allowed"`
	CreatedAtFrom    *time.Time `json:"created_at_from" form:"created_at_from" filter:"skip"`
	CreatedAtTo      *time.Time `json:"created_at_to" form:"created_at_to" filter:"skip"`
	UpdatedAtFrom    *time.Time `json:"updated_at_from" form:"updated_at_from" filter:"skip"`
	UpdatedAtTo      *time.Time `json:"updated_at_to" form:"updated_at_to" filter:"skip"`
}
