package modbus

import (
	"fmt"
	"os"

	"github.com/goburrow/modbus"
)

func main() {

	client := modbus.TCPClient("10.23.3.117:502")
	// Read input register 9
	results, err := client.ReadInputRegisters(8, 1)
	checkError(err)

	fmt.Println(string(results))

	// // Modbus RTU/ASCII
	// // Default configuration is 19200, 8, 1, even
	// client = modbus.RTUClient("/dev/ttyS0")
	// results, err = client.ReadCoils(2, 1)
	// checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
