//https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/
package main

import (
	"fmt"

	gkeevee "github.com/kadnan/gKeeVee/gKeeVee"
)

func main() {
	// _, err := gkeevee.Load("mytesting.db")
	// if err != nil {
	// 	println(err)
	// }
	// gkeevee.Set("FNAME", "Adu")
	// gkeevee.Set("LNAME", "SIDDI")
	// gkeevee.Set("CITY", "KARACHI")
	// r, _ := gkeevee.Save(f)
	val, err := gkeevee.Get("FNAME")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
