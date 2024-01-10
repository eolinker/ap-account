package store

import "gitlab.eolink.com/apinto/common/store"

type DepartmentStore interface {
	store.IBaseStore[Department]
}
type DepartmentMemberStore interface {
	store.IBaseStore[DepartmentMember]
}
