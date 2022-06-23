package view

import "context"

func Closure(builder func(context.Context) View) View {
	return &closureView{
		builder: builder,
	}
}

func For(count uint, builder func(int) View) View {
	return Closure(func(ctx context.Context) View {
		items := make([]View, 0, count)

		for index := 0; index < int(count); index++ {
			items = append(items, builder(index))
		}

		return Group(items...)
	})
}

func If(value bool, content View) View {
	return Closure(func(ctx context.Context) View {
		if value {
			return content
		} else {
			return nil
		}
	})
}

type closureView struct {
	builder func(context.Context) View
}

func (cv *closureView) Body(context context.Context) View {
	return cv.builder(context)
}
