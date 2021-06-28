# Index

- 删除单条
- 批量删除
- 软删除
  - 查找被软删的记录
  - 硬删除
- 存储 Uninx 时间戳

# 删除单条

- 删除单条数据时，需要指定删除对象的主键(生成的 SQL 包含了主键作为条件的语句)，否则会触发批量删除

```
// mockData := &models.User{ID: 9}
// if err := Db.Delete(mockData).Error; err != nil {
// 	log.Fatalln(err)
// }
// DELETE from users where id = 9;

// 根据主键删除
// result := Db.Delete(&models.User{}, 10)
result := Db.Delete(&models.User{}, []int{11, 12})
fmt.Println(result.RowsAffected)
```

# 批量删除

没有任何条件的批量删除，GORM 不会执行任何操作，并返回 `gorm.ErrMissingWhereClause`

```
result := Db.Where("age = ?", "21").Delete(&models.User{})
fmt.Println(result.RowsAffected)

// 没有任何条件的批量删除，GORM 不会执行任何操作，并返回  gorm.ErrMissingWhereClause
// Db.Delete(&models.User{}).Error
```

# 软删除

当模型中包含了一个 `gorm.DeletedAt ` 类型的字段时，模型自动获得软删除能力

## 查找被软删的数据

```
user := &models.User{}
Db.Unscoped().Find(user, 13)
```

## 硬删除

```
Db.Unscoped().Delete(&Model{}, PrimaryKey)
```
