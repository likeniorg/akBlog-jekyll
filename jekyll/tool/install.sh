#!/bin/bash

echo "安装所需环境" 
sudo apt install  git ruby-dev npm ruby bundler gem

# 更换清华源
gem sources --add https://mirrors.tuna.tsinghua.edu.cn/rubygems/ --remove https://rubygems.org/

#gem install jekyll --user-install

# 设置本地路径
bundle config set --local path 'vendor/bundle'

# 更换清华源
bundle config mirror.https://rubygems.org https://mirrors.tuna.tsinghua.edu.cn/rubygems

# 安装所需依赖
bundle 

# 编译
bundle exec jekyll build

# 删除不必要文件
rm -rf _site/*.xml rm -rf _site/*.js rm -rf _site/*.json

# 将生成的网页移动到文件夹
cp -r _site/* ../web/

# 忽略配置文件上传github
echo "config/*" > ../.gitignore
echo "akBlogInfo.tar.gz" >> ../.gitignore
echo "web/*" >> ../.gitignore
echo "jekyll/_posts" >> ../.gitignore
echo "jekyll/header.jpeg" >> ../.gitignore

