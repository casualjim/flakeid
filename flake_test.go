package flakeid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlakeIdGenerator(t *testing.T) {
	id1 := DefaultWorkID()
	id2 := DefaultWorkID()

	assert.Equal(t, id1, id2)
	total := 1000 * 1000
	data := make(map[string]int)

	flake := NewFlake()
	for i := 0; i < total; i++ {
		id, err := flake.Next()
		if assert.NoError(t, err) {
			assert.Equal(t, 0, data[id])
			data[id] = 1
		}
	}

	assert.Len(t, data, total)
}
