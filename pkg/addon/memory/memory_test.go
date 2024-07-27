package memory_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/ReanSn0w/gew/v4/pkg/addon/memory"
	"github.com/ReanSn0w/gew/v4/pkg/view"
)

func TestMemory_Get(t *testing.T) {
	checked := false

	StringBuilder(
		Text("test")(
			memory.Update("key", func(i interface{}, b bool) interface{} {
				return "value"
			}),
			view.Context(func(ctx context.Context) context.Context {
				data, _ := memory.Extract(ctx).Get("key")

				if data.(string) == "value" {
					checked = true
				}

				return ctx
			}),
		),
	)

	if !checked {
		t.Error("memory get failed")
	}
}

func TestMemory_Update(t *testing.T) {
	checked := false

	StringBuilder(
		Text("test")(
			memory.Update("key", func(i interface{}, b bool) interface{} {
				return "value"
			}),
			memory.Update("key", func(i interface{}, b bool) interface{} {
				return "value value"
			}),
			view.Context(func(ctx context.Context) context.Context {
				data, _ := memory.Extract(ctx).Get("key")

				if data.(string) == "value value" {
					checked = true
				}

				return ctx
			}),
		),
	)

	if !checked {
		t.Error("memory update failed")
	}
}

func TestMemory_Delete(t *testing.T) {
	checked := false

	StringBuilder(
		Text("test")(
			memory.Update("key", func(i interface{}, b bool) interface{} {
				return "value"
			}),
			memory.Delete("key"),
			view.Context(func(ctx context.Context) context.Context {
				_, ok := memory.Extract(ctx).Get("key")

				if !ok {
					checked = true
				}

				return ctx
			}),
		),
	)

	if !checked {
		t.Error("memory delete failed")
	}
}

// MARK: - TextBuilder

func Text(text string) view.Use {
	return view.External(text)
}

func StringBuilder(item view.View) string {
	buffer := new(bytes.Buffer)

	view.Builder(context.TODO(), view.NewView(item)(
		memory.Use(),
	), func(ctx context.Context, i interface{}) {
		buffer.WriteString(i.(string))
	})

	return buffer.String()
}
