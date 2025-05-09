package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error, context string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in %s: %s\n", context, err.Error())
		os.Exit(1)
	}
}


func main(){
	// ①実行の際に指定したhost:portでbind
	//ipアドレスとポートを与えることができる
	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
    }
	service := os.Args[1]

	// host:portの文字列をtcpaddrに変換
	// ②tcpAddrにipとポートが入ったオブジェクトが格納される
	// net.ResolveTCPAddrは「文字列→構造体」変換
	tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
	checkError(err,"tcpAddr")

	fmt.Printf("service:%s\n",service)
	fmt.Printf("tcpAddr:%s\n",tcpAddr)
	
	// ③サーバ側に接続
	conn, err := net.DialTCP("tcp",nil,tcpAddr)
	checkError(err,"conn")

	fmt.Printf("conn:%s\n",conn)

	// ④ソケットにデータの書き込み
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err,"conn write")
	res := make([]byte,1024)

	// fmt.Printf("ress:%s",res)
	
	// ⑤ソケットからデータの読み込み
	// サーバからのhttpレスポンスの内容がresに入る
	// よみ取ったバイト数がlenに入る
	len, err := conn.Read(res)
	checkError(err,"conn read")
	fmt.Println("response:",string(res[:len]))

	// ⑥接続の切断
    conn.Close()
}