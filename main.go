package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// API 基础配置
const (
	BaseURL     = "https://www.md2wechat.cn/api"
	ConvertPath = "/convert"
	APIKeyEnv   = "MD2WECHAT_API_KEY"
)

// 支持的主题列表
var supportedThemes = []string{"default", "bytedance", "chinese", "apple"}

// ConvertRequest 转换请求结构体
type ConvertRequest struct {
	Markdown string `json:"markdown"`
	Theme    string `json:"theme,omitempty"`
}

// ConvertResponseData 转换响应数据结构体
type ConvertResponseData struct {
	HTML              string `json:"html"`
	Theme             string `json:"theme"`
	WordCount         int    `json:"wordCount"`
	EstimatedReadTime int    `json:"estimatedReadTime"`
}

// ConvertResponse 转换响应结构体
type ConvertResponse struct {
	Code int                  `json:"code"`
	Msg  string               `json:"msg"`
	Data *ConvertResponseData `json:"data,omitempty"`
}

func main() {
	// 创建 MCP 服务器
	s := server.NewMCPServer(
		"微信 Markdown 编辑器 MCP 服务器",
		"1.0.0",
	)

	// 添加 convert_markdown 工具
	convertTool := mcp.NewTool("convert_markdown",
		mcp.WithDescription("将 Markdown 文本转换为微信公众号格式的 HTML。支持多种主题风格，包括默认、字节范、中国风和苹果范等。"),
		mcp.WithString("markdown",
			mcp.Required(),
			mcp.Description("要转换的 Markdown 文本内容。支持标题、段落、加粗、斜体、代码块、链接、图片、列表、引用等语法。"),
		),
		mcp.WithString("theme",
			mcp.Description("主题样式，可选值：default（默认微信经典风格）、bytedance（科技现代风格）、chinese（古典雅致风格）、apple（视觉渐变风格）。默认为 default。"),
		),
	)

	// 添加工具处理器
	s.AddTool(convertTool, convertMarkdownHandler)

	// 启动 stdio 服务器
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("服务器错误: %v\n", err)
	}
}

// convertMarkdownHandler 处理 Markdown 转换请求
func convertMarkdownHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 检查 API Key
	apiKey := os.Getenv(APIKeyEnv)
	if apiKey == "" {
		return nil, fmt.Errorf("未找到 API Key，请设置环境变量 %s", APIKeyEnv)
	}

	// 解析参数
	markdown, err := request.RequireString("markdown")
	if err != nil {
		return nil, fmt.Errorf("markdown 参数必须是字符串类型: %v", err)
	}

	if markdown == "" {
		return nil, errors.New("markdown 内容不能为空")
	}

	// 处理主题参数
	theme := request.GetString("theme", "default") // 使用默认值
	if theme != "" && !isValidTheme(theme) {
		return nil, fmt.Errorf("无效的主题: %s，支持的主题: %v", theme, supportedThemes)
	}

	// 构建请求
	reqData := ConvertRequest{
		Markdown: markdown,
		Theme:    theme,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("构建请求数据失败: %v", err)
	}

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", BaseURL+ConvertPath, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建 HTTP 请求失败: %v", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-Key", apiKey)

	// 发送请求
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送 HTTP 请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	// 解析响应
	var convertResp ConvertResponse
	if err := json.Unmarshal(body, &convertResp); err != nil {
		return nil, fmt.Errorf("解析响应 JSON 失败: %v", err)
	}

	// 检查 API 错误
	if convertResp.Code != 0 {
		return nil, fmt.Errorf("API 错误 (状态码: %d): %s", convertResp.Code, convertResp.Msg)
	}

	// 检查响应数据
	if convertResp.Data == nil {
		return nil, errors.New("API 响应数据为空")
	}

	// 构建成功响应
	result := map[string]interface{}{
		"success":           true,
		"html":              convertResp.Data.HTML,
		"theme":             convertResp.Data.Theme,
		"wordCount":         convertResp.Data.WordCount,
		"estimatedReadTime": convertResp.Data.EstimatedReadTime,
		"message":           "Markdown 转换成功",
	}

	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("构建响应 JSON 失败: %v", err)
	}

	return mcp.NewToolResultText(string(resultJSON)), nil
}

// isValidTheme 检查主题是否有效
func isValidTheme(theme string) bool {
	for _, validTheme := range supportedThemes {
		if theme == validTheme {
			return true
		}
	}
	return false
}
