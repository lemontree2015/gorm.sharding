# gorm.sharding
没找到golang gorm合适的分库分表组件，自己写了个供参考。

# 用到的组件
-  [Snowflake](https://github.com/bwmarrin/snowflake) (雪花算法生成唯一ID)
-  [Gorm](https://github.com/jinzhu/gorm) (golang orm组件)

# 使用说明
1.拉取代码

```sh
$ git clone https://github.com/lemontree2015/gorm.sharding.git
```

2.安装拓展
```sh
$ cd gorm-sharding
$ go mod tidy
$ go mod vendor
```

3.手动创建测试分库分表

4.修改LoadDb数据库链接

5.多数源配制方法可参考我提交的另一个仓库，[go-admin-template](https://github.com/lemontree2015/go-admin-template.git)