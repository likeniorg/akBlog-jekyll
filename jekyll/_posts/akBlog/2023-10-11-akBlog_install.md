---
title: akBlog  install
date: 2023-07-06 17:31:00 +8000
categories: [linux,akBlog]
tags: [akBlog_install]    
---
# akBlog利用jekyll生成静态网站
## 配置环境
### jekyll环境配置
进入jekyll文件夹并执行 install.sh
```bash
    cd akBlog/jekyll

    #需要使用sudo apt安装ruby环境
    ./install
```

## 项目编译
项目用到gin框架，国内直接访问容易被墙,解决方法:
* 翻墙
* 使用国内代理
```bash
    # 使用国内代理
    cd akBlog
    ./cmd/goProxy.sh

    # 项目编译
    go build .

    # 执行程序
    ./akBlog
```
