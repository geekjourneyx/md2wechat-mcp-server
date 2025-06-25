# 变更日志

本文件记录了 md2wechat-mcp-server 项目的所有重要变更。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
项目遵循 [语义化版本](https://semver.org/lang/zh-CN/) 规范。

## [1.0.0] - 2025-06-25

### 🎉 首次发布

#### 新增
- **核心功能**: Markdown 到微信公众号 HTML 格式转换
- **MCP 集成**: 完整的 Model Context Protocol 支持
- **多主题支持**: 
  - `default` - 默认微信经典风格
  - `bytedance` - 字节范（科技现代风格）
  - `chinese` - 中国风（古典雅致风格）
  - `apple` - 苹果范（视觉渐变风格）
- **统计功能**: 自动字数统计和阅读时间预估
- **API Key 认证**: 安全的 API 访问控制
- **Markdown 语法支持**:
  - 基础语法: 标题、段落、加粗、斜体、链接、图片、列表、引用等
  - 扩展语法: 代码块、表格、脚注、GFM 提示框等
- **错误处理**: 完善的错误处理和异常恢复机制
- **Claude Desktop 集成**: 开箱即用的 Claude 集成配置

#### 技术规格
- **语言**: Go 1.23.2+
- **依赖**: MCP Go v0.32.0
- **性能**: 
  - 转换速度 < 100ms
  - 支持最大 100KB 文档
  - API 限制: 100 请求/分钟
- **许可证**: MIT License

#### 文档
- 完整的 README.md 文档
- API 使用说明
- 安装和配置指南
- Claude Desktop 集成教程

---

## 版本说明

- **[1.0.0]**: 首个稳定版本，包含所有核心功能
- 后续版本将在此基础上持续改进和添加新功能

---

[1.0.0]: https://github.com/geekjourneyx/md2wechat-mcp-server/releases/tag/v1.0.0 
