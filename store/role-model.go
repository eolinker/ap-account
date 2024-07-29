package store

import "time"

type Role struct {
	Id          int64     `gorm:"column:id;type:int(11);not null;comment:id;primary_key;comment:主键ID;"`
	UUID        string    `gorm:"type:varchar(36);not null;column:uuid;uniqueIndex:uuid;comment:UUID;"`
	Name        string    `gorm:"type:varchar(100);not null;column:name;comment:角色名称"`
	Group       string    `gorm:"type:varchar(100);not null;column:group;comment:角色分组"`
	Description string    `gorm:"size:255;not null;column:description;comment:描述"`
	Permit      []string  `gorm:"type:text;not null;column:permit;comment:权限列表;serializer:json"`
	CreateAt    time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_at;comment:创建时间"`
	UpdateAt    time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:update_at;comment:修改时间"`
	Default     bool      `gorm:"type:tinyint(1);not null;column:default;comment:是否默认角色"`
}

func (r *Role) IdValue() int64 {
	return r.Id
}

func (r *Role) TableName() string {
	return "role"
}

type RoleMember struct {
	Id     int64  `gorm:"column:id;type:int(11);not null;comment:id;primary_key;comment:主键ID;"`
	Role   string `gorm:"type:varchar(100);not null;column:role;comment:角色ID"`
	User   string `gorm:"type:varchar(100);not null;column:user;comment:用户ID"`
	Target string `gorm:"type:varchar(100);not null;column:target;comment:目标ID"`
}

func (r *RoleMember) TableName() string {
	return "role_member"
}

func (r *RoleMember) IdValue() int64 {
	return r.Id
}
