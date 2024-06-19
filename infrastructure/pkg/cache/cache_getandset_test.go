package cache

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/dgraph-io/ristretto"
)

var config = bigcache.Config{
	Shards:             1024,             // Number of cache shards
	LifeWindow:         10 * time.Minute, // Duration for entries to stay in cache
	CleanWindow:        1 * time.Second,  // Interval between removing expired entries
	MaxEntriesInWindow: 1000 * 10 * 60,   // Max number of entries in life window
	MaxEntrySize:       500,              // Max size of each entry
	StatsEnabled:       false,            // Enable statistics
	Verbose:            true,             // Verbose mode
	HardMaxCacheSize:   1024,             // Hard max cache size in MB (1GB)
	Logger:             nil,              // Default logger
}

// Benchmark for BigCache Set operation
func BenchmarkBigCacheSet(t *testing.B) {
	cache, _ := bigcache.New(context.Background(), config)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		cache.Set(strconv.Itoa(i), []byte("value"))
	}
}

// Benchmark for BigCache Get operation
func BenchmarkBigCacheGet(t *testing.B) {
	cache, _ := bigcache.New(context.Background(), config)
	for i := 0; i < t.N; i++ {
		cache.Set(strconv.Itoa(i), []byte("value"))
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		cache.Get(strconv.Itoa(i))
	}
}

// Benchmark for Ristretto Set operation
func BenchmarkRistrettoSet(t *testing.B) {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		cache.Set(strconv.Itoa(i), "value", 1)
	}
}

// Benchmark for Ristretto Get operation
func BenchmarkRistrettoGet(t *testing.B) {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	for i := 0; i < t.N; i++ {
		cache.Set(strconv.Itoa(i), "value", 1)
	}
	cache.Wait()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		cache.Get(strconv.Itoa(i))
	}
}

// New benchmark test to compare sequential and random access patterns

// Benchmark for sequential BigCache Get operation
func BenchmarkBigCacheSequentialGet(t *testing.B) {
	cache, _ := bigcache.New(context.Background(), config)
	for i := 0; i < t.N; i++ {
		cache.Set(strconv.Itoa(i), []byte("value"))
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		cache.Get(strconv.Itoa(i))
	}
}

// Benchmark for random Ristretto Get operation
func BenchmarkRistrettoRandomGet(t *testing.B) {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	for i := 0; i < t.N; i++ {
		cache.Set(strconv.Itoa(i), "value", 1)
	}
	cache.Wait()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		cache.Get(strconv.Itoa(i + t.N))
	}
}
