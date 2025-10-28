package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	numKeys := []Key{"one", "two", "three", "four", "five", "six", "seven"}

	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(5)

		for i, key := range numKeys {
			c.Set(key, i)
		}

		// первых двух быть не должно
		val, ok := c.Get(numKeys[0])
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get(numKeys[1])
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("overflow logic", func(t *testing.T) {
		c := NewCache(5)

		nextKeys := []Key{"six", "seven"}
		for i, key := range numKeys {
			c.Set(key, i)
		}

		// Должно быть - 5,4,3,2,1
		// поднимем 2 элемента
		val, ok := c.Get(numKeys[4])
		require.True(t, ok)
		require.Equal(t, 4, val)

		wasInCache := c.Set(numKeys[3], 500)
		require.True(t, wasInCache)
		// Должно стать 4,5,3,2,1

		// Добавим еще парочку
		for i, key := range nextKeys {
			c.Set(key, i)
		}

		// один и два должны покинуть чат
		val, ok = c.Get(numKeys[0])
		require.False(t, ok)
		require.Equal(t, nil, val)

		val, ok = c.Get(numKeys[1])
		require.False(t, ok)
		require.Equal(t, nil, val)
	})

	t.Run("clear logic", func(t *testing.T) {
		c := NewCache(5)

		for i, key := range numKeys {
			c.Set(key, i)
		}

		c.Clear()

		// Ничего из добавленного нет
		for _, key := range numKeys {
			val, ok := c.Get(key)
			require.False(t, ok)
			require.Equal(t, nil, val)
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	var sets, gets, hits, misses int64

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
			atomic.AddInt64(&sets, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			if _, ok := c.Get(Key(strconv.Itoa(rand.Intn(1_000_000)))); ok {
				atomic.AddInt64(&hits, 1)
			} else {
				atomic.AddInt64(&misses, 1)
			}
			atomic.AddInt64(&gets, 1)
		}
	}()

	wg.Wait()

	t.Logf("done: sets=%d, gets=%d (hits=%d, misses=%d)",
		atomic.LoadInt64(&sets),
		atomic.LoadInt64(&gets),
		atomic.LoadInt64(&hits),
		atomic.LoadInt64(&misses),
	)

	if atomic.LoadInt64(&hits)+atomic.LoadInt64(&misses) != atomic.LoadInt64(&gets) {
		t.Fatalf("inconsistent read counters: hits+misses=%d, gets=%d",
			atomic.LoadInt64(&hits)+atomic.LoadInt64(&misses),
			atomic.LoadInt64(&gets),
		)
	}
}
