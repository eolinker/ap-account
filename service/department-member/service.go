package department_member

import (
	"context"
	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IMemberService interface {
	Members(ctx context.Context) ([]*Member, error)
	AddMembers(ctx context.Context, departmentId string, userIds ...string) error
	RemoveMembers(ctx context.Context, departmentId string, userIds ...string) error
	usage.IUserUsageService
}

func init() {
	autowire.Auto[IMemberService](func() reflect.Value {
		return reflect.ValueOf(new(imlMemberService))
	})
}
