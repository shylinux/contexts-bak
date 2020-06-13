Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "main": {name: "main", index: [
            "web.code.inner",
            "web.code.git.status",
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
        "trans": {name: "翻译", action: [
            {name: "english", help: "英文", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "button", name: "翻译"},
            ], index: "web.wiki.alpha.trans"},
            {name: "chinese", help: "中文", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "text", name: "method", value: "line"},
                {type: "button", name: "翻译"},
            ], index: "web.wiki.alpha.trans"},
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
