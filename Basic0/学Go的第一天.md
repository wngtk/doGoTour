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

我是这么做的，添加下面的语句到 $HOME/.profile.
```sh
# set PATH so it includes sys's /usr/local/go/bin if it exists
if [ -d "/usr/local/go/bin" ]; then
    PATH="/usr/local/go/bin:$PATH"
    export GO111MODULE=on
    export GOPROXY=https://goproxy.cn,direct
fi
```

## go

* Return Multiple values.
* Named return values.
* `:=` short assignment statement can be used in place of a `var` declaration
  with implicit type.

go 和 C++ 11 一样支持类型推导。 := 是一个赋值语句，可以放在使用 var 声明变量的地方。在声明变量的
同时初始化变量可以省略变量的类型。:= 语句就是把类型和 var 都省略了。

:= 只能在函数内使用。因为在函数的外面每一条语句都要以一个关键字开头。其实 C/C++ 的语句在函数外也是这样的，看起来 Go 的语法虽然和 C/C++ 不一样，但是语法本质上还是语句和块组成，不过多了语法糖。
> Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available. https://go.dev/tour/basics/10

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