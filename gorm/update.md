# Index

- 保存所有字段
- 更新单列
- 更新多列
- 更新选定字段
- 高级选项
  - 使用 SQL 表达式
  - 根据子查询更新
  - 不使用 Hook 和时间追踪

# 保存所有字段

Save 方法用于更新/创建记录，当传入的模型实例中没有主键时，表示创建，如果有主键，则更新所有字段，即使字段是零值;
需要注意的是，由于是更新所有字段，所以最好先将数据查询出来，再去更改指定的字段数据，否则会导致一些自动写入的字段，值为零值，从而导致数据无法写入；

```
user := &models.User{}
Db.Find(user, 14)
user.Age = 0
user.Name = "马泽法克"
Db.Save(user)


// 该模型存在 CreatedAt 字段，此处由 CreatedAt 是零值，所以导致无法写入
Db.Save(&models.User{ID: 14})
```

# 更新单列

`Update` 方法用于更新单列，需要注意的是：

- 必须指定条件，否则返回 `ErrMissingWhereClause` 错误;
- 零值字段会自动忽略;
- `Model` 方法传入的模型实例中如果包含主键，则作为条件之一。

```
Db.Model(&models.User{}).Where("id=?", 14).Update("sex", 1)
Db.Model(&models.User{ID: 14}).Update("sex", 2)
```

# 更新多列

Updates 用于更新多列，需要注意的是：

- 支持传入 `struct`或`map[string]interface{}`;
  - 当传入 `struct` 时，忽略零值字段;
  - `map` 则允许写入零值字段。
- 必须指定条件，否则返回 `ErrMissingWhereClause` 错误;

```
user := &models.User{ID: 14}

// 使用 map 批量更新，零值字段不会被忽略
Db.Model(user).Updates(map[string]interface{}{"name": "马泽法克", "age": 0})

// 使用结构体进行批量更新时，零值字段在更新时被忽略
result := Db.Model(user).Updates(&models.User{
  Name: "123",
  Age:  0,
})
// 受影响记录数
fmt.Println(result.RowsAffected)
```

# 更新选定字段

- `Select(query interface{}, args ...interface{})` 用于选中指定的字段。

需要注意的是，使用 `struct` 进行 `Select` 时，`struct` 不会忽略零值字段；因此如果使用 `*` 来匹配，会导致一些需要自动写入的字段因为写入了零值数据而报错。

- `Omit(columns ...string)` 用于排除指定的字段。

```
user := &models.User{
  ID: 14,
}

// UPDATE users SET name='马泽法克1' WHERE id=14;
Db.Model(user).Select("name").Updates(&models.User{Name: "马泽法克1", Age: 20})

// 报错
Db.Model(user).Select("*").Updates(&models.User{Name: "马泽法克", Age: 0})

// 排除 name 字段
Db.Model(user).Select("*").Omit("name").Updates(map[string]interface{}{"name": "马泽法克11", "age": 20})
```

# 高级选项

## SQL 表达式

GORM 允许使用 SQL 表达式更新列，例如：

```
// product 的 ID 是 `3`
db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;

db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;
```

并且 GORM 也允许使用 SQL 表达式、自定义数据类型的 Context Valuer 来更新，例如：

```
// 根据自定义数据类型创建
type Location struct {
    X, Y int
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

db.Model(&User{ID: 1}).Updates(User{
  Name:  "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// UPDATE `user_with_points` SET `name`="jinzhu",`location`=ST_PointFromText("POINT(100 100)") WHERE `id` = 1
```

## 根据子查询更新

```
db.Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

db.Table("users as u").Where("name = ?", "jinzhu").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))

db.Table("users as u").Where("name = ?", "jinzhu").Updates(map[string]interface{}{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})
```

## 不使用 Hook 和时间追踪

如果需要在更新时跳过 Hook 方法且不追踪更新时间，可以使用 `UpdateColumn`、`UpdateColumns`，其用法类似于 `Update`、`Updates`;

```
// 更新单个列
db.Model(&user).UpdateColumn("name", "hello")
// UPDATE users SET name='hello' WHERE id = 111;

// 更新多个列
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE id = 111;

// 更新选中的列
db.Model(&user).Select("name", "age").UpdateColumns(User{Name: "hello", Age: 0})
// UPDATE users SET name='hello', age=0 WHERE id = 111;
```
