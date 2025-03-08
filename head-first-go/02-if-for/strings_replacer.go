package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	broken := "G# r#cks!"

	replacer := strings.NewReplacer("#", "o") // 返回 *strings.Replacer 类型
	fmt.Println(reflect.TypeOf(replacer))
	
    fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
