# Cobra

```bash
# 库
go get -u github.com/spf13/cobra@latest

# 命令行工具
go install github.com/spf13/cobra-cli@latest

# 初始化
# 必须先执行 `go mod init <MODNAME>`, 初始化一个GO模块，然后才能执行以下命令
cobra-cli init

# 添加子命令
cobra-cli add serve

# 运行
go build -o my-cli
./my-cli
./my-cli serve
```

# 参数
```bash
go doc pflag.FlagSet


# 交互式组件
go get -u github.com/AlecAivazis/survey/v2

```

# 其他

https://typonotes.com/tags/cobra/

https://typonotes.com/posts/2023/03/02/cobra-book/
