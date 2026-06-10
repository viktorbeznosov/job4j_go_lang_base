package tracker

import (
	"github.com/google/uuid"
)

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type FindUsecase struct{}

func (u FindUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter item term")
	term := in.Get()
	item := tracker.FindItem(term)
	if item != nil {
		out.Out(item.toString())
	} else {
		out.Out("Item not found")
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter item ID")
	id := in.Get()
	out.Out("Enter item new name")
	name := in.Get()
	if tracker.UpdateItem(id, name) {
		out.Out("Item updated")
	} else {
		out.Out("Item not found")
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter item ID")
	id := in.Get()
	isDeleted := tracker.DeleteItem(id)
	if isDeleted {
		out.Out("Item is deleted")
	} else {
		out.Out("Item not found")
	}
}
