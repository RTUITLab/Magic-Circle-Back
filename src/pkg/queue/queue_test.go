package queue_test

import (
	"testing"

	"github.com/0B1t322/Magic-Circle/pkg/queue"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run(
		"workTest",
		func(t *testing.T) {
			var q queue.StringQueue = []string{"1", "2", "3"}

			g, err := q.Get()
			assert.NoError(t, err)
			assert.Equal(t, g, "1")
			t.Logf("%v", q)

			g, err = q.Get()
			assert.NoError(t, err)
			assert.Equal(t, g, "2")
			t.Logf("%v", q)

			g, err = q.Get()
			assert.NoError(t, err)
			assert.Equal(t, g, "3")
			t.Logf("%v", q)

			_, err = q.Get()
			assert.ErrorIs(t, err, queue.QueueIsEmpty)
			t.Logf("%v", q)
		},
	)

	t.Run(
		"cast test with ignore",
		func(t *testing.T) {
			var q queue.StringQueue = []string{"1", "2", "a","3"}
			intQueue, err := queue.StringQueueToIntQueue(queue.StringQueueToIntOpts{IfNotIntElemIgnore: true})(q)
			assert.NoError(t, err)


			for i, elem := range intQueue {
				assert.Equal(t, i+1, elem)
			}
		},
	)
}