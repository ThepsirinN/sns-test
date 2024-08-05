package constant

const (
	CONFIG_PATH = "./env/"
	SECRET_PATH = "./env/"
	CONFIG_FILE = "config.env"
	SECRET_FILE = "secret.env"
)

const (
	ENV_LOCAL      = "local"
	ENV_DEVELOP    = "develop"
	ENV_PRODUCTION = "production"
)

const (
	TRACE_ID_KEY = "traceId"
	SPAN_ID_KEY  = "spanId"
)

const (
	FRIEND_STATUS_PENDING                         = 1
	FRIEND_STATUS_SUCCESS                         = 2
	FRIEND_STATUS_PENDING_APPROVE_MESSAGE         = "รอการอนุมัติจากฝั่งตรงข้าม"
	FRIEND_STATUS_PENDING_WAITING_APPROVE_MESSAGE = "รอการตอบรับเป็นเพื่อน"
	CACHE_FRIEND_REQUEST_KEY                      = "FRIEND_REQUEST::"
	CACHE_FRIEND_LIST_KEY                         = "FRIEND_LIST::"
)
