package utils

const staticRouteAPI string = "http://localhost:3001/api/v1/"

func GetDocsRoute() string {
	return staticRouteAPI + "docs/"
}

func GetSignerRoute() string {
	return staticRouteAPI + "signer/"
}

func GetSignersRoute() string {
	return staticRouteAPI + "signers/"
}
