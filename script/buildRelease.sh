#!/bin/bash
full_ref="$1"  #refs/tags/vx.x.x-alpha
version="${full_ref#refs/tags/}"
platforms=(
#    "linux/amd64"
    "linux/386"
#    "linux/arm64"
#    "linux/arm"
#    "windows/amd64"
#    "windows/386"
#    "windows/arm64"
#    "windows/arm"
#    "darwin/amd64"
#    "darwin/arm64"
)

extensions=(
#    ""      # linux/amd64
    ""      # linux/386
#    ""      # linux/arm64
#    ""      # linux/arm
#    ".exe"  # windows/amd64
#    ".exe"  # windows/386
#    ".exe"  # windows/arm64
#    ".exe"  # windows/arm
#    ""      # darwin/amd64
#    ""      # darwin/arm64
)

for i in "${!platforms[@]}"; do
    IFS="/" read -r os arch <<< "${platforms[$i]}"
    extension="${extensions[$i]}"
    output="./output/dmInit-${os}-${arch}-${version}${extension}"

    echo "Building for ${os}/${arch}..."
    GOOS=$os GOARCH=$arch go build -o "$output" ./dmInit.go
done
echo "Build completed!"
echo "start build template!"
mkdir -p "./output/basic"
chmod +x "./output/dmInit-linux-386-${version}"
cd "./output/basic" || exit
"../dmInit-linux-386-${version}"
pwd
zip -r dc_migrations_template.zip *
cd ..
rm -rf ./basic
rm -rf "./dmInit-linux-386-${version}"
echo "zip completed"

