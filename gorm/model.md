# Index

- 模型定义
- 约定
- 高级选项
  - 字段级权限控制
  - 创建/更新/删除时间追踪（纳秒、毫秒、秒、Time）
  - 嵌入结构体
  - 关联标签（参见模型关联）
- 字段标签

# 模型定义

GORM 中，模型的定义以 `struct` 表现。

```
type User struct {
	ID        uint
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
```

# 约定

- 表名是结构体的复数形式，例如：User -> users;
- 列名是字段的蛇形小写;
- ID 字段为表的主键;
- `CreatedAt time.Time` 字段用于自动存储记录的创建时间，time.Time 表示表内字段类型为 time;
- `UpdatedAt int` 字段用于自动存储记录的修改时间，int 类型表示表内字段类型为 int，因此存入的数据为时间戳;
- `DeletedAt gorm.DeletedAt` 如果该字段存在，则用于自动存储记录的删除时间（软删除）;

# 高级选项

## 字段级权限控制

默认情况下，可导出的字段在使用 GORM 进行 CURD 时拥有全部的权限；同时，GORM 提供了通过标签控制字段的权限。这样就可以控制指定字段的行为。

- `<-:create` 允许读和创建；
- `<-:update` 允许读和更新；
- `<-` 允许读写（创建和更新）；
- `<-:false` 允许读，禁止写；
- `->` 只读（除非有自定义配置）；
- `->;<-:create` 允许读和写；
- `->` 只读（除非有自定义配置）；
- `->:false;<-:create` 仅创建（禁止从 db 读）；
- `-` 读写均忽略该字段，注意，使用 GORM Migrator 创建表时，不会创建被忽略的字段。

## 创建/更新/删除时间追踪（纳秒、毫秒、秒、Time）

GORM 约定使用 `CreatedAt、UpdatedAt` 追踪创建/更新时间。因此，如果定义了这种字段，则 GORM 在创建、更新时会自动填充当前时间。

可以通过为字段指定 `autoCreateTime`、`autoUpdateTime` 标签来指定其他字段作为追踪创建/更新时间的字段。

如果想要保存 unix 时间戳，而不是 time，则只需要将字段类型从 `time.Time` 类型改为 int 即可。

使用毫秒、纳秒时间戳：需要将字段数据类型改为 `int64`，同时使用字段标签 `autoUpdateTime:milli`、`autoUpdateTime:nano`...

```
type User struct {
	ID          int
	Name string `gorm:"size:30"`
	Updated     int64  `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created     int    `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
}
```

GORM 中约定字段类型为 `gorm.DeletedAt` 的字段用于记录删除时间，当模型中包含了该类型的字段时，则开启软删除；如果想保存 unix 时间戳，则需要使用 `soft_delete.DeletedAt` 类型，而不是 `gorm.DeletedAt`。

```
import "gorm.io/plugin/soft_delete"

type User struct {
	// DeletedAt gorm.DeletedAt
  // DeletedAt soft_delete.DeletedAt // 删除时间存储时间戳
}
```

## 嵌入结构体

对于匿名字段，GORM 会将其字段包含在父结构体

```
type User struct {
  gorm.Model
  Name string
}
// 等效于
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
```

对于正常字段，可以为其添加 `embedded` 标签将其嵌入：

```
type Author struct {
    Name  string
    Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
// 等效于
type Blog struct {
  ID    int64
    Name  string
    Email string
  Upvotes  int32
}
```

并可以通过标签 `embeddedPrefix` 来为其在表中添加字段前缀

```
type Blog struct {
  ID      int
  Author  Author `gorm:"embedded;embeddedPrefix:author_"`
  Upvotes int32
}
// 等效于
type Blog struct {
  ID          int64
    AuthorName  string
    AuthorEmail string
  Upvotes     int32
}
```

# 字段标签

| 名称                   | 说明                                                                                                                                                                                                                                                                                                                                   |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| column                 | 指定表中列名                                                                                                                                                                                                                                                                                                                           |
| type                   | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：not null、size, autoIncrement… 像 varbinary(8) 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT |
| size                   | 指定列大小，例如 `size:30`                                                                                                                                                                                                                                                                                                             |
| primaryKey             | 指定为主键                                                                                                                                                                                                                                                                                                                             |
| unique                 | 列唯一                                                                                                                                                                                                                                                                                                                                 |
| default                | 指定默认值                                                                                                                                                                                                                                                                                                                             |
| precision              | 指定列的精度                                                                                                                                                                                                                                                                                                                           |
| scale                  | 指定列大小                                                                                                                                                                                                                                                                                                                             |
| not null               | 指定列为 NOT NULL                                                                                                                                                                                                                                                                                                                      |
| autoIncrement          | 指定列自增                                                                                                                                                                                                                                                                                                                             |
| autoIncrementIncrement | 自动步长，控制连续记录之间的间隔                                                                                                                                                                                                                                                                                                       |
| embedded               | 嵌套字段                                                                                                                                                                                                                                                                                                                               |
| embeddedPrefix         | 嵌入字段的列名前缀                                                                                                                                                                                                                                                                                                                     |
| autoCreateTime         | 创建时追踪当前时间，对于 int 字段，它会追踪秒级时间戳，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano                                                                                                                                                                                                         |
| autoUpdateTime         | 创建/更新时追踪当前时间，对于 int 字段，它会追踪秒级时间戳，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli                                                                                                                                                                                                   |
| index                  | 根据参数创建索引，多个字段使用相同的名称则创建复合索引                                                                                                                                                                                                                                                                                 |
| uniqueIndex            | 与 index 相同，但创建的是唯一索引                                                                                                                                                                                                                                                                                                      |
| check                  | 创建检查约束，例如 check:age > 13                                                                                                                                                                                                                                                                                                      |
| <-                     | 设置字段写入的权限， <-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限                                                                                                                                                                                                                                        |
| ->                     | 设置字段读的权限，->:false 无读权限                                                                                                                                                                                                                                                                                                    |
| -                      | 忽略该字段，- 无读写权限                                                                                                                                                                                                                                                                                                               |
| comment                | 迁移时为字段添加注释                                                                                                                                                                                                                                                                                                                   |
