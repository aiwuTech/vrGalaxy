#!/bin/sh
#手动做成 by mint
#时间：2014/02/27 00:02
#作用：go语言环境变量加载

#提示没有在pkgconfig路径找到libzmq.pc文件，主要是pkgconfig路径的问题，只要配置一下pkgconfig目录给用户环境变量PKG_CONFIG_PATH即可．


export GOPATH=$HOME/develop/work/vrGalaxy
export GOBIN=$GOPATH/bin
export GOROOT=$HOME/develop/tools/go-1.3
export PATH=$PATH:$GOROOT/bin:$GOBIN
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
