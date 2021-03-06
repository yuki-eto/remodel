// Code generated by generate_code script - DO NOT EDIT.
package entity

import (
	"encoding/json"
	"time"

	"github.com/juju/errors"
	"go.knocknote.io/rapidash"
)

type UserFriend struct {
	ID          uint64
	UserID      uint64
	OtherUserID uint64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type UserFriends []*UserFriend

func (e *UserFriend) EncodeRapidash(enc rapidash.Encoder) error {
	if e.ID != 0 {
		enc.Uint64("id", e.ID)
	}
	enc.Uint64("id", e.ID)
	enc.Uint64("user_id", e.UserID)
	enc.Uint64("other_user_id", e.OtherUserID)
	enc.TimePtr("created_at", e.CreatedAt)
	enc.TimePtr("updated_at", e.UpdatedAt)
	return enc.Error()
}

func (e *UserFriends) EncodeRapidash(enc rapidash.Encoder) error {
	for _, v := range *e {
		if err := v.EncodeRapidash(enc.New()); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (e *UserFriend) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.UserID = dec.Uint64("user_id")
	e.OtherUserID = dec.Uint64("other_user_id")
	e.CreatedAt = dec.TimePtr("created_at")
	e.UpdatedAt = dec.TimePtr("updated_at")
	return dec.Error()
}

func (e *UserFriends) DecodeRapidash(dec rapidash.Decoder) error {
	count := dec.Len()
	*e = make([]*UserFriend, count)
	for i := 0; i < count; i++ {
		var v UserFriend
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return errors.Trace(err)
		}
		(*e)[i] = &v
	}
	return nil
}

func (e *UserFriend) Struct() *rapidash.Struct {
	s := rapidash.NewStruct("user_friends")
	s.FieldUint64("id")
	s.FieldUint64("user_id")
	s.FieldUint64("other_user_id")
	s.FieldTime("created_at")
	s.FieldTime("updated_at")
	return s
}

func (e *UserFriend) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{"id": e.ID}
	if e.CreatedAt != nil {
		m["createdAt"] = e.CreatedAt.Unix()
	}
	if e.UpdatedAt != nil {
		m["updatedAt"] = e.UpdatedAt.Unix()
	}
	b, err := json.Marshal(m)
	return b, errors.Trace(err)
}
