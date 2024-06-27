package cache

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/dgraph-io/ristretto"

	"rank-master-back/infrastructure/pkg/encoding"
)

var RistrettoCacheNotFound = errors.New("ristretto: key not found")

type ristrettoCache struct {
	client            *ristretto.Cache
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewRistrettoCache create a memory cache
func NewRistrettoCache(keyPrefix string, encoding encoding.Encoding, newObject func() interface{}) ICache {
	// see: https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/
	//		https://www.start.io/blog/we-chose-ristretto-cache-for-go-heres-why/
	config := &ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	}
	store, _ := ristretto.NewCache(config)
	return &ristrettoCache{
		client:    store,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}
}

// Set data
func (m *ristrettoCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := encoding.Marshal(m.encoding, val)
	if err != nil {
		return fmt.Errorf("encoding.Marshal error: %v, key=%s, val=%+v ", err, key, val)
	}
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}
	ok := m.client.SetWithTTL(cacheKey, buf, 0, expiration)
	if !ok {
		return errors.New("SetWithTTL failed")
	}

	return nil
}

// Get data
func (m *ristrettoCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	data, ok := m.client.Get(cacheKey)
	if !ok {
		return RistrettoCacheNotFound
	}

	if string(data.([]byte)) == NotFoundPlaceholder {
		return ErrPlaceholder
	}

	err = encoding.Unmarshal(m.encoding, data.([]byte), val)
	if err != nil {
		return fmt.Errorf("encoding.Unmarshal error: %v, key=%s, cacheKey=%s, type=%v, json=%+v ",
			err, key, cacheKey, reflect.TypeOf(val), string(data.([]byte)))
	}
	return nil
}

// Del delete data
func (m *ristrettoCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	key := keys[0]
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("build cache key error, err=%v, key=%s", err, key)
	}
	m.client.Del(cacheKey)
	return nil
}

// MultiSet multiple set data
func (m *ristrettoCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	var err error
	for key, value := range valueMap {
		err = m.Set(ctx, key, value, expiration)
		if err != nil {
			return err
		}
	}
	return nil
}

// MultiGet multiple get data
func (m *ristrettoCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
	valueMap := reflect.ValueOf(value)
	var err error
	for _, key := range keys {
		object := m.newObject()
		err = m.Get(ctx, key, object)
		if err != nil {
			continue
		}
		valueMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(object))
	}

	return nil
}

// SetCacheWithNotFound set not found
func (m *ristrettoCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	ok := m.client.SetWithTTL(cacheKey, []byte(NotFoundPlaceholder), 0, DefaultNotFoundExpireTime)
	if !ok {
		return errors.New("SetWithTTL failed")
	}

	return nil
}

// LPush
func (m *ristrettoCache) LPush(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	data, ok := m.client.Get(cacheKey)
	if !ok {
		return RistrettoCacheNotFound
	}

	if string(data.([]byte)) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	var list []interface{}
	err = encoding.Unmarshal(m.encoding, data.([]byte), val)
	if err != nil {
		return fmt.Errorf("encoding.Unmarshal error: %v, key=%s, cacheKey=%s, type=%v, json=%+v ",
			err, key, cacheKey, reflect.TypeOf(val), string(data.([]byte)))
	}
	vals := []interface{}{val}
	// 将新值添加到列表的前面
	vals = append(vals, list...)
	buf, err := encoding.Marshal(m.encoding, vals)
	if err != nil {
		return err
	}
	ok = m.client.SetWithTTL(cacheKey, buf, 0, expiration)
	if !ok {
		return errors.New("SetWithTTL failed")
	}
	return nil
}
