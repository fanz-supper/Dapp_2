# 🛍 Go 语言商城项目学习介绍

本项目是一个基于 **Go 语言 (Golang)** 的商城系统，采用 **MVC** 分层架构设计，旨在学习和体会** Go 🆚 Java ** 分层之间的异同！！

---

## 📂 项目目录结构

```
├── go.mod
├── go.sum
├── main.go              // 程序入口
└── pkg                  // 业务代码
    ├── database         // 数据库层
    │   ├── database.go
    │   └── models       // 数据模型（实体）
    │       ├── order.go
    │       ├── payment.go
    │       ├── product.go
    │       └── user.go
    ├── handler          // 表现层（HTTP 接口层）
    │   ├── base_handler.go
    │   ├── order
    │   ├── payment
    │   ├── product
    │   ├── server.go
    │   └── user
    │       └── handler.go
    └── service          // 业务逻辑层
        ├── base_service.go
        ├── order
        ├── payment
        ├── product
        └── user
            ├── add_user.go
            └── service.go
```

---

## 📌 各层作用说明

### 1. **入口层 (main.go)**

* 项目的启动文件
* 主要功能：

  * 初始化配置（数据库连接、环境变量）
  * 加载服务（handler、service）
  * 启动 HTTP 服务

---

### 2. **Database 层 (pkg/database)**

* 负责 **数据库连接** 与 **ORM 映射**。
* 文件说明：

  * `database.go`：统一管理数据库连接池（Postgres / MySQL），提供全局 DB 实例。
  * `models/`：存放数据库实体（Model），与数据库表一一对应。

    * `order.go`：订单表模型
    * `payment.go`：支付表模型
    * `product.go`：商品表模型
    * `user.go`：用户表模型

👉 **作用**：让上层 service 直接调用模型，而不用关心 SQL 细节。

---

### 3. **Service 层 (pkg/service)**

* 项目的 **业务逻辑层**，核心处理层。
* 文件说明：

  * `base_service.go`：基础服务，包含公共依赖（DB、配置等）。
  * `order/`：订单相关业务逻辑（如创建订单、查询订单）。
  * `payment/`：支付相关业务逻辑（如支付校验、订单结算）。
  * `product/`：商品业务逻辑（如商品上架、库存管理）。
  * `user/`：用户业务逻辑（如注册、登录、修改密码）。

    * `add_user.go`：专门处理用户注册逻辑。
    * `service.go`：用户服务的通用方法集合。

👉 **作用**：对数据库模型进行业务封装，提供复用逻辑，避免 handler 直接操作 DB。

---

### 4. **Handler 层 (pkg/handler)**

* **接口层（表现层）**，与客户端直接交互。
* 文件说明：

  * `base_handler.go`：封装公共 Handler 方法（如错误返回、统一响应格式）。
  * `server.go`：路由注册与服务启动（Gin/Chi 等 Web 框架）。
  * `order/`：订单相关接口（如 `/api/orders`）。
  * `payment/`：支付接口（如 `/api/payments`）。
  * `product/`：商品接口（如 `/api/products`）。
  * `user/handler.go`：用户接口（如 `/api/users/register`、`/api/users/login`）。

👉 **作用**：接收 HTTP 请求，调用 service 层方法处理，再返回响应。

---
