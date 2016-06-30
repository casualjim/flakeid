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

/*
func TestConcurrentFlakeIdGenerator(t *testing.T) {
	result := make(chan []string)
	done := make(chan [][]string)
	runtime.GOMAXPROCS(4)
	const (
		nrIDs     = 10000
		nrWorkers = 500
	)

	go func() {
		var res [][]string
		for vv := range result {
			res = append(res, vv)
		}
		done <- res
	}()

	for i := 0; i < nrWorkers; i++ {
		go func() {
			var partial []string
			for j := 0; j < nrIDs; j++ {
				partial = append(partial, MustNewID())
			}
			result <- partial
		}()
	}

	collected := <-done
	set := make(map[string]struct{})
	assert.Len(t, collected, nrWorkers)
	for _, part := range collected {
		assert.Len(t, part, nrIDs)
		for _, id := range part {
			_, known := set[id]
			if known {
				t.Log("found a duplicate ID!!!!")
				t.FailNow()
			}
			set[id] = struct{}{}
		}
	}
}
*/
