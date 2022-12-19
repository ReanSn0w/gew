package view

import "context"

// Closure возвращает View, функции переданной в нее в качестве аргумента.
//
// Через данную функцию реализованы такие функции как:
// For
func Closure(builder func(context.Context) View) ModificationApplyer {
	return NewModificationApplyer(&closureView{
		builder: builder,
	})
}

// For возвращает View, который содержит count элементов, созданных функцией builder.
func For(count int, builder func(int) View) ModificationApplyer {
	// Исключение для нулевого количества элементов.
	// Так же оно сработает в случае если count < 0.
	if count < 1 {
		return Group()
	}

	return Closure(func(ctx context.Context) View {
		items := make([]View, 0, count)

		for index := 0; index < int(count); index++ {
			items = append(items, builder(index))
		}

		return Group(items...)
	})
}

type closureView struct {
	builder func(context.Context) View
}

func (cv *closureView) Body(context context.Context) View {
	return cv.builder(context)
}
