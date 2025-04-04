package db

import (
	"bytes"
	"encoding/binary"
	"log"
)

func float32SliceToBytes(floats []float32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, floats)
	if err != nil {
		log.Println("failed to convert embedding to []byte:", err)
		return nil
	}
	return buf.Bytes()
}

func bytesToFloat32Slice(b []byte) ([]float32, error) {
	buf := bytes.NewReader(b)
	vec := make([]float32, len(b)/4)
	err := binary.Read(buf, binary.LittleEndian, &vec)
	return vec, err
}
