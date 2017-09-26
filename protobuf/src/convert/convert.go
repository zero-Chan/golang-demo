package convert

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/proto"

	"example/proto"
)

func TestJSONMarshal(data *example.Person, timer int) (time.Duration, error) {
	start := time.Now()

	for i := 0; i < timer; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			return time.Duration(0), err
		}
	}

	end := time.Now()

	return end.Sub(start), nil
}

func TestJSONUnmarshal(data []byte, timer int) (time.Duration, error) {

	v := &example.Person{}

	start := time.Now()

	for i := 0; i < timer; i++ {
		err := json.Unmarshal(data, v)
		if err != nil {
			return time.Duration(0), err
		}
	}

	end := time.Now()

	return end.Sub(start), nil
}

func TestXMLMarshal(data *example.Person, timer int) (time.Duration, error) {
	start := time.Now()

	for i := 0; i < timer; i++ {
		_, err := xml.Marshal(data)
		if err != nil {
			return time.Duration(0), err
		}
	}

	end := time.Now()

	return end.Sub(start), nil
}

func TestXMLUnmarshal(data []byte, timer int) (time.Duration, error) {

	v := &example.Person{}

	start := time.Now()

	for i := 0; i < timer; i++ {
		err := xml.Unmarshal(data, v)
		if err != nil {
			return time.Duration(0), err
		}
	}

	end := time.Now()

	return end.Sub(start), nil
}

func TestProtobufMarshal(data *example.Person, timer int) (time.Duration, error) {
	start := time.Now()

	for i := 0; i < timer; i++ {
		_, err := proto.Marshal(data)
		if err != nil {
			return time.Duration(0), err
		}
	}

	end := time.Now()

	return end.Sub(start), nil
}

func TestProtobufUnmarshal(data []byte, timer int) (time.Duration, error) {
	v := &example.Person{}

	start := time.Now()

	for i := 0; i < timer; i++ {
		err := proto.Unmarshal(data, v)
		if err != nil {
			return time.Duration(0), err
		}
	}

	end := time.Now()

	return end.Sub(start), nil
}

func TestJSONWriteFile(data *example.Person, timer int) (filesize int64, err error) {
	databuf, err := json.Marshal(data)

	os.Remove("test.json")

	file, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return 0, nil
	}

	for i := 0; i < timer; i++ {
		n, err := file.Write(databuf)
		if err != nil {
			return 0, err
		} else if n != len(databuf) {
			return 0, fmt.Errorf("Write databuf to test.json fail.")
		}
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}

func TestXMLWriteFile(data *example.Person, timer int) (filesize int64, err error) {
	databuf, err := xml.Marshal(data)

	os.Remove("test.xml")

	file, err := os.OpenFile("test.xml", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return 0, nil
	}

	for i := 0; i < timer; i++ {
		n, err := file.Write(databuf)
		if err != nil {
			return 0, err
		} else if n != len(databuf) {
			return 0, fmt.Errorf("Write databuf to test.json fail.")
		}
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}

func TestProtobufWriteFile(data *example.Person, timer int) (filesize int64, err error) {
	databuf, err := proto.Marshal(data)

	os.Remove("test.protobuf")

	file, err := os.OpenFile("test.protobuf", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return 0, nil
	}

	for i := 0; i < timer; i++ {
		n, err := file.Write(databuf)
		if err != nil {
			return 0, err
		} else if n != len(databuf) {
			return 0, fmt.Errorf("Write databuf to test.json fail.")
		}
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}
