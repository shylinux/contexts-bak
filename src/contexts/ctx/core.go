package ctx

import (
	"errors"
	"io"
	"strings"
	"time"
	"toolkit"
)

func (c *Context) Register(s *Context, x Server, args ...interface{}) *Context {
	name, force := s.Name, false
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case bool:
			force = arg
		case string:
			name, s.Name = arg, arg
		}
	}

	if c.contexts == nil {
		c.contexts = make(map[string]*Context)
	}
	if x, ok := c.contexts[name]; ok && !force {
		panic(errors.New(c.Name + "上下文中已存在模块:" + x.Name))
	}

	c.contexts[name] = s
	s.context = c
	s.Server = x
	return s
}
func (c *Context) Plugin(s *Context, args []string) string {
	c.Register(s, nil)
	m := &Message{code: 0, time: time.Now(), source: s, target: s, Meta: map[string][]string{}}
	kit.DisableLog = true
	m.Option("log.disable", true)
    m.Option("cli.modal", "action")

	if len(args) == 0 {
		m.Echo("%s: %s\n\n", s.Name, s.Help)
		m.Echo("命令列表:\n")
		for k, v := range s.Commands {
			if !strings.HasPrefix(k, "_") {
				m.Echo("  %s: %s\n    %v\n\n", k, v.Name, v.Help)
			}
		}
		m.Echo("配置列表:\n")
		for k, v := range s.Configs {
			if !strings.HasPrefix(k, "_") {
				m.Echo("--%s(%v): %s\n", k, kit.Formats(v.Value), v.Help)
			}
		}
	} else {
		if Index.Begin(Pulse, args...); Index.Start(Pulse, args...) {
		}
		m.Cmd(args)
	}
    for _, v := range m.Meta["result"] {
        m.Show(v)
    }
	return ""
}
func (c *Context) Spawn(m *Message, name string, help string) *Context {
	s := &Context{Name: name, Help: help, root: c.root, context: c, message: m,
		Caches:   map[string]*Cache{},
		Configs:  map[string]*Config{},
		Commands: map[string]*Command{},
	}

	if m.target = s; c.Server != nil {
		c.Register(s, c.Server.Spawn(m, s, m.Meta["detail"]...))
	} else {
		c.Register(s, nil)
	}
	return s
}
func (c *Context) Begin(m *Message, arg ...string) *Context {
	if len(arg) > 0 {
		m.Set("detail", arg)
	}

	module := c.Name
	if c.context != nil && c.context.Caches != nil && c.context.Caches["module"] != nil {
		module = c.context.Caches["module"].Value + "." + c.Name
	}

	c.Caches["module"] = &Cache{Name: "module", Value: module, Help: "模块域名"}
	c.Caches["status"] = &Cache{Name: "status(begin/start/close)", Value: "begin", Help: "模块状态, begin: 初始完成, start: 正在运行, close: 运行结束"}
	c.Caches["stream"] = &Cache{Name: "stream", Value: "", Help: "模块数据"}

	c.message = m
	c.requests = append(c.requests, m)
	m.source.sessions = append(m.source.sessions, m)
	c.exit = make(chan bool, 3)

	/*
		m.Log("begin", "%d context %v %v", m.Capi("ncontext", 1), m.Meta["detail"], m.Meta["option"])
		for k, x := range c.Configs {
			if x.Hand != nil {
				m.Log("begin", "%s config %v", k, m.Conf(k, x.Value))
			}
		}
	*/

	if c.Server != nil {
		c.Server.Begin(m, m.Meta["detail"]...)
	}
	return c
}
func (c *Context) Start(m *Message, arg ...string) bool {
	sync := false
	if len(arg) > 0 && arg[0] == "sync" {
		sync, arg = true, arg[1:]
	}
	if len(arg) > 0 {
		m.Set("detail", arg)
	}

	c.requests = append(c.requests, m)
	m.source.sessions = append(m.source.sessions, m)

	if m.Hand = true; m.Cap("status") == "start" {
		return true
	}

	m.Gos(m, func(m *Message) {
		m.Log(m.Cap("status", "start"), "%d server %v %v", m.Capi("nserver", 1), m.Meta["detail"], m.Meta["option"])

		c.message = m
		if c.exit <- false; c.Server == nil || c.Server.Start(m, m.Meta["detail"]...) {
			c.Close(m, m.Meta["detail"]...)
			c.exit <- true
		}
	}, func(m *Message) {
		c.Close(m, m.Meta["detail"]...)
		c.exit <- true
	})

	if sync {
		for !<-c.exit {
		}
		return true
	}
	return <-c.exit
}
func (c *Context) Close(m *Message, arg ...string) bool {
	if len(c.requests) == 0 {
		return true
	}

	if m.target == c {
		for i := len(c.requests) - 1; i >= 0; i-- {
			if msg := c.requests[i]; msg.code == m.code {
				if c.Server == nil || c.Server.Close(m, arg...) {
					m.Log("close", "request %d/%d", i, len(c.requests)-1)
					msg.Free()
					for j := i; j < len(c.requests)-1; j++ {
						c.requests[j] = c.requests[j+1]
					}
					c.requests = c.requests[:len(c.requests)-1]
				}
			}
		}
	}

	if len(c.requests) > 0 {
		return false
	}

	if m.Cap("status") == "start" {
		m.Log(m.Cap("status", "close"), "%d server %v", m.root.Capi("nserver", -1), arg)
		for _, msg := range c.sessions {
			if msg.Cap("status") != "close" {
				msg.target.Close(msg, arg...)
			}
		}
	}

	if c.context != nil {
		m.Log("close", "%d context %v", m.root.Capi("ncontext", -1), arg)
		delete(c.context.contexts, c.Name)
		c.exit <- true
	}
	return true
}

