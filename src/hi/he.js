Volcanos("onimport", {help: "导入数据", list: [
	{type: "button", name: "list"},
	{type: "button", name: "back"},
	{type: "button", name: "show"},
], _init: function(can, msg, cb, target) { 
	can.page.Modify(can, can._legend, "示例")
	msg.Push("hi", "hello")
	msg.Push("he", "world")
	msg.Echo("hello world")
	can.onmotion.clear(can)
	can.onappend.table(can, msg)
	can.onappend.board(can, msg)
	can.base.isFunc(cb) && cb(msg)
}})
Volcanos("onaction", {help: "操作数据", list: ["show", "some"],
	_trans: {show: "展示", some: "其它"},
	show: function(event, can) { var msg = can.request(event)
		msg.Push("value", 200)
		msg.Push("value", 300)
		msg.Push("value", 300)
		msg.Push("value", 300)
		msg.Push("value", 400)
		msg.Option("height", 400)
		can.onappend._output(can.sup, msg, "/plugin/story/pie.js")
	},
	some: function(event, can) { var msg = can.request(event)
		msg.Push("hi", "hello")
		msg.Push("hi", "hello")
		msg.Push("hi", "hello")
		msg.Push("hi", "hello")
		msg.Push("hi", "hello")
		can.onappend.table(can, msg)
	},
})
