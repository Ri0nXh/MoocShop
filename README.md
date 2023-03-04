# MoocShop
> 慕课视频cap老师的电商项目课程为基础，使用Gin框架构建的项目


# 目录结构说明
```
├─configs           # 配置文件目录
├─internal          # 内部目录
│  ├─api            # api接口
│  ├─dao            # dao层，与数据库交互
│  ├─initialize     # 服务器初始化
│  ├─model          # 模型层
│  └─router         # 路由层
├─logs              # 日志目录
└─wiki              # sql等相关目录
```


# 商品库存问题
**锁实现超卖问题**

最先使用锁来实现超卖，测试平均性能如下, QPS 在1100左右
```
Running 30s test @ http://192.168.1.16:8000/api/v1/order/create
  6 threads and 600 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   573.75ms   43.06ms 682.95ms   83.25%
    Req/Sec   230.12    240.84     0.99k    79.88%
  31005 requests in 30.08s, 4.91MB read
  Socket errors: connect 0, read 0, write 513, timeout 0
Requests/sec:   1030.67
Transfer/sec:    167.31KB
```

**悲观锁实现超卖问题**

这里应用mysql数据库本身的特性，即对于`UPDATE、DELETE、INSERT`语句，`InnoDB`会自动将涉及的数据集添加排他锁(X锁)。将`update`商品数量的sql语句上移，并用`update`操作后返回的结果来判断是否更新成功，来实现数据库悲观锁

