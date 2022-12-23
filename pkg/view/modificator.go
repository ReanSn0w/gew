package view

import "context"

// Modificator это функция, которая принимает View и возвращает View.
// те, она может модифицировать View в своем теле.
type Modificator func(View) View

// NewModificationApplyer Возвращает функцию, которая соответствует интерфейсу View.
//
// Она предназначена для удоного применения модификаторов к View.
func NewModificationApplyer(view View) ModificationApplyer {
	return func(modificators ...Modificator) View {
		for _, modificator := range modificators {
			if modificator == nil {
				continue
			}

			view = modificator(view)
		}

		return view
	}
}

// ModificationApplyer это функция, которая применяет модификации ко View.
type ModificationApplyer func(modificators ...Modificator) View

// Body реализует интерфейс View.
func (ma ModificationApplyer) Body(ctx context.Context) View {
	return ma()
}

// MARK: - Context modificators

// Context - Добавляет значения в контекст построения View.
func Context(key, value interface{}) Modificator {
	return func(view View) View {
		return &contexted{
			content: view,
			prepare: func(ctx context.Context) context.Context {
				return context.WithValue(ctx, key, value)
			},
		}
	}
}

// ContextModificator - позволяет изменить контекст построения View.
func ContextModificator(prepare func(ctx context.Context) context.Context) Modificator {
	return func(view View) View {
		return &contexted{
			content: view,
			prepare: prepare,
		}
	}
}

type contexted struct {
	prepare func(context.Context) context.Context
	content View
}

func (cv *contexted) Body(context context.Context) View {
	return cv.content
}

// BuildModificator - группа модификаторов для управления построением

// Модификатор для опционального применения модификаторов
func If(condition bool, modificators ...Modificator) Modificator {
	if condition {
		return func(v View) View {
			for _, modificator := range modificators {
				v = modificator(v)
			}

			return v
		}
	}

	return func(v View) View {
		return v
	}
}

// Hidden - модификатор позволяет скрыть элемент View из построения.
func Hidden(condition bool) Modificator {
	return func(view View) View {
		if !condition {
			return view
		}

		return nil
	}
}

// Replace - заменяет элемент View на другой.
func Replace(view View) Modificator {
	return func(_ View) View {
		return view
	}
}
