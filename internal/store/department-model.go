package store

import "time"

type Department struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Oid        string    `gorm:"column:oid;type:VARCHAR(36);NOT NULL;comment: 组织id;unique_index:oid;"`
	Name       string    `gorm:"column:name;type:VARCHAR(50);NOT NULL;comment: 组织名称;"`
	ParentOid  string    `gorm:"column:parent_oid;type:VARCHAR(36);NOT NULL;comment: 父组织id;"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;NOT NULL; comment: 创建时间"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;NOT NULL;comment: 更新时间"`
}

func (o *Department) TableName() string {
	return "department"
}

func (o *Department) IdValue() int64 {
	return o.Id
}

type DepartmentMember struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Oid        string    `gorm:"column:oid;type:VARCHAR(36);NOT NULL;comment: 组织id;index:oid;unique_index:oid_uid;"`
	Uid        string    `gorm:"column:uid;type:VARCHAR(36);NOT NULL;comment: 用户id;index:uid;unique_index:oid_uid;"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;NOT NULL; comment: 创建时间"`
}

func (o *DepartmentMember) TableName() string {
	return "department_member"
}

func (o *DepartmentMember) IdValue() int64 {
	return o.Id
}
