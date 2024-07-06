package anbernicrc

import (
	"os"
)

var crcTable [256]uint32

func init() {
	for i := 0; i < 256; i++ {
		crc := uint32(i)
		for j := 0; j < 8; j++ {
			if (crc & 1) == 1 {
				crc = 0xEDB88320 ^ (crc >> 1)
			} else {
				crc = crc >> 1
			}
		}
		crcTable[i] = crc
	}
}

func crcToBytes(crc uint32) []byte {
	result := make([]byte, 4)
	result[0] = byte(crc & 0xFF)
	result[1] = byte((crc >> 8) & 0xFF)
	result[2] = byte((crc >> 16) & 0xFF)
	result[3] = byte((crc >> 24) & 0xFF)
	return result
}

func AppendCRC(filename string) error {

	crc, err := FileCRC(filename)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	defer file.Close()
	if _, err := file.Seek(0, 2); err != nil {
		return err
	}

	_, err = file.Write(crcToBytes(crc))
	if err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}

func FileCRC(filename string) (crc uint32, err error) {
	var file *os.File
	file, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	buffer := make([]byte, 4096)

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			return 0, err
		}
		for i := 0; i < n; i++ {
			crc = crcTable[(byte(crc)^buffer[i])] ^ (crc >> 8)
		}
		if err != nil {
			break
		}
	}

	return
}
