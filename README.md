# contexts

contexts 通过集群化、模块化、自动化的方式，用一个 20M 的程序文件， 就可以在各种设备上，一键搭起完整的云计算与云研发服务。

## 安装项目

### 镜像方式
```sh
ctx_temp=$(mktemp); curl -o $ctx_temp -fsSL https://shylinux.com; source $ctx_temp binary
```

### 源码方式
```sh
ctx_temp=$(mktemp); curl -fsSL https://shylinux.com -o $ctx_temp; source $ctx_temp source
```

## 使用项目
如果是本地启动，直接免登录打开网页，http://localhost:9020 ，
如果是远程或容器启动，输入相应地址后，输入默认的 username: root password: root 即可登录

