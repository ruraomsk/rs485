package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/ruraomsk/rs232/transport"
)

var send = []byte{0xf0, 0x02, 0xfe, 0, 0}

func Start() {
	count := byte(0)
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		count++
		if count > 250 {
			count = 0
		}
		send[3] = count
		err := transport.SendToServer(send)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				transport.Reconect()
				continue
			}
			fmt.Println(err)
			transport.Reconect()
			continue
		}
		for i := 0; i < 6; i++ {
			buf, err := transport.GetFromServer()
			if err != nil {
				fmt.Println(err)
				transport.Reconect()
				break
			}
			fmt.Printf("%d:% 02x\n", i, buf)
		}
		fmt.Println("==========================================")
	}

}
