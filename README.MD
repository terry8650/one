# 基于GoFrame 和layuiAdmin做的一个后台管理，我是初学者，大佬请忽略

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

系统环境：linux or windows 7以上，安装go环境，redis,mysql,nginx等

感谢项目(排名不分先后)

- https://layui.dev
- https://goframe.org
- github.com/casbin/casbin/

- github.com/dobyte/gf-casbin
- https://gitee.com/tiger1103/gfast
- github.com/goflyfox/gtoken

管理系统基于casbin的权限管理，前后端分离，前端使用layuiadmin，可以详细配置每个角色的接口和菜单权限，layuiadmin前端代码请拉取adminsingle的代码