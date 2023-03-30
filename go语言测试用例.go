package main

import (
    "fmt"
    "sync"
)

// 定义一个字符串类型
type String struct {
    data string  // 字符串数据
}

// 创建一个字符串
func CreateString(data string) *String {
    return &String{data}
}

// 打印字符串
func (s *String) Print() {
    fmt.Println(s.data)
}

// 连接两个字符串
func ConcatenateStrings(s1, s2 *String) *String {
    return &String{s1.data + s2.data}
}

// 反转字符串
func ReverseString(s *String) *String {
    data := []rune(s.data)
    for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
    return &String{string(data)}
}

// 定义链表节点类型
type Node struct {
    string *String  // 指向字符串的指针
    next *Node      // 指向下一个节点的指针
}

// 创建一个链表节点
func CreateNode(s *String) *Node {
    return &Node{s, nil}
}

// 在链表末尾插入一个节点
func AppendNode(head **Node, node *Node) {
    if *head == nil {
        *head = node
    } else {
        current := *head
        for current.next != nil {
            current = current.next
        }
        current.next = node
    }
}

// 打印链表中的所有字符串
func PrintList(head *Node) {
    current := head
    for current != nil {
        current.string.Print()
        current = current.next
    }
}

func main() {
    // 创建一个等待组，用于等待所有goroutine完成
    var wg sync.WaitGroup

    // 创建一个管道，用于传递字符串
    ch := make(chan *String)

    // 创建链表
    var head *Node

    // 输入字符串
    var data string
    fmt.Print("Please input a string: ")
    fmt.Scanln(&data)
    string1 := CreateString(data)
    node1 := CreateNode(string1)
    AppendNode(&head, node1)

    // 输入字符串
    fmt.Print("Please input another string: ")
    fmt.Scanln(&data)
    string2 := CreateString(data)
    node2 := CreateNode(string2)
    AppendNode(&head, node2)

    // 开始处理字符串
    for current := head; current != nil; current = current.next {
        // 增加等待组计数器
        wg.Add(1)

        // 开启goroutine进行字符串处理
        go func(s *String) {
            defer wg.Done()

            // 打印字符串
            s.Print()

            // 连接字符串
            concatenated := ConcatenateStrings(s, string1)
            ch <- concatenated

            // 反转字符串
            reversed := ReverseString(s)
            ch <- reversed
        }(current.string)
    }

    // 关闭管道
    close(ch)

    // 从管道中读取处理后的字符串
    for s :=
