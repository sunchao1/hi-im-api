package cmd

// M2 minimal command set migrated from beehive lib/comm/mesg.go.
// ACK / NTF pairs are documented inline; expand in later milestones.

const (
	CMD_UNKNOWN = 0

	// 0x01xx lifecycle
	CMD_ONLINE     = 0x0101 // ONLINE request (server-bound); ACK: CMD_ONLINE_ACK
	CMD_ONLINE_ACK = 0x0102
	CMD_PING       = 0x0105 // client heartbeat (server-bound); ACK: CMD_PONG
	CMD_PONG       = 0x0106

	// 0x03xx group chat
	CMD_GROUP_CREAT     = 0x0301 // create group; ACK: CMD_GROUP_CREAT_ACK
	CMD_GROUP_CREAT_ACK = 0x0302
	CMD_GROUP_JOIN      = 0x0305 // join group; ACK: CMD_GROUP_JOIN_ACK
	CMD_GROUP_JOIN_ACK = 0x0306
	CMD_GROUP_CHAT     = 0x030B // group message; ACK: CMD_GROUP_CHAT_ACK
	CMD_GROUP_CHAT_ACK = 0x030C

	// 0x04xx chatroom
	CMD_ROOM_JOIN     = 0x0405 // join room; ACK: CMD_ROOM_JOIN_ACK
	CMD_ROOM_JOIN_ACK = 0x0406
	CMD_ROOM_CHAT     = 0x040B // room message; ACK: CMD_ROOM_CHAT_ACK
	CMD_ROOM_CHAT_ACK = 0x040C
)
