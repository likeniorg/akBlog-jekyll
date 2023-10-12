# akBlog项目架构

## 项目概述
使用Go语言创建akBlog,用于分享技术和创意内容。

## 技术栈
- 前端：ruby-jekyll
- 后端：Golang
- 框架：Gin
- 操作系统：Linux
- 安装脚本：Bash

## 架构图

- **web/**：前端页面源码
  - **files/**：上传和下载文件目录

- **jekyll/**： 将Markdown文件转换为HTML
    - **_tabs**： 左侧导航栏
    - **_data**： 用于存放数据文件
    - **_posts**：用于存放博客文章
    - **_site**： Jekyll生成的静态网站文件，用来部署网站
    - **vendor**：存放第三方的依赖库或插件
 
- **cmd/**：可执行脚本，用于辅助管理网站
    - **goProxy.sh**：切换Go代理源为中国源
    - **init.sh**：恢复初始化环境
    - **outputConfig.sh**：导出配置文件

- **app/**：源代码


## 安全构思
```
# 配置文件涉及证书，不允许其他用户读写
	os.Mkdir("config/cert/", 0700) 
	os.Mkdir("config/cert/adminCa", 0700)

```
# webpage()
避免非法路径导入其他文件

## 未完成功能
* config.go 更改配置文件功能
* compress.go 通过Go程序来压缩文件导入导出配置
* public.go-Command() 应该单独记录所有执行命令
* 管理员后端代码不完善，新增后端访问安全级别(需要证书)
* 启动时检测端口是否占用
* crud.go 修改和删除功能未完善
