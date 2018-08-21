# SenseAir_k30_go

## A Native Golang driver for the SenseAir K30 CO2 sensor.

## Usage

Create a K30 Object:

    k30 := K30{}

Initialize it with the proper device (for the Raspberry Pi 3, this works):

    dev := "/dev/i2c-1"
	r := k30.K30Init(dev)

Returns 0 on success, -1 on failure.

To read values:

    reading := k30.K30ReadValue()
	

returned value will be the CO2 reading in Parts Per Million (PPM) or -1 on failure.

It's a good idea to call

    defer k30.Dev.Close()

after the call to init so that the device will be closed after you're done with it. 

## Error Handling

There is none at this point. You have to handle them. but if the return is -1, it's likely that an error occured. It is very likely that subsequent readings will succeed, but is not guaranteed. I've left the print statements in at this point to help with error diagnosis.