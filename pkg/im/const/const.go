package imconst

// TTL and batch constants migrated from beehive lib/comm/const.go.

const (
	CHAT_SID_TTL            = 300
	CHAT_OP_TTL             = 30
	CHAT_NID_TTL            = 30
	CHAT_BAT_NUM            = 1000
	CHAT_ROOM_GROUP_MAX_NUM = 10000
)

const (
	TIME_MIN  = 60
	TIME_HOUR = 3600
	TIME_DAY  = 86400
	TIME_WEEK = 7 * 86400
	TIME_YEAR = 365 * TIME_DAY
)

const (
	NOLOCK = 0
	RDLOCK = 1
	WRLOCK = 2
)

const (
	LSND_TYPE_UNKNOWN = 0
	LSND_TYPE_TCP     = 1
	LSND_TYPE_WS      = 2
)

const (
	PROC_STATUS_EXIT = 0
	PROC_STATUS_EXEC = 1
	PROC_STATUS_BUSY = 2
)

const (
	SECTION_SID_NUM = 100000
)
