package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func Test_Add(t *testing.T) {

	rsl := base.Add(1, 2)
	expected := 3

	assert.Equal(t, rsl, expected)

}
