package main

// // redis 协议解析的示例
// func main() {
// 	conn, err := net.Dial("tcp", "localhost:6379")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 向 Redis 发送一个命令
// 	cmd := "set k-name bitcask-kv\r\n"
// 	conn.Write([]byte(cmd))

// 	// 解析 Redis 的响应
// 	reader := bufio.NewReader(conn)
// 	res, err := reader.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(res)
// }
