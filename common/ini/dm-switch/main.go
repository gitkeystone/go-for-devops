package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const IniPath = "/etc/ansible/db_config.ini"

// const IniPath = "./db_config.ini"

func main() {
	// 加载并读取INI文件
	cfg, err := ini.Load(IniPath)
	if err != nil {
		fmt.Printf("加载配置文件失败: %v", err)
		return
	}

	// 获取节
	dm := cfg.Section("dm")

	// 修改配置: 切换端口切换
	// 1. 手动切换端口
	// 读取旧端口
	//port := dm.Key("port").MustInt()

	//if port == 31821 {
	//	port = 32366
	//} else if port == 32366 {
	//	port = 31821
	//} else {
	//	port = 31821
	//}

	// 配置新端口
	//dm.Key("port").SetValue(strconv.Itoa(port))

	// 2. 通过数据库服务名 GroupName 查询端口
	// 获取新端口
	port, err := getPrimaryModePort()
	if err != nil {
		fmt.Printf("获取 DM 实例端口失败: %v", err)
		return
	} else {
		// 配置新端口
		dm.Key("port").SetValue(port)
	}

	// 保存修改
	if err = cfg.SaveTo(IniPath); err != nil {
		fmt.Printf("保存修改失败: %v", err)
	}
}