func (m *Message) Assert(e interface{}, msg ...string) bool {
	switch v := e.(type) {
	case nil:
		return true
	case *Message:
		if v.Result(0) != "error: " {
			return true
		}
		e = errors.New(strings.Join(v.Meta["result"], ""))
	default:
		if kit.Right(v) {
			return true
		}
	}

	switch e.(type) {
	case error:
	default:
		e = errors.New(kit.Format(msg))
	}

	kit.Log("error", "%v", e)
	panic(e)
}
func (m *Message) TryCatch(msg *Message, safe bool, hand ...func(msg *Message)) *Message {
	defer func() {
		switch e := recover(); e {
		case io.EOF:
		case nil:
		default:
			m.Log("bench", "chain: %s", msg.Format("chain"))
			m.Log("bench", "catch: %s", e)
			m.Log("bench", "stack: %s", msg.Format("stack"))

			if m.Log("error", "catch: %s", e); len(hand) > 1 {
				m.TryCatch(msg, safe, hand[1:]...)
			} else if !safe {
				m.Assert(e)
			}
		}
	}()

	if len(hand) > 0 {
		hand[0](msg)
	}
	return m
}
func (m *Message) Gos(msg *Message, hand ...func(msg *Message)) *Message {
	go func() {
		msg.Option("routine", m.Capi("ngo", 1))
		m.TryCatch(msg, true, hand...)
	}()
	return m
}

func (m *Message) Spawn(arg ...interface{}) *Message {
	temp := false
	c := m.target
	if len(arg) > 0 {
		switch v := arg[0].(type) {
		case *Context:
			c = v
		case *Message:
			c = v.target
		case string:
			temp = kit.Right(v)
		}
	}

	msg := &Message{
		time:    time.Now(),
		code:    m.Capi("nmessage", 1),
		source:  m.target,
		target:  c,
		message: m,
		root:    m.root,
	}

	if temp {
		return msg
	}

	m.messages = append(m.messages, msg)
	return msg
}
func (m *Message) Sess(key string, arg ...interface{}) *Message {
	if key == "" {
		return m.Spawn()
	}

	spawn := true
	if len(arg) > 0 {
		switch v := arg[0].(type) {
		case bool:
			spawn, arg = v, arg[1:]
		}
	}

	if len(arg) > 0 {
		if m.Sessions == nil {
			m.Sessions = make(map[string]*Message)
		}

		switch value := arg[0].(type) {
		case *Message:
			m.Sessions[key] = value
			return m.Sessions[key]
		case *Context:
			m.Sessions[key] = m.Spawn(value)
			return m.Sessions[key]
		case string:
			root := len(arg) < 3 || kit.Right(arg[2])

			method := "find"
			if len(arg) > 1 {
				method = kit.Format(arg[1])
			}

			switch method {
			case "find":
				m.Sessions[key] = m.Find(value, root)
			case "search":
				m.Sessions[key] = m.Search(value, root)[0]
			}
			return m.Sessions[key]
		case nil:
			delete(m.Sessions, key)
			return nil
		}
	}

	for msg := m; msg != nil; msg = msg.message {
		if x, ok := msg.Sessions[key]; ok {
			if spawn {
				x = m.Spawn(x.target)
				x.callback = func(sub *Message) *Message { return sub }
			}
			return x
		}
	}

	return nil
}
func (m *Message) Call(cb func(msg *Message) (sub *Message), arg ...interface{}) *Message {
	if m == nil {
		return m
	}
	if m.callback = cb; len(arg) > 0 || len(m.Meta["detail"]) > 0 {
		m.Log("call", m.Format("detail", "option"))
		m.Cmd(arg...)
	}
	return m
}
func (m *Message) Back(ms ...*Message) *Message {
	if m.callback == nil {
		return m
	}

	if len(ms) == 0 {
		ms = append(ms, m.Spawn(m.source).Copy(m, "append").Copy(m, "result"))
	}

	ns := []*Message{}

	for _, msg := range ms {
		if msg.Hand {
			m.Log("back", msg.Format("ship", "result", "append"))
		} else {
			m.Log("back", msg.Format("ship", "detail", "option"))
		}

		if sub := m.callback(msg); sub != nil && m.message != nil && m.message != m {
			ns = append(ns, sub)
		}
	}

	if len(ns) > 0 {
		m.message.Back(ns...)
	}
	return m
}
func (m *Message) CallBack(sync bool, cb func(msg *Message) (sub *Message), arg ...interface{}) *Message {
	if !sync {
		return m.Call(cb, arg...)
	}

	wait := make(chan *Message, 10)
	// m.Gos(m, func(m *Message) {
	m.Call(func(sub *Message) *Message {
		msg := cb(sub)
		wait <- m
		return msg
	}, arg...)
	// })

	select {
	case <-time.After(kit.Duration(m.Confx("call_timeout"))):
		m.Log("sync", m.Format("timeout", "detail", "option"))
		m.Echo("time out %v", m.Confx("call_timeout"))
	case <-wait:
	}
	return m
}
func (m *Message) Free(cbs ...func(msg *Message) (done bool)) *Message {
	if len(cbs) == 0 {
		for i := len(m.freeback) - 1; i >= 0; i-- {
			m.Log("free", "%d/%d", i, len(m.freeback)-1)
			if !m.freeback[i](m) {
				break
			}
			m.freeback = m.freeback[:i]
		}
		return m
	}

	m.freeback = append(m.freeback, cbs...)
	return m
}