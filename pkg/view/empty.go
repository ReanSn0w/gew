package view

import "context"

func Empty() View {
	return &EmptyView{}
}

type EmptyView struct{}

func (ev *EmptyView) Body(context context.Context) View {
	return nil
}
