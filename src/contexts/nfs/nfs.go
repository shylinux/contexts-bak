package nfs

import (
	"contexts/ctx"
	"crypto/md5"
	"toolkit"

	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/url"
	"regexp"
	"strings"

	"bufio"
	"github.com/nsf/termbox-go"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sort"
	"time"
)

type NFS struct {
	io  io.ReadWriter
	in  *os.File
	out *os.File

	send chan *ctx.Message
	echo chan *ctx.Message
	hand map[int]*ctx.Message

	*ctx.Context
}

func dir(m *ctx.Message, name string, level int, deep bool, dir_type string, trip int, dir_reg *regexp.Regexp, fields []string, format string) {
	back, e := os.Getwd()
	m.Assert(e)

	if fs, e := ioutil.ReadDir(name); m.Assert(e) {
		for _, f := range fs {
			if f.Name() == "." || f.Name() == ".." {
				continue
			}

			if strings.HasPrefix(f.Name(), ".") && dir_type != "all" {
				continue
			}

			f, e := os.Stat(path.Join(name, f.Name()))
			if e != nil {
				m.Log("info", "%s", e)
				continue
			}

			if !(dir_type == "file" && f.IsDir() || dir_type == "dir" && !f.IsDir()) && (dir_reg == nil || dir_reg.MatchString(f.Name())) {
				for _, field := range fields {
					switch field {
					case "time":
						m.Add("append", "time", f.ModTime().Format(format))
					case "type":
						if m.Assert(e) && f.IsDir() {
							m.Add("append", "type", "dir")
						} else {
							m.Add("append", "type", "file")
						}
					case "full":
						if f.IsDir() {
							m.Add("append", "full", path.Join(back, name, f.Name())+"/")
						} else {
							m.Add("append", "full", path.Join(back, name, f.Name()))
						}
					case "path":
						if f.IsDir() {
							m.Add("append", "path", path.Join(back, name, f.Name())[trip:]+"/")
						} else {
							m.Add("append", "path", path.Join(back, name, f.Name())[trip:])
						}
					case "tree":
						if level == 0 {
							m.Add("append", "tree", f.Name())
						} else {
							m.Add("append", "tree", strings.Repeat("| ", level-1)+"|-"+f.Name())
						}
					case "filename":
						if f.IsDir() {
							m.Add("append", "filename", f.Name()+"/")
						} else {
							m.Add("append", "filename", f.Name())
						}
					case "size":
						m.Add("append", "size", f.Size())
					case "line":
						nline := 0
						if f.IsDir() {
							d, e := ioutil.ReadDir(path.Join(name, f.Name()))
							m.Assert(e)
							nline = len(d)
						} else {
							f, e := os.Open(path.Join(name, f.Name()))
							m.Assert(e)
							defer f.Close()

							bio := bufio.NewScanner(f)
							for bio.Scan() {
								bio.Text()
								nline++
							}
						}
						m.Add("append", "line", nline)
					case "hash", "hashs":
						if f.IsDir() {
							d, e := ioutil.ReadDir(path.Join(name, f.Name()))
							m.Assert(e)
							meta := []string{}
							for _, v := range d {
								meta = append(meta, fmt.Sprintf("%s%d%s", v.Name(), v.Size(), v.ModTime()))
							}
							sort.Strings(meta)

							h := sha1.Sum([]byte(strings.Join(meta, "")))
							m.Add("append", "hash", hex.EncodeToString(h[:]))
							break
						}

						f, e := ioutil.ReadFile(path.Join(name, f.Name()))
						m.Assert(e)
						h := sha1.Sum(f)

						if field == "hash" {
							m.Add("append", "hash", hex.EncodeToString(h[:]))
						} else {
							m.Add("append", "hash", hex.EncodeToString(h[:4]))
						}
					}
				}
			}
			if f.IsDir() && deep {
				dir(m, path.Join(name, f.Name()), level+1, deep, dir_type, trip, dir_reg, fields, format)
			}
		}
	}
}
func open(m *ctx.Message, name string, arg ...int) (string, *os.File, error) {
	if !path.IsAbs(name) {
		pwd := m.Confv("pwd").([]interface{})
		for _, v := range pwd {
			p := path.Join(v.(string), name)
			if len(arg) > 0 {
				name = p
				break
			}
			if s, e := os.Stat(p); e == nil && !s.IsDir() {
				name = p
				break
			}
		}
	}

	flag := os.O_RDONLY
	if len(arg) > 0 {
		flag = arg[0]
	}

	f, e := os.OpenFile(name, flag, 0660)
	if e == nil {
		m.Log("info", "open %s", name)
		return name, f, e
	}
	m.Log("warn", "%v", e)
	return name, f, e
}

