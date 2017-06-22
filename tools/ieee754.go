package tools

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strconv"
)

func HexToIEEE754(input string, output *string, byteOrder string) error {
	log.Printf("[HexToIEEE754] input -> 0x%s, output -> %s, byte order -> %s", input, *output, byteOrder)
	if input == "" {
		return nil
	}
	buf, err := strconv.ParseInt(input, 16, 32)
	buff := int32(buf)
	if err != nil {
		return fmt.Errorf("[HexToIEEE754] Input [0x% X] translate error: %v", input, err)
	}
	inputByte := bytes.NewBuffer([]byte{})
	switch byteOrder {
	case "bigend":
		binary.Write(inputByte, binary.BigEndian, &buff)
		*output = fmt.Sprintf("%f", math.Float32frombits(binary.BigEndian.Uint32(inputByte.Bytes())))
	default:
		binary.Write(inputByte, binary.LittleEndian, &buff)
		*output = fmt.Sprintf("%f", math.Float32frombits(binary.LittleEndian.Uint32(inputByte.Bytes())))
	}
	log.Printf("[HexToIEEE754] input bytes -> 0x%X, output -> %s", inputByte.Bytes(), *output)
	return nil
}

func IEEE754ToHex(input string, output *string, byteOrder string) error {
	log.Printf("[IEEE754ToHex] input -> 0x%s, output -> %s, byte order -> %s", input, *output, byteOrder)
	f, err := strconv.ParseFloat(input, 32)
	ff := float32(f)
	if err != nil {
		return fmt.Errorf("[IEEE754ToHex] Input [0x% X] translate error: %v", input, err)
	}
	outputBytes := make([]byte, 4)
	switch byteOrder {
	case "bigend":
		binary.BigEndian.PutUint32(outputBytes, math.Float32bits(ff))
		*output = string(outputBytes[:4])
	default:
		binary.LittleEndian.PutUint32(outputBytes, math.Float32bits(ff))
		*output = string(outputBytes[:4])
	}
	log.Printf("[IEEE754ToHex] input -> %s, output bytes -> 0x%X", input, outputBytes)
	return nil
}
