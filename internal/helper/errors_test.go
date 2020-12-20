package helper_test

import (
	"errors"
	"github.com/satriahrh/adam-smith/internal/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorParseFromDictionary(t *testing.T) {
	errorDictionary := []helper.ErrorDictionaryItem{
		{0, "expected", errors.New("expected error")},
	}
	t.Run("UnknownError", func(t *testing.T) {
		assert.EqualError(t,
			helper.ErrorParseFromDictionary(errors.New("you won't be expecting this"), errorDictionary),
			"unknown error",
		)
	})
	t.Run("ExpectedError", func(t *testing.T) {
		assert.EqualError(t,
			helper.ErrorParseFromDictionary(errors.New("expected error yes it is"), errorDictionary),
			"expected error",
		)
	})
}
