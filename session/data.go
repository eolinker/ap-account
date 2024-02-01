package session

type SessionData struct {
	UID        string `json:"uid"`
	LoginTime  int64  `json:"login_time"`
	ExpireTime int64  `json:"expire_time"`
	Valid      bool   `json:"valid"`
}
