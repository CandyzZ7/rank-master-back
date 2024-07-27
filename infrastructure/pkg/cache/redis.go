package cache

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"rank-master-back/infrastructure/pkg/encoding"
)

// RedisNotFound no hit cache
const RedisNotFound = redis.Nil

// redisCache redis cache object
type redisCache struct {
	client            *redis.Redis
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewRedisCache new a cache, client parameter can be passed in for unit testing
func NewRedisCache(client *redis.Redis, keyPrefix string, encoding encoding.Encoding, newObject func() interface{}) ICache {
	return &redisCache{
		client:    client,
		KeyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}
}

// Set one value
func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := encoding.Marshal(c.encoding, val)
	if err != nil {
		return fmt.Errorf("encoding.Marshal error: %v, key=%s, val=%+v ", err, key, val)
	}

	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	err = c.client.SetexCtx(ctx, cacheKey, string(buf), int(expiration))
	if err != nil {
		return fmt.Errorf("c.client.Set error: %v, cacheKey=%s", err, cacheKey)
	}
	return nil
}

// Get one value
func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	str, err := c.client.GetCtx(ctx, cacheKey)
	// NOTE: don't handle the case where redis value is nil
	// but leave it to the upstream for processing
	if err != nil {
		return err
	}

	// prevent Unmarshal from reporting an error if data is empty
	if str == "" {
		return nil
	}
	if str == NotFoundPlaceholder {
		return ErrPlaceholder
	}
	err = encoding.Unmarshal(c.encoding, []byte(str), val)
	if err != nil {
		return fmt.Errorf("encoding.Unmarshal error: %v, key=%s, cacheKey=%s, type=%v, json=%+v ",
			err, key, cacheKey, reflect.TypeOf(val), str)
	}
	return nil
}

// MultiSet set multiple values
func (c *redisCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	if len(valueMap) == 0 {
		return nil
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	// the key-value is paired and has twice the capacity of a map
	paris := make([]interface{}, 0, 2*len(valueMap))
	for key, value := range valueMap {
		buf, err := encoding.Marshal(c.encoding, value)
		if err != nil {
			fmt.Printf("encoding.Marshal error, %v, value:%v\n", err, value)
			continue
		}
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			fmt.Printf("BuildCacheKey error, %v, key:%v\n", err, key)
			continue
		}
		paris = append(paris, []byte(cacheKey))
		paris = append(paris, buf)
	}
	_, err := c.client.MsetCtx(ctx, paris...)
	if err != nil {
		return err
	}
	for i := 0; i < len(paris); i = i + 2 {
		switch paris[i].(type) {
		case []byte:
			err := c.client.ExpireCtx(ctx, string(paris[i].([]byte)), int(expiration))
			if err != nil {
				return err
			}
		default:
			fmt.Printf("redis expire is unsupported key type: %+v\n", reflect.TypeOf(paris[i]))
		}
	}
	return nil
}

// MultiGet get multiple values
func (c *redisCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
		}
		cacheKeys[index] = cacheKey
	}
	values, err := c.client.MgetCtx(ctx, cacheKeys...)
	if err != nil {
		return fmt.Errorf("c.client.MGet error: %v, keys=%+v", err, cacheKeys)
	}

	// Injection into map via reflection
	valueMap := reflect.ValueOf(value)
	for i, v := range values {
		if v == "" {
			continue
		}
		object := c.newObject()
		err = encoding.Unmarshal(c.encoding, []byte(v), object)
		if err != nil {
			fmt.Printf("unmarshal data error: %+v, key=%s, cacheKey=%s type=%v\n", err, keys[i], cacheKeys[i], reflect.TypeOf(value))
			continue
		}
		valueMap.SetMapIndex(reflect.ValueOf(cacheKeys[i]), reflect.ValueOf(object))
	}
	return nil
}

// Del delete multiple values
func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	cacheKeys := make([]string, len(keys))
	for index, key := range keys {
		cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
		if err != nil {
			continue
		}
		cacheKeys[index] = cacheKey
	}
	_, err := c.client.DelCtx(ctx, cacheKeys...)
	if err != nil {
		return fmt.Errorf("c.client.Del error: %v, keys=%+v", err, cacheKeys)
	}
	return nil
}

// SetCacheWithNotFound set value for notfound
func (c *redisCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}

	return c.client.SetexCtx(ctx, cacheKey, NotFoundPlaceholder, int(DefaultNotFoundExpireTime))
}

// LPush
func (c *redisCache) LPush(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	buf, err := encoding.Marshal(c.encoding, val)
	if err != nil {
		return fmt.Errorf("encoding.Marshal error: %v, key=%s, val=%+v ", err, key, val)
	}

	cacheKey, err := BuildCacheKey(c.KeyPrefix, key)
	if err != nil {
		return fmt.Errorf("BuildCacheKey error: %v, key=%s", err, key)
	}
	_, err = c.client.LpushCtx(ctx, cacheKey, buf)
	if err != nil {
		return fmt.Errorf("c.client.Set error: %v, cacheKey=%s", err, cacheKey)
	}
	err = c.client.ExpireCtx(ctx, cacheKey, int(expiration))
	if err != nil {
		return err
	}
	return nil
}

// BuildCacheKey construct a cache key with a prefix
func BuildCacheKey(keyPrefix string, key string) (string, error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey := key
	if keyPrefix != "" {
		cacheKey = strings.Join([]string{keyPrefix, key}, ":")
	}

	return cacheKey, nil
}
