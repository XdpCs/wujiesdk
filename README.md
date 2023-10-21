# wujiesdk

![GitHub watchers](https://img.shields.io/github/watchers/XdpCs/wujiesdk?style=social)
![GitHub stars](https://img.shields.io/github/stars/XdpCs/wujiesdk?style=social)
![GitHub forks](https://img.shields.io/github/forks/XdpCs/wujiesdk?style=social)
![GitHub last commit](https://img.shields.io/github/last-commit/XdpCs/wujiesdk?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/XdpCs/wujiesdk?style=flat-square)
![GitHub license](https://img.shields.io/github/license/XdpCs/wujiesdk?style=flat-square)

无界ai sdk https://apifox.com/apidoc/shared-ecc069df-a9d5-4c86-b723-6dcd5cc79f81

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

- [x] 用户开放
    - [x] 获取用户积分可用余额
    - [x] 向指定用户(手机号)发起积分兑换
- [x] Ai作画
    - [x] 获取预设资源
    - [x] 获取模型列表
    - [x] 作画结果查询
    - [x] 作画参数查询接口, 一次最多可查询6个key的作画参数
    - [x] 作画成功后的图片详情查询
    - [x] 获取单个模型排队信息
    - [x] 获取风格模型的预设资源
    - [x] 发起AI作画
    - [x] 加速作画
    - [x] 单张图片进行超分辨处理
    - [x] 超分结果批量查询
    - [x] 撤销作画
    - [x] 计算作画成本
    - [x] 提交描述词优化任务
    - [x] 图片年轻化接口
    - [x] 作画咒语查询
    - [x] 查询描述词优化任务结果
- [x] 化身训练
    - [x] 化身训练图片检测
    - [x] 化身详情查询
    - [x] 化身删除
    - [x] 创建化身
- [x] 化身作画
    - [x] 化身作画接口
    - [x] 化身作画资源选项接口
- [x] 咒语解析
    - [x] 发起咒语解析
    - [x] 查询解析结果
- [x] 魔法骰子
    - [x] 魔法骰子主题列表
    - [x] 魔法骰子生成
- [ ] 视频生成视频
    - [x] 发起视频生视频
    - [x] 视频生成成功后的视频详情查询
    - [ ] 计算视频生视频成本
    - [ ] 获取视频生视频模型列表及价格表
    - [ ] 视频生视频模型排队情况查询
    - [ ] 视频生成结果查询
- [ ] Ai专业版作画
    - [x] 专业版发起AI作画
    - [x] 专业版作画轮询接口
    - [ ] 专业版作画查询接口
    - [ ] ControlNet type preprocessor model 参数依赖关系查询
    - [ ] 获取专业版模型列表
    - [ ] 账户时长余额
    - [ ] 账户账单
    - [ ] Ai实验室
        - [ ] 分割一切
        - [ ] 矢量图
        - [ ] 实验室-下拉选项列表
        - [ ] 实验室-作画详情
        - [ ] 一镜到底
        - [ ] 删除实验室作品

## 例子

```go
package main

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

```