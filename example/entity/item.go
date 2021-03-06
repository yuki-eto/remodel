// Code generated by generate_code script - DO NOT EDIT.
package entity

import (
	"encoding/json"

	"github.com/juju/errors"
	"go.knocknote.io/rapidash"
)

type Item struct {
	ID       uint64 `csv:"Id"`
	Type     string `csv:"Type"`
	Rarity   string `csv:"Rarity"`
	Name     string `csv:"Name"`
	MaxCount uint16 `csv:"MaxCount"`
}

type Items []*Item

func (e *Item) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.Type = dec.String("type")
	e.Rarity = dec.String("rarity")
	e.Name = dec.String("name")
	e.MaxCount = dec.Uint16("max_count")
	return dec.Error()
}

func (e *Items) DecodeRapidash(dec rapidash.Decoder) error {
	count := dec.Len()
	*e = make([]*Item, count)
	for i := 0; i < count; i++ {
		var v Item
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return errors.Trace(err)
		}
		(*e)[i] = &v
	}
	return nil
}

func (e *Item) Struct() *rapidash.Struct {
	s := rapidash.NewStruct("items")
	s.FieldUint64("id")
	s.FieldString("type")
	s.FieldString("rarity")
	s.FieldString("name")
	s.FieldUint16("max_count")
	return s
}

func (e *Item) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":       e.ID,
		"maxCount": e.MaxCount,
		"name":     e.Name,
		"rarity":   e.Rarity,
		"type":     e.Type,
	}
	b, err := json.Marshal(m)
	return b, errors.Trace(err)
}
