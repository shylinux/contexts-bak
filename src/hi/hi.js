Volcanos("onimport", {help: "导入数据", list: [], _init: function(can, msg, cb, target) {
	can.onappend.parse(can, [
		{type: "head", style: {height: 320}, list: [
			{type: "head", list: ["基础配置", "高级配置", "导航配置"]},
			{type: "left", list: [{type: "tabs", list: [
				{name: "地址", list: [
					{name: "山东省", list: ["邹城市", "济南市", "青岛市"]},
					{name: "江苏省", list: [
						{type: "plugin", name: "徐州市", index: "cli.runtime"},
						{type: "plugin", name: "南京市", index: "cli.qrcode"},
					]},
					{name: "北京市", list: ["海淀区", "朝阳区", "西城区", "东城区"]},
					{name: "广东省", list: [
						{name: "深圳市", list: ["宝安区", "南山区", "福田区", "龙华区"]}, "茂名市",
					]},
				]},		
				{name: "组件", list: ["2", "2", "3"]},		
				{name: "数据"},		
			]}]},
			{type: "main"},
			"foot",
		]},
		{type: "head", style: {height: 320}, list: [
			"head",
			{type: "left", list: [{type: "tabs", style: "left", list: [
				{name: "布局", list: ["空中楼阁1", "空中楼阁2", "空中楼阁3", "空中楼阁4", "空中楼阁5"]},		
				{name: "列表", list: ["分页器1", "分页器2", "分页器3", "分页器4", "分页器"]},		
				{name: "表单"},		
				{name: "按钮"},		
				{name: "输入"},		
			]}]},
			{type: "main"},
			"foot",
		]},
		{type: "head", style: {height: 200}, list: [
			"head",
			{type: "left", list: [{type: "tabs", style: "void", list: [
				{name: "布局", list: ["空中楼阁1", "空中楼阁2", "空中楼阁3", "空中楼阁4"]},		
				{name: "列表", list: ["分页器1", "分页器2", "分页器3"]},		
				{name: "表单"},		
				{name: "按钮"},		
				{name: "输入"},		
			]}]},
			{type: "main"},
			"foot",
		]}
	], target)
}})
