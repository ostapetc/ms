package service

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type Ds interface {
}

func TestIDGenerator_Generate(t *testing.T) {
	config := map[string]int{
		"users": 100,
		"posts": 500,
	}

	wg := sync.WaitGroup{}
	wg.Add(600)

	generator := new(IDGenerator)

	for object, count := range config {
		for c := 0; c < count; c++ {
			go func(object string) {
				generator.Generate(object)
				wg.Done()
			}(object)
		}
	}

	wg.Wait()

	for object, count := range config {
		assert.Equal(t, int64(count), generator.Get(object))
	}
}
