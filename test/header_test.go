package test

import (
	"encoding/hex"
	"testing"

	"github.com/sunchao1/hi-im-api/pkg/im/header"
)

// goldenHex matches beehive MesgHeadHton for a fixed MesgHeader sample.
const goldenHex = "00000101000000640123456789abcdeffedcba9876543210123456781111222233334444aaaabbbbccccdddd5555666677778888"

func TestHeaderGoldenBeehive(t *testing.T) {
	h := &header.Header{
		Cmd:    0x0101,
		Length: 100,
		Sid:    0x0123456789ABCDEF,
		Cid:    0xFEDCBA9876543210,
		Nid:    0x12345678,
		Seq:    0x1111222233334444,
		Dsid:   0xAAAABBBBCCCCDDDD,
		Dseq:   0x5555666677778888,
	}

	buf, err := h.Pack()
	if err != nil {
		t.Fatalf("Pack: %v", err)
	}
	if len(buf) != header.Size {
		t.Fatalf("Size = %d, want %d", len(buf), header.Size)
	}
	if hex.EncodeToString(buf) != goldenHex {
		t.Fatalf("golden mismatch:\n got %s\nwant %s", hex.EncodeToString(buf), goldenHex)
	}

	decoded, err := header.Unmarshal(buf)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if *decoded != *h {
		t.Fatalf("roundtrip mismatch: got %+v want %+v", decoded, h)
	}
}

func TestHeaderValidate(t *testing.T) {
	h := &header.Header{Nid: 1, Sid: 0}
	if err := h.Validate(0); err != nil {
		t.Fatalf("Validate(0): %v", err)
	}
	if err := h.Validate(header.RequireSid); err != header.ErrInvalidSid {
		t.Fatalf("Validate(RequireSid) = %v, want ErrInvalidSid", err)
	}

	h.Nid = 0
	if err := h.Validate(0); err != header.ErrInvalidNid {
		t.Fatalf("Validate zero nid = %v, want ErrInvalidNid", err)
	}
}
