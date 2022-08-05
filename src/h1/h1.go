package h1

import (
	"shylinux.com/x/ice"
)

type h1 struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"示例"`
}

func (s h1) List(m *ice.Message, arg ...string) {
	s.Zone.List(m, arg...)
}

func init() { ice.Cmd("web.code.h1.h1", h1{}) }
