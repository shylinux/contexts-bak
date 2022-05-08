Volcanos("onimport", {help: "导入数据", list: [], _init: function(can, msg, cb, target) {
	can.onmotion.hidden(can, can._legend)
	can.onmotion.hidden(can, can._option)
	can.onmotion.hidden(can, can._status)

	can.ConfHeight(can.ConfHeight()+can.Conf(html.MARGIN_Y)-(can.user.isWindows? 17: 0))

	can.onengine.listen(can, "menu", function(msg) { can.user.toast(can, msg.Option(html.ITEM)) })
	can.onengine.listen(can, "高级配置", function(msg) { can.user.toast(can, msg.Option(html.ITEM)) })
	can.onengine.listen(can, "h1", function(msg) { can.user.toast(can, "h1") })
	can.require(["hi.css"])

	can.onappend.parse(can, can.onappend._parse(can, `
head
	基础配置 type menu
	高级配置 type menu
		h1
		h2
		h3
		h4
	导航配置 type menu
left
	tabs
		地址
			山东省
				邹城市
				济南市
				青岛市
			江苏省
				徐州市 index cli.runtime
				南京市 index web.code.git.repos action auto
			北京市
				海淀区
				朝阳区
				西城区
				东城区
			广东省
				深圳市
					宝安区
					南山区
					福田区
					龙华区
				茂名市
		组件
			列表 index cli.runtime
			卡片 index cli.qrcode
		数据
main
`), target)
}}, [""])
