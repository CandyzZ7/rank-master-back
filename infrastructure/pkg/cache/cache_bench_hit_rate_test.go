package cache

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/dgraph-io/ristretto"
)

// TestBigCacheHitRate tests the hit rate of BigCache
func TestBigCacheHitRate(t *testing.T) {
	config := bigcache.Config{
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
	cache, _ := bigcache.New(context.TODO(), config)
	hit, miss := 0, 0
	totalOps := 1000000

	for i := 0; i < totalOps/2; i++ {
		cache.Set(strconv.Itoa(i), []byte("value"))
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < totalOps; i++ {
		key := strconv.Itoa(rand.Intn(totalOps))
		if _, err := cache.Get(key); err == nil {
			hit++
		} else {
			miss++
		}
	}

	t.Logf("BigCache Hit Rate: %.2f%%", float64(hit)/float64(hit+miss)*100)
}

// TestRistrettoHitRate tests the hit rate of Ristretto
func TestRistrettoHitRate(t *testing.T) {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	hit, miss := 0, 0
	totalOps := 1000000

	for i := 0; i < totalOps/2; i++ {
		cache.Set(strconv.Itoa(i), "value", 1)
	}

	cache.Wait()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < totalOps; i++ {
		key := strconv.Itoa(rand.Intn(totalOps))
		if _, found := cache.Get(key); found {
			hit++
		} else {
			miss++
		}
	}

	t.Logf("Ristretto Hit Rate: %.2f%%", float64(hit)/float64(hit+miss)*100)
}
