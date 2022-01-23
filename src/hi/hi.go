package hi

import (
	"shylinux.com/x/ice"
)

type hi struct {
	ice.Zone

	list string `name:"hi zone id auto insert" help:"hi"`
}

func (h hi) List(m *ice.Message, arg ...string) {
	h.Zone.List(m, arg...)
}

func init() { ice.Cmd("web.code.hi.hi", hi{}) }
