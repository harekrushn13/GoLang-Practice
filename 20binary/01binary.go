package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type packet struct {
	Sensid uint32
	Locid  uint16
	Tstamp uint32
	Temp   int16
}

func main() {
	buf := make([]byte, 10)
	ts := uint32(time.Now().Unix())

	binary.LittleEndian.PutUint16(buf[0:], 0xa20c) // sensorID
	//binary.BigEndian.PutUint16(buf[0:], 0xa20c) // sensorID
	binary.BigEndian.PutUint16(buf[2:], 0x04af) // locationID
	binary.BigEndian.PutUint32(buf[4:], ts)     // timestamp
	binary.BigEndian.PutUint16(buf[8:], 479)    // temp

	fmt.Println(buf)
	fmt.Printf("% x\n", buf)

	sensorID := binary.LittleEndian.Uint16(buf[0:])
	//sensorID := binary.BigEndian.Uint16(buf[0:])
	locID := binary.BigEndian.Uint16(buf[2:])
	tstamp := binary.BigEndian.Uint32(buf[4:])
	temp := binary.BigEndian.Uint16(buf[8:])
	fmt.Printf("sid: %0#x, locID %0#x ts: %0#x, temp:%d\n",
		sensorID, locID, tstamp, temp)

	dataIn := packet{
		Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12,
	}
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, dataIn)
	if err != nil {
		fmt.Println(err)
		return
	}

	var dataOut packet
	err = binary.Read(buff, binary.BigEndian, &dataOut)
	if err != nil {
		fmt.Println("failed to Read:", err)
		return
	}
	fmt.Println(dataOut)
}
