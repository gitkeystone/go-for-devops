package main

import "fmt"

func main() {
    fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
    fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
    fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
    fmt.Printf("%%.1f: %.1f\n", 12.3456) // 默认无填充，效果等同于最小宽度为0，对其方式：靠左，与负数一样
    fmt.Printf("%%.2f: %.2f\n", 12.3456)
}
