package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func RobustWiFiLogin() {
	fmt.Println("🚀 开始WiFi自动登录...")

	// 配置HTTP客户端
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}

	// 登录数据
	data := url.Values{
		"username":    {"JSfangke112"},
		"password":    {"jrG39MyG!#@"},
		"RedirectUrl": {""},
		"anonymous":   {"DISABLE"},
		"anonymousurl": {""},
		"accesscode":  {""},
		"accesscode1": {"DISABLE"},
		"checkbox":    {"on"},
		"checkbox1":   {"on"},
	}

	// 发送请求
	resp, err := client.PostForm("https://2.1.1.1:8443/login", data)

	// 处理响应
	if err != nil {
		// 检查是否是那种"成功但连接中断"的错误
		if strings.Contains(err.Error(), "malformed MIME header") ||
		   strings.Contains(err.Error(), "connection broken") {
			fmt.Println("✅ WiFi登录成功! (服务器响应异常但认证已完成)")
		} else {
			fmt.Printf("❌ 登录失败: %v\n", err)
			return
		}
	} else {
		defer resp.Body.Close()

		if resp.StatusCode == 302 {
			fmt.Println("✅ WiFi登录成功! (302 重定向)")
		} else if resp.StatusCode == 200 {
			fmt.Println("✅ WiFi登录成功! (200 OK)")
		} else {
			fmt.Printf("⚠️  未知响应状态: %d\n", resp.StatusCode)
		}
	}

	fmt.Println("🎉 登录流程完成!")

	// 等待并测试网络
	fmt.Println("⏳ 测试网络连接...")
	time.Sleep(2 * time.Second)

	// 简单测试网络
	testClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}

	if _, err := testClient.Get("https://www.baidu.com"); err == nil {
		fmt.Println("✅ 网络连接正常!")
	} else {
		fmt.Println("⚠️  网络连接测试失败，但认证可能已成功")
	}
}

func main() {
	RobustWiFiLogin()
}