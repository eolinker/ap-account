package store

import "gitlab.eolink.com/apinto/common/store"

type OrganizationStore interface {
	store.IBaseStore[Department]
}
type OrganizationUserStore interface {
	store.IBaseStore[DepartmentUser]
}
