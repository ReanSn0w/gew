package view

import "context"

func Group(elements ...View) View {
	return &group{elements: elements}
}

type group struct {
	elements []View
}

func (gv *group) Body(context context.Context) View {
	return nil
}
