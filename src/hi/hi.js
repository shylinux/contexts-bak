Volcanos("onimport", {help: "导入数据", list: [], _init: function(can, msg, cb, target) {
	msg.Push("hi", "hello")
	msg.Push("he", "hello")
	msg.Echo("hello world!\n")
	msg.Echo("hello world!\n")
	msg.Echo("hello world!\n")

	can.onmotion.clear(can), can.onappend.table(can, msg), can.onappend.board(can, msg)
	can.base.isFunc(cb) && cb(msg)
}})
