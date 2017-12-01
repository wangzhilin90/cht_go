#! /usr/bash

BUILD_HOME=$(cd `dirname $0`;pwd)
SOURCE_HOME=$BUILD_HOME/../src/cht
PACKAGE_HOME=$BUILD_HOME/../dependPackage

function install_go_environment()
{
	echo -e "\033[31m start install go env ... \033[0m"
	which go 
	if [ $? -ne 0 ];then
		echo -e "\033[31m start decompression go package ... \033[0m"
		cd $PACKAGE_HOME
		tar -zxvf go1.7.6.linux-amd64.tar.gz -C /usr/local
		mv /usr/local/go/bin/go /usr/bin
		export GOPATH="$SOURCE_HOME;$SOURCE_HOME/src/thirdparty"
		export GOROOT="/usr/local/go"
		echo $GOPATH
		echo $GOROOT
		echo -e "\033[31m end decompression go package ... \033[0m"
	fi
	echo -e "\033[31m end install go env.... \033[0m"
}

function compile()
{
	echo -e "\033[31m start compile ... \033[0m"
	cd $SOURCE_HOME
	if [ -f main.go ];then 
		go build main.go
	else
		echo -e "\033[31m comile failed,cant't find source file \033[0m"
		exit 1
	fi
	echo -e "\033[31m end compile ... \033[0m"
}

function install_tool()
{
	which dos2unix
	if [ $? -ne 0];then
		yum install -y dos2unix
	fi
	dos2unix build.sh
}

install_tool
install_go_environment
compile


		