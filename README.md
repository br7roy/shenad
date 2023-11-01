# shenad

## Build
**生成shenad**
```
sh configure && make
```

## start
### 1.配置sqllite文件位置
更改项目中配置文件`conf.toml`  
比如`food.sqlite`文件在/opt/food.sqlite下面  
那么需要配置
```toml
dbFilePath = "/opt/food.sqlite"
```

### 2.启动
**默认20000端口无需变更，配置文件和shenad在同级目录**
```shell
sh shenad

```

**如果配置文件和shenad在同级目录那么**
```shell
sh shenad -p 0.0.0.0:20000

```

**如果需要制定配置文件目录那么**
```shell
sh shenad -p 0.0.0.0:20000 -c /opt/conf.toml
```
##  预览
![预览图片](https://github.com/br7roy/shenad/blob/master/Xnip2023-11-02_00-13-39.jpg)

### LICENSE  
MIT
