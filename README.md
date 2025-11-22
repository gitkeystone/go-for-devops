# go-for-devops

```bash
go mod init WithCancel
go mod tidy
nano main.go
gofmt // alias gofmt='go fmt . && goimports -w .'
go install xxx
go run main.go
go build main.go
./main

go doc builtin.error
go doc context.context
```

# 常用后端包
- Gin -- 写API永远不后悔的框架
- GORM -- 那个骂你但一直离不开的ORM
- zap -- 真·高性能日志
- Testify -- 测试不再痛苦
- Cobra -- 真正能扩展的 CLI 框架
- Viper -- 配置管理不再“咬人”
- go-redis -- 稳如老狗的缓存客户端
- Go Kit -- 真·微服务架构利器
- Prometheus Client -- 度量从未如此轻松
- Wire -- 依赖注入，不再用魔法也能优雅
