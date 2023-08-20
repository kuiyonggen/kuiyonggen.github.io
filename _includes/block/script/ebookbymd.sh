#!/bin/bash
# ebookbymd.sh can convert a markdown file to both a epub and a azw3 files.
# dependencies: pandoc & calibre utility
mdFile=$1
if [ ! $mdFile ];then
    echo "Usage:"
    echo "./ebookbymd.sh <markdown file>"
    echo ""
    echo "markdown file name must follow the below rule:"
    echo "Title separated by the minus separator(-), eg: Effective-Go.md"
    echo ""
    echo "Example:"
    echo "./ebookbymd.sh ../../doc/md/Effectie-Go.md"
else
    mdFilename=$(basename $mdFile)
    filenameNoExt=${mdFilename%.md}
    title=$(sed "s/-/ /g" <<< "$filenameNoExt")
    epub=$filenameNoExt.epub
    azw3=$filenameNoExt.azw3
    pandoc -o $epub --toc-depth=5 --metadata title="$title" $mdFile
    ebook-convert $epub $azw3
fi
