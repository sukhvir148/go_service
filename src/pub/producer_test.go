package pub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	i, j := 1, 2
	producer := NewProducer(nil, nil, nil)
	expected := 3

	got := producer.Add(i, j)
	assert.Equal(t, got, expected)
}
