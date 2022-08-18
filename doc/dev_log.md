### 2022.08.02
之前的一个项目配置文件初始化全放置在 main.go 中，由于此次初始化配置项较多，若仍放在 main.go 中显得十分凌乱，遂采用 starter 模式来启动应用。
启动完成后暴露如下问题：
1. 发生 panic ,***程序崩溃***
2. panic 的异常是空指针异常，发生在 PusherList 执行 exec 函数阶段。

优先排查问题2，先解决本身的逻辑问题，通过日志及断点方式查看到 PusherList 已经初始化成功，并成功推入一个 DingTalkPusher 的元素。但当执行 push 方法时，该对象为nil，其 Type() 方法打印的值又确实为 dingtalk（这里有问题）。接着查看初始化时 PusherOfDingTalk 的值，该对象正常被初始化。也就是说被推入 PusherList 的元素为空的 DingTalkPusher 对象，而原本应该被推入 PusherList 的值 PusherOfDingTalk 是正常的，有没有可能本来注册到 PusherList 的值就是个空值，换句话说 PusherList 注册 DingTalkPusher 是在 DingTalkPusher 被初始化之前。事实也确实如此，主要原因还是在于对 init 函数的执行顺序不熟悉，**当当前文件依赖于其他包时，会优先执行其他包的init 函数**, 解决方法是在 DingTalkStarter 初始化 DingTalkPusher 完成后，直接将其注册到 PusherList 中。
```go
	global.PusherOfDingtalk = dt.NewDingTalkPusher(dtc)
	// 注册pusher
	sd.PusherList.RegisterPusher(global.PusherOfDingtalk)
```
更改后，还是发生了 panic，**但该 panic 没有使程序崩溃**，仍然是空指针异常。这次是 CommitJob 函数执行时报出的异常，由于有了前一次的经验，再次怀疑时 FlowController 没有初始化，结果也确实如此。
***那么奇怪的问题就来了，当注册启动器的初始化顺序如下时，FlowController 无法被初始化***

```go
	infra.Register(&starters.LogStarter{}, &starters.DingTalkStarter{}, &starters.RecvStarter{}, &starters.FlowControlStarter{})
```

调换顺序后，正常被初始化
```go
	infra.Register(&starters.FlowControlStarter{}, &starters.LogStarter{}, &starters.DingTalkStarter{}, &starters.RecvStarter{})
```
原因：starters.RecvStarter{} 阻塞了后面的初始化过程，因为它启动了监听服务。

回到问题一：程序崩溃，gin启动过程中已经使用recovery组件
```go
	r.Use(gin.Recovery())
```
也就说执行 push message 操作的时候，gin的主协程无法捕获这些panic，导致程序崩溃。
* 主程中的panic本身是会被gin拦截的
* 协程中的panic需要使用defer和recover进行保护


### 2022.08.03
1. 全接口测试
2. 编写 makefile

### 2022.08.04
1. pre-commit 集成 golangci-lint
2. 根据 golangci-lint 调整代码

### 2022.08.10
由于在请求处理最后阶段将任务 push 到队列中，最后返回操作是个同步阻塞操作，遂修改为异步处理。之前也想过这么去做，可是测试发现，当主 goroutine 完成后，会把子 goroutine 的任务停掉，遂放弃；但今天又进行了一番测试，发现是正常的，便又了如下疑问：***主 goroutine 的生命周期是否对子 goroutine 的生命周期产生印象？父子 goroutine 呢？***

之前都是在 main 函数中起 goroutine，发现 main 函数的执行周期直接影响其子协程的生命周期，便认为此是 golang 中父子 goroutine 的关系，但测试表明，***除了受 main 函数所在 goroutine 的生命周期影响外，其他 goroutine 的生命周期都是独立的，不同于linux里的进程依赖，golang里，协程都是互相独立的，没有依赖（父子）关系。main函数本身也运行在一个goroutine中，main是所有协程的被依赖者，这里是个特例。***