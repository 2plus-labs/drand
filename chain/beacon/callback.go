package beacon

import (
	"github.com/drand/drand/chain"
	"sync"
)

type CallbackCache struct {
	callbacks map[string]chan chain.Beacon
	state     sync.RWMutex
}

func NewCallbackCache() *CallbackCache {
	return &CallbackCache{
		callbacks: make(map[string]chan chain.Beacon, 0),
		state:     sync.RWMutex{},
	}
}

func (f *CallbackCache) Add(key string, value chan chain.Beacon) {
	f.state.Lock()
	defer f.state.Unlock()
	f.callbacks[key] = value
}

func (f *CallbackCache) Get(key string) chan chain.Beacon {
	f.state.RLock()
	defer f.state.RUnlock()
	if _, ok := f.callbacks[key]; ok {
		return f.callbacks[key]
	}
	return nil
}

func (f *CallbackCache) Delete(key string) {
	f.state.Lock()
	defer f.state.Unlock()
	if _, ok := f.callbacks[key]; ok {
		delete(f.callbacks, key)
	}
}

func (f *CallbackCache) Length() int {
	f.state.RLock()
	defer f.state.RUnlock()
	return len(f.callbacks)
}
