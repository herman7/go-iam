package main

import "fmt"

// 基类
type Bird struct {
	Type string
}

// 基类的方法Class
func (bird *Bird) Class() string {
	return bird.Type
}

// 多态通过接口实现
type Birds interface {
	Name() string
	Class() string
}

// 继承通过组合来实现
type Canary struct {
	Bird
	name string
}

func (c *Canary) Name() string {
	return c.name
}

type Crow struct {
	Bird
	name string
}

func (c *Crow) Name() string {
	return c.name
}

func NewCrow(name string) *Crow {
	return &Crow{
		Bird: Bird{
			Type: "Crow",
		},
		name: name,
	}
}

func NewCanary(name string) *Canary {
	return &Canary{
		Bird: Bird{
			Type: "Canary",
		},
		name: name,
	}
}

func BirdInfo(birds Birds) { // 多态通过接口实现
	fmt.Printf("I'm %s, I belong to %s bird class!\n", birds.Name(), birds.Class())
}

func main() {
	canary := NewCanary("CanaryA")
	crow := NewCrow("CrowA")
	BirdInfo(canary)
	BirdInfo(crow)
}
