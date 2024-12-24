## 快速开始
### 第一步  克隆项目
```sh
git clone https://github.com/niehengyang/ginTemp.git
```
### 第二步  修改配置
```sh
cd ginTemp
cp etc/bussiness-example.yml etc/bussiness.yml
# 修改bussiness.yml中的配置
```

### 第三步 构建、运行项目
```
cd ginTemp/apps/demo
go build -o demo main.go
```