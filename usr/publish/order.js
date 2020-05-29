Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "IDE": {name: "IDE", action: [
            {name: "editor", help: "编辑器", inputs: [
                {type: "text", name: "path", value: "hi.go", action: "auto"},
                {type: "button", name: "打开", action: "auto"},
                {type: "button", name: "返回"},
            ], index: "web.wiki.inner", feature: {display: "/plugin/inner.js"}},
            {name: "golang", help: "编译器", inputs: [
                {type: "text", name: "path", value: "go"},
                {type: "text", name: "path", value: "run"},
                {type: "text", name: "path", value: "hi.go"},
                {type: "button", name: "执行"},
            ], index: "cli.system", feature: {}},
            {name: "story", help: "故事会", inputs: [
                {type: "text", name: "story", value: "hi.go", action: "auto"},
                {type: "text", name: "key", action: "auto"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回", action: "auto"},
            ], index: "web.story", feature: {}},
        ]},

        "main": {name: "main", action: [
            {name: "hello", help: "world", inputs: [
                {type: "text", name: "path", value: "自然/编程/hi.shy", action: "auto"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回"},
            ], index: "web.wiki.word", feature: {display: "/plugin/local/wiki/word.js"}},
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
