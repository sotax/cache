package cache

import (
	"reflect"
	"sync"
	"testing"

	"github.com/VictoriaMetrics/fastcache"
)

func TestCache_Get(t *testing.T) {
	c := Cache{
		Expire:          10,
		NowRefreshCycle: 50000,
	}
	c.Init()
	key := "test_key01"
	value := []byte("test_value01")
	c.Set(key, value)

	type fields struct {
		Expire          int64
		Size            int32
		NowRefreshCycle int64
		now             int64
		cacheOnce       sync.Once
		cacheInst       *fastcache.Cache
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
		{
			name: "cache_expire_test",
			fields: fields{
				Size:            c.Size,
				Expire:          c.Expire,
				NowRefreshCycle: c.NowRefreshCycle,
				now:             c.now,
				cacheOnce:       c.cacheOnce,
				cacheInst:       c.cacheInst,
			},
			args: args{
				key: key,
			},
			want: value,
		},
		{
			name: "cache_expire_test",
			fields: fields{
				Size:            c.Size,
				Expire:          0, // no expire
				NowRefreshCycle: c.NowRefreshCycle,
				now:             c.now,
				cacheOnce:       c.cacheOnce,
				cacheInst:       c.cacheInst,
			},
			args: args{
				key: key,
			},
			want: value,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				Expire:          tt.fields.Expire,
				Size:            tt.fields.Size,
				NowRefreshCycle: tt.fields.NowRefreshCycle,
				now:             tt.fields.now,
				cacheOnce:       tt.fields.cacheOnce,
				cacheInst:       tt.fields.cacheInst,
			}
			if got := c.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
