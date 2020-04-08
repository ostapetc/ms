package service

import (
	"sync"
)

type IDGenerator struct {
	mx sync.Mutex
	id map[string]int64
}

func (i *IDGenerator) Generate(object string) int64 {
	i.mx.Lock()
	defer i.mx.Unlock()

	if i.id == nil {
		i.id = make(map[string]int64)
	}

	i.id[object]++

	return i.id[object]
}

func (i *IDGenerator) Get(object string) int64 {
	return i.id[object]
}
