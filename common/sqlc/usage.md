# use
1. compile - 此命令有助于检查 SQL 语法并报告任何类型错误
2. completion - 此命令用于为你的环境生成自动完成脚本。以下支持的环境：Bash、Fish、PowerShell 和 zsh
3. generate - 一个基于提供的 SQL 语句生成 .go 文件的命令。这将是我们在应用程序中大量使用的命令
4. init - 此命令是第一个用于初始化你的应用程序以开始使用此工具的命令



# Postgres
https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html


# sqlc: 格式
-- name: <方法名> :<命令类型>
<SQL 语句>


# 命令类型
标记            含义                返回类型            典型场景
:one            返回单行            (T, error)          SELECT 单条、INSERT ... RETURNING
:many           返回多行            ([]T, error)        SELECT 多条
:exec           执行不返回行        error               UPDATE、DELETE、INSERT（无 RETURNING）
:execrows       执行并返回影响行数  (int64, error)      需要知道更新/删除了多少条记录


# 注意
记得先在数据库种创建表，否则 sqlc 不会创建库和表
