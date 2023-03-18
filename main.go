package main
     
import "fmt"
 
func main() {
	for {
		var cmd string
		fmt.Println("请输入命令:")
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("命令列表:")
			fmt.Println("--quit")
			fmt.Println("--hello")
		} else if cmd == "quit" {
			break
		} else if cmd == "hello" {
			fmt.Println("hello")
		} else {
			fmt.Println("没有找到命令，需要帮助请输入 \"help\"")
		}
	}
}