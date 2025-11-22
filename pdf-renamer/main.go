package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func processFile(filename string) error {
	// 检查是否为PDF文件
	if strings.ToLower(filepath.Ext(filename)) != ".pdf" {
		return fmt.Errorf("不是PDF文件")
	}

	// 检查文件名中是否包含"行程单"
	if !strings.Contains(filename, "行程单") {
		return fmt.Errorf("不包含'行程单'")
	}

	// 分割文件名（不含扩展名）
	baseName := strings.TrimSuffix(filename, filepath.Ext(filename))
	parts := strings.Split(baseName, "-")

	// 检查是否有足够的段
	if len(parts) < 5 {
		return fmt.Errorf("文件名格式不正确")
	}

	// 删除第4段（索引为3）
	newParts := append(parts[:3], parts[4:]...)
	
	// 重新拼接文件名
	newBaseName := strings.Join(newParts, "-")
	newFilename := newBaseName + filepath.Ext(filename)

	// 重命名文件
	return os.Rename(filename, newFilename)
}

func main() {
	fmt.Println("PDF文件重命名工具")
	fmt.Println("==================")

	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败: %v\n", err)
		return
	}

	fmt.Printf("处理目录: %s\n\n", currentDir)

	// 读取目录中的所有文件
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
		return
	}

	// 计数器
	processedCount := 0

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		fmt.Printf("检查: %s\n", filename)

		err := processFile(filename)
		if err != nil {
			fmt.Printf("  ↳ 跳过: %v\n", err)
			continue
		}

		fmt.Printf("  ↳ 已重命名\n")
		processedCount++
	}

	fmt.Printf("\n处理完成! 共处理了 %d 个文件\n", processedCount)
	
	// 等待用户按键（Windows下有用）
	fmt.Println("按回车键退出...")
	fmt.Scanln()
}
