#!/bin/bash
git pull
cp -rf ../kyg/_site/* .
git status
git add .
git commit -am "new post"
git push
git status

