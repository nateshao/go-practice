/**
 * @date Created by 邵桐杰 on 2021/7/19 21:13
 * @微信公众号 千羽的编程时光
 * @个人网站 www.nateshao.cn
 * @博客 https://nateshao.gitee.io
 * @GitHub https://github.com/nateshao
 * Describe:
 */
package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":     "ccmouse",
		"course":   "golang",
		"username": "zhangsan",
		"site":     "cainiao",
	}
	fmt.Println(m)
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)
	for k, v := range m { // 遍历输出map m1---打印key和value	: 每一次输出的顺序不一致 原因：map是无序的，是hashmao
		fmt.Println("key =", k, " value = ", v)
	}

	fmt.Println("----------只打印key-----------")
	for k := range m {
		fmt.Println(k) // 只打印key
	}
	fmt.Println("----------只打印value-----------")
	for _, v := range m {
		fmt.Println(v) // 只打印value
	}
	fmt.Println("------------------------")
	coursename, ok := m["course"]
	fmt.Println(coursename, ok)

	fmt.Println("---------如果是拼错的情况下 - 输出空行----------")
	if caursename, ok := m["caurse"]; ok {
		fmt.Println("拼错的情况下 - 输出空行:", caursename, ok)
	} else {
		fmt.Println("key 不存在")
	}
	fmt.Println("-------ccmouse 删除前-------")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	fmt.Println("-------ccmouse 删除后-------")

	name, ok = m["name"] // 第二次赋值就不用 := 了
	fmt.Println(name, ok)
}
