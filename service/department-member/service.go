package department_member

import (
	"github.com/eolinker/ap-account/service/member"
	"github.com/eolinker/ap-account/store"
	"github.com/eolinker/go-common/autowire"
	"reflect"
)

type IMemberService member.IMemberService

func init() {
	autowire.Auto[IMemberService](func() reflect.Value {
		return reflect.ValueOf(new(member.Service[store.DepartmentMemberStore]))
	})
}
