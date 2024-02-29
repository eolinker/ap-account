package member

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gitlab.eolink.com/apinto/common/store"
	"gorm.io/gorm/clause"
)

var (
	_ IMemberStore = (*Store)(nil)
)

type Member struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL;comment:id;primary_key;"`
	Come       string    `gorm:"column:come;type:VARCHAR(36);NOT NULL;comment: 归属id;index:department;uniqueIndex:department_uid;"`
	Uid        string    `gorm:"column:uid;type:VARCHAR(36);NOT NULL;comment: 用户id;index:uid;uniqueIndex:department_uid;"`
	CreateTime time.Time `gorm:"column:create_at;type:timestamp;NOT NULL; comment: 创建时间"`
}

type IMemberStore interface {
	AddMember(ctx context.Context, cid string, uids ...string) error
	RemoveMember(ctx context.Context, cid string, uids ...string) error
	Members(ctx context.Context, cids []string, users []string) ([]*Member, error)
	Delete(ctx context.Context, cid ...string) error
	RemoveUser(ctx context.Context, uid ...string) error
}

type Store struct {
	db       store.IDB `autowired:""`
	name     string
	joins    string
	conflict clause.OnConflict
}

func (s *Store) RemoveUser(ctx context.Context, uid ...string) error {
	return s.db.DB(ctx).Table(s.name).Where("uid in (?)", uid).Delete(s.name).Error
}

func (s *Store) Delete(ctx context.Context, cid ...string) error {
	return s.db.DB(ctx).Table(s.name).Where("come in (?)", cid).Delete(s.name).Error
}

func (s *Store) AddMember(ctx context.Context, cid string, uids ...string) error {
	nt := time.Now()
	ms := make([]*Member, 0, len(uids))
	for _, uid := range uids {
		ms = append(ms, &Member{
			Id:         0,
			Come:       cid,
			Uid:        uid,
			CreateTime: nt,
		})
	}
	return s.db.DB(ctx).Table(s.name).Clauses(s.conflict).Create(ms).Error
}

func (s *Store) RemoveMember(ctx context.Context, cid string, uids ...string) error {
	return s.db.DB(ctx).Table(s.name).Where("come = ? and uid in (?)", cid, uids).Delete(s.name).Error
}

func (s *Store) OnComplete() {

	err := s.db.DB(context.Background()).Table(s.name).AutoMigrate(&Member{})
	if err != nil {
		panic(err)
	}

}
func (s *Store) Members(ctx context.Context, comes []string, users []string) ([]*Member, error) {
	result := make([]*Member, 0)
	where := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	if len(comes) > 0 {
		where = append(where, "member.come !='' and member.come in (?)")
		args = append(args, comes)
	}
	if len(users) > 0 {
		where = append(where, "user_info.uid in (?)")
		args = append(args, users)
	}
	tdb := s.db.DB(ctx).Table("user_info").Select("user_info.uid as uid,member.come as come,member.create_at as create_at").
		Joins(s.joins)

	if len(where) > 0 {
		tdb = tdb.Where(strings.Join(where, " and "), args...)
	}
	err := tdb.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewMemberStore(name string) *Store {
	tableName := fmt.Sprintf("%s_member", name)
	s := &Store{
		name:  tableName,
		joins: fmt.Sprintf("inner join %s as member on member.uid = user_info.uid and user_info.is_delete = 0", tableName),
		conflict: clause.OnConflict{
			Columns:   []clause.Column{{Name: "uid"}, {Name: "come"}},
			DoUpdates: clause.AssignmentColumns([]string{}),
		},
	}
	return s
}
