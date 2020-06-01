# contexts

contexts通过模块化、自动化，可以在各种设备上，快速搭建起云计算服务。

## 三端分立，各司其职
|英文名   |技术栈|中文名 |谐称     |缩写|简介|
|---------|------|-------|---------|----|----|
|volcanos |前端  |火山架 |我看不行 |FMS |a fieldset manager system
|icebergs |后端  |冰山架 |挨撕不可 |CMS |a cluster manager system
|intshell |终端  |神农架 |整个脚本 |PMS |a plugin manager system

## 安装部署

```sh
git clone https://github.com/shylinux/contexts
cd contexts
make
source etc/miss.sh
ish_miss_serve
curl localhost:9020
```
