# 微信 Markdown 编辑器 MCP 服务器

这是一个基于 [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go) 的 Go 服务器，用于将 Markdown 文本转换为微信公众号格式的 HTML。

## 功能特性

- 🚀 **快速转换**: 将 Markdown 文本转换为微信公众号兼容的 HTML 格式
- 🎨 **多种主题**: 支持 4 种主题风格（默认、字节范、中国风、苹果范）
- 📊 **详细统计**: 提供字数统计和预估阅读时间
- 🔒 **安全认证**: 支持 API Key 认证
- 💬 **智能集成**: 可与大语言模型（如 Claude、GPT）无缝集成

## 支持的 Markdown 语法

### 基础语法
- 标题 (H1-H6)
- 段落和换行
- **加粗文本** 和 *斜体文本*
- `内联代码`
- [链接](https://example.com)
- ![图片](https://example.com/image.jpg)
- 无序列表和有序列表
- > 引用块
- 分割线 (---)

### 扩展语法
- 代码块（支持语法高亮）
- 表格
- 脚注 [^1]
- GFM 提示框（NOTE、TIP、IMPORTANT、WARNING、CAUTION）

## 环境要求

- Go 1.21 或更高版本
- 微信 Markdown 编辑器 API Key

## 安装和配置

### 1. 克隆或下载代码

```bash
git clone <your-repo-url>
cd md2wechat-mcp-server
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 设置 API Key

获取微信 Markdown 编辑器的 API Key 后，设置环境变量：

```bash
# Linux/macOS
export MD2WECHAT_API_KEY="wme_your_api_key_here"

# Windows
set MD2WECHAT_API_KEY=wme_your_api_key_here
```

### 4. 构建和运行

```bash
# 构建
go build -o md2wechat-mcp-server

# 运行
./md2wechat-mcp-server
```

## MCP 工具说明

### convert_markdown

将 Markdown 文本转换为微信公众号格式的 HTML。

**参数：**

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| markdown | string | ✅ | 要转换的 Markdown 文本内容 |
| theme | string | ❌ | 主题样式，默认为 "default" |

**支持的主题：**

- `default` - 默认微信经典风格
- `bytedance` - 字节范（科技现代风格）
- `chinese` - 中国风（古典雅致风格）
- `apple` - 苹果范（视觉渐变风格）

**返回结果：**

```json
{
  "success": true,
  "html": "<section>...</section>",
  "theme": "default",
  "wordCount": 156,
  "estimatedReadTime": 1,
  "message": "Markdown 转换成功"
}
```

## 在大语言模型中使用

### Claude Desktop 配置

在 `claude_desktop_config.json` 中添加：

```json
{
  "mcpServers": {
    "md2wechat": {
      "command": "/path/to/md2wechat-mcp-server",
      "env": {
        "MD2WECHAT_API_KEY": "wme_your_api_key_here"
      }
    }
  }
}
```

### 使用示例

与 Claude 对话时，可以这样请求：

```
请帮我将以下 Markdown 文本转换为微信公众号格式：

# 我的文章标题

这是一个**重要**的段落。

## 子标题

- 列表项 1
- 列表项 2

使用中国风主题。
```

Claude 会调用 `convert_markdown` 工具，并返回转换后的 HTML 代码。

## 错误处理

服务器会处理以下常见错误：

- API Key 未设置或无效
- Markdown 内容为空
- 无效的主题参数
- 网络连接问题
- API 服务异常

## API 限制

- 单次请求 Markdown 内容不超过 100KB
- API Key 配额：100 请求/分钟
- 网络请求超时：30 秒

## 开发和贡献

### 项目结构

```
md2wechat-mcp-server/
├── main.go          # 主程序文件
├── go.mod           # Go 模块依赖
├── README.md        # 说明文档
└── api_doc.md       # 原始 API 文档
```

### 开发建议

1. 遵循 Go 代码规范
2. 添加适当的错误处理
3. 编写详细的注释
4. 测试 API 集成

## 许可证

本项目基于 MIT 许可证开源。

## 支持和反馈

如有问题或建议，请提交 Issue 或联系开发团队。

---

**注意**: 使用前请确保已正确配置 API Key，并遵守微信 Markdown 编辑器的使用条款。 
