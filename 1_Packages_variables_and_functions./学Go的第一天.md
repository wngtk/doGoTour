# 学 Go 的第一天

## 安装 Go

```sh
wget https://mirrors.ustc.edu.cn/golang/go1.20.2.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
```

Add /usr/local/go/bin to the PATH environment variable. You can do this by
adding the following line to your $HOME/.profile or /etc/profile (for a
system-wide installation):
```sh
export PATH=$PATH:/usr/local/go/bin
```

## go

* Return Multiple values.
* Named return values.
* `:=` short assignment statement can be used in place of a `var` declaration
  with implicit type.

go 和 C++ 11 一样支持类型推导。 := 是一个赋值语句，可以放在使用 var 声明变量的地方。在声明变量的
同时初始化变量可以省略变量的类型。:= 语句就是把类型和 var 都省略了。

:= 只能在函数内使用。

Go 语言支持 unsigned 的数据类型，对于整数还有个熟悉的 int。 int, uint, uintptr 都是根据平台来确定大小的数据类型。 
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Go 语言的变量默认初始化为 0。