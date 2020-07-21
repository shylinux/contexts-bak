Volcanos("onengine", { river: {
    "product": {name: "产品群", storm: {
        "lark": {name: "lark",  action: [
            {name: "聊天机器人", help: "lark", inputs: [
                {type: "text", name: "path", value: "icebergs/misc/lark/lark.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js", style: "word"}},
        ]},
        "share": {name: "share",  action: [
            {name: "云境", help: "shylinux/contexts", inputs: [
                {type: "text", name: "path", value: "learning/speak/20200724.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js", style: "word"}},
        ]},
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
        "context": {name: "context",  action: [
            {name: "context", help: "编程", inputs: [
                {type: "text", name: "path", value: "learning/自然/编程/hi.shy"},
                {type: "button", name: "查看", value: "auto"},
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
        "redis": {name: "redis",  action: [
            {name: "redis", help: "缓存", inputs: [
                {type: "text", name: "path", value: "usr/install/redis-5.0.4/src", action: "auto"},
                {type: "text", name: "file", value: "dict.h", action: "auto"},
                {type: "text", name: "line", value: "82", action: "auto"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.inner", feature: {display: "/plugin/local/code/inner.js", style: "editor"}},
        ]},
        "volcanos": {name: "volcanos",  action: [
            {name: "volcanos", help: "火山架", inputs: [
                {type: "text", name: "path", value: "usr/volcanos", action: "auto"},
                {type: "text", name: "file", value: "proto.js", action: "auto"},
                {type: "text", name: "line", value: "1", action: "auto"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.inner", feature: {display: "/plugin/local/code/inner.js", style: "editor"}},
        ]},
        "icebergs": {name: "icebergs",  action: [
            {name: "icebergs", help: "冰山", inputs: [
                {type: "text", name: "path", value: "usr/icebergs", action: "auto"},
                {type: "text", name: "file", value: "type.go", action: "auto"},
                {type: "text", name: "line", value: "1", action: "auto"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.inner", feature: {display: "/plugin/local/code/inner.js", style: "editor"}},
        ]},
        "intshell": {name: "intshell",  action: [
            {name: "icebergs", help: "冰山", inputs: [
                {type: "text", name: "path", value: "usr/intshell", action: "auto"},
                {type: "text", name: "file", value: "plug.sh", action: "auto"},
                {type: "text", name: "line", value: "1", action: "auto"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.code.inner", feature: {display: "/plugin/local/code/inner.js", style: "editor"}},
        ]},
    }},

    "main": {name: "main", storm: {
        "inner": {name: "inner", index: [
            "web.code.inner",
            "web.code.git.status",
            "web.code.git.total",
        ]},
        "pprof": {name: "pprof", index: [
            "web.code.pprof",
            "web.code.bench",
            "web.favor",
            "web.cache",
            "web.share",
        ]},
        "relay": {name: "relay", index: [
            "aaa.totp.get",
            "web.route",
            "web.space",
            "web.dream",
            "web.code.docker.container",
            "web.code.tmux.session",
        ]},

        "alpha": {name: "alpha", action: [
            {name: "wubi", help: "五笔", inputs: [
                {type: "text", name: "word", value: "shwq"},
                {type: "button", name: "查找"},
            ], index: "web.code.input.find"},

            {name: "wubi", help: "五笔", inputs: [
                {type: "text", name: "word", value: "想像"},
                {type: "text", name: "method", value: "line"},
                {type: "button", name: "查找"},
            ], index: "web.code.input.find"},


            {name: "english", help: "英文", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "button", name: "翻译"},
            ], index: "web.wiki.alpha.find"},
            {name: "english", help: "英文", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "button", name: "翻译"},
            ], index: "web.wiki.alpha.find2"},
            {name: "chinese", help: "中文", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "text", name: "method", value: "line"},
                {type: "button", name: "翻译"},
            ], index: "web.wiki.alpha.find"},
        ]},

        "office": {name: "office", index: [
            "web.team.plan",
            "web.wiki.feel",
            "web.wiki.data",
            "web.wiki.word",
            "web.wiki.draw.draw",
        ]},
        "docker": {name: "docker", index: [
            "web.code.docker.image",
            "web.code.docker.container",
            "web.code.docker.command",
        ]},
        "chrome": {name: "chrome", index: [
            "web.code.chrome.chrome",
            "web.code.chrome.bookmark",
        ]},
        "tmux": {name: "tmux", index: [
            "web.code.tmux.text",
            "web.code.tmux.buffer",
            "web.code.tmux.session",
        ]},

        "hello": {name: "hello", action: [
            {name: "some", help: "some", inputs: [
                {type: "text", name: "one"},
                {type: "button", name: "one"},
            ], engine: function(event, can, msg, pane, cmds, cb) {
                can.onappend.toast(can, "hello", "world");
                msg.Echo("hello world");
                typeof cb == "function" && cb(msg);
            }},
            {name: "hello", help: "world", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "one"},
            ], index: "cli.system"},
        ]},
    }},
}, })
