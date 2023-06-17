#!/bin/bash
COMMENT=$1
if [ -z "${COMMENT}" ];
then
    COMMENT="no comment!"
else
    echo ${COMMENT}
fi
git add .
git commit -am "${COMMENT}"
git push
