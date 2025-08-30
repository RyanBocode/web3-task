# Blog API

一个基于Go语言的博客API服务，使用Gin框架和GORM ORM。

## 功能特性

- 用户注册和登录（JWT认证）
- 博客文章的CRUD操作
- 评论系统
- **默认使用SQLite数据库**（无需额外配置）
- 支持MySQL数据库（可选）
- RESTful API设计

## 技术栈

- **框架**: Gin
- **ORM**: GORM
- **数据库**: **SQLite (默认)** / MySQL
- **认证**: JWT
- **密码加密**: bcrypt

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置环境变量（可选）

创建 `.env` 文件来覆盖默认配置：

```env
# 数据库配置 - 默认使用SQLite
DB_DRIVER=sqlite
SQLITE_PATH=blog.db

# 服务器配置
PORT=8080

# JWT配置
JWT_SECRET=your_secret_key
JWT_TTL_HOURS=24

# 环境配置
GIN_MODE=debug
```

**注意**: 如果不创建.env文件，应用将使用以下默认配置：
- 数据库: SQLite
- 数据库文件: `blog.db` (在当前目录)
- 端口: 8080

### 3. 运行应用

```bash
go run main.go
```

或者构建后运行：

```bash
go build
./blog
```

### 4. 访问API

服务将在 `http://localhost:8080` 启动

## SQLite数据库说明

### 默认配置
- **数据库类型**: SQLite
- **数据库文件**: `blog.db` (自动创建)
- **位置**: 项目根目录
- **迁移**: 自动创建表结构

### SQLite优势
- ✅ 无需安装数据库服务器
- ✅ 零配置，开箱即用
- ✅ 单文件数据库，便于部署
- ✅ 支持事务和关系查询
- ✅ 适合开发和中小型应用

### 数据库文件
启动应用后，会在项目根目录自动创建 `blog.db` 文件，包含：
- `users` 表 - 用户信息
- `posts` 表 - 博客文章
- `comments` 表 - 评论

## API端点

### 认证
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录

### 文章（公开）
- `GET /api/v1/posts` - 获取文章列表
- `GET /api/v1/posts/:id` - 获取单篇文章
- `GET /api/v1/posts/:id/comments` - 获取文章评论

### 文章（需要认证）
- `POST /api/v1/posts` - 创建文章
- `PUT /api/v1/posts/:id` - 更新文章
- `DELETE /api/v1/posts/:id` - 删除文章

### 评论（需要认证）
- `POST /api/v1/posts/:id/comments` - 创建评论

### 用户信息（需要认证）
- `GET /api/v1/me` - 获取当前用户信息

## 项目结构

```
blog/
├── main.go                 # 应用入口
├── go.mod                  # Go模块文件
├── blog.db                 # SQLite数据库文件（自动创建）
├── internal/               # 内部包
│   ├── config/            # 配置管理
│   ├── controllers/       # 控制器
│   ├── database/          # 数据库连接
│   ├── dto/               # 数据传输对象
│   ├── middleware/        # 中间件
│   ├── models/            # 数据模型
│   ├── responses/         # 响应处理
│   └── router/            # 路由配置
└── README.md              # 项目说明
```

## 数据库模型

- **User**: 用户信息
- **Post**: 博客文章
- **Comment**: 评论

## 开发说明

项目已修复所有格式问题，包括：
- 修复了所有Go文件的缩进和格式
- 添加了缺失的imports
- 修复了函数签名和结构体定义
- 统一了代码风格
- **优化了SQLite配置和日志输出**