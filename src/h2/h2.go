package h2

import (
	"shylinux.com/x/ice"
)

type h2 struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"示例"`
}

func (s h2) List(m *ice.Message, arg ...string) {
	s.Zone.List(m, arg...)
}

func init() { ice.Cmd("web.code.h2.h2", h2{}) }
