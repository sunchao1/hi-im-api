package test

import (
	"testing"

	imv1 "github.com/sunchao1/hi-im-api/gen/go/im/v1"
	"google.golang.org/protobuf/proto"
)

func TestProtoGroupChatRoundtrip(t *testing.T) {
	msg := &imv1.GroupChat{
		Uid:   1001,
		Gid:   2002,
		Level: 1,
		Time:  1719900000,
		Text:  "hello group",
		Data:  []byte("payload"),
	}
	buf, err := proto.Marshal(msg)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	out := &imv1.GroupChat{}
	if err := proto.Unmarshal(buf, out); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if !proto.Equal(msg, out) {
		t.Fatalf("roundtrip mismatch: got %+v want %+v", out, msg)
	}
}
