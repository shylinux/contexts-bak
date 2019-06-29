package ctx

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"errors"
	"io"
	"sort"
	"time"
	"toolkit"
)

func (c *Context) Register(s *Context, x Server, args ...interface{}) {
	force := false
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case bool:
			force = arg
		}
	}

	if c.contexts == nil {
		c.contexts = make(map[string]*Context)
	}
	if x, ok := c.contexts[s.Name]; ok && !force {
		panic(errors.New(c.Name + "上下文中已存在模块:" + x.Name))
	}

	c.contexts[s.Name] = s
	s.context = c
	s.Server = x
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

	m.GoFunc(m, func(m *Message) {
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

func (c *Context) Context() *Context {
	return c.context
}
func (c *Context) Message() *Message {
	return c.message
}
func (c *Context) Has(key ...string) bool {
	switch len(key) {
	case 2:
		if _, ok := c.Commands[key[0]]; ok && key[1] == "command" {
			return true
		}
		if _, ok := c.Configs[key[0]]; ok && key[1] == "config" {
			return true
		}
		if _, ok := c.Caches[key[0]]; ok && key[1] == "cache" {
			return true
		}
	case 1:
		if _, ok := c.Commands[key[0]]; ok {
			return true
		}
		if _, ok := c.Configs[key[0]]; ok {
			return true
		}
		if _, ok := c.Caches[key[0]]; ok {
			return true
		}
	}
	return false
}
func (c *Context) Sub(key string) *Context {
	return c.contexts[key]
}
func (c *Context) Travel(m *Message, hand func(m *Message, n int) (stop bool)) *Context {
	if c == nil {
		return nil
	}
	target := m.target

	cs := []*Context{c}
	for i := 0; i < len(cs); i++ {
		if m.target = cs[i]; hand(m, i) {
			return cs[i]
		}

		keys := []string{}
		for k, _ := range cs[i].contexts {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			cs = append(cs, cs[i].contexts[k])
		}
	}

	m.target = target
	return target
}
func (c *Context) BackTrace(m *Message, hand func(m *Message) (stop bool)) *Context {
	target := m.target

	for s := m.target; s != nil; s = s.context {
		if m.target = s; hand(m) {
			return s
		}
	}

	m.target = target
	return target
}

func (c *Context) Plugin(args []string) string {
	Index.Register(c, nil)
	m := &Message{code: 0, time: time.Now(), source: c, target: c, Meta: map[string][]string{}}
	if len(args) == 0 {
		m.Echo("%s: %s\n", c.Name, c.Help)
		for k, v := range c.Commands {
			m.Echo("%s: %s %v\n", k, v.Name, v.Help)
		}
	} else if cs, ok := c.Commands[args[0]]; ok {
		h := cs.Hand
		if e := h(m, c, args[0], args[1:]...); e != nil {
			m.Echo("error: ").Echo("%v\n", e)
		}
	} else {
		m.Echo("error: ").Echo("not found: %v\n", args[0])
	}
	return strings.Join(m.Meta["result"], "")
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
func (m *Message) Time(arg ...interface{}) string {
	t := m.time

	if len(arg) > 0 {
		if d, e := time.ParseDuration(arg[0].(string)); e == nil {
			arg = arg[1:]
			t.Add(d)
		}
	}

	str := m.Conf("time_format")
	if len(arg) > 1 {
		str = fmt.Sprintf(arg[0].(string), arg[1:]...)
	} else if len(arg) > 0 {
		str = fmt.Sprintf("%v", arg[0])
	}
	if str == "stamp" {
		return kit.Format(t.Unix())
	}
	return t.Format(str)
}
func (m *Message) Code() int {
	return m.code
}
func (m *Message) Source() *Context {
	return m.source
}
func (m *Message) Target() *Context {
	return m.target
}
func (m *Message) Message() *Message {
	return m.message
}
func (m *Message) Format(arg ...interface{}) string {
	if len(arg) == 0 {
		arg = append(arg, "time", "ship")
	}

	meta := []string{}
	for _, v := range arg {
		switch kit.Format(v) {
		case "summary":
			msg := arg[1].(*Message)
			ms := make([]*Message, 0, 1024)
			ms = append(ms, msg.message, msg)

			for i := 0; i < len(ms); i++ {
				msg := ms[i]
				if m.Add("append", "index", i); msg == nil {
					m.Add("append", "message", "")
					m.Add("append", "time", "")
					m.Add("append", "code", "")
					m.Add("append", "source", "")
					m.Add("append", "target", "")
					m.Add("append", "details", "")
					m.Add("append", "options", "")
					continue
				}

				if msg.message != nil {
					m.Add("append", "message", msg.message.code)
				} else {
					m.Add("append", "message", "")
				}
				m.Add("append", "time", msg.time.Format("15:04:05"))
				m.Add("append", "code", msg.code)
				m.Add("append", "source", msg.source.Name)
				m.Add("append", "target", msg.target.Name)
				m.Add("append", "details", fmt.Sprintf("%v", msg.Meta["detail"]))
				m.Add("append", "options", fmt.Sprintf("%v", msg.Meta["option"]))

				if i == 0 {
					continue
				}

				if len(ms) < 30 && len(arg) > 2 && arg[2] == "deep" {
					ms = append(ms, ms[i].messages...)
				}
			}
			m.Table()
		case "time":
			meta = append(meta, m.Time())
		case "code":
			meta = append(meta, kit.Format(m.code))
		case "ship":
			meta = append(meta, fmt.Sprintf("%s:%d(%s->%s)", m.Option("routine"), m.code, m.source.Name, m.target.Name))
		case "source":
			target := m.target
			m.target = m.source
			meta = append(meta, m.Cap("module"))
			m.target = target
		case "target":
			meta = append(meta, m.Cap("module"))

		case "detail":
			meta = append(meta, fmt.Sprintf("%v", m.Meta["detail"]))
		case "option":
			meta = append(meta, fmt.Sprintf("%v", m.Meta["option"]))
		case "append":
			meta = append(meta, fmt.Sprintf("%v", m.Meta["append"]))
		case "result":
			meta = append(meta, fmt.Sprintf("%v", m.Meta["result"]))

		case "full":
		case "chain":
			ms := []*Message{}
			if v == "full" {
				ms = append(ms, m)
			} else {
				for msg := m; msg != nil; msg = msg.message {
					ms = append(ms, msg)
				}
			}

			meta = append(meta, "\n")
			for i := len(ms) - 1; i >= 0; i-- {
				msg := ms[i]

				meta = append(meta, fmt.Sprintf("%s\n", msg.Format("time", "ship")))
				if len(msg.Meta["detail"]) > 0 {
					meta = append(meta, fmt.Sprintf("  detail: %d %v\n", len(msg.Meta["detail"]), msg.Meta["detail"]))
				}
				if len(msg.Meta["option"]) > 0 {
					meta = append(meta, fmt.Sprintf("  option: %d %v\n", len(msg.Meta["option"]), msg.Meta["option"]))
					for _, k := range msg.Meta["option"] {
						if v, ok := msg.Data[k]; ok {
							meta = append(meta, fmt.Sprintf("    %s: %v\n", k, kit.Format(v)))
						} else if v, ok := msg.Meta[k]; ok {
							meta = append(meta, fmt.Sprintf("    %s: %d %v\n", k, len(v), v))
						}
					}
				}
				if len(msg.Meta["append"]) > 0 {
					meta = append(meta, fmt.Sprintf("  append: %d %v\n", len(msg.Meta["append"]), msg.Meta["append"]))
					for _, k := range msg.Meta["append"] {
						if v, ok := msg.Data[k]; ok {
							meta = append(meta, fmt.Sprintf("    %s: %v\n", k, kit.Format(v)))
						} else if v, ok := msg.Meta[k]; ok {
							meta = append(meta, fmt.Sprintf("    %s: %d %v\n", k, len(v), v))
						}
					}
				}
				if len(msg.Meta["result"]) > 0 {
					meta = append(meta, fmt.Sprintf("  result: %d %v\n", len(msg.Meta["result"]), msg.Meta["result"]))
				}
			}
		case "stack":
			pc := make([]uintptr, 100)
			pc = pc[:runtime.Callers(6, pc)]
			frames := runtime.CallersFrames(pc)

			for {
				frame, more := frames.Next()
				file := strings.Split(frame.File, "/")
				name := strings.Split(frame.Function, "/")
				meta = append(meta, fmt.Sprintf("\n%s:%d\t%s", file[len(file)-1], frame.Line, name[len(name)-1]))
				if !more {
					break
				}
			}

		default:
			meta = append(meta, kit.FileName(kit.Format(v), "time"))
		}
	}
	return strings.Join(meta, " ")
}
func (m *Message) Tree(code int) *Message {
	ms := []*Message{m}
	for i := 0; i < len(ms); i++ {
		if ms[i].Code() == code {
			return ms[i]
		}
		ms = append(ms, ms[i].messages...)
	}
	return nil
}

func (m *Message) Add(meta string, key string, value ...interface{}) *Message {
	if m.Meta == nil {
		m.Meta = make(map[string][]string)
	}
	if _, ok := m.Meta[meta]; !ok {
		m.Meta[meta] = make([]string, 0, 3)
	}

	switch meta {
	case "detail", "result":
		m.Meta[meta] = append(m.Meta[meta], key)
		m.Meta[meta] = append(m.Meta[meta], kit.Trans(value...)...)

	case "option", "append":
		if _, ok := m.Meta[key]; !ok {
			m.Meta[key] = make([]string, 0, 3)
		}
		m.Meta[key] = append(m.Meta[key], kit.Trans(value...)...)

		for _, v := range m.Meta[meta] {
			if v == key {
				return m
			}
		}
		m.Meta[meta] = append(m.Meta[meta], key)

	default:
		m.Log("error", "add meta error %s %s %v", meta, key, value)
	}

	return m
}
func (m *Message) Set(meta string, arg ...interface{}) *Message {
	switch meta {
	case "detail", "result":
		if m != nil && m.Meta != nil {
			delete(m.Meta, meta)
		}
	case "option", "append":
		if len(arg) > 0 {
			delete(m.Meta, kit.Format(arg[0]))
		} else {
			for _, k := range m.Meta[meta] {
				delete(m.Data, k)
				delete(m.Meta, k)
			}
			delete(m.Meta, meta)
		}
	default:
		m.Log("error", "set meta error %s %s %v", meta, arg)
	}

	if args := kit.Trans(arg...); len(args) > 0 {
		m.Add(meta, args[0], args[1:])
	}
	return m
}
func (m *Message) Put(meta string, key string, value interface{}) *Message {
	switch meta {
	case "option", "append":
		if m.Set(meta, key); m.Data == nil {
			m.Data = make(map[string]interface{})
		}
		m.Data[key] = value

	default:
		m.Log("error", "put data error %s %s %v", meta, key, value)
	}
	return m
}
func (m *Message) Get(key string, arg ...interface{}) string {
	if meta, ok := m.Meta[key]; ok && len(meta) > 0 {
		index := 0
		if len(arg) > 0 {
			index = kit.Int(arg[0])
		}

		index = (index+2)%(len(meta)+2) - 2
		if index >= 0 && index < len(meta) {
			return meta[index]
		}
	}
	return ""
}
func (m *Message) Has(key ...string) bool {
	switch len(key) {
	case 1:
		if _, ok := m.Data[key[0]]; ok {
			return true
		}
		if _, ok := m.Meta[key[0]]; ok {
			return true
		}
	}
	return false
}
func (m *Message) CopyTo(msg *Message, arg ...string) *Message {
	msg.Copy(m, "append").Copy(m, "result")
	return m
}
func (m *Message) Copy(msg *Message, arg ...string) *Message {
	if msg == nil || m == msg {
		return m
	}

	for i := 0; i < len(arg); i++ {
		meta := arg[i]

		switch meta {
		case "target":
			m.target = msg.target
		case "callback":
			m.callback = msg.callback
		case "detail", "result":
			if len(msg.Meta[meta]) > 0 {
				m.Add(meta, msg.Meta[meta][0], msg.Meta[meta][1:])
			}
		case "option", "append":
			if msg.Meta == nil {
				msg.Meta = map[string][]string{}
			}
			if msg.Meta[meta] == nil {
				break
			}
			if i == len(arg)-1 {
				arg = append(arg, msg.Meta[meta]...)
			}

			for i++; i < len(arg); i++ {
				if v, ok := msg.Data[arg[i]]; ok {
					m.Put(meta, arg[i], v)
				} else if v, ok := msg.Meta[arg[i]]; ok {
					m.Set(meta, arg[i], v) // TODO fuck Add
				}
			}
		default:
			if msg.Hand {
				meta = "append"
			} else {
				meta = "option"
			}

			if v, ok := msg.Data[arg[i]]; ok {
				m.Put(meta, arg[i], v)
			}
			if v, ok := msg.Meta[arg[i]]; ok {
				m.Add(meta, arg[i], v)
			}
		}
	}

	return m
}
func (m *Message) CopyFuck(msg *Message, arg ...string) *Message {
	if m == msg {
		return m
	}

	for i := 0; i < len(arg); i++ {
		meta := arg[i]

		switch meta {
		case "target":
			m.target = msg.target
		case "callback":
			m.callback = msg.callback
		case "detail", "result":
			if len(msg.Meta[meta]) > 0 {
				m.Add(meta, msg.Meta[meta][0], msg.Meta[meta][1:])
			}
		case "option", "append":
			if msg.Meta == nil {
				msg.Meta = map[string][]string{}
			}
			if msg.Meta[meta] == nil {
				break
			}
			if i == len(arg)-1 {
				arg = append(arg, msg.Meta[meta]...)
			}

			for i++; i < len(arg); i++ {
				if v, ok := msg.Data[arg[i]]; ok {
					m.Put(meta, arg[i], v)
				} else if v, ok := msg.Meta[arg[i]]; ok {
					m.Add(meta, arg[i], v) // TODO fuck Add
				}
			}
		default:
			if msg.Hand {
				meta = "append"
			} else {
				meta = "option"
			}

			if v, ok := msg.Data[arg[i]]; ok {
				m.Put(meta, arg[i], v)
			}
			if v, ok := msg.Meta[arg[i]]; ok {
				m.Add(meta, arg[i], v)
			}
		}
	}

	return m
}
func (m *Message) Echo(str string, arg ...interface{}) *Message {
	if len(arg) > 0 {
		return m.Add("result", fmt.Sprintf(str, arg...))
	}
	return m.Add("result", str)
}
func (m *Message) Auto(arg ...string) *Message {
	for i := 0; i < len(arg); i += 3 {
		m.Add("append", "value", arg[i])
		m.Add("append", "name", arg[i+1])
		m.Add("append", "help", arg[i+2])
	}
	return m
}

func (m *Message) Insert(meta string, index int, arg ...interface{}) string {
	if m.Meta == nil {
		m.Meta = make(map[string][]string)
	}
	m.Meta[meta] = kit.Array(m.Meta[meta], index, arg)

	if -1 < index && index < len(m.Meta[meta]) {
		return m.Meta[meta][index]
	}
	return ""
}
func (m *Message) Detail(arg ...interface{}) string {
	noset, index := true, 0
	if len(arg) > 0 {
		switch v := arg[0].(type) {
		case int:
			noset, index, arg = false, v, arg[1:]
		}
	}
	if noset && len(arg) > 0 {
		index = -2
	}

	return m.Insert("detail", index, arg...)
}
func (m *Message) Detaili(arg ...interface{}) int {
	return kit.Int(m.Detail(arg...))
}
func (m *Message) Details(arg ...interface{}) bool {
	return kit.Right(m.Detail(arg...))
}
func (m *Message) Result(arg ...interface{}) string {
	noset, index := true, 0
	if len(arg) > 0 {
		switch v := arg[0].(type) {
		case int:
			noset, index, arg = false, v, arg[1:]
		}
	}
	if noset && len(arg) > 0 {
		index = -2
	}

	return m.Insert("result", index, arg...)
}
func (m *Message) Resulti(arg ...interface{}) int {
	return kit.Int(m.Result(arg...))
}
func (m *Message) Results(arg ...interface{}) bool {
	return kit.Right(m.Result(arg...))
}
func (m *Message) Option(key string, arg ...interface{}) string {
	if len(arg) > 0 {
		m.Insert(key, 0, arg...)
		if _, ok := m.Meta[key]; ok {
			m.Add("option", key)
		}
	}

	for msg := m; msg != nil; msg = msg.message {
		if !msg.Has(key) {
			continue
		}
		for _, k := range msg.Meta["option"] {
			if k == key {
				return msg.Get(key)
			}
		}
	}
	return ""
}
func (m *Message) Optioni(key string, arg ...interface{}) int {
	return kit.Int(m.Option(key, arg...))

}
func (m *Message) Options(key string, arg ...interface{}) bool {
	return kit.Right(m.Option(key, arg...))
}
func (m *Message) Optionv(key string, arg ...interface{}) interface{} {
	if len(arg) > 0 {
		switch arg[0].(type) {
		case nil:
		// case []string:
		// 	m.Option(key, v...)
		// case string:
		// 	m.Option(key, v)
		default:
			m.Put("option", key, arg[0])
		}
	}

	for msg := m; msg != nil; msg = msg.message {
		if msg.Meta == nil || !msg.Has(key) {
			continue
		}
		for _, k := range msg.Meta["option"] {
			if k == key {
				if v, ok := msg.Data[key]; ok {
					return v
				}
				return msg.Meta[key]
			}
		}
	}
	return nil
}
func (m *Message) Optionx(key string, arg ...string) interface{} {
	value := m.Conf(key)
	if value == "" {
		value = m.Option(key)
	}

	if len(arg) > 0 {
		value = fmt.Sprintf(arg[0], value)
	}
	return value
}
func (m *Message) Magic(begin string, chain interface{}, args ...interface{}) interface{} {
	auth := []string{"bench", "session", "user", "role", "componet", "command"}
	key := []string{"bench", "sessid", "username", "role", "componet", "command"}
	aaa := m.Sess("aaa", false)
	for i, v := range auth {
		if v == begin {
			h := m.Option(key[i])
			if v == "user" {
				h, _ = kit.Hash("username", m.Option("username"))
			}

			data := aaa.Confv("auth", []string{h, "data"})

			if kit.Format(chain) == "" {
				return data
			}

			if len(args) > 0 {
				value := kit.Chain(data, chain, args[0])
				aaa.Conf("auth", []string{m.Option(key[i]), "data"}, value)
				return value
			}

			value := kit.Chain(data, chain)
			if value != nil {
				return value
			}

			if i < len(auth)-1 {
				begin = auth[i+1]
			}
		}
	}
	return nil
}
func (m *Message) Current(text string) string {
	cs := []string{}
	if pod := kit.Format(m.Magic("session", "current.pod")); pod != "" {
		cs = append(cs, "context", "ssh", "remote", "'"+pod+"'")
	}
	if ctx := kit.Format(m.Magic("session", "current.ctx")); ctx != "" {
		cs = append(cs, "context", ctx)
	}
	if cmd := kit.Format(m.Magic("session", "current.cmd")); cmd != "" {
		cs = append(cs, cmd)
	}
	m.Log("info", "%s %s current %v", m.Option("username"), m.Option("sessid"), cs)
	cs = append(cs, text)
	return strings.Join(cs, " ")
}
func (m *Message) Append(key string, arg ...interface{}) string {
	if len(arg) > 0 {
		m.Insert(key, 0, arg...)
		if _, ok := m.Meta[key]; ok {
			m.Add("append", key)
		}
	}

	ms := []*Message{m}
	for i := 0; i < len(ms); i++ {
		ms = append(ms, ms[i].messages...)
		if !ms[i].Has(key) {
			continue
		}
		for _, k := range ms[i].Meta["append"] {
			if k == key {
				return ms[i].Get(key)
			}
		}
	}
	return ""
}
func (m *Message) Appendi(key string, arg ...interface{}) int64 {
	i, _ := strconv.ParseInt(m.Append(key, arg...), 10, 64)
	return i
}
func (m *Message) Appends(key string, arg ...interface{}) bool {
	return kit.Right(m.Append(key, arg...))
}
func (m *Message) Appendv(key string, arg ...interface{}) interface{} {
	if len(arg) > 0 {
		m.Put("append", key, arg[0])
	}

	ms := []*Message{m}
	for i := 0; i < len(ms); i++ {
		ms = append(ms, ms[i].messages...)
		if !ms[i].Has(key) {
			continue
		}
		for _, k := range ms[i].Meta["append"] {
			if k == key {
				if v, ok := ms[i].Data[key]; ok {
					return v
				}
				return ms[i].Meta[key]
			}
		}
	}
	return nil
}
func (m *Message) Table(cbs ...interface{}) *Message {
	if len(m.Meta["append"]) == 0 {
		return m
	}

	// 遍历函数
	if len(cbs) > 0 {
		switch cb := cbs[0].(type) {
		case func(map[string]string) bool:
			nrow := len(m.Meta[m.Meta["append"][0]])
			line := map[string]string{}
			for i := 0; i < nrow; i++ {
				for _, k := range m.Meta["append"] {
					line[k] = m.Meta[k][i]
				}
				if !cb(line) {
					break
				}
			}
			return m
		case func(map[string]string):
			nrow := len(m.Meta[m.Meta["append"][0]])
			for i := 0; i < nrow; i++ {
				line := map[string]string{}
				for _, k := range m.Meta["append"] {
					line[k] = m.Meta[k][i]
				}
				cb(line)
			}
			return m
		case func(int, map[string]string):
			nrow := len(m.Meta[m.Meta["append"][0]])
			for i := 0; i < nrow; i++ {
				line := map[string]string{}
				for _, k := range m.Meta["append"] {
					line[k] = m.Meta[k][i]
				}
				cb(i, line)
			}
			return m
		}
	}

	//计算列宽
	space := m.Confx("table_space")
	depth, width := 0, map[string]int{}
	for _, k := range m.Meta["append"] {
		if len(m.Meta[k]) > depth {
			depth = len(m.Meta[k])
		}
		width[k] = kit.Width(k, len(space))
		for _, v := range m.Meta[k] {
			if kit.Width(v, len(space)) > width[k] {
				width[k] = kit.Width(v, len(space))
			}
		}
	}

	// 回调函数
	var cb func(maps map[string]string, list []string, line int) (goon bool)
	if len(cbs) > 0 {
		cb = cbs[0].(func(maps map[string]string, list []string, line int) (goon bool))
	} else {
		row := m.Confx("table_row_sep")
		col := m.Confx("table_col_sep")
		compact := kit.Right(m.Confx("table_compact"))
		cb = func(maps map[string]string, lists []string, line int) bool {
			for i, v := range lists {
				if k := m.Meta["append"][i]; compact {
					v = maps[k]
				}

				if m.Echo(v); i < len(lists)-1 {
					m.Echo(col)
				}
			}
			m.Echo(row)
			return true
		}
	}

	// 输出表头
	row := map[string]string{}
	wor := []string{}
	for _, k := range m.Meta["append"] {
		row[k], wor = k, append(wor, k+strings.Repeat(space, width[k]-kit.Width(k, len(space))))
	}
	if !cb(row, wor, -1) {
		return m
	}

	// 输出数据
	for i := 0; i < depth; i++ {
		row := map[string]string{}
		wor := []string{}
		for _, k := range m.Meta["append"] {
			data := ""
			if i < len(m.Meta[k]) {
				data = m.Meta[k][i]
			}

			row[k], wor = data, append(wor, data+strings.Repeat(space, width[k]-kit.Width(data, len(space))))
		}
		if !cb(row, wor, i) {
			break
		}
	}

	return m
}
func (m *Message) Sort(key string, arg ...string) *Message {
	cmp := "str"
	if len(arg) > 0 {
		cmp = arg[0]
	}

	number := map[int]int{}
	table := []map[string]string{}
	m.Table(func(line map[string]string, lists []string, index int) bool {
		if index != -1 {
			table = append(table, line)
			switch cmp {
			case "int":
				number[index] = kit.Int(line[key])
			case "int_r":
				number[index] = -kit.Int(line[key])
			case "time":
				number[index] = kit.Time(line[key])
			case "time_r":
				number[index] = -kit.Time(line[key])
			}
		}
		return true
	})

	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			result := false
			switch cmp {
			case "str":
				if table[i][key] > table[j][key] {
					result = true
				}
			case "str_r":
				if table[i][key] < table[j][key] {
					result = true
				}
			default:
				if number[i] > number[j] {
					result = true
				}
			}

			if result {
				table[i], table[j] = table[j], table[i]
				number[i], number[j] = number[j], number[i]
			}
		}
	}

	for _, k := range m.Meta["append"] {
		delete(m.Meta, k)
	}

	for _, v := range table {
		for _, k := range m.Meta["append"] {
			m.Add("append", k, v[k])
		}
	}
	return m
}
func (m *Message) Parse(arg interface{}) string {
	switch str := arg.(type) {
	case string:
		if len(str) > 1 && str[0] == '$' {
			return m.Cap(str[1:])
		}
		if len(str) > 1 && str[0] == '@' {
			if v := m.Option(str[1:]); v != "" {
				return v
			}
			if v := kit.Format(m.Magic("bench", str[1:])); v != "" {
				return v
			}
			v := m.Conf(str[1:])
			return v
		}
		return str
	}
	return ""
}
func (m *Message) ToHTML(style string) string {
	cmd := strings.Join(m.Meta["detail"], " ")
	result := []string{}
	if len(m.Meta["append"]) > 0 {
		result = append(result, fmt.Sprintf("<table class='%s'>", style))
		result = append(result, "<caption>", cmd, "</caption>")
		m.Table(func(maps map[string]string, list []string, line int) bool {
			if line == -1 {
				result = append(result, "<tr>")
				for _, v := range list {
					result = append(result, "<th>", v, "</th>")
				}
				result = append(result, "</tr>")
				return true
			}
			result = append(result, "<tr>")
			for _, v := range list {
				result = append(result, "<td>", v, "</td>")
			}
			result = append(result, "</tr>")
			return true
		})
		result = append(result, "</table>")
	} else {
		result = append(result, "<pre><code>")
		result = append(result, fmt.Sprintf("%s", m.Find("shy", false).Conf("prompt")), cmd, "\n")
		result = append(result, m.Meta["result"]...)
		result = append(result, "</code></pre>")
	}
	return strings.Join(result, "")
}

func (m *Message) Gdb(arg ...interface{}) interface{} {
	if g := m.Sess("gdb", false); g != nil {
		if gdb, ok := g.target.Server.(DEBUG); ok {
			return gdb.Wait(m, arg...)
		}
	}
	return nil
}
func (m *Message) Log(action string, str string, arg ...interface{}) *Message {
	if m.Options("log.disable") {
		return m
	}

	if l := m.Sess("log", false); l != nil {
		if log, ok := l.target.Server.(LOGGER); ok {
			if action == "error" {
				log.Log(m, "error", "chain: %s", m.Format("chain"))
			}
			log.Log(m, action, str, arg...)
			if action == "error" {
				log.Log(m, "error", "stack: %s", m.Format("stack"))
			}
			return m
		}
	} else {
		log.Printf(str, arg...)
	}

	if action == "error" {
		kit.Log("error", fmt.Sprintf("chain: %s", m.Format("chain")))
		kit.Log("error", fmt.Sprintf("%s %s %s", m.Format(), action, fmt.Sprintf(str, arg...)))
		kit.Log("error", fmt.Sprintf("stack: %s", m.Format("stack")))
	}

	return m
}
func (m *Message) Show(args ...interface{}) *Message {
	if m.Option("cli.modal") == "action" {
		fmt.Printf(kit.Format(args...))
	} else if kit.STDIO != nil {
		kit.STDIO.Show(args...)
	}
	return m
}
func (m *Message) Assert(e interface{}, msg ...string) bool {
	switch v := e.(type) {
	case nil:
		return true
	case *Message:
		if v.Result(0) != "error: " {
			return true
		}
		e = v.Result(1)
		e = errors.New(v.Result(1))
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

	m.Log("error", "%v", e)
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
				msg.Assert(e)
			}
		}
	}()

	if len(hand) > 0 {
		hand[0](msg)
	}
	return m
}
func (m *Message) GoFunc(msg *Message, hand ...func(msg *Message)) *Message {
	go func() {
		msg.Option("routine", m.Capi("ngo", 1))
		// msg.Log("info", "%v safe go begin", ngo)
		// kit.Log("error", "%s ngo %s start", msg.Format(), ngo)
		m.TryCatch(msg, true, hand...)
		// kit.Log("error", "%s ngo %s end", msg.Format(), ngo)
		// msg.Log("info", "%v safe go end", ngo)
	}()
	return m
}
func (m *Message) GoLoop(msg *Message, hand ...func(msg *Message)) *Message {
	m.GoFunc(msg, func(msg *Message) {
		for {
			hand[0](msg)
		}
	})
	return m
}
func (m *Message) Start(name string, help string, arg ...string) bool {
	return m.Set("detail", arg).target.Spawn(m, name, help).Begin(m).Start(m)
}
func (m *Message) Close(arg ...string) bool {
	return m.Target().Close(m, arg...)
}
func (m *Message) Wait() bool {
	if m.target.exit != nil {
		return <-m.target.exit
	}
	return true
}

