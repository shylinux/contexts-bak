Volcanos("onengine", { river: {
    "serivce": {name: "运营群", storm: {
        "wx": {name: "wx",  action: [
            {name: "微信公众号", help: "wx", index: "web.wiki.word", args: ["icebergs/misc/wx/wx.shy"]},
        ]},
        "mp": {name: "mp",  action: [
            {name: "微信小程序", help: "mp", index: "web.wiki.word", args: ["icebergs/misc/mp/mp.shy"]},
        ]},
        "lark": {name: "lark",  action: [
            {name: "飞书机器人", help: "lark", index: "web.wiki.word", args: ["icebergs/misc/lark/lark.shy"]},
        ]},
        "share": {name: "share",  action: [
            {name: "系统上下文", help: "shylinux/contexts", index: "web.wiki.word", args: ["learning/社会/管理/20200724.shy"]},
        ]},
    }},
    "product": {name: "产品群", storm: {
        "office": {name: "office", index: [
            "web.wiki.draw.draw",
            "web.team.plan",
            "web.wiki.word",
            "web.wiki.data",
            "web.wiki.feel",
        ]},
        "english": {name: "english",  action: [
            {name: "english", help: "英汉", index: "web.wiki.alpha.find", args: ["hi"]},
            {name: "chinese", help: "汉英", index: "web.wiki.alpha.find", args: ["你好", "line"]},
            {name: "wubi", help: "五笔", index: "web.code.input.find", args: ["wqvb"]},
            {name: "wubi", help: "五笔", index: "web.code.input.find", args: ["你好", "line"]},
        ]},
        "context": {name: "context",  action: [
            {name: "context", help: "编程", index: "web.wiki.word", args: ["learning/社会/管理/context.shy"]},
        ]},
        "learning": {name: "learning",  action: [
            {name: "study", help: "学习", index: "web.wiki.word", args: ["learning/study.shy"]},
            {name: "tmux", help: "粘贴", index: "web.code.tmux.text"},
            {name: "golang", help: "编程", index: "web.wiki.word", args: ["golang-story/src/main.shy"]},
        ]},
    }},
    "project": {name: "研发群", storm: {
        "cli": {name: "cli",  action: [
            {name: "tmux", help: "命令行", index: "web.wiki.word", args: ["icebergs/misc/tmux/tmux.shy"]},
            {name: "git", help: "代码库", index: "web.wiki.word", args: ["icebergs/misc/git/git.shy"]},
            {name: "vim", help: "编辑器", index: "web.wiki.word", args: ["icebergs/misc/vim/vim.shy"]},
            {name: "zsh", help: "命令行", index: "web.wiki.word", args: ["icebergs/misc/zsh/zsh.shy"]},
        ]},
        "web": {name: "web",  action: [
            {name: "HTML5", help: "浏览器", index: "web.wiki.word", args: ["icebergs/misc/chrome/chrome.shy"]},
        ]},
        "linux": {name: "linux",  action: [
            {name: "linux", help: "系统", index: "web.wiki.word", args: ["linux-story/src/main.shy"]},
        ]},
        "nginx": {name: "nginx",  action: [
            {name: "nginx", help: "代理", index: "web.wiki.word", args: ["nginx-story/src/main.shy"]},
        ]},
        "golang": {name: "golang",  action: [
            {name: "golang", help: "编程", index: "web.wiki.word", args: ["golang-story/src/main.shy"]},
        ]},
        "redis": {name: "redis",  action: [
            {name: "redis", help: "缓存", index: "web.wiki.word", args: ["redis-story/src/main.shy"]},
        ]},
        "mysql": {name: "mysql",  action: [
            {name: "mysql", help: "数据", index: "web.wiki.word", args: ["mysql-story/src/main.shy"]},
        ]},
        "context": {name: "context",  action: [
            {name: "think", help: "智库", index: "web.wiki.word", args: ["learning/"]},
            {name: "index", help: "索引", index: "web.wiki.word", args: ["learning/index.shy"]},
            {name: "context", help: "编程", index: "web.wiki.word", args: ["learning/自然/编程/context.shy"]},
        ]},
    }},
    "profile": {name: "测试群", storm: {
        "pprof": {name: "pprof", index: [
            "web.code.pprof",
            "web.code.bench",
            "web.favor",
            "web.cache",
            "web.share",
        ]},
    }},
    "operate": {name: "运维群", storm: {
        "docker": {name: "docker", index: [
            "web.code.docker.image",
            "web.code.docker.container",
            "web.code.docker.command",
        ]},
        "relay": {name: "relay", index: [
            "aaa.totp.get",
            "web.route",
            "web.space",
            "web.dream",
            "web.code.docker.container",
            "web.code.tmux.session",
        ]},
        "os": {name: "os",  action: [
            {name: "操作系统", help: "os", index: "web.wiki.word", args: ["learning/自然/编程/system.shy"]},
        ]},
    }},
}, })
