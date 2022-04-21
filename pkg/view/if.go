package view

import "context"

func If(condition bool, content View) View {
	return &ifView{
		condition: condition,
		content:   content,
	}
}

type ifView struct {
	condition bool
	content   View
}

func (iv *ifView) Body(context context.Context) View {
	if iv.condition {
		return iv.content
	} else {
		return Empty()
	}
}
