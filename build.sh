#build go project for multiple platforms
platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386" "linux/arm" "linux/arm64")

for platform in ${platforms[@]}
do
    split=(${platform//\// })
    GOOS=${split[0]}
    GOARCH=${split[1]}
    output_name="$1_$GOOS-$GOARCH"
    #output to dist folder
    output_path=./dist/$output_name
    if [ $GOOS = "windows" ]; then
        output_path=$output_path.exe
    fi
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_path main.go
    echo "Built $output_name successfully!"
    if [ $? -ne 0 ]; then
           echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
#set all as executable
chmod +x ./dist/*
#number of packages built
platformLength=${#platforms[@]}
echo "Built $platformLength packages successfully!"