package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/levigross/grequests"
)

func main() {
	var err error

	// ENV
	apiUrl := os.Getenv("API_URL")
	imageContent := os.Getenv("IMAGE_PATH")

	// 检测是否是URL, 如果不是URL, 则将图像转换成Base64格式
	if !isURL(imageContent) {
		imageContent, err = ImageToBase64Encode(imageContent)
		if err != nil {
			log.Fatal(err)
		}
	}

	resp, err := imageToText(apiUrl, imageContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func imageToText(apiUrl, imageContent string) (string, error) {
	// 请求选项
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":    "application/json",
			"Request-Timeout": "10",
		},
		JSON: map[string]any{
			"file":     imageContent, // Base64编码的文件内容或者文件URL
			"fileType": 1,            // 文件类型，1表示图像文件
		},
	}

	// POST 请求
	resp, err := grequests.Post(apiUrl, grequests.FromRequestOptions(ro))
	if err != nil {
		return "", fmt.Errorf("unable to make request: %s", err.Error())
	}
	if !resp.Ok {
		return "", fmt.Errorf("request did not return OK: %s", resp.String())
	}

	// 响应结果
	return resp.String(), nil
}

func ImageToBase64Encode(imagePath string) (string, error) {
	fileBytes, err := os.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("can not read file <%s>: %s", imagePath, err.Error())
	}
	return base64.StdEncoding.EncodeToString(fileBytes), nil
}

func isURL(in string) bool {
	re := regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
	if re.MatchString(in) {
		return true
	}

	return false
}
