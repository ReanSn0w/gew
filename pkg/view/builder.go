package view

import "context"

// ViewBuilder это функция, которая принимает элементы View и возвращает ModificationApplyer.
type Builder func(elements ...View) ModificationApplyer

// Body реализует интерфейс View.
func (vb Builder) Body(ctx context.Context) View {
	return vb()
}

// If функция реализует логику для выполения логических операция со View.
func If(val bool) Builder {
	if val {
		return Group
	}

	return emptyGroup
}

// MARK: - Group

// Group объединяет элементы View в один.
func Group(elements ...View) ModificationApplyer {
	return NewModificationApplyer(group(elements))
}

func emptyGroup(elements ...View) ModificationApplyer {
	return NewModificationApplyer(group{})
}

type group []View

func (g group) Body(ctx context.Context) View {
	return nil
}
