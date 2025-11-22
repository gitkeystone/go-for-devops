package main

import "fmt"

func readData(shouldError bool) (string, error) {
	if shouldError {
		return "", fmt.Errorf("读取数据时，发生错误")
	}
	return "成功读取到的数据", nil
}

func main() {
	//data, err := readData(true)
	//if err != nil {
	//	fmt.Println("错误", err)
	//	return
	//}
	//fmt.Println("数据：", data)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到异常:", r)
		}
	}()

	panic("模拟一个严重错误")
	fmt.Println("这行代码不会被执行")
}
