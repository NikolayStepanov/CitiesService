package json_csv_convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type CSVJSONConverter interface {
	JSONtoCSVСonverter
	CSVtoJSONСonverter
}

type CSVJSONConvert struct {
}
type JSONtoCSVСonverter interface {
	JSONtoCSVСonverteObject(jsonObject []byte, fieldSequence []string) []string
	JSONtoCSVСonverte(jsonObject []byte, fieldSequence []string) [][]string
}

type CSVtoJSONСonverter interface {
	CSVtoJSONСonverteObject(csvObject []string, fieldSequence []string) []byte
	CSVtoJSONСonverte(csvObject [][]string, fieldSequence []string) []byte
}

func (C *CSVJSONConvert) GetFields(s any) []string {
	val := reflect.ValueOf(s)
	fields := make([]string, val.Type().NumField())
	for i := 0; i < val.Type().NumField(); i++ {
		fields[i] = val.Type().Field(i).Tag.Get("json")
	}
	return fields
}
func (C *CSVJSONConvert) JSONtoCSVСonverte(jsonObjects []byte, fieldSequence []string) [][]string {
	var jsonArrayObject []interface{}
	err := error(nil)
	resultRecords := [][]string{}
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonObjects, &jsonArrayObject)

	for _, jsonObject := range jsonArrayObject {
		oneJsonByte, _ := json.Marshal(jsonObject)
		resultRecord := C.JSONtoCSVСonverteObject(oneJsonByte, fieldSequence)
		resultRecords = append(resultRecords, resultRecord)
	}
	return resultRecords
}

func (C *CSVJSONConvert) CSVtoJSONСonverte(csvObjects [][]string, fieldSequence []string) []byte {
	resultRecord := []byte{}
	err := error(nil)
	var buffer bytes.Buffer

	buffer.WriteString("[")
	for i, record := range csvObjects {
		jsonRecord := C.CSVtoJSONСonverteObject(record, fieldSequence)
		buffer.WriteString(string(jsonRecord))
		if i < len(csvObjects)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	resultRecord, _ = json.MarshalIndent(rawMessage, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return resultRecord
}

func (C *CSVJSONConvert) CSVtoJSONСonverteObject(csvObject []string, fieldSequence []string) []byte {
	resultRecord := []byte{}
	err := error(nil)
	var buffer bytes.Buffer

	buffer.WriteString("{")
	for i, y := range csvObject {
		//fmt.Println(y)
		buffer.WriteString(`"` + fieldSequence[i] + `":`)
		_, fErr := strconv.ParseFloat(y, 32)
		_, bErr := strconv.ParseBool(y)
		if fErr == nil {
			buffer.WriteString(y)
		} else if bErr == nil {
			buffer.WriteString(strings.ToLower(y))
		} else {
			buffer.WriteString((`"` + y + `"`))
		}
		if i < len(csvObject)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	rawMessage := json.RawMessage(buffer.String())
	resultRecord, _ = json.MarshalIndent(rawMessage, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return resultRecord
}

func (C *CSVJSONConvert) JSONtoCSVСonverteObject(jsonObject []byte, fieldSequence []string) []string {
	var f interface{}
	err := error(nil)
	resultRecord := []string{}

	err = json.Unmarshal(jsonObject, &f)
	if err != nil {
		panic(err)
	}

	m := f.(map[string]interface{})
	for _, key := range fieldSequence {
		switch valueType := m[key].(type) {
		case string:
			resultRecord = append(resultRecord, m[key].(string))
		case int:
			str := strconv.Itoa(m[key].(int))
			resultRecord = append(resultRecord, str)
		case float64:
			str := strconv.FormatFloat(m[key].(float64), 'f', -1, 64)
			resultRecord = append(resultRecord, str)
		default:
			fmt.Println(valueType, "is of a type I don't know")
		}
	}
	return resultRecord
}
