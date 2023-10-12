# 编译
bundle exec jekyll build

#删除不必要文件
rm -rf _site/*.xml rm -rf _site/*.js rm -rf _site/*.json

#将生成的网页移动到文件夹
cp -r _site/* ../web/
# 忽略配置文件上传github
echo "config/*" > ../.gitignore
echo "web/*" >> ../.gitignore
echo "akBlogInfo.tar.gz" >> ../.gitignore
