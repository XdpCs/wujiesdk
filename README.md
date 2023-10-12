# wujiesdk

![GitHub watchers](https://img.shields.io/github/watchers/XdpCs/wujiesdk?style=social)
![GitHub stars](https://img.shields.io/github/stars/XdpCs/wujiesdk?style=social)
![GitHub forks](https://img.shields.io/github/forks/XdpCs/wujiesdk?style=social)
![GitHub last commit](https://img.shields.io/github/last-commit/XdpCs/wujiesdk?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/XdpCs/wujiesdk?style=flat-square)
![GitHub license](https://img.shields.io/github/license/XdpCs/wujiesdk?style=flat-square)

无界ai sdk https://torna.wujiebantu.com/#/share/7YX0lXxR/yq2RExzg

## 安装

`go get`

```shell
go get -u github.com/XdpCs/wujiesdk
```

`go mod`

```shell
require github.com/XdpCs/wujiesdk latest
```

## 完成情况

- [x] 用户账户
- [x] Ai作画
    - [x] 作画流程
    - [x] 单独超分
    - [x] 作画队列
    - [ ] 定制API
- [ ] 化身训练
- [ ] 化身作画
- [ ] 咒语生成
- [ ] 咒语解析
- [ ] 魔法骰子
- [ ] 视频生成视频
- [ ] Ai专业版作画
    - [ ] 作画流程
        - [x] 专业版发起AI作画
        - [x] 专业版作画轮询接口
        - [ ] 专业版作画查询接口
        - [ ] ControlNet type preprocessor model 参数依赖关系查询
        - [ ] 获取专业版模型列表

## 例子

```go
package main

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
```