package store

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
	"reflect"
	"time"
)

type Role struct {
	Id         int64     `gorm:"type:int(11);size:11;not null;auto_increment;primary_key;column:id;comment:主键ID;"`
	Rid        string    `gorm:"size:36;not null;column:rid;uniqueIndex:rid;comment:角色id;uniqueIndex:rid_uid;"`
	Name       string    `gorm:"size:36;not null;column:name;uniqueIndex:name;comment:角色名称;uniqueIndex:rid_uid;"`
	CreateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_time;comment:创建时间"`
}

func (u *Role) TableName() string {
	return "role"
}
func (u *Role) IdValue() int64 {
	return u.Id
}

type IRoleStore interface {
	store.IBaseStore[Role]
}
type imlRoleStore struct {
	store.BaseStore[Role]
}

func init() {
	autowire.Auto[IRoleStore](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleStore))
	})
}
