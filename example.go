package main

import (
	"fmt"

	gkeevee "github.com/kadnan/gKeeVee/gKeeVee"
)

func main() {
	f, err := gkeevee.Load("mytesting.db")
	if err != nil {
		println(err)
	}
	gkeevee.Set("FNAME", "ADNAN")
	gkeevee.Set("LNAME", "SIDDIQI")

	val, err := gkeevee.Get("FNAME")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	_, err = gkeevee.Save(f)
}
