package entity

import (
	"fmt"
	"time"
)

const (
	CollectionOrganizerWedding string = "organizerWedding"
)

type OrganizerWeddingDataID string

func NewOrganizerWeddingDataID(RandomID string) (OrganizerWeddingDataID, error) {

	var obj = OrganizerWeddingDataID(fmt.Sprintf("OrganizerWedding-%s", RandomID))

	return obj, nil
}

func (r OrganizerWeddingDataID) String() string {
	return string(r)
}

type OrganizerWeddingData struct {
	ID            OrganizerWeddingDataID `json:"id" bson:"_id" form:"id"`
	Name          string                 `json:"name" bson:"name" form:"name"`
	MinRangePrice string                 `json:"min_range_price" bson:"min_range_price" form:"min_range_price"`
	MaxRangePrice string                 `json:"max_range_price" bson:"max_range_price" form:"max_range_price"`
	Phone         string                 `json:"phone" bson:"phone" form:"phone"`
	PIC           string                 `json:"pic" bson:"pic" form:"pic"`
	CreatedAt     time.Time              `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at" bson:"updated_at" form:"updated_at"`
}

type EditOrganizerWeddingData struct {
	ID            OrganizerWeddingDataID `json:"id" bson:"_id" form:"id"`
	Name          *string                `json:"name" bson:"name" form:"name"`
	MinRangePrice *string                `json:"min_range_price" bson:"min_range_price" form:"min_range_price"`
	MaxRangePrice *string                `json:"max_range_price" bson:"max_range_price" form:"max_range_price"`
	Phone         *string                `json:"phone" bson:"phone" form:"phone"`
	PIC           *string                `json:"pic" bson:"pic" form:"pic"`
	UpdatedAt     time.Time              `json:"updated_at" bson:"updated_at" form:"updated_at"`
}

type FindOrganizerWeddingData struct {
	Name          *string    `json:"name" bson:"name" form:"name"`
	MinRangePrice *string    `json:"min_range_price" bson:"min_range_price" form:"min_range_price"`
	MaxRangePrice *string    `json:"max_range_price" bson:"max_range_price" form:"max_range_price"`
	Phone         *string    `json:"phone" bson:"phone" form:"phone"`
	PIC           *string    `json:"pic" bson:"pic" form:"pic"`
	CreatedAtFrom *time.Time `json:"created_at_from" form:"created_at_from" filter:"gte"`
	CreatedAtTo   *time.Time `json:"created_at_to" form:"created_at_to"`
	UpdatedAtFrom *time.Time `json:"updated_at_from" form:"updated_at_from"`
	UpdatedAtTo   *time.Time `json:"updated_at_to" form:"updated_at_to"`
}
