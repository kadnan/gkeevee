// Package gkeevee implements routines for data retrival
package gkeevee

import (
	"errors"
	"io"
	"os"

	"github.com/vmihailenco/msgpack"
)

const fileNotFound string = "-1"
const dbFileNotFound string = "-2"
const decodingFailed string = "-3"
const resetFailed string = "-4"

var u = make(map[string]string) //to store values
var isLoaded = false            //to check whether file was loaded or not before calling any route

// Get function retrieves the value of the given key. If failed, it returns error.
func Get(key string) (string, error) {
	if !isLoaded {
		return "0", errors.New(fileNotFound)
	}
	v, found := u[key]
	if found {
		return v, nil
	}
	return "0", errors.New(fileNotFound)
}

//Set function assigns a value to the given key. If successful it returns 1
func Set(key string, value string) (int8, error) {
	if !isLoaded {
		return -1, errors.New(dbFileNotFound)
	}
	u[key] = value
	return 1, nil
}

// Save function saves the value in the db file.
func Save(fileHandle *os.File) (int8, error) {
	if !isLoaded {
		return -1, errors.New(dbFileNotFound)
	}

	defer fileHandle.Close()

	if len(u) > 0 {
		b, err := msgpack.Marshal(u)

		if err != nil {
			return 0, errors.New(decodingFailed)
		}

		msgpack.Unmarshal([]byte(b), &u)
		fileHandle.Write(b)
	}

	return 1, nil
}

// Load function opens the file and return the handler
func Load(path string) (*os.File, error) {
	db, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	// Read file, line by line
	var text = make([]byte, 1024)

	for {
		_, err := db.Read(text)

		if err == io.EOF {
			break
		}
	}

	if len(text) != 0 {
		err := db.Truncate(0)
		if err != nil {
			return nil, errors.New(resetFailed)
		}
		_, err = db.Seek(0, 0)
		if err != nil {
			return nil, errors.New(resetFailed)
		}
		data := string(text)
		msgpack.Unmarshal([]byte(data), &u)
	}
	isLoaded = true
	return db, nil
}
