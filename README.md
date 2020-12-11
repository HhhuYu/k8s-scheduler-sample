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

## Deploy

```sh
kubectl apply -f ./deploy
```

## Test

```sh
kubectl apply -f ./Test
```

## Build

```sh
make local
```