func (nfs *NFS) Recv(line string) (field string, value string) {
	m := nfs.Context.Message()

	m.Log("recv", "%d [%s]", len(line), line)
	m.Capi("nread", len(line)+1)

	word := strings.Split(line, ": ")
	field, e := url.QueryUnescape(word[0])
	m.Assert(e)
	if len(word) == 1 {
		return
	}
	if len(word[1]) == 0 {
		return
	}

	value, e = url.QueryUnescape(word[1])
	m.Assert(e)
	return
}
func (nfs *NFS) Send(meta string, arg ...interface{}) *NFS {
	m := nfs.Context.Message()

	line := "\n"
	if meta != "" {
		if text, ok := arg[0].(string); ok && meta == "result" && len(text) > 1024 {
			text := arg[0].(string)
			for i := 0; i < len(text); i += 1024 {
				j := i + 1024
				if j >= len(text) {
					j = len(text)
				}
				line = fmt.Sprintf("%s: %s\n", url.QueryEscape(meta), url.QueryEscape(kit.Format(text[i:j])))
				n, e := fmt.Fprint(nfs.io, line)
				m.Assert(e)
				m.Capi("nwrite", n)
				m.Log("send", "%d [%s]", len(line), line)
			}
			return nfs
		} else {
			line = fmt.Sprintf("%s: %s\n", url.QueryEscape(meta), url.QueryEscape(kit.Format(arg[0])))
		}
	}

	n, e := fmt.Fprint(nfs.io, line)
	m.Assert(e)
	m.Capi("nwrite", n)
	m.Log("send", "%d [%s]", len(line), line)

	return nfs
}

func (nfs *NFS) Spawn(m *ctx.Message, c *ctx.Context, arg ...string) ctx.Server {
	if len(arg) > 0 && (arg[0] == "scan" || arg[0] == "open" || arg[0] == "append") {
		c.Caches = map[string]*ctx.Cache{
			"pos":    &ctx.Cache{Name: "pos", Value: "0", Help: "pos"},
			"size":   &ctx.Cache{Name: "size", Value: "0", Help: "size"},
			"nread":  &ctx.Cache{Name: "nread", Value: "0", Help: "nread"},
			"nwrite": &ctx.Cache{Name: "nwrite", Value: "0", Help: "nwrite"},
		}
		c.Configs = map[string]*ctx.Config{}
	} else {
		c.Caches = map[string]*ctx.Cache{
			"nsend":  &ctx.Cache{Name: "消息发送数量", Value: "0", Help: "消息发送数量"},
			"nrecv":  &ctx.Cache{Name: "消息接收数量", Value: "0", Help: "消息接收数量"},
			"nread":  &ctx.Cache{Name: "nread", Value: "0", Help: "nread"},
			"nwrite": &ctx.Cache{Name: "nwrite", Value: "0", Help: "nwrite"},
		}
		c.Configs = map[string]*ctx.Config{}
	}

	return &NFS{Context: c}

}
func (nfs *NFS) Begin(m *ctx.Message, arg ...string) ctx.Server {
	return nfs
}
func (nfs *NFS) Start(m *ctx.Message, arg ...string) bool {

	if len(arg) > 0 && (arg[0] == "open" || arg[0] == "append") {
		nfs.out = m.Optionv("out").(*os.File)
		nfs.in = m.Optionv("in").(*os.File)
		m.Cap("stream", arg[1])

		if s, e := nfs.in.Stat(); m.Assert(e) {
			if m.Capi("size", int(s.Size())); arg[0] == "append" {
				m.Capi("pos", int(s.Size()))
			}
		}
		return false
	}

	if len(arg) > 0 && arg[0] == "scan" {
		m.Cap("stream", arg[1])
		nfs.Caches["ninput"] = &ctx.Cache{Value: "0"}
		nfs.Caches["noutput"] = &ctx.Cache{Value: "0"}
		nfs.Caches["termbox"] = &ctx.Cache{Value: "0"}
		nfs.Configs["input"] = &ctx.Config{Value: []interface{}{}}
		nfs.Configs["output"] = &ctx.Config{Value: []interface{}{}}
		nfs.Configs["prompt"] = &ctx.Config{Value: ""}

		// 终端控制
		if nfs.in = m.Optionv("bio.in").(*os.File); m.Has("bio.out") {
			if nfs.out = m.Optionv("bio.out").(*os.File); m.Conf("runtime", "host.GOOS") != "windows" && !m.Options("daemon") {
				m.Conf("term", "use", nfs.Term(m, "init") != nil)
				defer nfs.Term(m, "exit")
				kit.STDIO = nfs

			} else if m.Options("daemon") {
				return false
			}
		}

		// 语句堆栈
		stack := &kit.Stack{}
		stack.Push(m.Option("stack.key", "source"), m.Options("stack.run", true), m.Optioni("stack.pos", 0))
		m.Optionv("bio.ctx", m.Target())
		m.Optionv("bio.stack", stack)

		line, bio := "", bufio.NewScanner(nfs)
		for nfs.prompt(); ; nfs.prompt() {

			// 读取数据
			for bio.Scan() {
				if text := bio.Text(); text == "" {
					continue
				} else if !strings.HasSuffix(text, "\\") {
					line += text
					break
				} else {
					line += strings.TrimSuffix(text, "\\")
				}
			}
			if line == "" {
				break
			}
			m.Log("debug", "%s %d %d [%s]", "input", m.Capi("ninput", 1), len(line), line)
			m.Confv("input", -2, map[string]interface{}{"time": time.Now().Unix(), "line": line})

			// 解析数据
			for i := m.Capi("ninput") - 1; i < m.Capi("ninput"); i++ {
				line = m.Conf("input", []interface{}{i, "line"})
				m.Optionv("input", m.Confv("input"))
				m.Optioni("stack.pos", i)

				// 执行语句
				msg := m.Cmd("yac.parse", line+"\n")
				nfs.print(m.Conf("prompt"), line)
				nfs.print(msg.Meta["result"]...)

				// 切换模块
				if v := msg.Optionv("bio.ctx"); v != nil {
					m.Optionv("bio.ctx", v)
				}

				// 跳转语句
				if msg.Appends("bio.pos0") {
					i = int(msg.Appendi("bio.pos0")) - 1
					msg.Append("bio.pos0", "")
				}

				// 结束脚本
				if msg.Appends("bio.end") {
					m.Copy(msg, "append")
					m.Copy(msg, "result")
					msg.Appends("bio.end", "")
					return true
				}
			}
			line = ""
		}
		return true
	}

	m.Cap("stream", m.Option("ms_source"))
	nfs.io, _ = m.Optionv("io").(io.ReadWriter)
	nfs.send = make(chan *ctx.Message, 10)
	nfs.echo = make(chan *ctx.Message, 10)
	nfs.hand = map[int]*ctx.Message{}

	// 消息发送队列
	m.GoLoop(m, func(m *ctx.Message) {
		msg, code, meta, body := m, 0, "detail", "option"
		select {
		case msg = <-nfs.send: // 发送请求
			code = msg.Code()
			nfs.hand[code] = msg
		case msg = <-nfs.echo: // 发送响应
			code, meta, body = msg.Optioni("remote_code"), "result", "append"
		}

		nfs.Send("code", code)
		for _, v := range msg.Meta[meta] {
			nfs.Send(meta, v)
		}
		for _, k := range msg.Meta[body] {
			for _, v := range msg.Meta[k] {
				nfs.Send(k, v)
			}
		}
		nfs.Send("")
	})

	// 消息接收队列
	msg, code, head, body := m, "0", "result", "append"
	bio := bufio.NewScanner(nfs.io)
	bio.Buffer(make([]byte, m.Confi("buf_size")), m.Confi("buf_size"))
	for bio.Scan() {

		m.TryCatch(m, true, func(m *ctx.Message) {
			switch field, value := nfs.Recv(bio.Text()); field {
			case "code":
				msg, code = m.Sess("ms_target"), value
				msg.Meta = map[string][]string{}

			case "detail":
				head, body = "detail", "option"
				msg.Add(field, value)

			case "result":
				head, body = "result", "append"
				msg.Add(field, value)

			case "":
				m.Log("recv", "time %v", time.Now().Format(m.Conf("time", "format")))
				if head == "detail" { // 接收请求
					msg.Detail(-1, "_route")
					msg.Option("remote_code", code)
					m.Gos(msg, func(msg *ctx.Message) {
						msg.Call(func(msg *ctx.Message) *ctx.Message {
							nfs.echo <- msg
							return nil
						})
					})
				} else { // 接收响应
					m.Set("option", "code", code).Gos(msg, func(msg *ctx.Message) {
						if h, ok := nfs.hand[kit.Int(m.Option("code"))]; ok {
							h.CopyFuck(msg, "result").CopyFuck(msg, "append").Back(h)
						}
					})
				}
				msg, code, head, body = nil, "0", "result", "append"

			default:
				msg.Add(body, field, value)
			}
		}, func(m *ctx.Message) {
			for bio.Scan() {
				if text := bio.Text(); text == "" {
					break
				}
			}
		})
	}

	m.Sess("tcp", false).Close()
	return true
}
func (nfs *NFS) Close(m *ctx.Message, arg ...string) bool {
	return true
}

