package store

import "time"

type UserGroup struct {
	Id         int64     `gorm:"type:BIGINT(20);size:20;not null;auto_increment;primary_key;column:id;comment:主键ID;"`
	Gid        string    `gorm:"size:36;not null;column:gid;uniqueIndex:gid;comment:分组id;uniqueIndex:gid_uid;"`
	Name       string    `gorm:"size:36;not null;column:name;uniqueIndex:name;comment:分组名称;uniqueIndex:gid_uid;"`
	CreateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_time;comment:创建时间"`
}

func (u *UserGroup) TableName() string {
	return "user_group"
}
func (u *UserGroup) IdValue() int64 {
	return u.Id
}

type UserGroupMember struct {
	Id         int64     `gorm:"type:BIGINT(20);size:20;not null;auto_increment;primary_key;column:id;comment:主键ID;"`
	Gid        string    `gorm:"size:36;not null;column:gid;uniqueIndex:gid;comment:分组id;uniqueIndex:gid_uid;"`
	Uid        string    `gorm:"size:36;not null;column:uid;uniqueIndex:uid;comment:用户id;uniqueIndex:gid_uid;"`
	CreateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_time;comment:创建时间"`
}

func (u *UserGroupMember) TableName() string {
	return "user_group_member"
}
func (u *UserGroupMember) IdValue() int64 {
	return u.Id
}
