package example_test

import (
	"context"
	"testing"

	"github.com/ReanSn0w/gew/pkg/example"
	"github.com/ReanSn0w/gew/pkg/view"
)

func Test_BuildEmpty(t *testing.T) {
	b := example.NewBuilder()

	res := b.Build(context.Background(), view.Empty())
	if string(res) != "" {
		t.Log("error: пустой view сгенерировал результат")
		t.FailNow()
	}
}

func Test_BuildText(t *testing.T) {
	data := []string{"1", "2", "3", "4", "5"}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		res := b.Build(context.Background(), view.Text(data[index]))

		if string(res) != data[index] {
			t.Log("error: значение не сошлось с указанным в массиве")
			t.FailNow()
		}
	}
}

func Test_BuildGroup(t *testing.T) {
	data := []string{"11111", "22222", "33333", "11223", "32132"}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		views := make([]view.View, 0, len(data[index]))
		for _, el := range data[index] {
			views = append(views, view.Text(string(el)))
		}

		res := b.Build(context.Background(), view.Group(views...))
		if string(res) != data[index] {
			t.Log("error: не удалось построить строку через группу")
			t.FailNow()
		}
	}
}

func Test_BuildEmptyGroup(t *testing.T) {
	b := example.NewBuilder()

	res := b.Build(context.Background(), view.Group())
	if string(res) != "" {
		t.Log("error: построение пустой группы завершилось неудачей")
	}
}

func Test_BuildFor(t *testing.T) {
	data := []string{"11111", "22222", "33333", "11223", "32132"}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		res := b.Build(context.Background(), view.For(uint(len(data[index])), func(i int) view.View {
			return view.Text(string(data[index][i]))
		}))
		if string(res) != data[index] {
			t.Log("error: не удалось построить строку через группу")
			t.FailNow()
		}
	}
}

func Test_BuildIf(t *testing.T) {
	type TestItem struct {
		state bool
		value string
	}

	data := []TestItem{
		{
			state: true,
			value: "some string 1",
		},
		{
			state: false,
			value: "some string 2",
		},
	}
	b := example.NewBuilder()

	for _, item := range data {
		res := b.Build(context.Background(), view.If(item.state, view.Text(item.value)))
		if item.state {
			if string(res) != item.value {
				t.Log("error: результат отличается от необходимого в данно состоянии")
				t.FailNow()
			}
		} else {
			if string(res) != "" {
				t.Log("error: результат отличается от необходимого в данно состоянии")
				t.FailNow()
			}
		}
	}
}

func Test_BuildClosure(t *testing.T) {
	data := []string{"1", "2", "3", "4", "5"}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		res := b.Build(context.Background(), view.Closure(func(context context.Context) view.View {
			return view.Text(data[index])
		}))

		if string(res) != data[index] {
			t.Log("error: значение не сошлось с указанным в массиве")
			t.FailNow()
		}
	}
}

func Test_BuildContexted(t *testing.T) {
	data := map[string]string{"key": "value", "key2": "value2"}
	b := example.NewBuilder()

	for key, value := range data {
		b.Build(context.Background(), view.Contexted(key, value, view.Closure(func(c context.Context) view.View {
			if c.Value(key) != value {
				t.Log("error: значение по ключу не доступно в контексте построения")
				t.FailNow()
			}

			return view.Empty()
		})))
	}
}
