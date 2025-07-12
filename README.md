# SuiDemo

🎯 一个使用 [Wails v3](https://wails.io) 构建的现代桌面应用模板，开箱即用地集成了以下功能：

- 🌍 多语言支持
- 🌗 黑暗 / 明亮主题切换
- 🗂️ SQLite 数据库读写示例（增删改查）
- 🧱 可扩展的前后端架构，适合二次开发

---

## 💻 技术栈

| 部分         | 技术                     |
|--------------|--------------------------|
| 前端         | Vue 3 + TypeScript       |
| 样式         | Tailwind CSS             |
| 国际化       | vue-i18n                 |
| 桌面框架     | Wails v3                 |
| 数据库       | SQLite（使用 Go 操作）   |

---

## 🚀 快速开始

### 环境准备

确保你已安装以下依赖：

- [Go 1.21+](https://golang.org/dl/)
- [Node.js 16+](https://nodejs.org)
- [Wails v3 CLI](https://wails.io/docs/gettingstarted/installation)

```bash
# 安装 Wails CLI
go install github.com/wailsapp/wails/v3/cmd/wails@latest

# 克隆并运行项目
git clone https://github.com/JinGongX/SuiDemo.git
cd SuiDemo

#启动开发模式
wails3 dev

#构建生产包
wails3 package
#构建后的应用在 bin 目录下可找到。
```

## 🧱 项目结构
```
SuiDemo/
├── frontend/             # Vue3 前端代码
│   ├── src/
│   │   ├── locales/         # 多语言资源文件
│   │   ├── components/   # Vue 组件
│   │   └── App.vue
├── services/             # Go 后端服务
├── main.go               # 应用入口
├── Taskfile.yml            # Wails 配置文件
└── go.mod
```
## 📜 许可证

Apache-2.0 License

## 📸 界面展示

![输入图片说明](effect/an.jpg)
![输入图片说明](effect/white.jpg)

## 🙌 鸣谢

[Wails v3](https://v3alpha.wails.io/)

## 💬 联系方式

如果你对这个项目感兴趣或有任何建议，欢迎提 issue 或发邮件联系我 ggfugg8@icloud.com