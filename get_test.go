package cacheoverlay

import (
	"github.com/dgraph-io/ristretto"
	"reflect"
	"testing"
)

func TestCacheOverlay_Get(t *testing.T) {
	type fields struct {
		cache     *ristretto.Cache
		retriever Retrieve
	}

	retriever := func(key string) (interface{}, error) {
		return "Hello " + key, nil
	}

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "not found",
			args: args{
				key: "world",
			},
			want:    "Hello world",
			wantErr: false,
		},
		{
			name: "found",
			args: args{
				key: "world",
			},
			want:    "Hello world",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co, _ := NewCacheOverlay(&ristretto.Config{
				NumCounters: 1e7,     // Num keys to track frequency of (10M).
				MaxCost:     1 << 30, // Maximum cost of cache (1GB).
				BufferItems: 64,      // Number of keys per Get buffer.
			}, retriever)
			got, err := co.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
