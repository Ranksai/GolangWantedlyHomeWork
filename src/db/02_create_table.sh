#!/bin/bash

#
# postgresql にデータ入れる
#
USER="postgres"
DATABASE="wantedlyhomework"
CREATEFILE="create_table.sql"
echo "create table"
cd /docker-entrypoint-initdb.d/
psql -h localhost -U ${USER} -d ${DATABASE} < ${CREATEFILE}
echo "finish"

