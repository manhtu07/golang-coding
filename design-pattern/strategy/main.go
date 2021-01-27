package main

import "fmt"

type cache struct {
	Value        map[int]string
	Capacity     int
	MaxCapacity  int
	EvictionAlgo evictionAlgo
}

func NewEvictionAlgo(e evictionAlgo) *cache {
	value := make(map[int]string)
	return &cache{
		EvictionAlgo: e,
		Capacity:     0,
		MaxCapacity:  3,
		Value:        value,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.EvictionAlgo = e
}

func (c *cache) add(key int, value string) {
	if c.Capacity == c.MaxCapacity {
		c.evict()
	}
	c.Value[key] = value
	c.Capacity++
}

func (c *cache) evict() {
	c.EvictionAlgo.evict(c)
	c.Capacity--
}

type evictionAlgo interface {
	evict(cache *cache)
}

type lru struct{}

func (l *lru) evict(cache *cache) {
	fmt.Println("lru is calling")
}

type lfu struct{}

func (l *lfu) evict(cache *cache) {
	fmt.Println("lfu is calling")
}

func main() {
	newCache := NewEvictionAlgo(&lru{})
	newCache.add(1, "a")
	newCache.add(2, "b")
	newCache.add(3, "c")
	newCache.add(4, "d")
	newCache.setEvictionAlgo(&lfu{})
	newCache.add(5, "e")
}
