package rediskey

import "fmt"

// Key templates migrated from beehive lib/comm/key.go and chatroom/models/key.go.
// String patterns are unchanged to stay binary/data compatible.

const (
	KeySidAttr       = "im:sid:%d:attr"
	KeyGidToNidZSet  = "chat:gid:%d:to:nid:zset"
	KeyGidToUidZSet  = "chat:gid:%d:to:uid:zset"
	KeyRidToNidZSet  = "room:rid:%d:to:nid:zset"
)

func SidAttrKey(sid uint64) string {
	return fmt.Sprintf(KeySidAttr, sid)
}

func GidToNidZSetKey(gid uint64) string {
	return fmt.Sprintf(KeyGidToNidZSet, gid)
}

func GidToUidZSetKey(gid uint64) string {
	return fmt.Sprintf(KeyGidToUidZSet, gid)
}

func RidToNidZSetKey(rid uint64) string {
	return fmt.Sprintf(KeyRidToNidZSet, rid)
}
