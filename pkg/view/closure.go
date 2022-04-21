package view

import "context"

func Closure(builder func(context.Context) View) View {
	return &closureView{
		builder: builder,
	}
}

type closureView struct {
	builder func(context.Context) View
}

func (cv *closureView) Body(context context.Context) View {
	return cv.builder(context)
}
