Volcanos("onengine", { river: {
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
