package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func checkError(err error, context string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in %s: %s\n", context, err.Error())
		os.Exit(1)
	}
}

func main(){
	// ①ソケットの作成とIP:portのbind
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
	checkError(err,"tcpAddr")
	
	// ②接続の待機
	listener, err := net.ListenTCP("tcp",tcpAddr)
	checkError(err,"listener")
	log.Println("normal socket\nlisten on port",service)

	for {
		// ③接続の受信
		conn,err := listener.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		// ④ソケットの読み込み
		req := make([]byte,1024)
		len, err := conn.Read(req)
		log.Println("riquest:",string(req[:len]))
		// ⑤ソケットの書き込み
        daytime := time.Now().String()
        conn.Write([]byte(daytime))
		// ⑥接続の切断
		conn.Close()
	}

}