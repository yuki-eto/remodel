// Code generated by generate_code script - DO NOT EDIT.
package entity

import (
	"encoding/json"
	"time"

	"github.com/juju/errors"
	"go.knocknote.io/rapidash"
)

type UserByte struct {
	ID        uint64
	UserID    uint64
	Bytes     []byte
	Tags      []string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type UserBytes []*UserByte

func (e *UserByte) EncodeRapidash(enc rapidash.Encoder) error {
	if e.ID != 0 {
		enc.Uint64("id", e.ID)
	}
	enc.Uint64("id", e.ID)
	enc.Uint64("user_id", e.UserID)
	enc.Bytes("bytes", e.Bytes)
	enc.Strings("tags", e.Tags)
	enc.TimePtr("created_at", e.CreatedAt)
	enc.TimePtr("updated_at", e.UpdatedAt)
	return enc.Error()
}

func (e *UserBytes) EncodeRapidash(enc rapidash.Encoder) error {
	for _, v := range *e {
		if err := v.EncodeRapidash(enc.New()); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (e *UserByte) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.UserID = dec.Uint64("user_id")
	e.Bytes = dec.Bytes("bytes")
	e.Tags = dec.Strings("tags")
	e.CreatedAt = dec.TimePtr("created_at")
	e.UpdatedAt = dec.TimePtr("updated_at")
	return dec.Error()
}

func (e *UserBytes) DecodeRapidash(dec rapidash.Decoder) error {
	count := dec.Len()
	*e = make([]*UserByte, count)
	for i := 0; i < count; i++ {
		var v UserByte
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return errors.Trace(err)
		}
		(*e)[i] = &v
	}
	return nil
}

func (e *UserByte) Struct() *rapidash.Struct {
	s := rapidash.NewStruct("user_bytes")
	s.FieldUint64("id")
	s.FieldUint64("user_id")
	s.FieldBytes("bytes")
	s.FieldSlice("tags", rapidash.StringType)
	s.FieldTime("created_at")
	s.FieldTime("updated_at")
	return s
}

func (e *UserByte) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"bytes": e.Bytes,
		"id":    e.ID,
		"tags":  e.Tags,
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