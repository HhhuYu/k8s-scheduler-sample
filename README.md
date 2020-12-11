# kubernetes-scheduler-framework-demo

## requirement

mertics-server

## plugin

- demoplugin  
主要用来测试
- nodelabel  
nodelabel selector
- socreplugin  
score计算优先级

## demoplugin

### clone

```bash
git clone https://github.com/HhhuYu/k8s-scheduler-sample
```

### 编译

```bash
go update
make clean
make local
```

### 制作镜像

```bash
docker build -t $ALIYUN/charstal/k8s-scheduler:v0.0.1 .
```

### 上传镜像

```bash
docker push $ALIYUN/charstal/metrics-server:v0.4.1 
```

## 测试结果

测试环境使用`kind(kubernete in docker)`搭建
k8s: v1.15.12

### k8s node

![](https://imgbed.momodel.cn/20201031170824.png)

### 配置插件

![](https://imgbed.momodel.cn/20201031171646.png)

### 部署

![](https://imgbed.momodel.cn/20201031171807.png)

```bash
# 查看schedule状态
 kubectl describe pods scheduler-framework-demo-786d8bb44f-l2vhw -n kube-system
```

![](https://imgbed.momodel.cn/20201031171930.png)

```bash
# 查看 schedule log
kubectl logs scheduler-framework-demo-786d8bb44f-l2vhw -n kube-system
```

![](https://imgbed.momodel.cn/20201031172115.png)

#### pod 效果

![](https://imgbed.momodel.cn/20201031172201.png)

![](https://imgbed.momodel.cn/20201031172214.png)

## scoreplugin

[metrics-server](https://github.com/kubernetes-sigs/metrics-server/tree/master)

[metrics-api](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/resource-metrics-api.md)

By default Usage is the mean from samples collected within the returned time window. The default time window is 1 minute.

如果收集不到metrics，可以尝试修改yaml文件，在`args`内添加配置



```bash
args:
    – /metrics-server
    – –kubelet-insecure-tls
    – –kubelet-preferred-address-types=InternalIP

```

### 操作

和上面一样

```bash
kubectl apply -f ./test/nginx-test.yaml
```

### 效果

![](https://imgbed.momodel.cn/20201110190900.png)