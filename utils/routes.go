package utils

import "fmt"

const staticRouteAPI string = "http://localhost:3001/api/v1/"

func GetDocsRoute() string {
	return fmt.Sprintf("%s%s", staticRouteAPI, "docs/")
}

func GetSignerRoute() string {
	return fmt.Sprintf("%s%s", staticRouteAPI, "signer/")
}

func GetSignersRoute() string {
	return fmt.Sprintf("%s%s", staticRouteAPI, "signers/")
}
