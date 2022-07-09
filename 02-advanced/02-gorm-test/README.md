# gorm

## 优点

1. 提高了开发效率。
2. 屏蔽 SQL 细节。可以自动对实体 Entity 对象与数据库中的 Table 进行字段与属性的映射；不用直接 SQL 编码
3. 屏蔽各种数据库之间的差异

## 缺点

1. orm 会牺牲程序的执行效率和会固定思维模式，因为在生成 SQL 语句为了屏蔽数据库差异不管语句有没有优化，在生成过程都需要时间
2. 太过依赖 orm 会导致 SQL 理解不够
3. 对于固定的 orm 依赖过重，导致切换到其他的 orm 代价高

## 如何正确看到 orm 和 SQL 之间的关系

1. SQL 为主，ORM 为辅
2. ORM 主要目的是为了增加代码的可维护性和开发效率

## Doc

https://gorm.io
https://gorm.io/zh_CN/
https://gorm.io/zh_CN/docs/