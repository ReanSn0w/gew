package view

import "context"

func Group(elements ...View) View {
	return &GroupView{Elements: elements}
}

type GroupView struct {
	Elements []View
}

func (gv *GroupView) Body(context context.Context) View {
	return nil
}
