#!/bin/bash
# usage:   ./gitcloneall.sh <GITHUB_API_TOKEN>
GH_TOKEN=$1
if [ -z ${GH_TOKEN} ]; then
    echo 'usage: ./gitcloneall.sh <GITHUB_API_TOKEN>'
else
    WORKSPACE_DIR=$(dirname $(dirname $(dirname $(dirname $(realpath  "$0")))))
    cd ${WORKSPACE_DIR}
    curl -s -H "Accept: application/vnd.github+json" -H “Authorization: Bearer ${GH_TOKEN}“ -H "X-Github-Api-Version: 2022-11-28" https://api.github.com/users/kuiyonggen/repos | jq -r '. [].ssh_url | select( . != "git@github.com:kuiyonggen/kuiyonggen.github.io.git" )' | xargs -L1 git clone
    ls
    cd -
fi
