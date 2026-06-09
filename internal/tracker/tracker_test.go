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
		tr.AddItem(item)

		res := tr.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]tracker.Item{item},
			tr.GetItems(),
		)
	})
}
