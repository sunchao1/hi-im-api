package httpx

// Response is the standard HTTP JSON envelope (beehive HttpResp).
type Response struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"errmsg"`
}
