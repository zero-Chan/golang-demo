package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/golang/protobuf/proto"

	"convert"
	"example/proto"
)

func main() {
	test := &example.Person{
		Name:  proto.String("Cza"),
		Id:    proto.Int32(1),
		Email: proto.String("chenziao@sunteng.com"),
	}

	err := MarshalCompare(test, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = UnmarshalCompare(test, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = SizeCompare(test, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, _ := proto.Marshal(test)
	fmt.Println(b)
}

func MarshalCompare(data *example.Person, timer int) error {
	protoExpend, err := convert.TestProtobufMarshal(data, timer)
	if err != nil {
		return err
	}

	jsonExpend, err := convert.TestJSONMarshal(data, timer)
	if err != nil {
		return err
	}

	xmlExpend, err := convert.TestXMLMarshal(data, timer)
	if err != nil {
		return err
	}

	fmt.Printf("Marshal Compare: timer[%d], protobuf[%s] json[%s] xml[%s]\n", timer, protoExpend, jsonExpend, xmlExpend)
	return nil
}

func UnmarshalCompare(data *example.Person, timer int) error {
	protodata, err := proto.Marshal(data)
	if err != nil {
		return err
	}

	protoExpend, err := convert.TestProtobufUnmarshal(protodata, timer)
	if err != nil {
		return err
	}

	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	jsonExpend, err := convert.TestJSONUnmarshal(jsondata, timer)
	if err != nil {
		return err
	}

	xmldata, err := xml.Marshal(data)
	if err != nil {
		return err
	}

	xmlExpend, err := convert.TestXMLUnmarshal(xmldata, timer)
	if err != nil {
		return err
	}

	fmt.Printf("Unmarshal Compare: timer[%d], protobuf[%s] json[%s] xml[%s]\n", timer, protoExpend, jsonExpend, xmlExpend)
	return nil
}

func SizeCompare(data *example.Person, timer int) error {
	protoSize, err := convert.TestProtobufWriteFile(data, timer)
	if err != nil {
		return err
	}

	jsonSize, err := convert.TestJSONWriteFile(data, timer)
	if err != nil {
		return err
	}

	xmlSize, err := convert.TestXMLWriteFile(data, timer)
	if err != nil {
		return err
	}

	fmt.Printf("Size Compare: timer[%d], protobuf[%dB] json[%dB] xml[%dB]\n", timer, protoSize, jsonSize, xmlSize)
	return nil
}
