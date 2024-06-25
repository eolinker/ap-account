package member

import (
	"time"

	"github.com/eolinker/ap-account/store/member"
)

type Member struct {
	Come       string
	UID        string
	CreateTime time.Time
}

func toModel(e *member.Member) *Member {
	return &Member{
		Come:       e.Come,
		UID:        e.Uid,
		CreateTime: e.CreateTime,
	}
}
func UserID(m *Member) string {
	return m.UID
}
func Cid(m *Member) string {
	return m.Come
}
