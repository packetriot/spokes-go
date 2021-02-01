package json

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

func Decode(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

func DecodeFile(path string, obj interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		err = Decode(file, obj)
		file.Close()
	}
	return err
}

func EncodeToBytes(obj interface{}) []byte {
	buffer := bytes.Buffer{}
	encoder := json.NewEncoder(&buffer)

	err := encoder.Encode(obj)
	if err != nil {
		return nil
	}

	return buffer.Bytes()
}

func Encode(w io.Writer, obj interface{}) error {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(obj); err != nil {
		return err
	}
	return nil
}

func EncodePretty(w io.Writer, obj interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(obj); err != nil {
		return err
	}
	return nil
}

func EncodePrettyFile(path string, obj interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(obj); err != nil {
		return err
	}
	return nil
}

func EncodeFile(path string, obj interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		err = Encode(file, obj)
		file.Close()
	}
	return err
}
