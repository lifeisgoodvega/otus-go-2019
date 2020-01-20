package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	list := List{}
	assert.Equal(t, (*Item)(nil), list.First())
	assert.Equal(t, (*Item)(nil), list.Last())
}

func TestOneElement(t *testing.T) {
	list := List{}
	list.PushBack(1)

	assert.Equal(t, list.First().Value(), 1)
	assert.Equal(t, list.Last().Value(), 1)
	assert.Equal(t, list.Len(), 1)

	list.Remove(*list.First())
	assert.Equal(t, (*Item)(nil), list.First())
	assert.Equal(t, (*Item)(nil), list.Last())
	assert.Equal(t, 0, list.Len())
}

func TestRemoveMiddleElement(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	assert.Equal(t, list.First().Value(), 1)
	assert.Equal(t, list.Last().Value(), 3)
	assert.Equal(t, list.Len(), 3)

	i := list.First().Next()
	assert.Equal(t, i.Value(), 2)

	list.Remove(*i)
	assert.Equal(t, list.Len(), 2)
}

func TestRemoveFirstElement(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	i := list.First()
	assert.Equal(t, i.Value(), 1)

	list.Remove(*i)
	assert.Equal(t, list.Len(), 2)
	assert.Equal(t, list.First().Value(), 2)
}

func TestRemoveLastElement(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	i := list.Last()
	assert.Equal(t, i.Value(), 3)

	list.Remove(*i)
	assert.Equal(t, list.Len(), 2)
	assert.Equal(t, list.Last().Value(), 2)
}

func TestInvalidate(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	i := list.Last()
	assert.Equal(t, i.Value(), 3)

	assert.NoError(t, list.Remove(*i))
	assert.Error(t, list.Remove(*i), "Item was deleted, therefore invalidated")
}

func TestPushFront(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushFront(0)

	assert.Equal(t, list.First().Value(), 0)
	assert.Equal(t, list.Last().Value(), 1)
	assert.Equal(t, list.Len(), 2)
}

func TestItemPrev(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	i := list.Last()
	for k := 3; k > 0; k-- {
		assert.Equal(t, k, i.Value())
		i = i.Prev()
	}

	assert.Equal(t, (*Item)(nil), i)

	i = list.Last()
	for k := 3; k > 0; k-- {
		prev := i.Prev()
		list.Remove(*i)
		i = prev
	}

	assert.Equal(t, 0, list.Len())
}

func TestErrAlienItem(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list2 := List{}
	list2.PushBack(1)

	assert.Error(t, list.Remove(*list2.First()), "Item must belong to list, which method is called")
}
