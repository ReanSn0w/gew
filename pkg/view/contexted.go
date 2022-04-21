package view

import "context"

func Contexted(key interface{}, value interface{}, content View) View {
	return &ContextedView{
		key:     key,
		value:   value,
		content: content,
	}
}

type ContextedView struct {
	key     interface{}
	value   interface{}
	content View
}

func (cv *ContextedView) Body(context context.Context) View {
	return cv.content
}

func (cv *ContextedView) UpdateContext(parent context.Context) context.Context {
	return context.WithValue(parent, cv.key, cv.value)
}
