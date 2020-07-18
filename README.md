# contexts

contexts通过模块化、集群化、自动化，可以在各种设备上，快速组建起云计算服务。

## install

### install by binary
``` sh
$ curl https://shylinux.com/publish/ice.sh | sh
$ ctx_self=http://:9020 bin/ice.sh start serve
ice>
```
open http://localhost:9020

### install by docker
```
$ docker run -p 9020:9020 -it shylinux/contexts
ice>
```
open http://localhost:9020

### install by source
```
$ git clone https://github.com/shylinux/contexts
$ cd contexts && source etc/miss.sh
$ make
$ ish_miss_serve
ice>
```
open http://localhost:9020

