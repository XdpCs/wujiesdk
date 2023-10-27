package main

// @Title        main.go
// @Description
// @Create       XdpCs 2023-10-12 12:44
// @Update       XdpCs 2023-10-20 20:25

import (
	"context"

	"github.com/XdpCs/wujiesdk"
)

func main() {
	c, err := wujiesdk.NewCredentials("appID", "PrivateKey")
	if err != nil {
		panic(err)
	}

	client := wujiesdk.NewDefaultClient(c)
	ca := wujiesdk.NewCaller(client)
	_, _, err = ca.CancelImage(context.Background(), "2087C400944DF2D6B25BED29C910B1B8")
	if err != nil {
		panic(err)
	}
}
