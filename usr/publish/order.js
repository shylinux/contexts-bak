Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "main": {name: "main", index: [
            "web.wiki.inner",
            "web.code.git.status",
        ]},
        "task": {name: "task", index: [
            "web.team.plan",
            "web.wiki.word",
        ]},
        "relay": {name: "relay", index: [
            "aaa.totp.get",
            "web.route",
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
