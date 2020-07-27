Volcanos("onengine", { river: {
    "serivce": {name: "运营群", storm: {
        "wx": {name: "wx",  action: [
            {name: "微信公众号", help: "wx", inputs: [
                {type: "text", name: "path", value: "icebergs/misc/wx/wx.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "mp": {name: "mp",  action: [
            {name: "微信小程序", help: "mp", inputs: [
                {type: "text", name: "path", value: "icebergs/misc/mp/mp.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "lark": {name: "lark",  action: [
            {name: "飞书机器人", help: "lark", inputs: [
                {type: "text", name: "path", value: "icebergs/misc/lark/lark.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js", style: "word"}},
        ]},
        "share": {name: "share",  action: [
            {name: "云境科技", help: "shylinux/contexts", inputs: [
                {type: "text", name: "path", value: "learning/speak/20200724.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js", style: "word"}},
        ]},
        "company": {name: "company",  action: [
            {name: "公司", help: "company", inputs: [
                {type: "text", name: "path", value: "learning/社会/管理/company.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
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
        "chrome": {name: "chrome", index: [
            "web.code.chrome.chrome",
            "web.code.chrome.bookmark",
        ]},
        "english": {name: "english",  action: [
            {name: "english", help: "英汉", inputs: [
                {type: "text", name: "word", value: "hi"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.alpha.find", feature: {}},
            {name: "chinese", help: "汉英", inputs: [
                {type: "text", name: "word", value: "你好"},
                {type: "text", name: "method", value: "line"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.alpha.find", feature: {}},
            {name: "wubi", help: "五笔", inputs: [
                {type: "text", name: "word", value: "wqvb"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.input.find", feature: {}},
            {name: "wubi", help: "五笔", inputs: [
                {type: "text", name: "word", value: "你好"},
                {type: "text", name: "method", value: "line"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.input.find", feature: {}},
        ]},
        "context": {name: "context",  action: [
            {name: "context", help: "编程", inputs: [
                {type: "text", name: "path", value: "learning/自然/编程/hi.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
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
        "html": {name: "html",  action: [
            {name: "spide", help: "爬虫", inputs: [
                {type: "text", name: "name", value: "icebergs"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.git.spide", feature: {display: "/plugin/story/spide.js"}},
            {name: "trend", help: "趋势", inputs: [
                {type: "text", name: "name", value: "icebergs"},
                {type: "text", name: "begin_time", value: "@date"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.git.trend", feature: {display: "/plugin/story/trend.js"}},
            {name: "draw", help: "绘图", inputs: [
                {type: "text", name: "path", value: "hi.svg"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.draw.draw", feature: {display: "/plugin/local/wiki/draw.js"}},
            {name: "HTML5", help: "网页", inputs: [
                {type: "text", name: "path", value: "icebergs/misc/chrome/chrome.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "nginx": {name: "nginx",  action: [
            {name: "nginx", help: "代理", inputs: [
                {type: "text", name: "path", value: "nginx-story/src/main.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "golang": {name: "golang",  action: [
            {name: "golang", help: "编程", inputs: [
                {type: "text", name: "path", value: "golang-story/src/main.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "redis": {name: "redis",  action: [
            {name: "redis", help: "缓存", inputs: [
                {type: "text", name: "path", value: "redis-story/src/main.shy", action: "auto"},
                {type: "button", name: "查看", value: "auto"},
                {type: "button", name: "返回"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
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
            {name: "操作系统", help: "os", inputs: [
                {type: "text", name: "path", value: "learning/自然/编程/system.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
    }},
}, })
