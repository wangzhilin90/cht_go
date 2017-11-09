#! /bin/bash

TEST_SUITE_PATH=`find -name *_test.go`
dir=""
file=""

function get_file_path()
{
        for i in $TEST_SUITE_PATH
        do
        dir=`dirname $i`
        filename=`basename $i`
        cd $dir
        go test -v | grep -i pass
        if [ $? -ne 0 ]
        then
        echo -e "\033[33m 失败用例是:$filename \033[0m"
        fi
        cd -
        done
}

function run_all_suite()
{
        get_file_path
}

run_all_suite
