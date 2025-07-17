package store

type AuthDriver struct {
	Id     int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Uuid   string `gorm:"column:uuid;type:varchar(36);not null;uniqueIndex:uuid;comment:登录类型唯一标识"` // 登录类型唯一标识
	Config string `gorm:"column:config;type:text;not null;comment:登录类型配置"`                         // 登录类型配置
	Enable bool   `gorm:"column:enable;type:tinyint(1);not null;default:1;comment:是否启用"`           // 是否启用
}

func (a *AuthDriver) TableName() string {
	return "auth_driver"
}

func (a *AuthDriver) IdValue() int64 {
	return a.Id
}
