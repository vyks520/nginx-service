# **windows安装Nginx服务**

## 使用说明
如Nginx安装在D:/nginx目录  
将“nginx-service.exe”复制到“D:/nginx”目录下（只编译了64位程序，32位需自己编译）  
然后执行下面命令安装，注意命令行需以管理员身份运行  

安装服务
```cmd
cd /d D:/nginx
nginx-service.exe install
```
启动服务
```cmd
net start nginx-service
```
卸载服务
```cmd
net stop nginx-service
nginx-service.exe remove
```
重载nginx配置（登录身份是本地系统用户会导致权限不足重载失败）
```cmd
nginx-service.exe reload
```
