package draft_test

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	. "github.com/kamilsk/passport/pkg/draft"
	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"constant", 5},
		{"random", rand.New(rand.NewSource(time.Now().UnixNano())).Int()},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Len(t, Sequence(tc.size), tc.size)
		})
	}
}

func TestSafe(t *testing.T) {
	tests := []struct {
		name   string
		action func() error
		closer func(error)
	}{
		{"with error", func() error { return errors.New("error") }, func(err error) { assert.Error(t, err) }},
		{"with panic", func() error { panic(errors.New("panic")) }, func(err error) { assert.Error(t, err) }},
		{"without anything", func() error { return nil }, func(err error) { assert.NoError(t, err) }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.NotPanics(t, func() { Safe(tc.action, tc.closer) })
		})
	}
}
