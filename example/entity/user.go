// Code generated by generate_code script - DO NOT EDIT.
package entity

import (
	"json"
	"time"

	"github.com/juju/errors"
	"go.knocknote.io/rapidash"
)

type User struct {
	ID            uint64
	Uuid          string
	AccessToken   string
	OutsideUserID string
	Name          string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type Users []*User

func (e *User) EncodeRapidash(enc rapidash.Encoder) error {
	if e.ID != 0 {
		enc.Uint64("id", e.ID)
	}
	enc.Uint64("id", e.ID)
	enc.String("uuid", e.Uuid)
	enc.String("access_token", e.AccessToken)
	enc.String("outside_user_id", e.OutsideUserID)
	enc.String("name", e.Name)
	enc.TimePtr("created_at", e.CreatedAt)
	enc.TimePtr("updated_at", e.UpdatedAt)
	return enc.Error()
}

func (e *Users) EncodeRapidash(enc rapidash.Encoder) error {
	for _, v := range *e {
		if err := v.EncodeRapidash(enc.New()); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (e *User) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.Uuid = dec.String("uuid")
	e.AccessToken = dec.String("access_token")
	e.OutsideUserID = dec.String("outside_user_id")
	e.Name = dec.String("name")
	e.CreatedAt = dec.TimePtr("created_at")
	e.UpdatedAt = dec.TimePtr("updated_at")
	return dec.Error()
}

func (e *Users) DecodeRapidash(dec rapidash.Decoder) error {
	count := dec.Len()
	*e = make([]*User, count)
	for i := 0; i < count; i++ {
		var v User
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return errors.Trace(err)
		}
		(*e)[i] = &v
	}
	return nil
}

func (e *User) Struct() *rapidash.Struct {
	s := rapidash.NewStruct("users")
	s.FieldUint64("id")
	s.FieldString("uuid")
	s.FieldString("access_token")
	s.FieldString("outside_user_id")
	s.FieldString("name")
	s.FieldTime("created_at")
	s.FieldTime("updated_at")
	return s
}

func (e *User) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"accessToken": e.AccessToken,
		"name":        e.Name,
		"uuid":        e.Uuid,
	}
	if e.CreatedAt != nil {
		m["createdAt"] = e.CreatedAt.Unix()
	}
	if e.UpdatedAt != nil {
		m["updatedAt"] = e.UpdatedAt.Unix()
	}
	b, err := json.Marshal(m)
	return b, errors.Trace(err)
}
