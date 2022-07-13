Volcanos("onimport", {help: "导入数据", _init: function(can, msg) {
	can.onaction.show(can, msg)
	msg.Echo("hello world")
	msg.Echo("hello world")
	msg.Dump(can)
}})
Volcanos("onaction", {help: "操作数据", show: function(can, msg) {
	msg.Push("hi", "hello")
	msg.Echo("hello world")
}})
