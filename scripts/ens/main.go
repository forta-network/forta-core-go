package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/forta-network/forta-core-go/registry"
	"github.com/kelseyhightower/envconfig"
)

const (
	ModeDev  = "dev"
	ModeProd = "prod"
	ModeAll  = "all"
)

type Env struct {
	ProdAPI string `envconfig:"PROD_API" default:"https://rpc.ankr.com/polygon"`
	DevAPI  string `envconfig:"DEV_API" default:"https://rpc.ankr.com/polygon_mumbai"`
	Mode    string `envconfig:"MODE" default:"dev"`
}

var env Env

var devConfig = registry.ClientConfig{
	ENSAddress: "0x5f7c5bbBa72e1e1fae689120D76D2f334A390Ae9",
	Name:       "registry-client",
}

var prodConfig = registry.ClientConfig{
	ENSAddress: "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7",
	Name:       "registry-client",
}

func main() {
	envconfig.MustProcess("", &env)

	devConfig.JsonRpcUrl = env.DevAPI
	prodConfig.JsonRpcUrl = env.ProdAPI

	if env.Mode == ModeDev || env.Mode == ModeAll {
		c, err := registry.NewClient(context.Background(), devConfig)
		if err != nil {
			panic(err)
		}
		contracts := c.RegistryContracts()
		b, _ := json.MarshalIndent(contracts, "", "  ")
		fmt.Println(string(b))
	}

	if env.Mode == ModeProd || env.Mode == ModeAll {
		c, err := registry.NewClient(context.Background(), prodConfig)
		if err != nil {
			panic(err)
		}
		contracts := c.RegistryContracts()
		b, _ := json.MarshalIndent(contracts, "", "  ")
		fmt.Println(string(b))
	}
}
