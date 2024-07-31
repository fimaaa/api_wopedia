package entity

import (
	"fmt"
	"time"
)

const (
	CollectionEventWedding string = "eventWedding"
)

type EventWeddingDataID string

func NewEventWeddingDataID(RandomID string) (EventWeddingDataID, error) {

	var obj = EventWeddingDataID(fmt.Sprintf("EventWedding-%s", RandomID))

	return obj, nil
}

func (r EventWeddingDataID) String() string {
	return string(r)
}

type EventWeddingData struct {
	ID           EventWeddingDataID `json:"id" bson:"_id" form:"id"`
	Name         string             `json:"name" bson:"name" form:"name"`
	Phone        string             `json:"phone" bson:"phone" form:"phone"`
	TotalInvited int                `json:"total_invited" bson:"total_invited" form:"total_invited"`
	City         string             `json:"city" bson:"city" form:"city"`
	Address      string             `json:"address" bson:"address" form:"address"`
	Latitude     string             `json:"latitude" bson:"latitude" form:"latitude"`
	Longitude    string             `json:"longitude" bson:"longitude" form:"longitude"`
	DateStart    string             `json:"date_start" bson:"date_start" form:"date_start"`
	DateEnd      string             `json:"date_end" bson:"date_end" form:"date_end"`
	IDWeddingOrg string             `json:"id_wedding_org" bson:"id_wedding_org" form:"id_wedding_org"`
	IDUser       string             `json:"id_user" bson:"id_user" form:"id_user"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at" form:"updated_at"`
	Status       string             `json:"status" bson:"status" form:"status"`
}

type EditEventWeddingData struct {
	ID           EventWeddingDataID `json:"id" bson:"_id" form:"id"`
	Name         *string            `json:"name" bson:"name" form:"name"`
	Phone        *string            `json:"phone" bson:"phone" form:"phone"`
	City         *string            `json:"city" bson:"city" form:"city"`
	TotalInvited *int               `json:"total_invited" bson:"total_invited" form:"total_invited"`
	Address      *string            `json:"address" bson:"address" form:"address"`
	Latitude     *string            `json:"latitude" bson:"latitude" form:"latitude"`
	Longitude    *string            `json:"longitude" bson:"longitude" form:"longitude"`
	DateStart    *string            `json:"date_start" bson:"date_start" form:"date_start"`
	DateEnd      *string            `json:"date_end" bson:"date_end" form:"date_end"`
	IDWeddingOrg *string            `json:"id_wedding_org" bson:"id_wedding_org" form:"id_wedding_org"`
	IDUser       *string            `json:"id_user" bson:"id_user" form:"id_user"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at" form:"updated_at"`
	Status       *string            `json:"status" bson:"status" form:"status"`
}

type FindEventWeddingData struct {
	Name            *string    `json:"name" bson:"name" form:"name"`
	Phone           *string    `json:"phone" bson:"phone" form:"phone"`
	MinTotalInvited *int       `json:"min_total_invited" bson:"min_total_invited" form:"min_total_invited" filter:"skip"`
	MaxTotalInvited *int       `json:"max_total_invited" bson:"max_total_invited" form:"max_total_invited" filter:"skip"`
	City            *string    `json:"city" bson:"city" form:"city"`
	Address         *string    `json:"address" bson:"address" form:"address"`
	DateStart       *string    `json:"date_start" bson:"date_start" form:"date_start"`
	DateEnd         *string    `json:"date_end" bson:"date_end" form:"date_end"`
	IDWeddingOrg    *string    `json:"id_wedding_org" bson:"id_wedding_org" form:"id_wedding_org"`
	IDUser          *string    `json:"id_user" bson:"id_user" form:"id_user"`
	CreatedAtFrom   *time.Time `json:"created_at_from" form:"created_at_from" filter:"skip"`
	CreatedAtTo     *time.Time `json:"created_at_to" form:"created_at_to" filter:"skip"`
	UpdatedAtFrom   *time.Time `json:"updated_at_from" form:"updated_at_from" filter:"skip"`
	UpdatedAtTo     *time.Time `json:"updated_at_to" form:"updated_at_to" filter:"skip"`
	Status          *string    `json:"status" bson:"status" form:"status"`
}
