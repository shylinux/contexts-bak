package wuhan

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/team"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "wuhan", Help: "武汉加油",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"wuhan": {Name: "wuhan", Help: "武汉", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"wuhan": {Name: "wuhan", Help: "武汉", Meta: kit.Dict(
			"display", "china",
		), Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			data := kit.UnMarshal(m.Cmdx(ice.WEB_SPIDE, "shy", "raw", "GET", "https://view.inews.qq.com/g2/getOnsInfo?name=wuwei_ww_area_counts"))
			m.Option("title", kit.Format("疫情分布\n%s", m.Time()))

			suspect := map[string]int{}
			confirm := map[string]int{}
			dead := map[string]int{}
			deal := map[string]int{}

			kit.Fetch(kit.UnMarshal(kit.Value(data, "data").(string)), func(index int, value map[string]interface{}) {
				area := kit.Select(kit.Format(value["country"]), kit.Format(value["area"]))
				suspect[area] += kit.Int(value["suspect"])
				confirm[area] += kit.Int(value["confirm"])
				dead[area] += kit.Int(value["dead"])
				deal[area] += kit.Int(value["deal"])
			})

			for k, v := range suspect {
				m.Push("name", k)
				m.Push("suspect", v)
				m.Push("confirm", confirm[k])
				m.Push("dead", dead[k])
				m.Push("deal", deal[k])
				m.Push("value", confirm[k])
			}
			m.Sort("confirm", "int_r")
		}},
	},
}

func init() { team.Index.Register(Index, nil) }
