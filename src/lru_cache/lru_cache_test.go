package lrucache

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LRUCacheTestSuite struct {
	suite.Suite
	cache *LRUCache
}

func (s *LRUCacheTestSuite) SetupTest() {
	s.cache = NewLRUCache(2)
}

func (s *LRUCacheTestSuite) TestBasicOperations() {
	s.cache.Put(1, 1)
	s.cache.Put(2, 2)

	val, ok := s.cache.Get(1)
	s.True(ok)
	s.Equal(1, val)

	// Test eviction
	s.cache.Put(3, 3) // This should evict key 2

	_, ok = s.cache.Get(2)
	s.False(ok)

	val, ok = s.cache.Get(3)
	s.True(ok)
	s.Equal(3, val)
}

func (s *LRUCacheTestSuite) TestUpdatingExistingKey() {
	s.cache.Put(1, 1)
	s.cache.Put(1, 2)

	val, ok := s.cache.Get(1)
	s.True(ok)
	s.Equal(2, val)
}

func (s *LRUCacheTestSuite) TestLRUEvictionOrder() {
	cache := NewLRUCache(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4) // This should evict key 1

	_, ok := cache.Get(1)
	s.False(ok)

	cache.Get(2)    // This should move 2 to the front
	cache.Put(5, 5) // This should evict key 3

	_, ok = cache.Get(3)
	s.False(ok)
}

func (s *LRUCacheTestSuite) TestCacheCapacity() {
	cache := NewLRUCache(1)

	cache.Put(1, 1)
	cache.Put(2, 2)

	_, ok := cache.Get(1)
	s.False(ok)

	val, ok := cache.Get(2)
	s.True(ok)
	s.Equal(2, val)
}

func (s *LRUCacheTestSuite) TestLargeCache() {
	cache := NewLRUCache(1000)

	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}

	for i := 0; i < 1000; i++ {
		val, ok := cache.Get(i)
		s.True(ok)
		s.Equal(i, val)
	}

	cache.Put(1000, 1000) // This should evict key 0

	_, ok := cache.Get(0)
	s.False(ok)

	val, ok := cache.Get(1000)
	s.True(ok)
	s.Equal(1000, val)
}

func TestLRUCacheSuite(t *testing.T) {
	suite.Run(t, new(LRUCacheTestSuite))
}
