package helper_test

import (
	"github.com/satriahrh/adam-smith/internal/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringContainsFrom(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		assert.True(t, helper.StringContainsFrom("halo", "halo", 0))
	})
	t.Run("NotEqual", func(t *testing.T) {
		assert.False(t, helper.StringContainsFrom("halo", "loha", 0))
	})
	t.Run("SubStringExceedingFromZero", func(t *testing.T) {
		assert.False(t, helper.StringContainsFrom("halo", "haloa", 0))
	})
	t.Run("SubStringExceeding", func(t *testing.T) {
		assert.False(t, helper.StringContainsFrom("halo", "halo", 1))
	})
	t.Run("NotMatchedMiddle", func(t *testing.T) {
		assert.False(t, helper.StringContainsFrom("halo", "ho", 1))
	})
	t.Run("MatchedLast", func(t *testing.T) {
		assert.True(t, helper.StringContainsFrom("halo", "lo", 2))

	})
	t.Run("MatchedFirst", func(t *testing.T) {
		assert.True(t, helper.StringContainsFrom("halo", "ha", 0))
	})
}
