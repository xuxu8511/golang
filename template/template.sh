#!/bin/bash

#$1 dir name
#$2 struct name

if [ "$1" == "" ] || [ "$2" == "" ]; then 
  echo parm1: $1 parm2: $2
  exit 0
fi

cp command.go $1_command.go;
cp template_session.go $1_session.go;
cp template_config.go $1_config.go;
cp template_protocol.go $1_protocol.go
cp template.go $1.go;
mkdir ../src/$1;
sed -i "s/Template/$2/g" $1*;
sed -i "s/template/$1/g" $1*;
mv $1* ../src/$1/;
cp template_config.xml $1_config.xml;
cp template_log.xml $1_log.xml;
sed -i "s/template/$1/g" $1*;
mv $1* ../config/;