var Index = &ctx.Context{Name: "nfs", Help: "存储中心",
	Caches: map[string]*ctx.Cache{
		"nfile": &ctx.Cache{Name: "nfile", Value: "0", Help: "已经打开的文件数量"},
	},
	Configs: map[string]*ctx.Config{
		"term": &ctx.Config{Name: "term", Value: map[string]interface{}{
			"use": "false",
			"mouse": map[string]interface{}{
				"resize": false,
			},
			"width": 80, "height": "24",

			"left": 0, "top": 0, "right": 80, "bottom": 24,
			"cursor_x": 0, "cursor_y": 0, "fgcolor": -1, "bgcolor": -1,
			"prompt": "", "wrap": "false",
			"scroll_count": "5",
			"scroll_lines": "5",
			"begin_row":    0, "begin_col": 0,

			"shadow":      "hello",
			"show_shadow": "false",

			"rest_fg": "0",
			"rest_bg": "7",
			"pick_fg": "0",
			"pick_bg": "7",
			"pick":    "",

			"help_index":     0,
			"help_state":     "command",
			"help_next_auto": "=",
			"help_stack":     []interface{}{},
			"help_table":     map[string]interface{}{},
		}, Help: "二维码的默认大小"},
		"auto": &ctx.Config{Name: "auto", Value: map[string]interface{}{
			"!": map[string]interface{}{
				"state":     "message",
				"next_auto": ":",
				"color":     2, "cmd": "message",
				"table": "message", "field": "code",
				"format": "%s(%s) %s->%s %s %s", "fields": []interface{}{"code", "time", "source", "target", "details", "options"},
			},
			"~": map[string]interface{}{
				"state": "context", "next_auto": ":",
				"color": 2, "cmd": "context",
				"table": "context", "field": "name",
				"format": "%s(%s) %s %s", "fields": []interface{}{"name", "status", "stream", "help"},
			},
			"": map[string]interface{}{
				"state": "command", "next_auto": "=",
				"color": 3, "cmd": "command",
				"table": "command", "field": "key",
				"format": "%s %s", "fields": []interface{}{"key", "name"},
			},
			":": map[string]interface{}{
				"state": "command", "next_auto": "=",
				"color": 3, "cmd": "command",
				"table": "command", "field": "key",
				"format": "%s %s", "fields": []interface{}{"key", "name"},
			},
			"=": map[string]interface{}{
				"cmd":    "help",
				"format": "%s %s %s ", "fields": []interface{}{"value", "name", "help"},
				"color": 3, "table": "command", "field": "value",
				"state": "argument", "next_auto": "=",
			},
			"@": map[string]interface{}{
				"state": "config", "next_auto": "@",
				"color": 4, "cmd": "config",
				"table": "config", "field": "key",
				"format": "%s(%s) %s", "fields": []interface{}{"key", "value", "name"},
			},
			"$": map[string]interface{}{
				"state": "cache", "next_auto": "$",
				"color": 7, "cmd": "cache",
				"table": "cache", "field": "key",
				"format": "%s(%s) %s", "fields": []interface{}{"key", "value", "name"},
			},
		}, Help: "读取文件的缓存区的大小"},

		"buf_size": &ctx.Config{Name: "buf_size", Value: "81920", Help: "读取文件的缓存区的大小"},
		"grep": &ctx.Config{Name: "grep", Value: map[string]interface{}{
			"list": []interface{}{},
		}, Help: "dir命令输出文件名的类型, name: 文件名, tree: 带缩进的文件名, path: 相对路径, full: 绝对路径"},
		"git": &ctx.Config{Name: "git", Value: map[string]interface{}{
			"args":   []interface{}{"-C", "@git_dir"},
			"info":   map[string]interface{}{"cmds": []interface{}{"log", "status", "branch"}},
			"update": map[string]interface{}{"cmds": []interface{}{"stash", "pull", "pop"}},
			"pop":    map[string]interface{}{"args": []interface{}{"stash", "pop"}},
			"commit": map[string]interface{}{"args": []interface{}{"commit", "-am"}},
			"branch": map[string]interface{}{"args": []interface{}{"branch", "-v"}},
			"status": map[string]interface{}{"args": []interface{}{"status", "-sb"}},
			"log":    map[string]interface{}{"args": []interface{}{"log", "-n", "@table.limit", "--skip", "@table.offset", "pretty", "date"}},
			"trans": map[string]interface{}{
				"date":   "--date=format:%m/%d %H:%M",
				"pretty": "--pretty=format:%h %ad %an %s",
			},
		}, Help: "命令集合"},
		"dir_fields": &ctx.Config{Name: "dir_fields(time/type/name/size/line/hash)", Value: "time size line filename", Help: "dir命令输出文件名的类型, name: 文件名, tree: 带缩进的文件名, path: 相对路径, full: 绝对路径"},
		"dir_type":   &ctx.Config{Name: "dir_type(file/dir/both/all)", Value: "both", Help: "dir命令输出的文件类型, file: 只输出普通文件, dir: 只输出目录文件, 否则输出所有文件"},
		"dir": &ctx.Config{Name: "dir", Value: map[string]interface{}{
			"temp":  "var/tmp/file",
			"trash": "var/tmp/trash",
		}, Help: ""},
		"pwd": &ctx.Config{Name: "pwd", Value: []interface{}{"var", "usr", "bin", "etc", ""}, Help: "当前目录"},
	},
	Commands: map[string]*ctx.Command{
		"_init": &ctx.Command{Name: "_init", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			m.Conf("pwd", -2, m.Conf("runtime", "boot.ctx_home"))
			m.Conf("pwd", -2, m.Conf("runtime", "boot.ctx_root"))
			return
		}},
		"_exit": &ctx.Command{Name: "_init", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if m.Confs("term", "use") {
				termbox.Close()
			}
			return
		}},
		"path": &ctx.Command{Name: "path filename", Help: "查找文件路径", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) == 0 {
				return
			}

			m.Confm("pwd", func(index int, value string) bool {
				p := path.Join(value, arg[0])
				if _, e := os.Stat(p); e == nil {
					m.Echo(p)
					return true
				}
				return false
			})
			return
		}},
		"pwd": &ctx.Command{Name: "pwd [all] | [[index] path] ", Help: "当前目录，all: 查看所有, index path: 设置路径, path: 设置当前路径", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) > 0 && arg[0] == "all" {
				m.Cmdy("nfs.config", "pwd")
				return
			}

			index := 0
			if len(arg) > 1 {
				index, arg = kit.Int(arg[0]), arg[1:]
			}
			for i, v := range arg {
				m.Log("info", "pwd %s %s", index+i, v)
				m.Confv("pwd", index+i, v)
			}

			if p := m.Conf("pwd", index); path.IsAbs(p) {
				m.Echo("%s", p)
			} else if wd, e := os.Getwd(); m.Assert(e) {
				m.Echo("%s", path.Join(wd, p))
			}
			return
		}},
		"dir": &ctx.Command{Name: "dir [path [fields...]]", Help: []string{"查看目录, path: 路径, fields...: 查询字段, time|type|full|path|tree|filename|size|line|hash",
			"dir_deep: 递归查询", "dir_type both|file|dir|all: 文件类型", "dir_reg reg: 正则表达式", "dir_sort field order: 排序"},
			Form: map[string]int{"dir_deep": 0, "dir_type": 1, "dir_reg": 1, "dir_sort": 2, "dir_sed": -1},
			Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
				if len(arg) == 0 {
					arg = append(arg, "")
				}

				p := arg[0]
				p0 := p + ".tmp0"
				p1 := p + ".tmp1"
				if args := kit.Trans(m.Optionv("dir_sed")); len(args) > 0 {
					if _, e := os.Stat(p1); e == nil {
						return e
					}
					out, e := os.Create(p1)
					defer os.Remove(p1)
					defer out.Close()
					m.Log("info", "open %v", p1)

					if _, e := os.Stat(p0); e != nil {
						m.Cmd("nfs.copy", p0, p)
					}
					in, e := os.Open(p0)
					defer in.Close()
					m.Log("info", "open %v", p0)

					switch args[0] {
					case "set":
						defer os.Rename(p1, p0)
						defer os.Remove(p0)

						index := kit.Int(args[1])
						for bio, i := bufio.NewScanner(in), 0; bio.Scan(); i++ {
							if i == index {
								out.WriteString(args[2])
							} else {
								out.WriteString(bio.Text())
							}
							out.WriteString("\n")
						}
						return e

					case "add":
						out0, _ := os.OpenFile(p0, os.O_WRONLY|os.O_APPEND, 0666)
						defer out0.Close()
						out0.WriteString("\n")

					case "put":
						defer os.Rename(p0, p)
					}
					return e
				}

				wd, e := os.Getwd()
				m.Assert(e)
				trip := len(wd) + 1

				rg, e := regexp.Compile(m.Option("dir_reg"))

				m.Confm("pwd", func(index int, value string) bool {
					// p := path.Join(value, m.Option("dir_root"), arg[0])
					p := path.Join(value, arg[0])
					if s, e := os.Stat(p); e == nil {
						if s.IsDir() {
							dir(m, p, 0, kit.Right(m.Has("dir_deep")), m.Confx("dir_type"), trip, rg,
								strings.Split(m.Confx("dir_fields", strings.Join(arg[1:], " ")), " "),
								m.Conf("time", "format"))
						} else {
							if s.Size() < int64(m.Confi("buf_size")) {
								p0 := p + ".tmp0"
								f, e := os.Open(p0)
								if e != nil {
									f, e = os.Open(p)
									m.Log("info", "open %v", p)
								} else {
									p = p0
									m.Log("info", "open %v", p0)
								}
								for bio := bufio.NewScanner(f); bio.Scan(); {
									m.Echo(bio.Text())
								}
								m.Append("file", p)
								m.Append("size", s.Size())
								m.Append("time", s.ModTime().Format(m.Conf("time", "format")))
							} else {
								m.Append("directory", p)
							}
						}
						return true
					}
					return false
				})

				if m.Has("dir_sort") {
					m.Sort(m.Meta["dir_sort"][0], m.Meta["dir_sort"][1:]...)
				}

				if len(m.Meta["append"]) == 1 {
					for _, v := range m.Meta[m.Meta["append"][0]] {
						m.Echo(v).Echo(" ")
					}
				} else {
					if !m.Appends("file") {
						m.Table()
					}
				}
				return
			}},
		"git": &ctx.Command{Name: "git sum", Help: "版本控制", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) > 0 && arg[0] == "sum" || len(arg) > 1 && arg[1] == "sum" {
				args := []string{"log"}
				if len(arg) > 1 && arg[1] == "sum" && arg[0] != "" {
					args = append(args, "-C", arg[0])
				}
				args = append(args, "--reverse", "--shortstat", "--pretty=commit: %ad", "--date=format:%Y-%m-%d %H:%M")
				if out, e := exec.Command("git", args...).CombinedOutput(); m.Assert(e) {
					for _, v := range strings.Split(string(out), "commit: ") {
						if l := strings.Split(v, "\n"); len(l) > 2 {
							fs := strings.Split(strings.TrimSpace(l[2]), ", ")
							hs := strings.Split(l[0], " ")
							m.Add("append", "date", hs[0])

							if adds := strings.Split(fs[1], " "); len(fs) > 2 {
								dels := strings.Split(fs[2], " ")
								m.Add("append", "adds", adds[0])
								m.Add("append", "dels", dels[0])
							} else if adds[1] == "insertions(+)" {
								m.Add("append", "adds", adds[0])
								m.Add("append", "dels", "0")
							} else {
								m.Add("append", "adds", "0")
								m.Add("append", "dels", adds[0])
							}
							m.Add("append", "time", hs[1])
						} else if len(l[0]) > 0 {
							hs := strings.Split(l[0], " ")
							m.Add("append", "date", hs[0])
							m.Add("append", "adds", 0)
							m.Add("append", "dels", 0)
							m.Add("append", "time", hs[1])
						}
					}
					m.Table()
				}
				return
			}

			if len(arg) == 0 {
				m.Cmdy("nfs.config", "git")
				return
			}

			if p := m.Cmdx("nfs.path", arg[0]); p != "" && !m.Confs("git", arg[0]) {
				m.Option("git_dir", p)
				arg = arg[1:]
			} else {
				wd, e := os.Getwd()
				m.Assert(e)
				m.Option("git_dir", wd)
			}

			cmds := []string{}
			if v := m.Confv("git", []string{arg[0], "cmds"}); v != nil {
				cmds = append(cmds, kit.Trans(v)...)
			} else {
				cmds = append(cmds, arg[0])
			}

			for _, cmd := range cmds {
				args := append([]string{}, kit.Trans(m.Confv("git", "args"))...)
				if v := m.Confv("git", []string{cmd, "args"}); v != nil {
					args = append(args, kit.Trans(v)...)
				} else {
					args = append(args, cmd)
				}
				args = append(args, arg[1:]...)

				for i, _ := range args {
					v := m.Parse(args[i])
					if v == args[i] || v == "" {
						args[i] = kit.Select(args[i], m.Conf("git", []string{"trans", args[i]}))
					} else {
						args[i] = v
					}
				}

				m.Cmd("cli.system", "git", args).Echo("\n\n").CopyTo(m)
			}
			return
		}},

		"temp": &ctx.Command{Name: "temp data", Help: "查找文件路径", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if f, p, e := kit.Create(path.Join(m.Conf("dir", "temp"), kit.Hashs("uniq"))); m.Assert(e) {
				defer f.Close()

				for _, v := range arg {
					f.WriteString(v)
				}
				m.Echo(p)
			}
			return
		}},
		"hash": &ctx.Command{Name: "hash file", Help: "文件哈希", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			p := m.Append("name", m.Cmdx("nfs.path", arg[0]))
			if s, e := os.Stat(p); e == nil && !s.IsDir() {
				m.Append("size", s.Size())
				if f, e := os.Open(p); e == nil {
					defer f.Close()

					md := md5.New()
					io.Copy(md, f)
					h := md.Sum(nil)
					m.Echo(hex.EncodeToString(h[:]))
				}
			}
			return
		}},
		"copy": &ctx.Command{Name: "copy to from", Help: "查找文件路径", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if to, p, e := kit.Create(arg[0]); m.Assert(e) {
				defer to.Close()

				for _, from := range arg[1:] {
					if f, e := os.Open(m.Cmdx("nfs.path", from)); e == nil {
						defer f.Close()

						if n, e := io.Copy(to, f); m.Assert(e) {
							m.Log("info", "copy %d to %s from %s", n, p, from)
						}
					}
				}
				m.Echo(p)
			}
			return
		}},
		"grep": &ctx.Command{Name: "grep", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			hold := false
			if len(arg) > 0 {
				switch arg[0] {
				case "add":
					m.Confv("grep", "list.-2", map[string]interface{}{
						"pos": 0, "offset": 0, "file": arg[1],
					})
					return

				case "head":
					m.Confm("grep", "list", func(index int, value map[string]interface{}) {
						if len(arg) == 1 {
							value["offset"] = 0
							value["pos"] = 0
						}
					})
					return

				case "tail":
					m.Confm("grep", "list", func(index int, value map[string]interface{}) {
						if len(arg) == 1 {
							value["pos"] = -1
							value["offset"] = 0
						}
					})
					return

				case "hold":
					hold, arg = true, arg[1:]
				}
			}

			m.Confm("grep", "list", func(index int, value map[string]interface{}) {
				f, e := os.Open(kit.Format(value["file"]))
				if e != nil {
					m.Log("warn", "%v", e)
					return
				}
				defer f.Close()

				// s, e := f.Stat()
				// m.Assert(e)

				begin, e := f.Seek(int64(kit.Int(value["pos"])), 0)
				if kit.Int(value["pos"]) == -1 {
					begin, e = f.Seek(0, 2)
				}
				m.Assert(e)

				n := 0
				offset := kit.Int(value["offset"])
				bio := bufio.NewScanner(f)
				for i := 0; i < m.Optioni("table.limit") && bio.Scan(); i++ {
					text := bio.Text()
					if len(arg) == 0 || strings.Contains(text, arg[0]) {
						m.Add("append", "index", index)
						m.Add("append", "file", path.Base(kit.Format(value["file"])))
						// m.Add("append", "pos",begin+int64(n))
						// m.Add("append", "len",len(text))
						// m.Add("append", "end",s.Size())
						m.Add("append", "line", offset)
						m.Add("append", "text", text)
					} else {
						i--
					}
					n += len(text) + 1
					offset += 1
				}

				if !hold {
					value["offset"] = offset
					value["pos"] = begin + int64(n)
				}
			})
			m.Table()
			return
		}},
		"trash": &ctx.Command{Name: "trash file", Help: "查找文件路径", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) == 0 {
				m.Cmdy("nfs.dir", m.Conf("dir", "trash"))
				return
			}
			m.Assert(os.Mkdir(m.Conf("dir", "trash"), 0777))
			m.Assert(os.Rename(arg[0], path.Join(m.Conf("dir", "trash"), arg[0])))
			return
		}},

		"load": &ctx.Command{Name: "load file [buf_size [pos]]", Help: "加载文件, buf_size: 加载大小, pos: 加载位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if p, f, e := open(m, arg[0]); e == nil {
				defer f.Close()

				size := kit.Int(m.Confx("buf_size", arg, 1))
				if size == -1 {
					if s, e := f.Stat(); m.Assert(e) {
						size = int(s.Size())
					}
				}
				buf := make([]byte, size)

				pos := kit.Int(kit.Select("0", arg, 2))
				if l, e := f.ReadAt(buf, int64(pos)); e == io.EOF || m.Assert(e) {
					m.Log("info", "load %s %d %d", p, l, pos)
					m.Echo(string(buf[:l]))
				}
			}
			return
		}},
		"save": &ctx.Command{Name: "save file string...", Help: "保存文件, file: 保存的文件, string: 保存的内容", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) == 1 && m.Has("data") {
				arg = append(arg, m.Option("data"))
			}

			if f, p, e := kit.Create(arg[0]); m.Assert(e) {
				defer f.Close()

				for _, v := range arg[1:] {
					if n, e := f.WriteString(v); m.Assert(e) {
						m.Log("info", "save %s %d", p, n)
					}
				}
				m.Echo(p)
			}
			return
		}},
		"import": &ctx.Command{Name: "import filename [index]", Help: "导入数据", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			p, f, e := open(m, arg[0])
			m.Assert(e)
			defer f.Close()

			s, e := f.Stat()
			m.Option("filepath", p)
			m.Option("filename", s.Name())
			m.Option("filesize", s.Size())
			m.Option("filetime", s.ModTime().Format(m.Conf("time", "format")))

			switch {
			case strings.HasSuffix(arg[0], ".json"):
				var data interface{}
				json.NewDecoder(f).Decode(&data)
				m.Put("option", "filedata", data).Cmdy("ctx.trans", "filedata", arg[1:]).CopyTo(m)

			case strings.HasSuffix(arg[0], ".csv"):
				r := csv.NewReader(f)
				if l, e := r.Read(); m.Assert(e) {
					m.Meta["append"] = l

					for l, e = r.Read(); e == nil; l, e = r.Read() {
						for i, v := range l {
							m.Add("append", m.Meta["append"][i], v)
						}
					}
				}
				m.Table()

			default:
				if b, e := ioutil.ReadAll(f); m.Assert(e) {
					m.Echo(string(b))
				}
			}
			return
		}},
		"export": &ctx.Command{Name: "export filename", Help: "导出数据", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			tp := false
			if len(arg) > 0 && arg[0] == "time" {
				tp, arg = true, arg[1:]
			}

			f, p, e := kit.Create(kit.Select(arg[0], m.Format(arg[0]), tp))
			m.Assert(e)
			defer f.Close()

			data := m.Optionv(kit.Select("data", arg, 1))
			if len(arg) > 0 && arg[0] == "all" {
				data, arg = m.Meta, arg[1:]
			}

			switch {
			case strings.HasSuffix(arg[0], ".json"):
				if data == nil && len(m.Meta["append"]) > 0 {
					lines := []interface{}{}
					nrow := len(m.Meta[m.Meta["append"][0]])
					for i := 0; i < nrow; i++ {
						line := map[string]interface{}{}
						for _, k := range m.Meta["append"] {
							line[k] = m.Meta[k][i]
						}

						lines = append(lines, line)
						data = lines
					}
				}

				en := json.NewEncoder(f)
				en.SetIndent("", "  ")
				en.Encode(data)

			case strings.HasSuffix(arg[0], ".csv"):
				fields := m.Meta["append"]
				if m.Options("fields") {
					fields = m.Meta["fields"]
				}

				if data == nil && len(m.Meta["append"]) > 0 {
					lines := []interface{}{}
					nrow := len(m.Meta[m.Meta["append"][0]])
					for i := 0; i < nrow; i++ {
						line := []string{}
						for _, k := range fields {
							line = append(line, m.Meta[k][i])
						}
						lines = append(lines, line)
						data = lines
					}
				}

				if data, ok := data.([]interface{}); ok {
					w := csv.NewWriter(f)
					w.Write(fields)
					for _, v := range data {
						w.Write(kit.Trans(v))
					}
					w.Flush()
				}

			case strings.HasSuffix(arg[0], ".png"):
				if data == nil {
					data = strings.Join(arg[1:], "")
				}

				if qr, e := qrcode.New(kit.Format(data), qrcode.Medium); m.Assert(e) {
					m.Assert(qr.Write(256, f))
				}

			default:
				f.WriteString(kit.Format(m.Meta["result"]))
			}

			m.Set("append").Add("append", "directory", p)
			m.Set("result").Echo(p)
			return
		}},
		"json": &ctx.Command{Name: "json str", Help: "导入数据", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			var data interface{}
			m.Assert(json.Unmarshal([]byte(arg[0]), &data))
			if b, e := json.MarshalIndent(data, "", "  "); m.Assert(e) {
				m.Echo(string(b))
			}
			return
		}},

		"open": &ctx.Command{Name: "open file", Help: "打开文件, file: 文件名", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if m.Has("io") {
			} else if p, f, e := open(m, arg[0], os.O_RDWR|os.O_CREATE); e == nil {
				m.Put("option", "in", f).Put("option", "out", f)
				arg[0] = p
			} else {
				return nil
			}

			m.Start(fmt.Sprintf("file%d", m.Capi("nfile")), fmt.Sprintf("file %s", arg[0]), "open", arg[0])
			m.Append("bio.ctx1", m.Cap("module"))
			m.Echo(m.Cap("module"))
			return
		}},
		"read": &ctx.Command{Name: "read [buf_size [pos]]", Help: "读取文件, buf_size: 读取大小, pos: 读取位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.in != nil {
				if len(arg) > 1 {
					m.Cap("pos", arg[1])
				}

				buf := make([]byte, kit.Int(m.Confx("buf_size", arg, 0)))
				if n, e := nfs.in.ReadAt(buf, int64(m.Capi("pos"))); e == io.EOF || m.Assert(e) {
					m.Capi("nread", n)
					if m.Capi("pos", n); n == 0 {
						m.Cap("pos", "0")
					}
				}
				m.Echo(string(buf))
			}
			return
		}},
		"write": &ctx.Command{Name: "write string [pos]", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.out != nil {
				if len(arg) > 1 {
					m.Cap("pos", arg[1])
				}

				if len(arg[0]) == 0 {
					m.Assert(nfs.out.Truncate(int64(m.Capi("pos"))))
					m.Cap("size", m.Cap("pos"))
					m.Cap("pos", "0")
				} else {
					n, e := nfs.out.WriteAt([]byte(arg[0]), int64(m.Capi("pos")))
					if m.Capi("nwrite", n); m.Assert(e) && m.Capi("pos", n) > m.Capi("size") {
						m.Cap("size", m.Cap("pos"))
					}
					nfs.out.Sync()
				}

				m.Echo(m.Cap("pos"))
			}
			return
		}},

		"printf": &ctx.Command{Name: "printf arg", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) {
				nfs.print(arg...)
			}
			return
		}},
		"prompt": &ctx.Command{Name: "prompt arg", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) {
				nfs.prompt(arg)
			}
			return
		}},
		"term": &ctx.Command{Name: "term action args...", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) {
				nfs.Term(m, arg[0], arg[1:])
			}
			return
		}},
		"action": &ctx.Command{Name: "action cmd", Help: "", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) {
				msg := m.Cmd("cli.source", arg)
				// nfs.print(msg.Conf("prompt"), arg, "\n")
				nfs.print(msg.Meta["result"]...)
			}
			return
		}},

		"source": &ctx.Command{Name: "source [script|stdio|snippet]", Help: "解析脚本, script: 脚本文件, stdio: 命令终端, snippet: 代码片段", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) == 0 {
				m.Cmdy("dir", "src", "time", "line", "path", "dir_deep", "dir_reg", ".*\\.(sh|shy|py)$")
				return
			}

			m.Optionv("bio.args", arg)
			if help := fmt.Sprintf("scan %s", arg[0]); arg[0] == "stdio" {
				m.Put("option", "bio.in", os.Stdin).Put("option", "bio.out", os.Stdout).Start(arg[0], help, "scan", arg[0])
				m.Wait()

			} else if p, f, e := open(m, arg[0]); e == nil {
				m.Put("option", "bio.in", f).Start(fmt.Sprintf("file%d", m.Capi("nfile", 1)), help, "scan", p)
				m.Wait()

			} else {
				m.Cmdy("yac.parse", strings.Join(arg, " ")+"\n")
			}
			return
		}},
		"arguments": &ctx.Command{Name: "arguments", Help: "脚本参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			args := kit.Trans(m.Optionv("bio.args"))
			if len(arg) == 0 {
				m.Set("result", args)
			} else {
				m.Echo(kit.Select("", args, kit.Int(arg[0])))
			}
			return
		}},

		"remote": &ctx.Command{Name: "remote listen|dial args...", Help: "启动文件服务, args: 参考tcp模块, listen命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if _, ok := m.Target().Server.(*NFS); m.Assert(ok) { //{{{
				m.Sess("tcp").Call(func(sub *ctx.Message) *ctx.Message {
					if sub.Has("node.port") {
						return sub
					}
					sub.Sess("ms_source", sub)
					sub.Sess("ms_target", m.Source())
					sub.Start(fmt.Sprintf("file%d", m.Capi("nfile", 1)), "远程文件")
					return sub
				}, arg)
			}
			return
		}},
		"send": &ctx.Command{Name: "send [file] args...", Help: "连接文件服务, args: 参考tcp模块, dial命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.io != nil {
				nfs.send <- m.Set("detail", arg)
			}
			return
		}},
	},
}

func init() {
	ctx.Index.Register(Index, &NFS{Context: Index})
}
