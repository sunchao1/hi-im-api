package header

// Header is the IM wire header (beehive MesgHeader), packed big-endian on the wire.
type Header struct {
	Cmd    uint32
	Length uint32
	Sid    uint64
	Cid    uint64
	Nid    uint32
	Seq    uint64
	Dsid   uint64
	Dseq   uint64
}

// Size is the on-wire header length in bytes (beehive MesgHeadHton layout).
const Size = 52
