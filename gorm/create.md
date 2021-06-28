# Index

- 基础
  - 创建记录
  - 指定字段创建记录
  - 批量插入
  - 使用 SQL
- 高级选项
  - 默认值
  - 关联创建

# 基础

## 创建记录

通过实例话结构体，并将结构体指针传入 `Create` 方法完成。

```
data := &model{...}
result := db.Create(data)

result.ID // 插入的数据主键
result.RowsAffected // 插入的数据记录条数
result.Error // error
```

## 指定字段创建记录

```
model := &model{...}
result := db.Select(...field string).Create(model)
// INSERT INTO `model` (`...`,`...`,`...`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
```

## 批量插入

通过将切片传递给 `Create` 方法，GORM 将生成一个单一的 SQL 来插入所有数据，并将回填主键的值，钩子方法也将被调用。

```
models := []Models{{...}, {...}}
db.Create(&models)

for _, model := range models {
    model.ID
}
```

## 根据 Map 创建

GORM 支持通过 `map[string]interface{}` 或 `[]map[string]interface{}{}` 来创建记录；
需要注意的是，通过 map 创建记录无法自动填充主键，association 不会被调用。

```
	Db.Model(&models.User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	Db.Model(&models.User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
```

# 使用 SQL

GORM 允许使用 SQL 表达式插入数据，有两种方法实现这个目标。根据 `map[string]interface{}` 或 自定义数据类型 创建，例如：

```
// 通过 map 创建记录
db.Model(User{}).Create(map[string]interface{}{
  "Name": "jinzhu",
  "Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));

// 通过自定义类型创建记录
type Location struct {
    X, Y int
}

// Scan 方法实现了 sql.Scanner 接口
func (loc *Location) Scan(v interface{}) error {
  // Scan a value into struct from database driver
}

func (loc Location) GormDataType() string {
  return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

type User struct {
  Name     string
  Location Location
}

db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)")
```

# 高级选项

## 关联创建

创建关联数据时，如果关联值是非零值，这些关联会被 upsert，且它们的 Hook 方法也会被调用

```
type UserCard struct {
  gorm.Model
  Number   string
  UserID   uint
}

type User struct {
  gorm.Model
  Name       string
  UserCard UserCard
}

db.Create(&User{
  Name: "jinzhu",
  UserCard: UserCard{Number: "xxx,yyy,zzz"}
})

// INSERT INTO `users` ...
// INSERT INTO `user_cards` ...

```

## 默认值

通过 `default` 为字段定义默认值,例如

```
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}
```

插入记录到数据库时，默认值 会被用于 填充值为 零值 的字段。
