#! /usr/bash

BUILD_HOME=$(cd `dirname $0`;pwd)
WORK_HOME=$BUILD_HOME/..
SOURCE_HOME=$BUILD_HOME/../src/cht
PACKAGE_HOME=$BUILD_HOME/../dependPackage
TARGET_BIN=$BUILD_HOME/../release/cht_bin

set +x

function install_go_environment()
{
        echo -e "\033[31m ---------- start install go env ---------- \033[0m"
        which go
        if [ $? -ne 0 ];then
                echo -e "\033[31m ---------- start decompression go package ---------- \033[0m"
                cd $PACKAGE_HOME
                tar -zxvf go1.7.6.linux-amd64.tar.gz -C /usr/local >/dev/null 2>&1
                ln -s /usr/local/go/bin/go /usr/bin/go
                echo -e "\033[31m ---------- end decompression go package ---------- \033[0m"
        fi
        export GOPATH="$WORK_HOME:$WORK_HOME/src/thirdparty"
        export GOROOT="/usr/local/go"
        export PATH=$PATH:$GOROOT/bin

        echo -e "\033[31m ---------- end install go env ---------- \033[0m"
}

function compile()
{
		rm  $TARGET_BIN
        echo -e "\033[31m ---------- start compile ---------- \033[0m"
        cd $SOURCE_HOME
        if [ -f main.go ];then
                go build -o $TARGET_BIN
        else
                echo -e "\033[31m ---------- comile failed,cant't find source file ---------- \033[0m"
                exit 1
        fi
        echo -e "\033[31m ---------- end compile ---------- \033[0m"
}


install_go_environment
compile
