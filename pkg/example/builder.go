package example

import (
	"bytes"
	"context"

	"github.com/ReanSn0w/gew/pkg/view"
)

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct{}

func (b *Builder) Build(context context.Context, item view.View) []byte {
	switch v := item.(type) {
	case *view.EmptyView, nil:
		return []byte{}
	case *view.TextView:
		return []byte(v.Text)
	case *view.GroupView:
		buffer := new(bytes.Buffer)

		for _, item := range v.Elements {
			buffer.Write(b.Build(context, item))
		}

		return buffer.Bytes()
	case *view.ContextedView:
		return b.Build(v.UpdateContext(context), v.Body(context))
	default:
		return b.Build(context, item.Body(context))
	}
}
