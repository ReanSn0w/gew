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

func (b *Builder) Build(ctx context.Context, item view.View) []byte {
	buffer := new(bytes.Buffer)

	view.Build(item, ctx, func(i interface{}, ctx context.Context) {
		val := b.externalBuilder(i)
		buffer.Write(val)
	})

	return buffer.Bytes()
}

func (b *Builder) externalBuilder(value interface{}) []byte {
	res, even := value.(string)

	if even {
		return []byte(res)
	} else {
		return []byte{}
	}
}

func Text(value string) view.View {
	return view.External(value)
}

func WrongView() view.View {
	return view.External(nil)
}
