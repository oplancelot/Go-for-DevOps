package main

//https://pkg.go.dev/flag#Var
import (
	"flag"
	"fmt"
	"net/url"
	"reflect"
)

// Value is the interface to the dynamic value stored in a flag. (The default value is represented as a string.)

// If a Value has an IsBoolFlag() bool method returning true, the command-line parser makes -name equivalent to -name=true rather than using the next command-line argument.

// Set is called once, in command line order, for each flag present. The flag package may call the String method with a zero-valued receiver, such as a nil pointer.
// Value 是存储在标志中的动态值的接口(默认值表示为字符串)

// 如果 Value 具有返回 true 的 IsBoolFlag () bool 方法，则命令行解析器使-name 等效于-name = true，而不使用下一个命令行参数。

// 对于存在的每个标志，Set 按命令行顺序调用一次。标志包可以调用具有零值接收器(如 nil 指针)的 String 方法。

// This code does the following:Defines a flag.Value type called URLValueCreates a flag called -url that reads in a valid URLUses the URLValue wrapper to store the URL in a *url.URL variableUses the reflect package to determine whether struct is emptyBy defining a Set() method on a type, as we did previously, you can read inany custom value.
type Value interface {
	String() string
	Set(string) error
}

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}
func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		//*取值
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

// 这里定义flag url
// golang&和*的区别
// &符号的意思是对变量取地址
// *符号的意思是对指针取值

// Var defines a flag with the specified name and usage string.
//
//	The type and value of the flag are represented by the first argument, of type Value, which typically holds a user-defined implementation of Value.
//
// For instance, the caller could create a flag that turns a comma-separated string into a slice of strings by giving the slice the methods of Value;
//
// in particular, Set would decompose the comma-separated string into the slice.
// Var用指定的名称和用法字符串定义一个标志。标志的类型和值由第一个参数表示，该参数类型为value，通常包含用户定义的value实现。
// 例如，调用者可以创建一个标志，将逗号分隔的字符串转换为字符串切片，方法是将Value;特别地，Set会将逗号分隔的字符串分解为切片。

func init() { flag.Var(&URLValue{u}, "url", "URL to parse") }
func main() {
	flag.Parse()

	if reflect.ValueOf(*u).IsZero() {
		panic("did not pass an URL")
	}
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
}

//go run url.go -url https://example.com/foo%2fbar
//{scheme: "https", host: "example.com", path: "/foo/bar"}
