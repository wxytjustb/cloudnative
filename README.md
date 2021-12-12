# cloudnative




k8s是模块八作业
使用ingress controller使用traefik

默认使用traefik helm部署，手动修改deployment的8000和8443端口的hostPort

```
- containerPort: 8000
  hostPort: 8000
  name: web
  protocol: TCP
- containerPort: 8443
  hostPort: 8443
  name: websecure
  protocol: TCP
```







