package h1

import (
	"shylinux.com/x/ice"
)

type h1 struct {
	ice.Hash

	list string `name:"list hash id auto create" help:"数据"`
}

func (h h1) List(m *ice.Message, arg ...string) {
	h.Hash.List(m, arg...)
}

func init() { ice.Cmd("web.code.h1.h1.go", h1{}) }

