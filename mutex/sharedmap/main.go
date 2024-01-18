package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"sync"
)

type shard struct {
	sync.RWMutex
	m map[string]any
}

type ShardedMap []*shard

func NewShardedMap(n int) ShardedMap {
	shards := make([]*shard, n)
	for i := 0; i < n; i++ {
		m := make(map[string]any)
		shards[i] = &shard{m: m}
	}
	return shards
}

func (sm ShardedMap) getShardIndex(key string) int {
	checksum := sha1.New().Sum([]byte(key))
	hash := int(checksum[17])
	return hash % len(sm)
}

func (sm ShardedMap) getShard(key string) *shard {
	index := sm.getShardIndex(key)
	return sm[index]
}

func (sm ShardedMap) Get(key string) any {
	s := sm.getShard(key)
	s.RLock()
	defer s.RUnlock()
	return s.m[key]
}

func (sm ShardedMap) Set(key string, value any) {
	s := sm.getShard(key)
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (sm ShardedMap) Keys() []string {
	keys := make([]string, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(sm))
	for _, s := range sm {
		go func(s *shard) {
			defer wg.Done()
			s.RLock()
			defer s.RUnlock()
			mu.Lock()
			defer mu.Unlock()
			for key := range s.m {
				keys = append(keys, key)
			}
		}(s)
	}
	wg.Wait()
	return keys
}

func main() {
	shardedMap := NewShardedMap(5)

	shardedMap.Set("alpha", 1)
	shardedMap.Set("beta", 2)
	shardedMap.Set("gamma", 3)

	fmt.Println(shardedMap.Get("alpha"))
	fmt.Println(shardedMap.Get("beta"))
	fmt.Println(shardedMap.Get("gamma"))

	for i := 0; i < 10000; i++ {
		shardedMap.Set(strconv.Itoa(i), i)
	}

	keys := shardedMap.Keys()
	// for _, k := range keys {
	// 	fmt.Println(k)
	// }
	fmt.Println("len: ", len(keys)) // see, there is a bug. fixed by adding a mutex.
}
