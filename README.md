# 配置环境

## jekyll环境配置
进入jekyll/文件夹并执行 ./tool/install.sh
```bash
    cd jekyll
   # 需要使用sudo apt安装ruby环境生成网页代码
    ./tool/install
```

## 项目编译
项目用到gin框架，国内直接访问容易被墙,解决方法:
* 翻墙
* 使用国内代理
```bash
# 使用国内代理
./cmd/goProxy.sh

# 项目编译
go build .

# 执行程序
./akBlog
```
## 后台访问
后台访问需要通过防火墙并且IP为服务器IP

## tool工具解读
```
# 编译_posts中的md文件，转换为前端页面
build.sh

# 删除依赖环境恢复默认
init.sh

# 安装jekyll环境
install.sh
```

## cmd工具解读
```
# 切换Go代理
goProxy.sh

# 恢复网站默认状态
init.sh

# outputConfig.sh
将写入的头像和文章打包压缩为akBlogInfo.tar.gz，然后删除原本路径中的信息
```
