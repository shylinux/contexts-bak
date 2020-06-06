Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "main": {name: "main", action: [
            {name: "IDE", help: "集成开发环境", inputs: [
                {type: "text", name: "path", value: "tmp", action: "auto"},
                {type: "text", name: "name", value: "hi.qrc", action: "auto"},
                {type: "text", name: "key", value: "", action: "auto"},
                {type: "button", name: "打开", action: "auto"},
                {type: "button", name: "返回"},
                {type: "button", name: "历史"},
                {type: "button", name: "复盘"},
            ], index: "web.wiki.inner", feature: {display: "/plugin/inner.js", style: "editor"}},
            {name: "word", help: "文档", inputs: [
                {type: "text", name: "path", value: "自然/编程/hi.shy", action: "auto"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js", style: "word"}},
        ]},
        "task": {name: "task", index: [
            "web.team.task",
            "web.team.plan",
        ]},
        "hello": {name: "应用1", action: [
            {name: "route", help: "路由", inputs: [
                {type: "text", name: "name", value: "", action: "auto"},
                {type: "text", name: "cmd"},
                {type: "button", name: "查看", action: "auto"},
            ], index: "web.route"},
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
