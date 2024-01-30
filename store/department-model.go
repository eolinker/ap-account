package store

import "time"

type Department struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	UUID       string    `gorm:"column:uuid;type:VARCHAR(36);NOT NULL;comment: 组织id;uniqueIndex:uuid;"`
	Name       string    `gorm:"column:name;type:VARCHAR(50);NOT NULL;comment: 组织名称;"`
	Parent     string    `gorm:"column:parent;type:VARCHAR(36);NOT NULL;comment: 父组织id;"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;NOT NULL; comment: 创建时间"`
}

func (o *Department) TableName() string {
	return "department"
}

func (o *Department) IdValue() int64 {
	return o.Id
}

type DepartmentMember struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Department string    `gorm:"column:department;type:VARCHAR(36);NOT NULL;comment: 组织id;index:department;uniqueIndex:department_uid;"`
	Uid        string    `gorm:"column:uid;type:VARCHAR(36);NOT NULL;comment: 用户id;index:uid;uniqueIndex:department_uid;"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;NOT NULL; comment: 创建时间"`
}

func (o *DepartmentMember) TableName() string {
	return "department_member"
}

func (o *DepartmentMember) IdValue() int64 {
	return o.Id
}

type UserMember struct {
	Uid        string `json:"uid"`
	Department string `json:"department"`
}