function j { 
    iscmd=1
    [ "$1" != "add" ] && [ "$1" != "quickadd" ] && [ "$1" != "delete" ] && [ "$1" != "list" ] && [ "$1" != "edit" ] && [ "$1" != "-h" ] && [ "$1" != "--help" ] \
       && [ "$1" != "completion" ] && [ "$1" != "__complete" ] && [ "$1" != "__completeNoDesc" ] && iscmd=0
    if [ $iscmd -eq 1 ]; then
        jgo $@
    else
        rsp=$(jgo $@)
        if [ $? -eq 0 ]; then
            if [ -n "$rsp" ]; then
                cd $rsp
            fi
        else
            echo $rsp
        fi
    fi  
}

j $@