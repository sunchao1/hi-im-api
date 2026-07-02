package header

import "encoding/binary"

// Unmarshal parses a wire header from buf (beehive MesgHeadNtoh).
func Unmarshal(buf []byte) (*Header, error) {
	if len(buf) < Size {
		return nil, ErrBufferTooShort
	}
	return &Header{
		Cmd:    binary.BigEndian.Uint32(buf[0:4]),
		Length: binary.BigEndian.Uint32(buf[4:8]),
		Sid:    binary.BigEndian.Uint64(buf[8:16]),
		Cid:    binary.BigEndian.Uint64(buf[16:24]),
		Nid:    binary.BigEndian.Uint32(buf[24:28]),
		Seq:    binary.BigEndian.Uint64(buf[28:36]),
		Dsid:   binary.BigEndian.Uint64(buf[36:44]),
		Dseq:   binary.BigEndian.Uint64(buf[44:52]),
	}, nil
}
