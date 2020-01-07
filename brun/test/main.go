package main

import (
	"encoding/json"
	"fmt"
	"github.com/vonnwang/resk/services"
)

func main() {

	d, e := json.Marshal(&services.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
