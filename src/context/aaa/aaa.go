package aaa // {{{
// }}}
import ( // {{{
	"context"

	"crypto/md5"
	"encoding/hex"
	"math/rand"

	"fmt"
	"strconv"
	"time"
)

// }}}

type AAA struct {
	sessions map[string]*ctx.Context
	*ctx.Context
}

func (aaa *AAA) session(meta string) string { // {{{
	bs := md5.Sum([]byte(fmt.Sprintln("%d%d%s", time.Now().Unix(), rand.Int(), meta)))
	sessid := hex.EncodeToString(bs[:])
	return sessid
}

// }}}

func (aaa *AAA) Spawn(m *ctx.Message, c *ctx.Context, arg ...string) ctx.Server { // {{{
	c.Caches = map[string]*ctx.Cache{}
	c.Configs = map[string]*ctx.Config{}

	s := new(AAA)
	s.Context = c
	return s
}

// }}}
func (aaa *AAA) Begin(m *ctx.Message, arg ...string) ctx.Server { // {{{
	aaa.Caches["group"] = &ctx.Cache{Name: "用户组", Value: "", Help: "用户组"}
	aaa.Caches["username"] = &ctx.Cache{Name: "用户名", Value: "", Help: "用户名"}
	aaa.Caches["password"] = &ctx.Cache{Name: "用户密码", Value: "", Help: "用户密码，加密存储", Hand: func(m *ctx.Message, x *ctx.Cache, arg ...string) string {
		if len(arg) > 0 { // {{{
			bs := md5.Sum([]byte(fmt.Sprintln("用户密码:%s", arg[0])))
			m.Assert(x.Value == "" || x.Value == hex.EncodeToString(bs[:]), "密码错误")
			m.Cap("expire", fmt.Sprintf("%d", time.Now().Unix()+int64(m.Confi("expire"))))
			return hex.EncodeToString(bs[:])
		}
		return x.Value
		// }}}
	}}

	aaa.Caches["sessid"] = &ctx.Cache{Name: "会话标识", Value: "", Help: "用户的会话标识"}
	aaa.Caches["expire"] = &ctx.Cache{Name: "会话超时", Value: "", Help: "用户的会话标识"}
	aaa.Caches["time"] = &ctx.Cache{Name: "登录时间", Value: fmt.Sprintf("%d", time.Now().Unix()), Help: "用户登录时间", Hand: func(m *ctx.Message, x *ctx.Cache, arg ...string) string {
		if len(arg) > 0 { // {{{
			return arg[0]
		}

		n, e := strconv.Atoi(x.Value)
		m.Assert(e)
		return time.Unix(int64(n), 0).Format("15:03:04")
		// }}}
	}}

	aaa.Owner = aaa.Context
	return aaa
}

// }}}
func (aaa *AAA) Start(m *ctx.Message, arg ...string) bool { // {{{
	if len(arg) > 1 {
		if m.Cap("sessid") == "" {
			m.Cap("sessid", aaa.session(arg[1]))
			m.Capi("nuser", 1)
		}
		m.Log("info", m.Source, "login %s %s", m.Cap("group", arg[0]), m.Cap("username", arg[1]))
		m.Cap("stream", m.Cap("username"))
	}

	return false
}

// }}}
func (aaa *AAA) Close(m *ctx.Message, arg ...string) bool { // {{{
	if m.Target == aaa.Context && aaa.Owner == aaa.Context {
		if m.Cap("username") != m.Conf("rootname") {
			m.Log("info", nil, "logout %s", m.Cap("group"), m.Cap("username"))
			m.Capi("nuser", -1)
			return true
		}
	}

	return false
}

// }}}

var Index = &ctx.Context{Name: "aaa", Help: "认证中心",
	Caches: map[string]*ctx.Cache{
		"nuser": &ctx.Cache{Name: "用户数量", Value: "0", Help: "用户数量"},
	},
	Configs: map[string]*ctx.Config{
		"rootname": &ctx.Config{Name: "根用户名", Value: "root", Help: "根用户名"},
		"expire":   &ctx.Config{Name: "会话超时(s)", Value: "120", Help: "会话超时"},
	},
	Commands: map[string]*ctx.Command{
		"login": &ctx.Command{Name: "login [sessid]|[[group] username password]]", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) string {
			m.Target, m.Master = c, c // {{{
			aaa := c.Server.(*AAA)

			switch len(arg) {
			case 0:
				m.Travel(c, func(m *ctx.Message) bool {
					m.Echo("%s(%s): %s\n", m.Target.Name, m.Cap("group"), m.Cap("time"))
					return true
				})
			case 1:
				if s, ok := aaa.sessions[arg[0]]; ok {
					if m.Target = s; int64(m.Capi("expire")) < time.Now().Unix() {
						s.Close(m)
						return ""
					}

					m.Source.Group, m.Source.Owner = m.Cap("group"), m.Target
					m.Log("info", m.Source, "logon %s", m.Cap("group"), m.Cap("username"))
					return m.Cap("username")
				}
			case 2, 3:
				group, username, password := arg[0], arg[0], arg[1]
				if len(arg) == 3 {
					username, password = arg[1], arg[2]
				}

				if username == m.Conf("rootname") {
					m.Set("detail", group, username).Target.Start(m)
				} else if msg := m.Find(username); msg == nil {
					m.Start(username, group, group, username)
				} else {
					m.Target = msg.Target
				}

				m.Cap("password", password)
				m.Source.Group, m.Source.Owner = m.Cap("group"), m.Target
				aaa.sessions[m.Cap("sessid")] = m.Target
				return m.Cap("sessid")
			}
			return ""
			// }}}
		}},
	},
	Index: map[string]*ctx.Context{
		"void": &ctx.Context{Name: "void",
			Commands: map[string]*ctx.Command{"login": &ctx.Command{}},
		},
	},
}

func init() {
	aaa := &AAA{}
	aaa.Context = Index
	ctx.Index.Register(Index, aaa)

	aaa.sessions = make(map[string]*ctx.Context)
}
