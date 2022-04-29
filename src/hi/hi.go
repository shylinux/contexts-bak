package hi

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
)

type hi struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"示例"`
}

func (h hi) Command(m *ice.Message, arg ...string) {
	if !m.PodCmd(ctx.COMMAND, arg) {
		m.Cmdy(ctx.COMMAND, arg)
	}
}
func (h hi) Run(m *ice.Message, arg ...string) {
	if m.Right(arg) && !m.PodCmd(arg) {
		m.Debug("what %v", arg)
		m.Cmdy(arg)
	}
}
func (h hi) List(m *ice.Message, arg ...string) {
	h.Zone.List(m, arg...)
}

func init() { ice.Cmd("web.code.hi.hi", hi{}) }
