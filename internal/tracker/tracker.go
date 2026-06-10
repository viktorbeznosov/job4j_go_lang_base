package tracker

import (
	"fmt"
	"strings"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.Items = append(t.Items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) FindItem(term string) *Item {
	for _, item := range t.Items {
		if strings.Contains(item.Name, term) {
			copyItem := item // Явное копирование
			return &copyItem
		}
	}
	return nil
}

func (t *Tracker) UpdateItem(id string, name string) bool {
	result := false
	for index, item := range t.Items {
		if item.ID == id {
			t.Items[index].Name = name
			result = true
		}
	}

	return result
}

func (t *Tracker) DeleteItem(id string) bool {
	result := false
	for index, item := range t.Items {
		if item.ID == id {
			t.Items = append(t.Items[:index], t.Items[index+1:]...)
			result = true
		}
	}
	return result
}
