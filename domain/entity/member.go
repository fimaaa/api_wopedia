package entity

import (
	"fmt"
	"strings"
	"time"

	"backend_base_app/domain/domerror"
	"backend_base_app/shared/util"

	"github.com/gosimple/slug"
)

const (
	CollectionMember string = "member"
)

type MemberDataID string

func NewMemberDataID(RandomID string) (MemberDataID, error) {

	var obj = MemberDataID(fmt.Sprintf("Member-%s", RandomID))

	return obj, nil
}

func (r MemberDataID) String() string {
	return string(r)
}

type MemberData struct {
	ID             MemberDataID `json:"id" bson:"_id" form:"id"`
	Username       string       `json:"username" bson:"username" form:"username"`
	Fullname       string       `json:"fullname" bson:"fullname" form:"fullname"`
	Password       string       `json:"password" bson:"password" form:"password"`
	MemberType     string       `json:"member_type" bson:"member_type" form:"member_type"`
	IsSuspend      bool         `json:"is_suspend" bson:"is_suspend" form:"is_suspend"`
	CreatedAt      time.Time    `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" bson:"updated_at" form:"updated_at"`
	TokenBroadcast string       `json:"token_broadcast" bson:"token_broadcast" form:"token_broadcast"`
	LastLogin      time.Time    `json:"last_login" bson:"last_login" form:"last_login"`
	DeviceId       string       `json:"id_device" bson:"id_device" form:"id_device"`

	// Info
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Email       string `json:"email" bson:"email"`
	MemberPhoto string `json:"photo_member" bson:"photo_member"`
}

type CreateMemberData struct {
	Username   string `json:"username"`
	Fullname   string `json:"fullname"`
	Password   string `json:"password"`
	MemberType string `json:"member_type"`

	// Info
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	MemberPhoto *string `json:"photo_member"`
}

type MemberReqAuth struct {
	Username       string `json:"username" form:"username"`
	Password       string `json:"password" form:"password"`
	TokenBroadcast string `json:"token_broadcast" form:"token_broadcast"`
	DeviceId       string `json:"id_device" form:"id_device"`
}

type MemberResAuth struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Fullname   string    `json:"fullname"`
	MemberType string    `json:"member_type"`
	IsSuspend  bool      `json:"is_suspend"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	LastLogin      time.Time `json:"last_login"`
	TokenBroadcast string    `json:"token_broadcast"`
	DeviceId       string    `json:"id_device"`

	// Info
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	MemberPhoto string `json:"photo_member"`
}

type MemberDataShown struct {
	ID             string    `json:"id" bson:"_id"`
	Username       string    `json:"username" bson:"username"`
	Fullname       string    `json:"fullname" bson:"fullname"`
	MemberType     string    `json:"member_type" bson:"member_type"`
	IsSuspend      bool      `json:"is_suspend" bson:"is_suspend"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
	TokenBroadcast string    `json:"token_broadcast" bson:"token_broadcast"`
	LastLogin      time.Time `json:"last_login" bson:"last_login"`
	DeviceId       string    `json:"id_device" bson:"id_device"`
	// Info
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Email       string `json:"email" bson:"email"`
	MemberPhoto string `json:"photo_member" bson:"photo_member"`
}

type MemberListShown struct {
	ID          string `json:"id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Fullname    string `json:"fullname" bson:"fullname"`
	MemberType  string `json:"member_type" bson:"member_type"`
	IsSuspend   bool   `json:"is_suspend" bson:"is_suspend"`
	MemberPhoto string `json:"photo_member" bson:"photo_member"`
}

type MemberDataFind struct {
	Username      string     `form:"username"`
	Fullname      string     `form:"fullname"`
	MemberType    string     `form:"member_type"`
	IsSuspend     *bool      `form:"is_suspend"`
	CreatedAtFrom *time.Time `form:"created_at_from"`
	CreatedAtTo   *time.Time `form:"created_at_to"`
	UpdatedAtFrom *time.Time `form:"updated_at_from"`
	UpdatedAtTo   *time.Time `form:"updated_at_to"`
	LastLoginFrom *time.Time `form:"last_login_from"`
	LastLoginTo   *time.Time `form:"last_login_to"`

	// Info
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (r CreateMemberData) ValidateCreate() error {

	if len(strings.TrimSpace(r.Username)) == 0 {
		return UsernameMustNotEmpty
	}
	if len(strings.TrimSpace(r.Password)) == 0 {
		return PasswordMustNotEmpty
	}
	if len(strings.TrimSpace(r.Fullname)) == 0 {
		return FullNameMustNotEmpty
	}
	if len(strings.TrimSpace(r.MemberType)) == 0 {
		return MemberTypeMustNotEmpty
	}
	if !((len(strings.TrimSpace(*r.Email)) >= 1) || (len(strings.TrimSpace(*r.PhoneNumber)) >= 1)) {
		return PhoneNumberOrEmailMustNotEmpty
	}

	return nil
}

func (r MemberData) ToShown() MemberDataShown {
	fmt.Println("MAP DATA TO SHOWN", r)
	return MemberDataShown{
		ID:         r.ID.String(),
		Username:   r.Username,
		Fullname:   r.Fullname,
		MemberType: r.MemberType,
		IsSuspend:  r.IsSuspend,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,

		LastLogin:      r.LastLogin,
		TokenBroadcast: r.TokenBroadcast,
		DeviceId:       r.DeviceId,

		// Info
		PhoneNumber: r.PhoneNumber,
		Email:       r.Email,
		MemberPhoto: r.MemberPhoto,
	}
}

func NewMemberData(req CreateMemberData) (*MemberData, error) {

	randomId := util.GenerateID()
	id, err := NewMemberDataID(randomId)
	if err != nil {
		return nil, err
	}

	var obj MemberData
	//automapper
	err = util.Automapper(req, &obj)
	if err != nil {
		return nil, err
	}
	//custom fields
	obj.ID = id
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()
	obj.IsSuspend = false

	obj.MemberType = slug.Make(strings.ToLower(obj.MemberType))

	err = req.ValidateCreate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

const UsernameMustNotEmpty domerror.ErrorType = "ER1000 username must not empty"      //
const FullNameMustNotEmpty domerror.ErrorType = "ER1000 username must not empty"      //
const PasswordMustNotEmpty domerror.ErrorType = "ER1000 password must not empty"      //
const MemberTypeMustNotEmpty domerror.ErrorType = "ER1000 member type must not empty" //
const PhoneNumberOrEmailMustNotEmpty domerror.ErrorType = "ER1000 Phone Number or Email must be filled"

//const UsernameMustNotEmpty domerror.ErrorType = "ER1000 username must not empty" //
