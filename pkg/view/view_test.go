package view_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func TestIfTrue(t *testing.T) {
	hello := view.If(true)(
		Text("Hello"),
	)

	val := StringBuilder(hello) // Hello

	if val != "Hello" {
		t.Error("If(true) not working")
	}
}

func TestIfFalse(t *testing.T) {
	hello := view.If(false)(
		Text("Hello"),
	)

	val := StringBuilder(hello) // ""

	if val != "" {
		t.Error("If(false) not working")
	}
}

func TestFor(t *testing.T) {
	hello := view.For(3, func(index int) view.View {
		return Text("Hello")
	})

	val := StringBuilder(hello) // HelloHelloHello

	if val != "HelloHelloHello" {
		t.Error("For not working")
	}
}

func TestForZero(t *testing.T) {
	hello := view.For(0, func(index int) view.View {
		return Text("Hello")
	})

	val := StringBuilder(hello) // ""

	if val != "" {
		t.Error("For(0) not working")
	}
}

func TestForNegative(t *testing.T) {
	hello := view.For(-1, func(index int) view.View {
		return Text("Hello")
	})

	val := StringBuilder(hello) // ""

	if val != "" {
		t.Error("For(-1) not working")
	}
}

func TestViewBuilderWithValues(t *testing.T) {
	hello := view.Group(
		Text("Hello"),
		Text("World"),
	)

	val := StringBuilder(hello) // HelloWorld

	if val != "HelloWorld" {
		t.Error("ViewBuilder not working")
	}
}

func TestContextModificator(t *testing.T) {
	res := false
	key := "k"

	hello := view.Group(
		Text("Hello"),
		Text("World")(
			view.ContextModificator(func(ctx context.Context) context.Context {
				val, ok := ctx.Value(&key).(string)

				if val == "value" && ok {
					res = true
				}

				return ctx
			}),
			view.Context(&key, "value"),
		),
	)

	_ = StringBuilder(hello) // HelloWorld

	if !res {
		t.Error("Context not working")
	}
}

func TestNil(t *testing.T) {
	hello := view.Group(
		Text("Hello"),
		nil,
		Text("World"),
	)

	val := StringBuilder(hello) // HelloWorld

	if val != "HelloWorld" {
		t.Error("nil not working")
	}
}

// MARK: - TextBuilder

func Text(text string) view.ModificationApplyer {
	return view.External(text)
}

func StringBuilder(item view.View) string {
	buffer := new(bytes.Buffer)

	view.Build(item, context.TODO(), func(item interface{}, ctx context.Context) {
		buffer.WriteString(item.(string))
	})

	return buffer.String()
}
