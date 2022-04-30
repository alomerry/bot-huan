package component

import (
	"bot-huan/core/log"
	"github.com/panjf2000/ants/v2"
)

var (
	pool *ants.Pool
)

func init() {
	initPool()
}

func initPool() {
	var err error
	pool, err = ants.NewPool(5, ants.WithPanicHandler(func(i interface{}) {
		log.GetLogger().Error().Msg("Panic")
	}))
	if err != nil {
		log.GetLogger().Err(err)
	}
}

func GetGoroutinePool() *ants.Pool {
	if pool == nil {
		initPool()
	}
	return pool
}

func SubmitTask(f func()) error {
	return GetGoroutinePool().Submit(f)
}
