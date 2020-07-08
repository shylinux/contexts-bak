Volcanos("onengine", { river: {
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
        "redis": {name: "redis",  action: [
            {name: "redis", help: "缓存", inputs: [
                {type: "text", name: "path", value: "redis-story/src/main.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "golang": {name: "golang",  action: [
            {name: "golang", help: "编程", inputs: [
                {type: "text", name: "path", value: "golang-story/src/main.shy"},
                {type: "button", name: "查看", value: "auto"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
        ]},
        "nginx": {name: "nginx",  action: [
            {name: "nginx", help: "代理", inputs: [
                {type: "text", name: "path", value: "nginx-story/src/main.shy"},
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
