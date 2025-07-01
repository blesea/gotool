package util

import (
	"math/rand"
	"time"
)

// RandomPool ...
type RandomPool[T any] struct {
	pool   []T
	index  int
	length int
	rng    *rand.Rand
}

// NewRandomPoolWithoutCopy 
func NewRandomPoolWithoutCopy[T any](items []T) *RandomPool[T] {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	rng.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	return &RandomPool[T]{
		pool:   items,
		index:  0,
		length: len(items),
		rng:    rng,
	}
}

// NewRandomPool ...
func NewRandomPool[T any](items []T) *RandomPool[T] {
	pool := make([]T, len(items))
	copy(pool, items)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})
	return &RandomPool[T]{
		pool:   pool,
		index:  0,
		length: len(pool),
		rng:    rng,
	}
}

// Next returns the next element in the pool in a round-robin fashion
func (rp *RandomPool[T]) Next() T {
	result := rp.pool[rp.index]
	rp.index = (rp.index + 1) % rp.length
	return result
}

// Reset reshuffles the pool and resets the index to 0
func (rp *RandomPool[T]) Reset() {
	rp.rng.Shuffle(len(rp.pool), func(i, j int) {
		rp.pool[i], rp.pool[j] = rp.pool[j], rp.pool[i]
	})
	rp.index = 0
}
