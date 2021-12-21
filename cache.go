//Package cache Cache with data expiration is based on fastcache
package cache

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/sotax/cache/base"
	"sync"
	"time"
	"unsafe"
)

const CACHE_DEFAULT_SIZE int32 = 10 * (1 << 20)
const ASYNC_REFRESH_NOW_DEFAULT_TIME int64 = 2000 // Asynchronously refresh the current time(now), unit: millisecond

// object defination
type Cache struct {
	Expire int64 // data expiration time, unit: ms. if the value is 0, it will not expire.
	Size   int32 // cache space size, unit: byte.
	// The cycle of asynchronously refresh the current time(now), unit: millisecond.
	// The larger the value, the worse the data expiration accuracy.
	NowRefreshCycle int64
	now             int64 // The current time's cache, unit: ms
	cacheOnce       sync.Once
	cacheInst       *fastcache.Cache // storage object
}

func (c *Cache) Init() {
	init := func() {
		// create
		if c.cacheInst == nil {
			var allocSize = CACHE_DEFAULT_SIZE
			if c.Size > 0 {
				allocSize = c.Size
			}
			c.cacheInst = fastcache.New(int(allocSize))
		}
		// params init
		c.now = time.Now().UnixNano() / 1e6  // millisecond
		c.NowRefreshCycle = ASYNC_REFRESH_NOW_DEFAULT_TIME
		c.asyncNow() // start now's refresh thread
	}
	// initialize only once
	c.cacheOnce.Do(init)

	return
}

func (c *Cache) Get(key string) []byte {
	// query
	data, exist := c.cacheInst.HasGet(nil, []byte(key))
	if !exist {
		return nil
	}
	// 前几个字节为expire过期时间
	headerSize := unsafe.Sizeof(c.Expire)
	if int(headerSize) > cap(data) { // 越界
		return nil
	}
	// check expiration while reading
	if c.Expire > 0 { // if the value is greater than 0, do the expiration check.
		expire := base.BytesToInt64(data[0:headerSize])
		if c.now > expire { // expired
			c.cacheInst.Del([]byte(key)) // delete
			return nil
		}
	}

	return data[headerSize:]
}

func (c *Cache) Set(key string, value []byte) {
	// storage format: header(expiration time, 8 bytes) + data
	c.cacheInst.Set([]byte(key), append(base.Int64ToBytes(c.now+c.Expire), value...))
}

// clear all data
func (c *Cache) Clear() {
	c.cacheInst.Reset()
}

// refresh the current time(now) asynchronously to improve performance
func (c *Cache) asyncNow() {
	go func() {
		for {
			if c.NowRefreshCycle <= 0 { // invalid value, so use default
				c.NowRefreshCycle = ASYNC_REFRESH_NOW_DEFAULT_TIME
			}
			c.now = time.Now().UnixNano() / 1e6 // 在用户空间完成未发生系统调用, 因使用vDSO技术
			time.Sleep(time.Duration(c.NowRefreshCycle) * time.Millisecond)
		}
	}()
}
