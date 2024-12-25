package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const ipfsAPIURL = "http://localhost:5001/api/v0/add" // IPFS 节点的 API 地址

// 上传文件到 IPFS
func uploadFile(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// 创建一个新的表单数据，multipart/form-data 格式
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件到表单
	filePart, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	// 将文件内容写入表单
	_, err = io.Copy(filePart, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %v", err)
	}

	// 关闭 writer，结束表单创建
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// 发送 POST 请求到 IPFS 节点
	resp, err := http.Post(ipfsAPIURL, writer.FormDataContentType(), &requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to send request to IPFS: %v", err)
	}
	defer resp.Body.Close()

	// 解析返回的 JSON 响应，获取 CID
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	// 返回文件的 CID
	cid, ok := result["Hash"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get CID from response")
	}

	return cid, nil
}

const ipfsCatAPIURL = "http://localhost:5001/api/v0/cat" // IPFS 节点的 API 地址

// 从 IPFS 下载文件
func downloadFile(cid, filePath string) error {
	// 构造请求 URL，包含文件的 CID
	url := fmt.Sprintf("%s?arg=%s", ipfsCatAPIURL, cid)

	// 创建一个 POST 请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// 发送 POST 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码是否为 200 OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: received status code %d", resp.StatusCode)
	}
	// 创建一个本地文件用于保存下载的内容
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// 将下载的内容写入本地文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}

func main() {
	// 要上传的文件路径
	filePath := "hipfs.txt" // 替换为你自己的文件路径

	// 上传文件并获取 CID
	cid, err := uploadFile(filePath)
	if err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}

	// 打印 CID
	fmt.Printf("File uploaded successfully! CID: %s\n", cid)

	// 下载文件并保存
	err = downloadFile(cid, "downloaded_file.txt")
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	// 打印成功信息
	fmt.Println("File downloaded successfully!")
}
