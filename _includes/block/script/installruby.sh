#!/bin/bash
# after installation of brew
# get the ruby version by the following command:
# brew search ruby
brew install ruby@3.2
# 添加 TUNA 源并移除默认源
gem sources --add https://mirrors.tuna.tsinghua.edu.cn/rubygems/ --remove https://rubygems.org/
# 列出已有源
gem sources -l
# 应该只有 TUNA 一个

bundle config mirror.https://rubygems.org https://mirrors.tuna.tsinghua.edu.cn/rubygems

# install jekyll
gem install jekyll bundler

