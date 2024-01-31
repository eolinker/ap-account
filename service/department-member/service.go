package department_member

import (
	"gitlab.eolink.com/apinto/aoaccount/service/member"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IMemberService member.IMemberService

func init() {
	autowire.Auto[IMemberService](func() reflect.Value {
		return reflect.ValueOf(new(member.Service[store.DepartmentMemberStore]))
	})
}
