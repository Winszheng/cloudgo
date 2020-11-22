# 开发简单web服务程序cloudgo

## 一、实现功能

### 基本功能

-   支持静态文件服务@
-   支持简单js访问@
-   提交表单，并输出一个表格(使用模板)@
-   测试：curl测试、ab测试@

### 拓展功能

-   访问权限限制
-   统一错误处理
-   日志文件输出
-   通过源码分析、解释一些关键功能的实现
-   对gorilla/mux库和zap库做源码分析，解释它们是如何实现拓展的原理，包括一些golang程序设计技巧(函数式编程)。

拓展功能的实现说明具体见博客，地址为：

本Readme仅介绍作业完成的内容与测试情况。

## 二、测试

### 1 curl测试

![image-20201122024339634](/home/wucp/git/cloudgo/img/image-20201122024339634.png)

可见能够提供静态文件服务。然而，我们不希望服务器上的所有资源都能够被访问，有些文件是我们不希望暴露给用户的：

创建文件sudo.txt，修改权限：

![image-20201122190223686](/home/wucp/git/cloudgo/img/image-20201122190223686.png)

结果可见无法访问：

![image-20201122190327239](/home/wucp/.config/Typora/typora-user-images/image-20201122190327239.png)

### 2 简单js访问

![image-20201122184529801](/home/wucp/.config/Typora/typora-user-images/image-20201122184529801.png)

### 3 使用模板，提交表单输出表格

![image-20201122200222673](/home/wucp/git/cloudgo/img/image-20201122200222673.png)

提交之后：

![image-20201122200246914](/home/wucp/.config/Typora/typora-user-images/image-20201122200246914.png)

### 4 ab测试

ubuntu20.04下安装apache2-utils：

```
sudo apt install apache2-utils
```

发出请求，其中-n表示执行的请求数量，-c表示并发请求个数，可见1000个请求全部完成：

![image-20201122192043016](/home/wucp/.config/Typora/typora-user-images/image-20201122192043016.png)

解释出现的参数：

-   Server Hostname: 服务器主机名
-   Server Port: 服务器端口
-   Document Path: 文件路径
-   Document Length: 文件大小
-   Concurrency Level: 并发等级
-   Time taken for tests: 整个 ab 测试消耗的总时间
-   Complete requests: 完成的请求数
-   Failed request: 失败的请求数
-   Write errors: 写入错误数
-   Total transferred: 传输的数据的字节数
-   HTML transferred: 传输的 HTML 文件的字节数
-   Requst per second: 平均每秒的请求个数
-   Time per request: 用户平均请求等待时间
-   Time per request(across all concurrent requests): Time per request / Concurrency Level
-   Transfer rate: 传输的平均速率
-   Connection Times: 表内描述了所有的过程中所消耗的最小、平均、中位、最长时间
-   Percentage of the requests served within a certain time: 每个百分段的请求完成所需的时间

![image-20201122192104826](/home/wucp/.config/Typora/typora-user-images/image-20201122192104826.png)

![image-20201122192121584](/home/wucp/.config/Typora/typora-user-images/image-20201122192121584.png)



