package main

import (
	"fmt"
	"sync"
)

func main() {
	cmap := NewConcurrentMap()
    var wg sync.WaitGroup
	//пишем конкурентно
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", i)
            cmap.Set(key, i)
        }(i)
    }

    wg.Wait()
	//читаем конкуретно
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("key%d", i)
        value, ok := cmap.Get(key)
        if ok {
            fmt.Printf("Key: %s, Value: %v\n", key, value)
        } else {
            fmt.Printf("Key: %s not found\n", key)
        }
    }
}

type ConcurrentMap struct {
	sync.RWMutex
	m map[string]interface{}
}


func NewConcurrentMap() *ConcurrentMap{
	return &ConcurrentMap{m:make(map[string]interface{})}
}

func (cm *ConcurrentMap) Set(key string, value interface{}){
	cm.Lock()
	defer cm.Unlock()
	cm.m[key] = value
}

func (cm *ConcurrentMap) Remove(key string){
	cm.Lock()
	defer cm.Unlock()
	delete(cm.m,key)
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool){
	cm.RLock()
	defer cm.RUnlock()
	value, ok := cm.m[key]
    return value, ok
}

func (cm *ConcurrentMap) Length() int {
    cm.RLock()
    defer cm.RUnlock()
    return len(cm.m)
}