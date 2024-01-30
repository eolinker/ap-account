package user_group

import "context"

type IUserGroupService interface {
	Crete(ctx context.Context, id, name string) error
	Edit(ctx context.Context, id, name string) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*UserGroup, error)
	GetList(ctx context.Context) ([]*UserGroup, error)
}

type IUserGroupMemberService interface {
	AddGroup(ctx context.Context, userID, groupID string) error
	RemoveGroup(ctx context.Context, userID, groupID string) error
	Members(ctx context.Context, gids ...string) ([]*Member, error)
}
