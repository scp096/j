# j
A linux cmdline tool to set jump path easily and faster.

## Install
#### Install the binary
	go install github.com/scp096/j
#### Edit the bashrc file
###### Add following content in ~/.bashrc
	function j { 
    iscmd=1
    [ "$1" != "add" ] && [ "$1" != "quickadd" ] && [ "$1" != "delete" ] && [ "$1" != "list" ] && [ "$1" != "edit" ] && [ "$1" != "-h" ] && [ "$1" != "--help" ] && iscmd=0 
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

## Usage
	1. To fast jump to a path:
		j [shortcut]
	2. To add a shortcut:
		j add [shortcut]=[path]
	3. To quickadd a shortcut:
		j quickadd [shortcut]
	4. To delete a shortcut:
		j delete [shortcut]
	5. To list all shortcut:
		j list
	6. To edit all shortcuts:
		j edit`