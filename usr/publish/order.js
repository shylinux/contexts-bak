Volcanos("onengine", { river: {
    "one": {name: "第一组", storm: {
        "hello": {name: "应用1", action: [
            {name: "some", help: "some", inputs: [
                {type: "text", name: "one"},
                {type: "button", name: "one"},
            ], engine: function(event, can, msg, pane, cmds, cb) {
                msg.Echo("hello world")
                typeof cb == "function" && cb(msg)
            }},
        ]},
        "world": {name: "应用2", action: [
            {name: "hello", help: "world", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "one"},
            ], group: "cli", index: "system"},
        ]},
        "word": {name: "word", action: [
            {name: "trans", help: "词汇", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "text", name: "method", value: ""},
                {type: "button", name: "翻译"},
            ], group: "web.wiki.alpha", index: "trans"},
        ]},
    }},
}, })
