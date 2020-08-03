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
            {name: "系统上下文", help: "shylinux/contexts", index: "web.wiki.word", args: ["learning/自然/编程/20200724.shy"]},
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
        "chrome": {name: "chrome", index: [
            "web.code.chrome.chrome",
            "web.code.chrome.bookmark",
        ]},
        "context": {name: "context",  action: [
            {name: "context", help: "编程", index: "web.wiki.word", args: ["learning/自然/编程/context.shy"]},
        ]},
    }},
    "project": {name: "研发群", storm: {
        "inner": {name: "inner", index: [
            "web.code.inner",
            "web.code.git.status",
            "web.code.git.total",
        ]},
        "relay": {name: "relay", index: [
            "aaa.totp.get",
            "web.route",
            "web.space",
            "web.dream",
            "web.code.docker.container",
            "web.code.tmux.session",
        ]},
        "tmux": {name: "tmux", index: [
            "web.code.tmux.text",
            "web.code.tmux.buffer",
            "web.code.tmux.session",
        ]},
        "vim": {name: "vim",  action: [
            {name: "git", help: "git", index: "web.wiki.word", args: ["icebergs/misc/git/git.shy"]},
            {name: "vim", help: "vim", index: "web.wiki.word", args: ["icebergs/misc/vim/vim.shy"]},
        ]},
        "html": {name: "html",  action: [
            {name: "draw", help: "绘图", index: "web.wiki.draw.draw", args: ["hi.svg"]},
            {name: "spide", help: "爬虫", index: "web.code.git.spide", args: ["icebergs"]},
            {name: "trend", help: "趋势", index: "web.code.git.trend", args: ["icebergs"]},
            {name: "HTML5", help: "网页", index: "web.wiki.word", args: ["icebergs/misc/chrome/chrome.shy"]},
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
        "context": {name: "context",  action: [
            {name: "think", help: "智库", index: "web.wiki.word", args: ["learning/"]},
            {name: "index", help: "索引", index: "web.wiki.word", args: ["learning/index.shy"]},
            {name: "context", help: "编程", index: "web.wiki.word", args: ["learning/自然/编程/hi.shy"]},
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
        "docker": {name: "docker", index: [
            "web.code.docker.image",
            "web.code.docker.container",
            "web.code.docker.command",
        ]},
    }},
    "operate": {name: "运维群", storm: {
        "os": {name: "os",  action: [
            {name: "操作系统", help: "os", index: "web.wiki.word", args: ["learning/自然/编程/system.shy"]},
        ]},
    }},
}, })
