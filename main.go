package main

import (
	"encoding/json"
	"fmt"

	"github.com/ZapSign/SdkGo/models"
)

func main() {
	a := models.Signer.CreateSigner(models.Signer{})
	agua, _ := json.Marshal(a)
	fmt.Println(string(agua))
}
