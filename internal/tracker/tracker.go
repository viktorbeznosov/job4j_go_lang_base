package tracker

import (
	"strings"
)

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) error {
	_, ok := t.indexOf(item.ID)
	if ok {
		return ErrAlreadyExists
	}
	t.Items = append(t.Items, item)
	return nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.Items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
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

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.Items[index] = item
	return nil
}

func (t *Tracker) DeleteItem(id string) error {
	index, ok := t.indexOf(id)
	if !ok {
		return ErrNotFound
	}
	t.Items = append(t.Items[:index], t.Items[index+1:]...)
	return nil
}
