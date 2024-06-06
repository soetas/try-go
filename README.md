# try-go

```sh
ldd *.exe

time 

```

```go
type number = float64

func init() {
  defer func() {
    // TODO ...
  }()

LABEL:
  // TODO ...
  
}

```

## 编码

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
go mod init <project-name>

go mod tidy

go env -w GO111MODULE=on

```

## 测试

### 单元测试

### 基准测试

### 模糊测试
