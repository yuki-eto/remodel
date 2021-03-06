// Code generated by generate_code script - DO NOT EDIT.
package model

import (
	"example/dao"
	"example/entity"
	"sort"

	"github.com/juju/errors"
)

type ItemImpl struct {
	itemDao dao.Item
}

func (m *ItemImpl) createInstance(e *entity.Item) *ItemInstance {
	return &ItemInstance{Item: e}
}

type ItemInstance struct {
	*entity.Item
}
type ItemsInstance struct {
	values []*ItemInstance
}

func NewItemsInstance() *ItemsInstance {
	return &ItemsInstance{values: []*ItemInstance{}}
}

func (i *ItemsInstance) Add(v *ItemInstance) {
	i.values = append(i.values, v)
}

func (i *ItemsInstance) FindByID(id uint64) *ItemInstance {
	for _, v := range i.values {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func (i *ItemsInstance) FilterBy(f func(*ItemInstance) bool) *ItemsInstance {
	instance := NewItemsInstance()
	for _, v := range i.values {
		if f(v) {
			instance.Add(v)
		}
	}
	return instance
}

func (i *ItemsInstance) Each(f func(*ItemInstance)) {
	for _, v := range i.values {
		f(v)
	}
}

func (i *ItemsInstance) EachWithError(f func(*ItemInstance) error) error {
	for _, v := range i.values {
		if err := f(v); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (i *ItemsInstance) First() *ItemInstance {
	if len(i.values) == 0 {
		return nil
	}
	return i.values[0]
}

func (i *ItemsInstance) At(idx int) *ItemInstance {
	if len(i.values) < idx {
		return nil
	}
	return i.values[idx]
}

func (i *ItemsInstance) Len() int {
	return len(i.values)
}

func (i *ItemsInstance) IsEmpty() bool {
	return i.Len() == 0
}

func (i *ItemsInstance) FilterByID(c uint64) *ItemsInstance {
	s := NewItemsInstance()
	for _, v := range i.values {
		if v.ID == c {
			s.Add(v)
		}
	}
	return s
}

func (i *ItemsInstance) SortByID(isDesc bool) *ItemsInstance {
	s := NewItemsInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].ID > s.values[j].ID
		}
		return s.values[i].ID < s.values[j].ID
	})
	return s
}

func (i *ItemsInstance) IDs() []uint64 {
	s := []uint64{}
	i.Each(func(v *ItemInstance) {
		s = append(s, v.ID)
	})
	return s
}

func (i *ItemsInstance) FilterByType(c string) *ItemsInstance {
	s := NewItemsInstance()
	for _, v := range i.values {
		if v.Type == c {
			s.Add(v)
		}
	}
	return s
}

func (i *ItemsInstance) SortByType(isDesc bool) *ItemsInstance {
	s := NewItemsInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].Type > s.values[j].Type
		}
		return s.values[i].Type < s.values[j].Type
	})
	return s
}

func (i *ItemsInstance) Types() []string {
	s := []string{}
	i.Each(func(v *ItemInstance) {
		s = append(s, v.Type)
	})
	return s
}

func (i *ItemsInstance) FilterByRarity(c string) *ItemsInstance {
	s := NewItemsInstance()
	for _, v := range i.values {
		if v.Rarity == c {
			s.Add(v)
		}
	}
	return s
}

func (i *ItemsInstance) SortByRarity(isDesc bool) *ItemsInstance {
	s := NewItemsInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].Rarity > s.values[j].Rarity
		}
		return s.values[i].Rarity < s.values[j].Rarity
	})
	return s
}

func (i *ItemsInstance) Rarities() []string {
	s := []string{}
	i.Each(func(v *ItemInstance) {
		s = append(s, v.Rarity)
	})
	return s
}

func (i *ItemsInstance) FilterByName(c string) *ItemsInstance {
	s := NewItemsInstance()
	for _, v := range i.values {
		if v.Name == c {
			s.Add(v)
		}
	}
	return s
}

func (i *ItemsInstance) SortByName(isDesc bool) *ItemsInstance {
	s := NewItemsInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].Name > s.values[j].Name
		}
		return s.values[i].Name < s.values[j].Name
	})
	return s
}

func (i *ItemsInstance) Names() []string {
	s := []string{}
	i.Each(func(v *ItemInstance) {
		s = append(s, v.Name)
	})
	return s
}

func (i *ItemsInstance) FilterByMaxCount(c uint16) *ItemsInstance {
	s := NewItemsInstance()
	for _, v := range i.values {
		if v.MaxCount == c {
			s.Add(v)
		}
	}
	return s
}

func (i *ItemsInstance) SortByMaxCount(isDesc bool) *ItemsInstance {
	s := NewItemsInstance()
	s.values = i.values
	sort.SliceStable(s.values, func(i, j int) bool {
		if isDesc {
			return s.values[i].MaxCount > s.values[j].MaxCount
		}
		return s.values[i].MaxCount < s.values[j].MaxCount
	})
	return s
}

func (i *ItemsInstance) MaxCounts() []uint16 {
	s := []uint16{}
	i.Each(func(v *ItemInstance) {
		s = append(s, v.MaxCount)
	})
	return s
}
