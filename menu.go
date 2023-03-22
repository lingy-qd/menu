package main

import (
	"fmt"
	"os"
)

type DataNode struct {
	cmd     string
	desc    string
	handler func()
	next    *DataNode
}

func FindCmd(head *DataNode, cmd string) *DataNode {
	if head == nil || len(cmd) == 0 {
		return nil
	}
	for p := head; p != nil; p = p.next {
		if p.cmd == cmd {
			return p
		}
	}
	return nil
}

func ShowAllCmd(head *DataNode) {
	fmt.Println("Menu List:")
	for p := head; p != nil; p = p.next {
		fmt.Println(p.cmd, " - ", p.desc)
	}
}

var head *DataNode

// some cmds
var cmds []*DataNode = []*DataNode{
	{"help", "this is help cmd!", Help, nil},
	{"version", "menu program V2.0", ShowVersion, nil},
	{"q", "exit menu", Exit, nil},
}

func main() {
	//将cmds插入到链表中
	head = &DataNode{}
	tail := head
	for _, cmd := range cmds {
		tail.next = cmd
		tail = tail.next
	}
	tail.next = nil
	head = head.next

	var cmd string
	for {
		fmt.Print("Input a cmd > ")
		fmt.Scanln(&cmd)
		p := FindCmd(head, cmd)
		if p == nil {
			fmt.Println("This is a wrong cmd!")
			continue
		}
		fmt.Println(p.cmd, " - ", p.desc)
		if p.handler != nil {
			p.handler()
		}
		fmt.Println()
	}
}

func Help() {
	ShowAllCmd(head)
}
func ShowVersion() {
	fmt.Println("menu v2.0")
}
func Exit() {
	os.Exit(0)
}
