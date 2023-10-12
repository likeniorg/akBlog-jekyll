#!/bin/bash
# Ca文件夹未创建
tar czvf akBlogInfo.tar.gz jekyll/header.jpeg  config/cert/Ca/* jekyll/_posts/* 
rm -rf jekyll/header.jpeg  config/cert/Ca/* jekyll/_posts/*
