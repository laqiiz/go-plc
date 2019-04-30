package slmp

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	serverIP := "10.23.3.117"
	serverPort := "1280"

	message := "5000 00 FF FF03 00 1200 1000 0114 0000 2C0100 A8 0300 0011 0031 0231"

	message = strings.Replace(message, " ", "", -1)
	fmt.Fprintf(os.Stdout, "client-command: %s\n", message)

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	src, err := hex.DecodeString(message)
	checkError(err)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write(src)

	readBuff := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readlen, err := conn.Read(readBuff)
	checkError(err)

	res := hex.EncodeToString(readBuff[:readlen])
	fmt.Println("server-response: " + res)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
