package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/ruraomsk/rs232/transport"
)

var replay = [][]byte{
	[]byte{0xf0, 0x02, 0x00, 0x00, 0},
	[]byte{0xf6, 0x03, 0x3f, 0xbb, 0x01, 0},
	[]byte{0xaf, 0x05, 0x0e, 0x02, 0x24, 0x49, 0x00, 0},
	[]byte{0xf0, 0x03, 0x0e, 0xfe, 0xd4, 0},
	[]byte{0xe9, 0x03, 0x0e, 0xfe, 0xd4, 0},
	[]byte{0xe4, 0x02, 0x0e, 0x00, 0},
}

func Start() {
	for {
		buf, err := transport.GetFromServer()
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				continue
			}
			fmt.Println(err)
			transport.Reconect()
			continue
		}
		fmt.Printf("% 02x\n", buf)
		for _, v := range replay {
			v[3] = buf[3]
			err := transport.SendToServer(v)
			if err != nil {
				if strings.Contains(err.Error(), "timeout") {
					transport.Reconect()
					break
				}
				fmt.Println(err)
				transport.Reconect()
				break
			}
			time.Sleep(20 * time.Millisecond)
		}
	}

}
