# callMysql

## 环境配置
- k8s1.27.2
- helm

## 安装

```shell
helm install app ./app-helm
```

## 验证

```shell
scp -r app-helm/ root@192.168.106.136:/root
```
