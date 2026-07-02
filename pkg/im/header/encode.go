package header

import "encoding/binary"

// Marshal writes h into buf using big-endian field order (beehive MesgHeadHton).
func (h *Header) Marshal(buf []byte) error {
	if len(buf) < Size {
		return ErrBufferTooShort
	}
	binary.BigEndian.PutUint32(buf[0:4], h.Cmd)
	binary.BigEndian.PutUint32(buf[4:8], h.Length)
	binary.BigEndian.PutUint64(buf[8:16], h.Sid)
	binary.BigEndian.PutUint64(buf[16:24], h.Cid)
	binary.BigEndian.PutUint32(buf[24:28], h.Nid)
	binary.BigEndian.PutUint64(buf[28:36], h.Seq)
	binary.BigEndian.PutUint64(buf[36:44], h.Dsid)
	binary.BigEndian.PutUint64(buf[44:52], h.Dseq)
	return nil
}

// Pack returns a newly allocated wire buffer for h.
func (h *Header) Pack() ([]byte, error) {
	buf := make([]byte, Size)
	if err := h.Marshal(buf); err != nil {
		return nil, err
	}
	return buf, nil
}
