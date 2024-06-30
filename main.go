package main

import (
	"fmt"
	"time"

	"github.com/goburrow/serial"
	"github.com/ruraomsk/rs232/client"
	"github.com/ruraomsk/rs232/server"
	"github.com/ruraomsk/rs232/transport"

	"syscall"
)

var port serial.Port
var config serial.Config
var err error

func main() {
	var uname syscall.Utsname

	if err = syscall.Uname(&uname); err != nil {
		panic("Error name linux" + err.Error())
	}
	b := make([]rune, 0)
	for _, v := range uname.Release {
		if v == 0 {
			break
		}
		b = append(b, rune(v))
	}
	release := string(b)
	fmt.Printf("System on %s\n", release)
	baud := 19200
	switch release {
	case "3.18.20":
		config = serial.Config{Address: "/dev/com1", BaudRate: baud, StopBits: 1, Parity: "N", Timeout: 200 * time.Millisecond}

	case "6.1.37-sunxi":
		config = serial.Config{Address: "/dev/ttyS2", BaudRate: baud, StopBits: 1, Parity: "N", Timeout: 1 * time.Second}
	case "5.15.0-113-generic":
		config = serial.Config{Address: "/dev/ttyUSB0", BaudRate: baud, StopBits: 1, Parity: "N", Timeout: 200 * time.Millisecond}
	}
	// config.RS485 = serial.RS485Config{Enabled: true} //, DelayRtsBeforeSend: time.Duration(100 * time.Millisecond), RxDuringTx: true}
	transport.Start(config)
	switch release {
	case "3.18.20":
		go client.Start()
	case "5.15.0-113-generic":
		go server.Start()
	}
	for {
		time.Sleep(time.Second)
	}
}
