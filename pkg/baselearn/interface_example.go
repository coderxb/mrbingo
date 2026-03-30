package baselearn

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义结构体实现接口
type Circle struct {
	Radius float64
}

// 实现接口Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现接口Perimeter方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}


// 接口示例函数
func interfaceExample()  {
	c := Circle{Radius: 5}
	// 直接使用 Circle 类型调用方法
	fmt.Printf("Circle Area: %.2f\n", c.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", c.Perimeter())

	// 将 Circle 赋值给 Shape 接口变量
	var s Shape = c
	fmt.Printf("Circle Area: %.2f\n", s.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", s.Perimeter())
}

/* 空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。 */
func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}