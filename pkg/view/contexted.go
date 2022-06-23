package view

import "context"

func Contexted(key interface{}, value interface{}, content View) View {
	return &contexted{
		content: content,
		prepare: func(ctx context.Context) context.Context {
			return context.WithValue(ctx, key, value)
		},
	}
}

func CustomContextPreparation(view View, prepare func(ctx context.Context) context.Context) View {
	return &contexted{
		content: view,
		prepare: prepare,
	}
}

type contexted struct {
	prepare func(context.Context) context.Context
	content View
}

func (cv *contexted) Body(context context.Context) View {
	return cv.content
}
