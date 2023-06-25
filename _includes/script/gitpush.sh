#!/bin/bash
COMMENT=$1
if [ -z "${COMMENT}" ];
then
    COMMENT="no comment!"
else
    echo ${COMMENT}
fi
REPO_DIR=$(dirname $(dirname $(dirname $(realpath  "$0"))))
cd ${REPO_DIR}
git add .
git commit -am "${COMMENT}"
git push
cd -