func (m *Message) Find(name string, root ...bool) *Message {
	if name == "" {
		return m.Spawn()
	}
	target := m.target.root
	if len(root) > 0 && !root[0] {
		target = m.target
	}

	cs := target.contexts
	for _, v := range strings.Split(name, ".") {
		if x, ok := cs[v]; ok {
			target, cs = x, x.contexts
		} else if target.Name == v {
			continue
		} else {
			m.Log("error", "context not find %s", name)
			return nil
		}
	}

	if len(root) > 1 && root[1] {
		m.target = target
		return m
	}

	return m.Spawn(target)
}
func (m *Message) Search(key string, root ...bool) []*Message {
	reg, e := regexp.Compile(key)
	m.Assert(e)

	target := m.target
	if target == nil {
		return []*Message{nil}
	}
	if len(root) > 0 && root[0] {
		target = m.target.root
	}

	cs := make([]*Context, 0, 3)
	target.Travel(m, func(m *Message, i int) bool {
		if reg.MatchString(m.target.Name) || reg.FindString(m.target.Help) != "" {
			m.Log("search", "%d %s match [%s]", len(cs), m.target.Name, key)
			cs = append(cs, m.target)
		}
		return false
	})

	ms := make([]*Message, len(cs))
	for i := 0; i < len(cs); i++ {
		ms[i] = m.Spawn(cs[i])
	}
	if len(ms) == 0 {
		ms = append(ms, nil)
	}

	return ms
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
func (m *Message) Match(key string, spawn bool, hand func(m *Message, s *Context, c *Context, key string) bool) *Message {
	if m == nil {
		return m
	}

	context := []*Context{m.target}
	for _, v := range []string{"aaa", "ssh", "cli", "nfs"} {
		if msg := m.Sess(v, false); msg != nil && msg.target != nil {
			context = append(context, msg.target)
		}
	}
	// if m.target.root != nil && m.target.root.Configs != nil && m.target.root.Configs["search"] != nil && m.target.root.Configs["search"].Value != nil {
	// 	target := m.target
	// 	for _, v := range kit.Trans(kit.Chain(m.target.root.Configs["search"].Value, "context")) {
	// 		if t := m.Find(v, true, true); t != nil {
	// 			kit.Log("error", "%v", t)
	// 			// 		// 	context = append(context, t.target)
	// 		}
	// 	}
	// 	m.target = target
	// }

	context = append(context, m.source)

	for _, s := range context {
		for c := s; c != nil; c = c.context {
			if hand(m, s, c, key) {
				return m
			}
		}
	}
	return m
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
func (m *Message) Backs(msg *Message) *Message {
	m.Back(msg)
	return msg
}
func (m *Message) CallBack(sync bool, cb func(msg *Message) (sub *Message), arg ...interface{}) *Message {
	if !sync {
		return m.Call(cb, arg...)
	}

	wait := make(chan *Message, 10)
	// m.GoFunc(m, func(m *Message) {
	m.Call(func(sub *Message) *Message {
		msg := cb(sub)
		m.Log("sync", m.Format("done", "result", "append"))
		wait <- m
		return msg
	}, arg...)
	// })

	m.Log("sync", m.Format("wait", "result", "append"))
	select {
	case <-time.After(kit.Duration(m.Conf("call_timeout"))):
		m.Log("sync", m.Format("timeout", "result", "append"))
	case <-wait:
	}
	return m
}
func (m *Message) Free(cbs ...func(msg *Message) (done bool)) *Message {
	if len(cbs) == 0 {
		for i := len(m.freedoms) - 1; i >= 0; i-- {
			m.Log("free", "%d/%d", i, len(m.freedoms)-1)
			if !m.freedoms[i](m) {
				break
			}
			m.freedoms = m.freedoms[:i]
		}
		return m
	}

	m.freedoms = append(m.freedoms, cbs...)
	return m
}

func (m *Message) Cmdp(t time.Duration, head []string, prefix []string, suffix [][]string) *Message {
	if head != nil && len(head) > 0 {
		m.Show(strings.Join(head, " "), "...\n")
	}

	for i, v := range suffix {
		m.Show(fmt.Sprintf("%v/%v %v...\n", i+1, len(suffix), v))
		m.CopyFuck(m.Cmd(prefix, v), "append")
		time.Sleep(t)
	}
	m.Show("\n")
	m.Table()
	return m
}
func (m *Message) Cmdm(args ...interface{}) *Message {
	m.Log("info", "current: %v", m.Magic("session", "current"))

	arg := []string{}
	if pod := kit.Format(m.Magic("session", "current.pod")); pod != "" {
		arg = append(arg, "context", "ssh", "remote", pod)
	}
	if ctx := kit.Format(m.Magic("session", "current.ctx")); ctx != "" {
		arg = append(arg, "context", ctx)
	}
	arg = append(arg, kit.Trans(args...)...)

	// 执行命令
	m.Spawn().Cmd(arg).CopyTo(m)
	// m.Magic("session", "current.ctx", msg.target.Name)
	return m
}
func (m *Message) Cmdy(args ...interface{}) *Message {
	m.Cmd(args...).CopyTo(m)
	return m
}
func (m *Message) Cmdx(args ...interface{}) string {
	msg := m.Cmd(args...)
	if msg.Result(0) == "error: " {
		return msg.Result(1)
	}
	return msg.Result(0)
}
func (m *Message) Cmds(args ...interface{}) bool {
	return m.Cmd(args...).Results(0)
}
func (m *Message) Cmd(args ...interface{}) *Message {
	if m == nil {
		return m
	}

	if len(args) > 0 {
		m.Set("detail", kit.Trans(args...))
	}
	key, arg := m.Meta["detail"][0], m.Meta["detail"][1:]

	msg := m
	if strings.Contains(key, ":") {
		ps := strings.Split(key, ":")
		msg, key, arg = msg.Sess("ssh"), "_route", append([]string{"sync", ps[0], ps[1]}, arg...)
		defer func() { m.Copy(msg, "append").Copy(msg, "result") }()
		m.Hand = true

	} else if strings.Contains(key, ".") {
		arg := strings.Split(key, ".")
		msg, key = msg.Sess(arg[0]), arg[1]
		msg.Option("remote_code", "")
	}
	if msg == nil {
		return msg
	}

	msg = msg.Match(key, true, func(msg *Message, s *Context, c *Context, key string) bool {
		msg.Hand = false
		if x, ok := c.Commands[key]; ok && x.Hand != nil {
			msg.TryCatch(msg, true, func(msg *Message) {
				msg.Log("cmd", "%s %s %v %v", c.Name, key, arg, msg.Meta["option"])

				for _, form := range []map[string]int{map[string]int{"page.limit": 1, "page.offset": 1}, x.Form} {

					if args := []string{}; form != nil {
						for i := 0; i < len(arg); i++ {
							if n, ok := form[arg[i]]; ok {
								if n < 0 {
									n += len(arg) - i
								}
								for j := i + 1; j <= i+n && j < len(arg); j++ {
									if _, ok := form[arg[j]]; ok {
										n = j - i - 1
									}
								}
								if i+1+n > len(arg) {
									msg.Add("option", arg[i], arg[i+1:])
								} else {
									msg.Add("option", arg[i], arg[i+1:i+1+n])
								}
								i += n
							} else {
								args = append(args, arg[i])
							}
						}
						arg = args
					}
				}

				target := msg.target
				msg.target = s

				msg.Hand = true
				switch v := msg.Gdb("command", key, arg).(type) {
				case string:
					msg.Echo(v)
				case nil:
					if msg.Options("auto_cmd") {
						if x.Auto != nil {
							x.Auto(msg, c, key, arg...)
						}
					} else {
						x.Hand(msg, c, key, arg...)
					}
				}
				if msg.target == s {
					msg.target = target
				}
			})
		}
		return msg.Hand
	})

	if !msg.Hand {
		msg.Log("error", "cmd run error %s", msg.Format())
	}
	return msg
}

func (m *Message) Confm(key string, args ...interface{}) map[string]interface{} {
	random := ""
	var chain interface{}
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case []interface{}:
			chain, args = arg, args[1:]
		case []string:
			chain, args = arg, args[1:]
		case string:
			switch arg {
			case "%", "*":
				random, args = arg, args[1:]
			default:
				chain, args = arg, args[1:]
			}
		}
	}

	var v interface{}
	if chain == nil {
		v = m.Confv(key)
	} else {
		v = m.Confv(key, chain)
	}

	table, _ := v.([]interface{})
	value, _ := v.(map[string]interface{})
	if len(args) == 0 {
		return value
	}

	switch fun := args[0].(type) {
	case func(int, string):
		for i, v := range table {
			fun(i, kit.Format(v))
		}
	case func(int, string) bool:
		for i, v := range table {
			if fun(i, kit.Format(v)) {
				break
			}
		}
	case func(string, string):
		for k, v := range value {
			fun(k, kit.Format(v))
		}
	case func(string, string) bool:
		for k, v := range value {
			if fun(k, kit.Format(v)) {
				break
			}
		}
	case func(map[string]interface{}):
		if len(value) == 0 {
			return nil
		}
		fun(value)
	case func(string, map[string]interface{}):
		switch random {
		case "%":
			n, i := rand.Intn(len(value)), 0
			for k, v := range value {
				if val, ok := v.(map[string]interface{}); i == n && ok {
					fun(k, val)
					break
				}
				i++
			}
		case "*":
			fallthrough
		default:
			for k, v := range value {
				if val, ok := v.(map[string]interface{}); ok {
					fun(k, val)
				}
			}
		}
	case func(string, int, map[string]interface{}):
		for k, v := range value {
			if val, ok := v.([]interface{}); ok {
				for i, v := range val {
					if val, ok := v.(map[string]interface{}); ok {
						fun(k, i, val)
					}
				}
			}
		}

	case func(string, map[string]interface{}) bool:
		for k, v := range value {
			if val, ok := v.(map[string]interface{}); ok {
				if fun(k, val) {
					break
				}
			}
		}
	case func(int, map[string]interface{}):
		for i := m.Optioni("page.begin"); i < len(table); i++ {
			if val, ok := table[i].(map[string]interface{}); ok {
				fun(i, val)
			}
		}
	}
	return value
}
func (m *Message) Confx(key string, args ...interface{}) string {
	value := kit.Select(m.Conf(key), m.Option(key))
	if len(args) == 0 {
		return value
	}

	switch arg := args[0].(type) {
	case []string:
		if len(args) > 1 {
			value = kit.Select(value, arg, args[1])
		} else {
			value = kit.Select(value, arg)
		}
		args = args[1:]
	case map[string]interface{}:
		value = kit.Select(value, kit.Format(arg[key]))
	case string:
		value = kit.Select(value, arg)
	case nil:
	default:
		value = kit.Select(value, args[0])
	}

	format := "%s"
	if args = args[1:]; len(args) > 0 {
		format, args = kit.Format(args[0]), args[1:]
	}
	arg := []interface{}{format, value}
	for _, v := range args {
		arg = append(arg, v)
	}

	return kit.Format(arg...)
}
func (m *Message) Confs(key string, arg ...interface{}) bool {
	return kit.Right(m.Confv(key, arg...))
}
func (m *Message) Confi(key string, arg ...interface{}) int {
	return kit.Int(m.Confv(key, arg...))
}
func (m *Message) Confv(key string, args ...interface{}) interface{} {
	if strings.Contains(key, ".") {
		target := m.target
		defer func() { m.target = target }()

		ps := strings.Split(key, ".")
		if msg := m.Sess(ps[0], false); msg != nil {
			m.target, key = msg.target, ps[1]
		}
	}

	var config *Config
	m.Match(key, false, func(m *Message, s *Context, c *Context, key string) bool {
		if x, ok := c.Configs[key]; ok {
			config = x
			return true
		}
		return false
	})

	if len(args) == 0 {
		if config == nil {
			return nil
		}
		return config.Value
	}

	if config == nil {
		config = &Config{}
		m.target.Configs[key] = config
	}

	switch config.Value.(type) {
	case string:
		config.Value = kit.Format(args...)
	case bool:
		config.Value = kit.Right(args...)
	case int:
		config.Value = kit.Int(args...)
	case nil:
		config.Value = args[0]
	default:
		for i := 0; i < len(args); i += 2 {
			if i < len(args)-1 {
				config.Value = kit.Chain(config.Value, args[i], args[i+1])
			} else {
				return kit.Chain(config.Value, args[i])
			}
		}
	}

	return config.Value
}
func (m *Message) Conf(key string, args ...interface{}) string {
	return kit.Format(m.Confv(key, args...))
}
func (m *Message) Caps(key string, arg ...interface{}) bool {
	if len(arg) > 0 {
		return kit.Right(m.Cap(key, arg...))
	}
	return kit.Right(m.Cap(key))
}
func (m *Message) Capi(key string, arg ...interface{}) int {
	n := kit.Int(m.Cap(key))
	if len(arg) > 0 {
		return kit.Int(m.Cap(key, n+kit.Int(arg...)))
	}
	return n
}
func (m *Message) Cap(key string, arg ...interface{}) string {
	var cache *Cache
	m.Match(key, false, func(m *Message, s *Context, c *Context, key string) bool {
		if x, ok := c.Caches[key]; ok {
			cache = x
			return true
		}
		return false
	})

	if len(arg) == 0 {
		if cache == nil {
			return ""
		}
		if cache.Hand != nil {
			return cache.Hand(m, cache)
		}
		return cache.Value
	}

	if cache == nil {
		cache = &Cache{}
		m.target.Caches[key] = cache
	}

	if cache.Hand != nil {
		cache.Value = cache.Hand(m, cache, kit.Format(arg...))
	} else {
		cache.Value = kit.Format(arg...)
	}
	return cache.Value
}
