// Package gkeevee implements routines for data retrival
package gkeevee

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vmihailenco/msgpack"
)

var u = make(map[string]string) //to store values
var isLoaded bool = false       //to check whether file was loaded or not before calling any route

// Get function retrieves the value of the given key. If failed, it returns error.
func Get(key string) (string, error) {
	if !isLoaded {
		return "-", errors.New("The db file was not loaded")
	}
	v, found := u[key]
	if found {
		v = u[key]
		return v, nil
	}
	return "-", errors.New("-1")
}

//Set function assigns a value to the given key. If successful it returns 1
func Set(key string, value string) (int8, error) {
	if !isLoaded {
		return -1, errors.New("The db file was not loaded")
	}
	u[key] = value
	return 1, nil
}

// Save function saves the value in the db file.
func Save(fileHandle *os.File) (int8, error) {
	if !isLoaded {
		return -1, errors.New("The db file was not loaded")
	}
	_, err := os.Stat(fileHandle.Name())

	if err != nil {
		fmt.Println(err.Error())
	}
	if os.IsNotExist(err) {
		return -1, errors.New("-1")
	}
	defer fileHandle.Close()

	if len(u) > 0 {
		b, _ := msgpack.Marshal(u)

		msgpack.Unmarshal([]byte(b), &u)
		fmt.Println(u)
		fileHandle.WriteString(string(b))
	}

	return 1, nil
}

// Load function opens the file and return the handler
func Load(path string) (*os.File, error) {
	db, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	// Read file, line by line
	var text = make([]byte, 1024)

	for {
		_, err := db.Read(text)

		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
			break
		}
	}
	if err != nil {
		return nil, err
	}

	if string(text) != "" {
		db.Truncate(0)
		db.Seek(0, 0)
		//db.WriteString("")
		data := string(text)
		data = strings.TrimSpace(data)
		msgpack.Unmarshal([]byte(data), &u)
		//fmt.Println(u)
	}
	isLoaded = true
	return db, nil
}
