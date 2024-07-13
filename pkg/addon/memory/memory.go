package memory

import (
	"context"

	"github.com/ReanSn0w/gew/v3/pkg/view"
)

var (
	memoryCtxKey = &memoryCtx{}
)

type (
	memoryCtx struct{}

	memory struct {
		data map[string]interface{}
	}

	Memory interface {
		Get(key string) (interface{}, bool)
		Update(key string, update func(interface{}, bool) interface{})
		Delete(key string)
	}
)

func Extract(ctx context.Context) Memory {
	mem, ok := ctx.Value(memoryCtxKey).(*memory)
	if !ok {
		return &memory{data: make(map[string]interface{})}
	}

	return mem
}

func Use() view.Mod {
	return view.Context(func(ctx context.Context) context.Context {
		return context.WithValue(ctx, memoryCtxKey, &memory{
			data: make(map[string]interface{}),
		})
	})
}

func Update(key string, update func(interface{}, bool) interface{}) view.Mod {
	return view.Context(func(ctx context.Context) context.Context {
		Extract(ctx).Update(key, update)
		return ctx
	})
}

func Delete(key string) view.Mod {
	return view.Context(func(ctx context.Context) context.Context {
		Extract(ctx).Delete(key)
		return ctx
	})
}

func (m *memory) Get(key string) (interface{}, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *memory) Update(key string, update func(interface{}, bool) interface{}) {
	val, ok := m.data[key]
	m.data[key] = update(val, ok)
}

func (m *memory) Delete(key string) {
	delete(m.data, key)
}
