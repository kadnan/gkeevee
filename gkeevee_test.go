package main

import (
	"testing"

	gkeevee "github.com/kadnan/gKeeVee/gKeeVee"
)

func TestLoad(t *testing.T) {
	f, _ := gkeevee.Load("my.db")

	if f.Name() != "my.db" {
		t.Errorf("TestLoad Failed")
	}
}

func TestSet(t *testing.T) {
	result, err := gkeevee.Set("Name", "Adnan")

	if err != nil && result == -1 {
		t.Errorf("Test Set Failed")
	}
}

//TestGet implements...
func TestGet(t *testing.T) {
	result, _ := gkeevee.Get("Adnan")

	if result != "0" {
		t.Errorf("TestGet Failed")
	}
}
