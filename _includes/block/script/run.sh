#!/bin/bash
LOCAL=$1
if [ ! $LOCAL ];then
    echo 'Run Installation.'
    REPO_DIR=$(dirname $(dirname $(dirname $(dirname $(realpath  "$0")))))
    cd ${REPO_DIR}
    cp -rf _includes/block assets/
    bundle install
    cd -
else
    echo 'Run local.'
fi
bundle exec jekyll serve --livereload --open-url --host 127.0.0.1 --port 8084
