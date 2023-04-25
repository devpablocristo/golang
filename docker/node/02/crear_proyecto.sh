#!/bin/bash  
#Con zsh, da error, read -p, " -p: no coprocess".

# Config 
TAG_IMG=tmp_img
NOM_CONT=tmp_cont

DIR_INST=inicio/instalador
DIR_CONF=inicio/config

TIMESTAMP=$(date +"%F-%H-%M-%S")
NOM_DIR_PROY="proyecto-$TIMESTAMP"
DIR_PROY=${PWD}/$NOM_DIR_PROY

DEPENDENCIAS="npm init -y \
&& npm install express"

docker build . -f $DIR_INST/Dockerfile.inst -t $TAG_IMG \
&& mkdir $DIR_PROY \
&& docker run --rm --name $NOM_CONT -v $DIR_PROY:/app $TAG_IMG sh -c "$DEPENDENCIAS" \
&& cp -r $DIR_CONF/. $DIR_PROY \
&& sudo chown -R pablo:pablo $DIR_PROY \
&& docker rmi $TAG_IMG -f \
&& docker rm -f $NOM_CONT