package store

import "time"

type UserAuth struct {
	Id          int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Uid         string    `gorm:"column:uid;type:VARCHAR(36);NOT NULL;comment: 用户id;index:uid;"`
	Driver      string    `gorm:"column:driver;type:VARCHAR(20);NOT NULL;comment: 登录类型;uniqueIndex:only;index:driver;"`
	Identifier  string    `gorm:"column:identifier;type:VARCHAR(255);NOT NULL; comment:手机号 邮箱 用户名或第三方应用的唯一标识;uniqueIndex:only"`
	Certificate string    `gorm:"column:certificate;type:VARCHAR(512);NOT NULL;comment: 密码凭证，站内的保存密码，站外的不保存或保存token;"`
	CreateTime  time.Time `gorm:"column:create_at;type:timestamp; NOT NULL; comment:创建时间;"`
	UpdateTime  time.Time `gorm:"column:update_at;type:timestamp; NOT NULL;comment: 更新时间;default:CURRENT_TIMESTAMP"`
}

func (u *UserAuth) TableName() string {
	return "user_auth"
}

func (u *UserAuth) IdValue() int64 {
	return u.Id
}

type UserInfo struct {
	Id        int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key"`
	Uid       string    `gorm:"column:uid;type:varchar(36);NOT NULL;comment: 用户id;;"`
	Status    uint8     `gorm:"column:status;type:TINYINT(3) UNSIGNED;NOT NULL;default:0;comment: 用户状态 0-unknown 1-active 2-inactive;"`
	Name      string    `gorm:"column:name;type:VARCHAR(50);NOT NULL;comment: 用户名;"`
	Gender    uint8     `gorm:"column:gender;type:TINYINT(2) UNSIGNED;NOT NULL;default:0;comment: 性别 0-unknown 1-male 2 female;"`
	Mobile    string    `gorm:"column:mobile;type:VARCHAR(16);NOT NULL;comment: 手机号;"`
	Email     string    `gorm:"column:email;type:VARCHAR(100);NOT NULL;comment: 邮箱;"`
	CreateAt  time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_at;comment:创建时间"`
	UpdateAt  time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:update_at;comment:修改时间" json:"update_at"`
	PushToken string    `gorm:"column:push_token;type:VARCHAR(50);NOT NULL;comment: 推送token;"`
	IsDeleted bool      `gorm:"type:tinyint(1);not null;column:is_delete;comment:是否删除"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

func (u *UserInfo) IdValue() int64 {
	return u.Id
}

type UserLoginLog struct {
	Id      int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Uid     string    `gorm:"column:uid;type:type:VARCHAR(36);NOT NULL;comment: 用户id;index:uid;"`
	Driver  string    `gorm:"column:driver;type:varchar(10); NOT NULL;comment: 登录类型;index:driver;"`
	Command uint8     `gorm:"column:command;type:TINYINT(3) UNSIGNED;NOT NULL;comment: 操作类型 1登陆成功  2登出成功 3登录失败 4登出失败"`
	Lastip  string    `gorm:"column:lastip;type:VARCHAR(32);NOT NULL;comment: 最后登录ip"`
	Text    string    `gorm:"column:text;type:VARCHAR(200);NOT NULL;charset=utf8mb4;comment: 操作内容"`
	Time    time.Time `gorm:"column:time;type:timestamp;NOT NULL;comment: 操作时间"`
}

func (u *UserLoginLog) TableName() string {
	return "user_login_log"
}

func (u *UserLoginLog) IdValue() int64 {
	return u.Id
}

type UserRegisterLog struct {
	Id             int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Uid            string    `gorm:"column:uid;type:type:VARCHAR(36);NOT NULL;comment: 用户id;uniqueIndex:uid;"`
	Driver         string    `gorm:"column:driver;type:varchar(10); NOT NULL;comment: 登录类型;index:driver;"`
	RegisterTime   time.Time `gorm:"column:register_time;type:timestamp;NOT NULL;comment: 注册时间"`
	RegisterIp     string    `gorm:"column:register_ip;type:VARCHAR(16);NOT NULL;comment: 注册ip"`
	RegisterClient string    `gorm:"column:operator;type:VARCHAR(32);NOT NULL;comment: 注册客户端;"`
}

func (u *UserRegisterLog) TableName() string {
	return "user_register_log"
}

func (u *UserRegisterLog) IdValue() int64 {
	return u.Id
}

type UserInfoUpdateLog struct {
	Id              int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Uid             string    `gorm:"column:uid;type:type:VARCHAR(36);NOT NULL;comment: 用户id;index:uid;"`
	AttributeName   string    `gorm:"column:attribute_name;type:VARCHAR(30);NOT NULL;comment: 属性名"`
	AttributeOldVal string    `gorm:"column:attribute_old_val;type:VARCHAR(30);NOT NULL;comment: 属性旧值"`
	AttributeNewVal string    `gorm:"column:attribute_new_val;type:VARCHAR(30);NOT NULL;comment: 属性新值"`
	Operator        string    `gorm:"column:operator;type:VARCHAR(36);NOT NULL;comment: 操作人"`
	UpdateTime      time.Time `gorm:"column:update_at;type:timestamp;NOT NULL;comment: 更新时间"`
}

func (u *UserInfoUpdateLog) TableName() string {
	return "user_info_update_log"
}

func (u *UserInfoUpdateLog) IdValue() int64 {
	return u.Id
}
