package view

import "context"

func External(content interface{}) View {
	return &ExternalContent{
		Content: content,
	}
}

type ExternalContent struct {
	Content interface{}
}

func (view *ExternalContent) Body(ctx context.Context) View {
	return nil
}
