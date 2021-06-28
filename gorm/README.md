# # Index

- 介绍
- 安装
- 连接数据库(connect.md) 待记录
- 模型定义(model.md)
- CURD
  - 增(create.md)
  - 删(delete.md)
  - 改(update.md)

# 介绍

- 全功能 ORM
- 关联 (Has One，Has Many，Belongs To，Many To Many，多态，单表继承)
- Create，Save，Update，Delete，Find 中钩子方法
- 支持 Preload、Joins 的预加载
- 事务，嵌套事务，Save Point，Rollback To Saved Point
- Context、预编译模式、DryRun 模式
- 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUD
- SQL 构建器，Upsert，数据库锁，Optimizer/Index/Comment Hint，命名参数，子查询
- 复合主键，索引，约束
- Auto Migration
- 自定义 Logger
- 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…

# 安装

1. `go get -u gorm.io/gorm`;
2. `gorm.io/driver/mysql`;
3. 根据需求安装其他数据库的驱动.
