package view

import (
	"context"
	"log"
)

// Text, Empty, Group
type View interface {
	Body(context.Context) View
}

func Build(item View, ctx context.Context, ext func(interface{}, context.Context)) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ERROR] gew: %v", item)
		}
	}()

	switch v := item.(type) {
	case nil:
		return
	case group:
		for _, item := range v {
			Build(item, ctx, ext)
		}
	case *contexted:
		newCtx := v.prepare(ctx)
		Build(v.Body(newCtx), newCtx, ext)
	case *ExternalContent:
		ext(v.Content, ctx)
	case ModificationApplyer:
		if v == nil {
			return
		}

		Build(v.Body(ctx), ctx, ext)
	default:
		Build(v.Body(ctx), ctx, ext)
	}
}
