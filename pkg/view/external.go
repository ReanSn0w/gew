package view

import "context"

func External(content interface{}) View {
	return &external{
		content: content,
	}
}

type external struct {
	content interface{}
}

func (view *external) Body(ctx context.Context) View {
	return nil
}
