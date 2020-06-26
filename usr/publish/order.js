Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "dream": {name: "dream", index: [
            "web.space",
            "web.dream",
            "web.code.docker.container",
            "web.code.tmux.session",
        ]},
        "main": {name: "main", index: [
            "web.code.inner",
            "web.code.git.status",
            "web.code.git.total",
        ]},
        "task": {name: "task", index: [
            "web.team.plan",
            "web.wiki.draw.draw",
            "web.wiki.word",
        ]},
        "relay": {name: "relay", index: [
            "aaa.totp.get",
            "web.route",
        ]},
        "性能": {name: "性能", index: [
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
        "tmux": {name: "tmux", index: [
            "web.code.tmux.text",
            "web.code.tmux.buffer",
            "web.code.tmux.session",
        ]},
        "wiki": {name: "wiki", index: [
            // "web.wiki.draw.draw",
            "web.wiki.feel",
            "web.wiki.data",
        ]},
        "trans": {name: "翻译", action: [
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
        "hello": {name: "应用1", action: [
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
