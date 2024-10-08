#!/bin/bash
# filename: installpgsql.sh
# usage: ./installpgsql.sh [PGSQL VERSION]
#
# example: 
#        Install the default version(15.3)
#        $ ./installpgsql.sh
#        Install the specific version
#        $ ./installpgsql.sh 15.1
#
# reference: https://www.postgresql.org/docs/current/install-short.html
VERSION=$1
if [ ! ${VERSION} ];then
    VERSION=15.3
    echo 'set version: '${VERSION}'.'
else
    echo ${VERSION}
fi
TARGET=postgresql-${VERSION}
curl -O https://ftp.postgresql.org/pub/source/v${VERSION}/${TARGET}.tar.gz
tar xzf ${TARGET}.tar.gz
cd ${TARGET}
./configure --without-readline --without-zlib
make -j8
sudo make install
cd -
rm -rf ${TARGET}*
sudo adduser --disabled-password --gecos "" postgres
sudo mkdir -p /usr/local/pgsql/data
#sudo chown postgres /usr/local/pgsql/data
#su - postgres
#/usr/local/pgsql/bin/initdb -D /usr/local/pgsql/data
#/usr/local/pgsql/bin/pg_ctl -D /usr/local/pgsql/data -l logfile start
#/usr/local/pgsql/bin/createdb test
#/usr/local/pgsql/bin/psql test
