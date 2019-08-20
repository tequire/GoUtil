package config

import "fmt"

type EnvironmentEnum string

const (
	LOCAL EnvironmentEnum = "localhost"
	DEV   EnvironmentEnum = DevIdentityServer
	PROD  EnvironmentEnum = ProdIdentityServer
)

func EnvironmentEnumFromIsProd(isProd bool) EnvironmentEnum {
	if isProd {
		return PROD
	}
	return DEV
}

func (e EnvironmentEnum) WithPort(port int) EnvironmentEnum {
	return EnvironmentEnum(fmt.Sprint(e, ":", port))
}

// DevIdentityServer is the URL for the dev identity server
const DevIdentityServer = "https://identity-dev.highered.global"

// ProdIdentityServer is the URL for the dev identity server
const ProdIdentityServer = "https://identity.highered.global"

// DevJobAdsAPI is the URL for the dev JobAdsAPI
const DevJobAdsAPI = "https://highered-identity-jobad-dev.azurewebsites.net"

// ProdJobAdsAPI is the URL for the prod JobAdsAPI
const ProdJobAdsAPI = "https://highered-jobadsapi-identity-prod.azurewebsites.net"

// DevHigheredAPI is the URL for the dev GetHighered-API
const DevHigheredAPI = "https://api.gethighered-dev.global"

// ProdHigheredAPI is the URL for the prod GetHighered-API
const ProdHigheredAPI = "https://api.gethighered.global"
