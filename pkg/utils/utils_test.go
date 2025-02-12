package utils

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoundDuration(t *testing.T) {
	t.Run("round duration should round to default 100ms", func(t *testing.T) {

		//given
		d := time.Duration(99111111111)

		// when
		rounded := RoundDuration(d)

		// then
		assert.Equal(t, "1m39.11s", rounded.String())
	})

	t.Run("round duration should round to given value", func(t *testing.T) {

		//given
		d := time.Duration(99111111111)

		// when
		rounded := RoundDuration(d, time.Minute)

		// then
		assert.Equal(t, "2m0s", rounded.String())
	})
}

func TestSanitizeName(t *testing.T) {
	t.Parallel()

	t.Run("name should not be changed", func(t *testing.T) {

		//given
		name := "abc-123"

		// when
		sanized := SanitizeName(name)

		// then
		assert.Equal(t, "abc-123", sanized)
	})

	t.Run("name should be shorted", func(t *testing.T) {

		//given
		name := "abc" + strings.Repeat("0123456789", 10)

		// when
		sanized := SanitizeName(name)

		// then
		assert.Equal(t, "abc"+strings.Repeat("0123456789", 6), sanized)
	})

	t.Run("name should be sanitized", func(t *testing.T) {

		//given
		name := "@#$%!abc()~+123{}<>;"

		// when
		sanized := SanitizeName(name)

		// then
		assert.Equal(t, "abc-123", sanized)
	})
}
