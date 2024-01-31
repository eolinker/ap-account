package account

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IAccountModule interface {
}

func init() {
	autowire.Auto[IAccountModule](func() reflect.Value {
		return reflect.ValueOf(new(imlAccountModule))
	})
}
