package SenseAir_K30_go

import (
	"fmt"
	"golang.org/x/exp/io/i2c"
)

type K30 struct {
	Dev     *i2c.Device
	co2Read []byte
}

const (
	CO2_ADDR = 0x68
)

//
// Opens a file system handle to the I2C device
//  and sets the device
//
func (k30 *K30) K30Init(channel string) int {
	k30.co2Read = make([]byte, 4)
	k30.co2Read[0] = 0x22
	k30.co2Read[1] = 0x00
	k30.co2Read[2] = 0x08
	k30.co2Read[3] = 0x2A
	response := make([]byte, 4)
	var err error
	device.Dev, err = i2c.Open(&i2c.Devfs{Dev: channel}, CO2_ADDR)
	if err != nil {
		panic(err)
	}
	return 0
} /* K30Init() */

//
// Read the sensor register value
// and translate into calibrated reading
//

func (k30 *K30) K30ReadValue() int {
	result := make([]byte, 4)

	err = k30.Dev.Write(k30.co2Read)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	inBuf := []byte{0x00, 0x00, 0x00, 0x00}
	time.Sleep(38 * time.Millisecond)
	err = k30.Dev.Read(inBuf)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	co2_value := 0
	co2_value |= int(inBuf[1]) & int(0xFF)
	co2_value = co2_value << 8
	co2_value |= int(inBuf[2]) & int(0xFF)
	sum := 0                                            //Checksum Byte
	sum = int(inBuf[0]) + int(inBuf[1]) + int(inBuf[2]) //Byte addition utilizes overflow
	if sum != int(inBuf[3]) {
		fmt.Println("Checksum Error")
		return -1
	}
	//fmt.Println("CO2: ", co2_value)
	return co2Value
}
