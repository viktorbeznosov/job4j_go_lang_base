package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func Test_Validate_Request_Empty(t *testing.T) {
	var empteRequest *base.ValidateRequest

	empteRequest = &base.ValidateRequest{}

	result := base.Validate(empteRequest)

	assert.Contains(t, result, "UserID is empty")
	assert.Contains(t, result, "Title is empty")
	assert.Contains(t, result, "Description is empty")
	assert.Equal(t, len(result), 3)
}

func Test_Validate_Request_Is_Ok(t *testing.T) {
	var request *base.ValidateRequest

	request = &base.ValidateRequest{
		UserID:      "1",
		Description: "test",
		Title:       "Test User",
	}

	result := base.Validate(request)

	assert.Empty(t, result)
}
