# try-go

```sh
ldd *.exe

time 

```

> init()函数的执行顺序：a) 同一文件内的多个init函数按声明顺序执行 b) 同一包下的多个go文件按文件名称排序后的顺序
> 依次执行各文件内的init函数 c) 导入多个包时，当不存在依赖时，按import导入的顺序执行；存在依赖时，最先依赖包内的
> init函数执行

```go
type number = float64

func init() {
  defer func() {
    // TODO ...
  }()

  goto LABEL

LABEL:
  // TODO ...
  
}

// 非侵入式接口

```

## 交叉编译

Go的多版本管理工具: [GVM](https://github.com/moovweb/gvm)

```sh
go tool compile -m main.go

GOOS=linux GOARCH=amd64 go build main.go

GODEBUG=gctrace=1 go run main.go

```

## 编码

> 整数环绕

字符集:

- ascii
- gb2312
- big5
- gbk
- unicode

### 定长编码

### 变长编码

| 编号  | 编码模板 |
| ---   |  ---    |
| 0 ~ 127 | 0??????? |
| 128 ~ 2047 | 110????? 10?????? |
| 2048 ~ 65535 | 1110???? 10?????? 10?????? |

> 对于UTF8编码，一个汉字占用3个字节

## 扩容

- 预估扩容容量
- 占用内存

> 内存管理模块和内存规格

## GMP

GMP(Goroutine协程、Processor处理器和Thread线程)

> 操作码(opcode)

## 协程

并发和并行

Go调度器的调度策略:
  
- 复用线程
- 利用并行
- 抢占
- 全局协程队列

## 内存

> 机器字长和内存对齐

## GC垃圾回收

## GO Modules

```sh
go get -u <package@version>

go mod help

go mod init <project-name>

go mod tidy

go mod graph

go mod download

go mod verify 

go mod why

go mod edit 

go mod vendor

go list -m -u all

go build 

go install

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn, direct

```

## Go Workspace

```sh
go work init

go work use <pkg>

```

```txt
require localPackage

replace localPackage => 

```

## 反射

```go


```

## 测试

### 单元测试

```sh
go test [-v]

```

### 基准测试

### 模糊测试

## ORM

> NoSQL(not just sql)

主键(primary key)和外键(foreign key)

```sql
CREATE DATABASE `db_name`;

SHOW DATABASES;

DROP DATABASE `db_name`;

USE `db_name`;

CREATE TABLE `table_name`(
  `id` INT AUTO_INCREMENT,
  `user_id` INT,
  `name` VARCHAR(30) NOT NULL,
  `gender` VARCHAR(10) DEFAULT 'female',
  `create_date` DATE UNIQUE,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`user_id`) REFERENCES `table_name`(`field_name`) ON DELETE CASCADE,
  FOREIGN KEY(`user_id`) REFERENCES `table_name`(`field_name`) ON DELETE SET NULL 
);

DESCRIBE `table_name`;

DROP TABLE `table_name`;

SELECT * FROM `table_name`;
SELECT `field_name`, `field_name`, ... FROM `table_name`;

SELECT `field_name`, `field_name`, ... 
FROM `table_name`
ORDER BY `field_name`, `field_name` ... 
DESC
LIMIT n;

ALTER TABLE `table_name` ADD `field_name` DECIMAL(3,2);
ALTER TABLE `table_name` DROP COLUMN `field_name`;

INSERT INTO `table_name` VALUES(v1, v2, ...);
INSERT INTO `table_name`(n1, n2, ...) VALUES(v1, v2, ...);

UPDATE `table_name` 
SET `field_name` = `field_value`, `field_name` = `field_value` 
WHERE `create_date` LIKE '____-12-__' OR `phone` LIKE '166%';

DELETE FROM `table_name`
WHERE `major` IN ('history', 'english', 'chinese') AND `score` <> 60;

-- aggregate function

SELECT COUNT(*) FROM `table_name`;
SELECT AVG(`field_name`) FROM `table_name`;
SELECT SUM(`field_name`) FROM `table_name`;
SELECT MAX(`field_name`) FROM `table_name`;
SELECT MIN(`field_name`) FROM `table_name`;

SELECT `field_name` AS `alias` FROM `table_name` UNION SELECT `field_name` FROM `table_name`;

SELECT `field_name` FROM `table_name` 
LEFT JOIN `table_name` ON `table_name`.`field_name` = `field_name`;

-- sub query

SELECT `field_name` FROM `table_name` 
WHERE `field_name` = (
  SELECT `field_name` FROM `table_name` 
  WHERE expr
);


```
