package store

import (
	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/store"
	"reflect"
	"time"
)

type Role struct {
	Id         int64     `gorm:"type:BIGINT(20);size:20;not null;auto_increment;primary_key;column:id;comment:主键ID;"`
	UUID       string    `gorm:"size:36;not null;column:uuid;uniqueIndex:uuid;comment:角色id;"`
	Name       string    `gorm:"size:255;not null;column:name;uniqueIndex:name;comment:角色名称;"`
	Creator    string    `gorm:"type:varchar(36);column:creator;comment:创建者"`
	CreateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_at;comment:创建时间"`
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
	store.Store[Role]
}

func init() {
	autowire.Auto[IRoleStore](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleStore))
	})
}
