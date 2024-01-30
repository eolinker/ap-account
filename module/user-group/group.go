package user_group

import "context"

type IUserGroupModel interface {
	List(ctx context.Context)
}
