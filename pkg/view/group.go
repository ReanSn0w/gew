package view

import (
	"context"
)

// MARK: - Group

// Group объединяет элементы View в один.
func Group(elements ...View) ModificationApplyer {
	return NewModificationApplyer(group(elements))
}

type group []View

func (g group) Body(ctx context.Context) View {
	return nil
}
