// Code generated by generate_code script - DO NOT EDIT.
package entity

import "go.knocknote.io/rapidash"

type Structable interface {
	Struct() *rapidash.Struct
}

func Structables() map[string]Structable {
	return map[string]Structable{
		"user_bytes":   new(UserByte),
		"user_friends": new(UserFriend),
		"users":        new(User),
	}
}
func ReadOnlyStructables() map[string]Structable {
	return map[string]Structable{"items": new(Item)}
}
