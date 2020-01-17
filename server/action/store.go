package action

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

// Get ...
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()

	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return errors.New("not found")
}

// Set ...
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

// Watch ...
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {

	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())

	fmt.Println("m: ", p.m)
	fmt.Println("filter: ", p.filter)

	ch := make(chan string, 10)

	p.mu.Lock()
	p.filter[id] = func(key string) {
		fmt.Println("key1: ", key)
		ch <- key
	}
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		fmt.Println("key2: ", key)
		*keyChanged = key
		return nil
	}

	return nil
}
