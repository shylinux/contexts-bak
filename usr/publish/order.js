Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "main": {name: "main", action: [
            {name: "route", help: "路由", index: "route"},
            {name: "status", help: "源码", index: "web.code.git.status"},
        ]},
        "word": {name: "word", action: [
            {name: "trans", help: "词汇", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "text", name: "method", value: ""},
                {type: "button", name: "翻译"},
            ], group: "web.wiki.alpha", index: "trans"},
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
        ]},
        "world": {name: "应用2", action: [
            {name: "hello", help: "world", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "one"},
            ], group: "cli", index: "system"},
        ]},
    }},
}, })
