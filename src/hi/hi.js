Volcanos("onimport", {help: "导入数据", list: [], _init: function(can, msg, cb, target) {
	msg.Echo("hello world")	
	can.onmotion.clear(can)
	can.onappend.table(can, msg)
	can.onappend.board(can, msg)
	can.base.isFunc(cb) && cb(msg)
}}, [""])
