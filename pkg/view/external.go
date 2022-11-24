package view

import "context"

func External(content interface{}) View {
	return &ExternalContent{
		content: content,
	}
}

type ExternalContent struct {
	content interface{}
}

func (view *ExternalContent) Body(ctx context.Context) View {
	return nil
}
