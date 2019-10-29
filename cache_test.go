package cacheoverlay

import (
	"github.com/dgraph-io/ristretto"
	"testing"
)

func TestNewCacheOverlay(t *testing.T) {
	type args struct {
		retriever Retrieve
	}
	retriever := func(string) (interface{}, error) {
		return "hello", nil
	}
	tests := []struct {
		name    string
		args    args
		want    *CacheOverlay
		wantErr bool
	}{
		{
			name:    "no retriever",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "retriever",
			args: args{retriever},
			want: &CacheOverlay{
				retriever: retriever,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCacheOverlay(&ristretto.Config{
				NumCounters: 1e7,     // Num keys to track frequency of (10M).
				MaxCost:     1 << 30, // Maximum cost of cache (1GB).
				BufferItems: 64,      // Number of keys per Get buffer.
			}, tt.args.retriever)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCacheOverlay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
