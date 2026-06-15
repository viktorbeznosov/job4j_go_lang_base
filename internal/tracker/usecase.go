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
	err := tracker.AddItem(Item{Name: name, ID: id})
	result := "Item added"
	if err != nil {
		result = err.Error()
	}

	out.Out(result)
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
		out.Out(ErrNotFound.Error())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter item ID")
	id := in.Get()
	out.Out("Enter item new name")
	name := in.Get()
	err := tracker.UpdateItem(Item{ID: id, Name: name})
	result := "Item updated"
	if err != nil {
		result = err.Error()
	}

	out.Out(result)
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter item ID")
	id := in.Get()
	err := tracker.DeleteItem(id)
	result := "Item is deleted"
	if err != nil {
		result = err.Error()
	}

	out.Out(result)
}
