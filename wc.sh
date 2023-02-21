# array containing dir paths "tools", "encryption"
declare -a arr=("tools" "encryption" "tools/scan" "tools/file" "tools/system")
#arr of all file paths in dirs 
declare -a files=()
#loop through dirs
for i in "${arr[@]}"
do
    #loop through files in dir
    for f in $i/*
    do
        #add file to files array
        files+=("$f")
    done
done
#add file "main.go" and "Admin/main.go" to files array
files+=("main.go" "Admin/main.go")

#get total wc of all files
total_wc=$(wc -l "${files[@]}" | tail -n 1 | awk '{print $1}')
echo "Total lines of code: $total_wc"
