package view

import (
	"context"
)

// External - служит для добавления своих компонентов в пайплайн построения view
func External(content interface{}) ModificationApplyer {
	return NewModificationApplyer(&ExternalContent{
		Content: content,
	})
}

type ExternalContent struct {
	Content interface{}
}

func (view *ExternalContent) Body(ctx context.Context) View {
	return nil
}
