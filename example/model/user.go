// Code generated by generate_code script - DO NOT EDIT.
package model

import (
	"example/dao"
	"example/entity"
	"sort"
	"time"

	"github.com/juju/errors"
)

type UserImpl struct {
	userDao dao.User
}

func (m *UserImpl) createInstance(e *entity.User) *UserInstance {
	return &UserInstance{
		User:    e,
		userDao: m.userDao,
	}
}

type UserInstance struct {
	*entity.User
	userDao dao.User
}

func (i *UserInstance) Save() error {
	if i.userDao == nil {
		return nil
	}
	return errors.Trace(i.userDao.Save(i.User))
}

func (i *UserInstance) Delete() error {
	if i.userDao == nil {
		return nil
	}
	return errors.Trace(i.userDao.Delete(i.User))
}

type UsersInstance struct {
	values []*UserInstance
}

func NewUsersInstance() *UsersInstance {
	return &UsersInstance{values: []*UserInstance{}}
}

func (i *UsersInstance) Add(v *UserInstance) {
	i.values = append(i.values, v)
}

func (i *UsersInstance) FindByID(id uint64) *UserInstance {
	for _, v := range i.values {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func (i *UsersInstance) FilterBy(f func(*UserInstance) bool) *UsersInstance {
	instance := NewUsersInstance()
	for _, v := range i.values {
		if f(v) {
			instance.Add(v)
		}
	}
	return instance
}

func (i *UsersInstance) Each(f func(*UserInstance)) {
	for _, v := range i.values {
		f(v)
	}
}

func (i *UsersInstance) EachWithError(f func(*UserInstance) error) error {
	for _, v := range i.values {
		if err := f(v); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (i *UsersInstance) First() *UserInstance {
	if len(i.values) == 0 {
		return nil
	}
	return i.values[0]
}

func (i *UsersInstance) At(idx int) *UserInstance {
	if len(i.values) < idx {
		return nil
	}
	return i.values[idx]
}

func (i *UsersInstance) Len() int {
	return len(i.values)
}

func (i *UsersInstance) FilterByID(c uint64) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.ID == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByID(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].ID > s.values[j].ID
		}
		return s.values[i].ID < s.values[j].ID
	})
	return s
}

func (i *UsersInstance) IDs() []uint64 {
	s := []uint64{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.ID)
	})
	return s
}

func (i *UsersInstance) FilterByUuid(c string) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.Uuid == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByUuid(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].Uuid > s.values[j].Uuid
		}
		return s.values[i].Uuid < s.values[j].Uuid
	})
	return s
}

func (i *UsersInstance) Uuids() []string {
	s := []string{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.Uuid)
	})
	return s
}

func (i *UsersInstance) FilterByAccessToken(c string) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.AccessToken == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByAccessToken(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].AccessToken > s.values[j].AccessToken
		}
		return s.values[i].AccessToken < s.values[j].AccessToken
	})
	return s
}

func (i *UsersInstance) AccessTokens() []string {
	s := []string{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.AccessToken)
	})
	return s
}

func (i *UsersInstance) FilterByOutsideUserID(c string) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.OutsideUserID == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByOutsideUserID(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].OutsideUserID > s.values[j].OutsideUserID
		}
		return s.values[i].OutsideUserID < s.values[j].OutsideUserID
	})
	return s
}

func (i *UsersInstance) OutsideUserIDs() []string {
	s := []string{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.OutsideUserID)
	})
	return s
}

func (i *UsersInstance) FilterByName(c string) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.Name == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByName(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].Name > s.values[j].Name
		}
		return s.values[i].Name < s.values[j].Name
	})
	return s
}

func (i *UsersInstance) Names() []string {
	s := []string{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.Name)
	})
	return s
}

func (i *UsersInstance) FilterByCreatedAt(c *time.Time) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.CreatedAt == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByCreatedAt(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].CreatedAt.Before(*s.values[j].CreatedAt)
		}
		return s.values[i].CreatedAt.After(*s.values[j].CreatedAt)
	})
	return s
}

func (i *UsersInstance) CreatedAts() []*time.Time {
	s := []*time.Time{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.CreatedAt)
	})
	return s
}

func (i *UsersInstance) FilterByUpdatedAt(c *time.Time) *UsersInstance {
	s := NewUsersInstance()
	for _, v := range i.values {
		if v.UpdatedAt == c {
			s.Add(v)
		}
	}
	return s
}

func (i *UsersInstance) SortByUpdatedAt(isDesc bool) *UsersInstance {
	s := NewUsersInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].UpdatedAt.Before(*s.values[j].UpdatedAt)
		}
		return s.values[i].UpdatedAt.After(*s.values[j].UpdatedAt)
	})
	return s
}

func (i *UsersInstance) UpdatedAts() []*time.Time {
	s := []*time.Time{}
	i.Each(func(v *UserInstance) {
		s = append(s, v.UpdatedAt)
	})
	return s
}
