package transport

import (
	"errors"
	"fmt"

	"github.com/goburrow/serial"
)

var port serial.Port
var config serial.Config
var err error

func Start(c serial.Config) {
	config = c
	port, err = serial.Open(&config)
	if err != nil {
		panic("Error " + err.Error())
	}
}
func Reconect() {
	port.Close()
	port, err = serial.Open(&config)
	if err != nil {
		panic("Error " + err.Error())
	}
}
func Crc(buffer []byte) byte {
	crc := 0
	for i := 0; i < len(buffer)-1; i++ {
		crc += int(buffer[i])
	}
	return byte(crc & 0xff)
}
func Is_Crc(buffer []byte) bool {
	crc := Crc(buffer)
	return crc == buffer[len(buffer)-1]
}

func GetFromServer() ([]byte, error) {
	body := make([]byte, 0)
	start := make([]byte, 2)
	_, err := port.Read(start)
	if err != nil {
		return body, err
	}
	body = append(body, start...)
	tail := make([]byte, start[1]+1)
	_, err = port.Read(tail)
	if err != nil {
		return body, err
	}
	body = append(body, tail...)
	is := Is_Crc(body)
	if !is {
		fmt.Printf("!!% 02x\n", body)
		return body, errors.New("bad CRC")
	}
	return body, nil
}

func SendToServer(buffer []byte) error {
	s := make([]byte, 0)
	for i := 0; i < len(buffer)-1; i++ {
		s = append(s, buffer[i])
	}
	s = append(s, Crc(buffer))
	n, err := port.Write(s)
	if err != nil {
		return err
	}
	if n != len(buffer) {
		return errors.New("bad write")
	}
	// fmt.Printf("->% 02x\n", s)
	return nil
}
