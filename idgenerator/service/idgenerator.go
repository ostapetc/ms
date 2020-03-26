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

//func init() {
//	// Log as JSON instead of the default ASCII formatter.
//	log.SetFormatter(&log.JSONFormatter{})
//
//	// Output to stdout instead of the default stderr
//	// Can be any io.Writer, see below for File example
//	log.SetOutput(os.Stdout)
//
//	// Only log the warning severity or above.
//	log.SetLevel(log.InfoLevel)
//}

//func requestHandler(ctx *fasthttp.RequestCtx) {
//	atomic.AddInt64(&ID, 1)
//
//	fmt.Fprint(ctx, ID)
//	log.Info("ID ", ID)
//}
