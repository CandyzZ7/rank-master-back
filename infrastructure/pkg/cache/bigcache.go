package cache

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/allegro/bigcache/v3"

	"rank-master-back/infrastructure/pkg/encoding"
)

type bigcacheCache struct {
	client            *bigcache.BigCache
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewBigcacheCache create a memory cache
func NewBigcacheCache(ctx context.Context, keyPrefix string, encoding encoding.Encoding, newObject func() interface{}) (*bigcacheCache, error) {
	config := bigcache.DefaultConfig(DefaultExpireTime)
	store, err := bigcache.New(ctx, config)
	if err != nil {
		return nil, err
	}
	return &bigcacheCache{
		client:    store,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}, nil
}

// Set data
func (m *bigcacheCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := encoding.Marshal(m.encoding, val)
	if err != nil {
		return fmt.Errorf("encoding.Marshal error: %v, key=%s, val=%+v ", err, key, val)
	}
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}
	err = m.client.Set(cacheKey, buf)
	if err != nil {
		return err
	}

	return nil
}

// Get data
func (m *bigcacheCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	data, err := m.client.Get(cacheKey)
	if err != nil {
		return err
	}

	if string(data) == NotFoundPlaceholder {
		return ErrPlaceholder
	}

	err = encoding.Unmarshal(m.encoding, data, val)
	if err != nil {
		return fmt.Errorf("encoding.Unmarshal error: %v, key=%s, cacheKey=%s, type=%v, json=%+v ",
			err, key, cacheKey, reflect.TypeOf(val), string(data))
	}
	return nil
}

// Del delete data
func (m *bigcacheCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	key := keys[0]
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("build cache key error, err=%v, key=%s", err, key)
	}
	err = m.client.Delete(cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// MultiSet multiple set data
func (m *bigcacheCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
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
func (m *bigcacheCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
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
func (m *bigcacheCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	err = m.client.Set(cacheKey, []byte(NotFoundPlaceholder))
	if err != nil {
		return err
	}

	return nil
}

// LPush
func (m *bigcacheCache) LPush(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	cacheKey, err := BuildCacheKey(m.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	data, err := m.client.Get(cacheKey)
	if err != nil {
		return err
	}

	if string(data) == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	var list []interface{}
	err = encoding.Unmarshal(m.encoding, data, &list)
	if err != nil {
		return fmt.Errorf("encoding.Unmarshal error: %v, key=%s, cacheKey=%s, type=%v, json=%+v ",
			err, key, cacheKey, reflect.TypeOf(val), string(data))
	}
	vals := []interface{}{val}
	// 将新值添加到列表的前面
	vals = append(vals, list...)
	buf, err := encoding.Marshal(m.encoding, vals)
	if err != nil {
		return err
	}
	err = m.client.Set(cacheKey, buf)
	if err != nil {
		return err
	}
	return nil
}
