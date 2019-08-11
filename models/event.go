package models

type SnEvent struct {
	EventType   string
	Cid         string
	Vid         string
	Fp          string
	Ip          string
	RemoteIP    string
	UserAgent   string
	User        User
	Ts          int64
	Device      Device
	CookieName  string
	CookieValue string
	params      map[string]string
}
