package member

import "gitlab.eolink.com/apinto/aoaccount/store/member"

type Member struct {
	Come string
	UID  string
}

func toModel(e *member.Member) *Member {
	return &Member{
		Come: e.Come,
		UID:  e.Uid,
	}
}
func UserID(m *Member) string {
	return m.UID
}
func Cid(m *Member) string {
	return m.Come
}
