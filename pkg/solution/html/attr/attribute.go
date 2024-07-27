package attr

import (
	"context"

	"github.com/ReanSn0w/gew/v4/pkg/view"
)

var (
	attributeCtxKey = &attributeCtx{}
)

type (
	attributeCtx = struct{}

	attributeStorage struct {
		data map[string]string
	}
)

func (a *attributeStorage) String() string {
	res := ""

	for k, v := range a.data {
		if v == "" {
			res += k + " "
		} else {
			res += k + "=\"" + v + "\" "
		}
	}

	return res
}

func Setup(ctx context.Context) context.Context {
	return context.WithValue(ctx, attributeCtxKey, &attributeStorage{data: map[string]string{}})
}

func Extract(ctx context.Context) (*attributeStorage, bool) {
	val := ctx.Value(attributeCtxKey)
	as, ok := val.(*attributeStorage)
	return as, ok
}

func Set(key, val string) view.Mod {
	return view.Context(func(ctx context.Context) context.Context {
		storage, ok := Extract(ctx)
		if !ok {
			ctx = Setup(ctx)
			storage, _ = Extract(ctx)
		}
		storage.data[key] += " " + val
		return ctx
	})
}

func Prepare(key string, prepare func(string) string) view.Mod {
	return view.Context(func(ctx context.Context) context.Context {
		storage, ok := Extract(ctx)
		if !ok {
			ctx = Setup(ctx)
			storage, _ = Extract(ctx)
		}

		storage.data[key] = prepare(storage.data[key])
		return ctx
	})
}

func Delete(key string) view.Mod {
	return view.Context(func(ctx context.Context) context.Context {
		storage, ok := Extract(ctx)
		if !ok {
			return ctx
		}
		delete(storage.data, key)
		return ctx
	})
}
