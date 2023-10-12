package main

// @Title        main.go
// @Description
// @Create       XdpCs 2023-10-12 12:44
// @Update       XdpCs 2023-10-12 12:44

import (
	"context"
	"fmt"

	"github.com/XdpCs/wujiesdk"
)

func main() {
	c, err := wujiesdk.NewCredentials("appID", "PrivateKey")
	if err != nil {
		panic(err)
	}

	client := wujiesdk.NewDefaultClient(c)
	ca := wujiesdk.NewCaller(client)
	code, _, err := ca.CancelImage(context.Background(), wujiesdk.NewCancelImageRequest("2087C400944DF2D6B25BED29C910B1B8"))
	if err != nil {
		fmt.Println(code)
	}
}
