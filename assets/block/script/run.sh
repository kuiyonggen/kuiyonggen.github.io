#!/bin/bash
LOCAL=$1
if [ ! $LOCAL ];then
    # 添加镜像源并移除默认源
    gem sources --add https://mirrors.tuna.tsinghua.edu.cn/rubygems/ --remove https://rubygems.org/
    # 列出已有源
    gem sources -l
    # 应该只有镜像源一个
    gem install jekyll bundler
    echo 'Run Installation.'
    bundle config mirror.https://rubygems.org https://mirrors.tuna.tsinghua.edu.cn/rubygems
    REPO_DIR=$(dirname $(dirname $(dirname $(dirname $(realpath  "$0")))))
    cd ${REPO_DIR}
    cp -rf _includes/block assets/
    bundle install
    cd -
else
    echo 'Run local.'
fi
bundle exec jekyll serve --livereload --open-url --host 127.0.0.1 --port 8084
