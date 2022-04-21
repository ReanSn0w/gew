package view

import "context"

func For(count uint, builder func(int) View) View {
	return &forView{
		count:   count,
		builder: builder,
	}
}

type forView struct {
	count   uint
	builder func(int) View
}

func (fv *forView) Body(context context.Context) View {
	items := make([]View, 0, fv.count)

	for index := 0; index < int(fv.count); index++ {
		items = append(items, fv.builder(index))
	}

	return Group(items...)
}
