package tracker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/tracker"
)

func Test_Check_Link_Leak(t *testing.T) {
	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		err := tr.AddItem(item)
		assert.NoError(t, err)

		res := tr.GetItems()
		res[0].Name = "Second Item"

		// GetItems возвращает копию, поэтому изменение res не влияет на трекер
		assert.Equal(t,
			[]tracker.Item{item},
			tr.GetItems(),
		)
	})
}

func Test_Item_Update(t *testing.T) {
	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tr.UpdateItem(item)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})

	t.Run("success update", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		err := tr.AddItem(item)
		assert.NoError(t, err)

		updatedItem := tracker.Item{
			ID:   "1",
			Name: "Updated Item",
		}
		err = tr.UpdateItem(updatedItem)
		assert.NoError(t, err)

		items := tr.GetItems()
		assert.Len(t, items, 1)
		assert.Equal(t, "Updated Item", items[0].Name)
	})
}

func Test_AddItem_Duplicate(t *testing.T) {
	t.Run("error add duplicate", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tr.AddItem(item)
		assert.NoError(t, err)

		// Попытка добавить дубликат с тем же ID
		err = tr.AddItem(item)
		assert.ErrorIs(t, err, tracker.ErrAlreadyExists)
	})
}

func Test_DeleteItem(t *testing.T) {
	t.Run("success delete", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		err := tr.AddItem(item) // ✅ Исправлено: проверка ошибки
		assert.NoError(t, err)

		err = tr.DeleteItem("1")
		assert.NoError(t, err)
		assert.Len(t, tr.GetItems(), 0)
	})

	t.Run("error delete - not found", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		err := tr.DeleteItem("999")
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})
}
