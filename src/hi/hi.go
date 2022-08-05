package hi

import (
	"shylinux.com/x/ice"
)

type hi struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"示例"`
}

func (s hi) Show(m *ice.Message, arg ...string) {
	m.Echo("hello world")
	m.StatusTime()
}

func (s hi) List(m *ice.Message, arg ...string) {
	s.Zone.ListPage(m, arg...)
	m.Echo("hello world")
}

func init() { ice.Cmd("web.code.hi.hi", hi{}) }
