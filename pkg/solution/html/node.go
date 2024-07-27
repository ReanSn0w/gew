package html

import (
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/v4/pkg/solution/html/attr"
	"github.com/ReanSn0w/gew/v4/pkg/view"
)

type (
	element interface {
		build(ctx context.Context, wr io.Writer)
	}
)

func NewNode(tag string, content ...view.View) view.Use {
	return view.NewView(
		view.External(
			&Node{
				tag:     tag,
				content: view.Group(content...),
			},
		),
	)
}

type Node struct {
	tag     string
	content view.View
}

func (n *Node) build(ctx context.Context, wr io.Writer) {
	if as, ok := attr.Extract(ctx); ok {
		wr.Write([]byte(fmt.Sprintf("<%v %v>", n.tag, as.String())))
	} else {
		wr.Write([]byte(fmt.Sprintf("<%v>", n.tag)))
	}

	Build(ctx, wr, n.content)

	wr.Write([]byte(fmt.Sprintf("</%v>", n.tag)))
}

func NewInlineNode(tag string) view.Use {
	return view.NewView(
		view.External(
			&InlineNode{
				tag: tag,
			},
		),
	)
}

type InlineNode struct {
	tag string
}

func (n *InlineNode) build(ctx context.Context, wr io.Writer) {
	if as, ok := attr.Extract(ctx); ok {
		wr.Write([]byte(fmt.Sprintf("<%v %v>", n.tag, as.String())))
	} else {
		wr.Write([]byte(fmt.Sprintf("<%v>", n.tag)))
	}
}

func NewPlainNode(data []byte) view.Use {
	return view.NewView(
		view.External(
			&PlainNode{
				data: data,
			},
		),
	)
}

type PlainNode struct {
	data []byte
}

func (p *PlainNode) build(ctx context.Context, wr io.Writer) {
	wr.Write(p.data)
}

func Build(ctx context.Context, writer io.Writer, v view.View) {
	view.Builder(ctx, v, func(ctx context.Context, i interface{}) {
		if el, ok := i.(element); ok {
			el.build(ctx, writer)
		}
	})
}
