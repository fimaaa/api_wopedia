package entity

import "fmt"

const (
	CollectionRelationUserWo string = "relationUserWo"
)

type RelationUserWoDataID string

func NewRelationUserWoDataID(RandomID string) (RelationUserWoDataID, error) {

	var obj = RelationUserWoDataID(fmt.Sprintf("RelationUserWo-%s", RandomID))

	return obj, nil
}

func (r RelationUserWoDataID) String() string {
	return string(r)
}

type RelationUserWoData struct {
	ID           RelationUserWoDataID `json:"id" bson:"_id" form:"id"`
	IDWeddingOrg string               `json:"id_wedding_org" bson:"id_wedding_org" form:"id_wedding_org"`
	IDUser       string               `json:"id_user" bson:"id_user" form:"id_user"`
}

type FindRelationUserWoData struct {
	IDWeddingOrg *string `json:"id_wedding_org" bson:"id_wedding_org" form:"id_wedding_org"`
	IDUser       *string `json:"id_user" bson:"id_user" form:"id_user"`
}
