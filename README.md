# go_schedule
## 简介
用go语言编写的分布式任务调度服务，主要特点如下:
1. 内存型任务调度。所有任务存储在内存中，快速的进行任务调度
2. 采用一致性哈希算法。每个节点负责部分任务
3. crontab语法。调度时间使用类crontab语法，灵活支持各种形式的调度时间
4. mongodb存储。面向文档的数据库支持任务字段的灵活扩展
5. 扩缩容自适应。利用zookeeper感知集群节点的变化，为每个节点重新分配任务

## 相关依赖
1. zookeeper
2. mongodb
3. grpc

## 模块介绍
 包 | 功能 
 --- | --- 
 logic/consistent_hash | 一致性哈希算法，包括哈希值计算，节点IP获取等功能
 logic/execute | 任务执行部分，设想是通过接口或者mq通知下游，本服务不负责具体执行
 logic/schedule | 调度部分，负责管理当前节点的所有任务，每个任务下次执行时间的计算以及执行
 logic/task | 对外接口，提供增删改查能力
 dao/mongodb | 封装mongodb能力
 dao/zookeeper | 封装zookeeper能力
 client | 服务的客户端，对任务进行增删该查

 ## 服务逻辑
 ### 服务启动
1. 初始化mongodb
2. 初始化zookeeper，创建必要的节点，将本节点ip注册到zookeeper，并监听集群节点的变化
3. 根据注册的节点计算节点哈希
4. 读mongodb任务表，将属于当前节点的任务初始化并开始调度
 ### 创建任务
1. 判断是否在处理扩缩容，若是，返回错误。扩缩容期间不允许任务增删改
2. 计算任务ID及其哈希值
3. 确定该任务对应的节点IP
4. 若任务是当前节点处理，则写mongodb，初始化并调度；若不是，则进行转发
 ### 修改任务
1. 判断是否在处理扩缩容，若是，返回错误。扩缩容期间不允许任务增删改
2. 计算任务的哈希值
3. 确定该任务对应的节点IP
4. 若任务是当前节点处理，则修改mongodb对应的任务，修改内存中的任务；若不是，则进行转发
 ### 删除任务
1. 判断是否在处理扩缩容，若是，返回错误。扩缩容期间不允许任务增删改
2. 计算任务的哈希值
3. 确定该任务对应的节点IP
4. 若任务是当前节点处理，则删除mongodb对应的任务，删除内存中的任务，停止调度；若不是，则进行转发
 ### 查询任务
查询mongodb，将任务返回
 ### 集群扩缩容
 1. 重新监听集群节点的变化
 2. 重新初始化当前节点的任务
 ### 任务初始化
 1. 利用zookeeper加锁，不允许在初始化过程中增删改
 2. 重新获取当前所有节点，并计算其哈希值
 3. 读mongodb任务表，若属于当前节点并且当前节点没有改任务，则创建；若不属于当前节点并且当前节点有该任务，则删除
