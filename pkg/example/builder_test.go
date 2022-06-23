package example_test

import (
	"context"
	"testing"

	"github.com/ReanSn0w/gew/pkg/example"
	"github.com/ReanSn0w/gew/pkg/view"
)

func Test_BuildEmpty(t *testing.T) {
	b := example.NewBuilder()

	res := b.Build(context.Background(), nil)
	if string(res) != "" {
		t.Log("error: пустой view сгенерировал результат")
		t.FailNow()
	}
}

func Test_BuildText(t *testing.T) {
	data := []string{"1", "2", "3", "4", "5"}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		res := b.Build(context.Background(), example.Text(data[index]))

		if string(res) != data[index] {
			t.Log("error: значение не сошлось с указанным в массиве")
			t.FailNow()
		}
	}
}

func Test_BuildGroup(t *testing.T) {
	data := []string{"11111", "22222", "33333", "11223", "32132", ""}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		views := make([]view.View, 0, len(data[index]))
		for _, el := range data[index] {
			views = append(views, example.Text(string(el)))
		}

		res := b.Build(context.Background(), view.Group(views...))
		resStr := string(res)
		if resStr != data[index] {
			t.Log("error: не удалось построить строку через группу")
			t.FailNow()
		}
	}
}

func Test_BuildClosure(t *testing.T) {
	data := []string{"1", "2", "3", "4", "5"}
	b := example.NewBuilder()

	for index := 0; index < len(data); index++ {
		res := b.Build(context.Background(), view.Closure(func(context context.Context) view.View {
			return example.Text(data[index])
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

			return nil
		})))
	}
}

func Test_BuildWrongView(t *testing.T) {
	b := example.NewBuilder()

	res := b.Build(context.Background(), example.WrongView())
	if len(res) != 0 {
		t.Log("error: в данном тесте bilder не должен возвращать значения")
		t.FailNow()
	}
}
