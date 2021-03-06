package view

import (
	"context"
)

// Text, Empty, Group
type View interface {
	Body(context.Context) View
}

func Build(item View, ctx context.Context, ext func(interface{}, context.Context)) {
	switch v := item.(type) {
	case nil:
		return
	case *group:
		for _, item := range v.elements {
			Build(item, ctx, ext)
		}
	case *contexted:
		newCtx := v.prepare(ctx)
		Build(v.Body(newCtx), newCtx, ext)
	case *external:
		ext(v.content, ctx)
	default:
		Build(v.Body(ctx), ctx, ext)
	}
}
