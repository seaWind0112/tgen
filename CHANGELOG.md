### tgen 更新记录

#### v0.0.8
##### 增加
-   试增加 grpc 支持

#### v0.0.7
##### 修复
-   SomeServer 实现中, 返回的 err 不为 nil, 也需要对 SomeResponse.Value 进行赋值

#### v0.0.6
##### 增加
-   [issue50](https://github.com/seaWind0112/tgen/issues/50)

#### v0.0.5
##### 修复
-   定义为 Oneway 的方法, 不需要 res 参数

#### v0.0.4
##### 修复
-   log 正确地换行

#### v0.0.3
##### 增加
-   可选是否生成 WebApi 相关的代码
-   可选是否生成 RpcClient 相关的代码

##### 修复
-   Server 的各个 Method 总是要求 params 和 res 两个参数

#### v0.0.2
##### 增加
-   `/health/ping` 接口

#### v0.0.1
-   初始版本
