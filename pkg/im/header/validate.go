package header

import "errors"

var (
	ErrBufferTooShort = errors.New("header: buffer too short")
	ErrInvalidNid     = errors.New("header: nid must be non-zero")
	ErrInvalidSid     = errors.New("header: sid must be non-zero")
)

// ValidateFlags controls optional sid validation (beehive IsValid flag).
type ValidateFlags uint32

const (
	RequireSid ValidateFlags = 1
)

// Validate checks nid/sid rules migrated from beehive MesgHeader.IsValid.
func (h *Header) Validate(flags ValidateFlags) error {
	if h.Nid == 0 {
		return ErrInvalidNid
	}
	if flags != 0 && h.Sid == 0 {
		return ErrInvalidSid
	}
	return nil
}
