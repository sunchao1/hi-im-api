package test

import (
	"testing"

	"github.com/sunchao1/hi-im-api/pkg/rediskey"
)

func TestRedisKeyFormat(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"sid attr", rediskey.SidAttrKey(42), "im:sid:42:attr"},
		{"gid to nid", rediskey.GidToNidZSetKey(1001), "chat:gid:1001:to:nid:zset"},
		{"gid to uid", rediskey.GidToUidZSetKey(1001), "chat:gid:1001:to:uid:zset"},
		{"rid to nid", rediskey.RidToNidZSetKey(2002), "room:rid:2002:to:nid:zset"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %q want %q", tt.got, tt.want)
			}
		})
	}
}
