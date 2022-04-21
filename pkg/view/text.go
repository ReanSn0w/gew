package view

import "context"

func Text(text string) View {
	return &TextView{Text: text}
}

type TextView struct {
	Text string
}

func (tv *TextView) Body(context context.Context) View {
	return Empty()
}
