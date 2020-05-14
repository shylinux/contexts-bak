Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "main": {name: "main", action: [
            {name: "free", help: "内存", inputs: [
                {type: "text", name: "name", value: "free"},
                {type: "text", name: "value", value: "-h"},
                {type: "button", name: "查看"},
            ], index: "cli.system"},
            {name: "df", help: "磁盘", inputs: [
                {type: "text", name: "name", value: "df"},
                {type: "text", name: "value", value: "-h"},
                {type: "button", name: "查看"},
            ], index: "cli.system"},
            {name: "ps", help: "进程", inputs: [
                {type: "text", name: "name", value: "ps"},
                {type: "button", name: "查看"},
            ], index: "cli.system"},
            {name: "ls", help: "文件", inputs: [
                {type: "text", name: "name", value: "ls"},
                {type: "text", name: "value", value: "-l"},
                {type: "button", name: "查看"},
            ], index: "cli.system"},
            {name: "ifconfig", help: "网卡", index: "tcp.ifconfig"},
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
        "word": {name: "word", action: [
            {name: "trans", help: "词汇", inputs: [
                {type: "text", name: "word", value: "miss"},
                {type: "text", name: "method", value: ""},
                {type: "button", name: "翻译"},
            ], group: "web.wiki.alpha", index: "trans"},
        ]},
    }},
}, })
