#!/bin/bash
VERSION="v0.1.0-alpha"
OUTPUT_FILE="./.github/workflows/release.yml"

if [ -f "$OUTPUT_FILE" ]; then
    rm "$OUTPUT_FILE"
fi
cat <<EOL > "$OUTPUT_FILE"
name: Release

on:
  push:
    tags:
      - '$VERSION'
permissions:
  contents: write
  packages: write
jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: '1.23'

      - name: Build binaries
        run: |
          chmod +x ./script/buildRelease.sh
          ./script/buildRelease.sh  \${{ github.ref }}

      - name: Create Release
        id: create_release
        uses: softprops/action-gh-release@v1
        with:
          tag_name: \${{ github.ref }}
          files: ./output/*
        env:
          GITHUB_TOKEN: \${{ secrets.GITHUB_TOKEN }}
EOL


echo "release.yml create success：$OUTPUT_FILE"
# Markdown
OUTPUT_FILE="doc/Initialize_Guide.md"

echo "# Quickly Initialize Your Project" > "$OUTPUT_FILE"

echo "## Recommend method" >> "$OUTPUT_FILE"
echo "### Wget" >> "$OUTPUT_FILE"
echo '```bash' >> "$OUTPUT_FILE"
echo "wget https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -O dc_migrations_template.zip && unzip dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"

echo "### CURL" >> "$OUTPUT_FILE"
echo '```bash' >> "$OUTPUT_FILE"
echo "curl  https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -o  dc_migrations_template.zip && unzip dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo '' >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"

echo "### Manual mode" >> "$OUTPUT_FILE"
echo "**You can manually download and extract it to the root directory of the project**" >> "$OUTPUT_FILE"
echo "[Download Template Zip](https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip)" >> "$OUTPUT_FILE"

echo "## Successful " >> "$OUTPUT_FILE"
echo "When you see the following directory structure, it means it has been successful" >> "$OUTPUT_FILE"
echo '```shell' >> "$OUTPUT_FILE"
echo "example/" >> "$OUTPUT_FILE"
echo "├── dmc.go // This is the command-line tool for gormMigration" >> "$OUTPUT_FILE"
echo "├── dc_migrations // This is the migration file directory for gormMigration" >> "$OUTPUT_FILE"
echo "└── register.go // This is the migration file registration file for gormMigration, which is automatically generated and maintained by gormMigration. Please do not manually modify it" >> "$OUTPUT_FILE"
echo "  └── 20230301_000000_create_users_table.go // This is the migration file for gormMigration" >> "$OUTPUT_FILE"
echo "├── go.mod" >> "$OUTPUT_FILE"
echo "├── go.sum" >> "$OUTPUT_FILE"
echo "└── ... you project files" >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"


echo "" >> "$OUTPUT_FILE"

#echo '# More quick initialization to your project' >> "$OUTPUT_FILE"
#
#echo "## Windows" >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### 64位 (amd64)" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -o dmInit-windows-amd64-${VERSION}.exe https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-windows-amd64-${VERSION}.exe
#.\dmInit-windows-amd64-${VERSION}.exe" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### 32位 (386)" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -o dmInit-windows-386-${VERSION}.exe https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-windows-386-${VERSION}.exe
#.\dmInit-windows-386-${VERSION}.exe" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### ARM" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -o dmInit-windows-arm-${VERSION}.exe https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-windows-arm-${VERSION}.exe
#.\dmInit-windows-arm-${VERSION}.exe" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### ARM64" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -o dmInit-windows-arm64-${VERSION}.exe https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-windows-arm64-${VERSION}.exe
# .\dmInit-windows-arm64-${VERSION}.exe
#" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
## Linux
#echo "## Linux" >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### 64位 (amd64)" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -L -o dmInit-linux-amd64-${VERSION} https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-linux-amd64-${VERSION}
#chmod +x dmInit-linux-amd64-${VERSION}
#./dmInit-linux-amd64-${VERSION}" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### ARM" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -L -o dmInit-linux-arm-${VERSION} https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-linux-arm-${VERSION}
#chmod +x dmInit-linux-arm-${VERSION}
#./dmInit-linux-arm-${VERSION}" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### ARM64" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -L -o dmInit-linux-arm64-${VERSION} https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-linux-arm64-${VERSION}
#chmod +x dmInit-linux-arm64-${VERSION}
#./dmInit-linux-arm64-${VERSION}" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
## macOS
#echo "## macOS" >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### 64位 (amd64)" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -L -o dmInit-darwin-amd64-${VERSION} https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-darwin-amd64-${VERSION}
#chmod +x dmInit-darwin-amd64-${VERSION}
#./dmInit-darwin-amd64-${VERSION}" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
#echo "### ARM64" >> "$OUTPUT_FILE"
#echo '```bash' >> "$OUTPUT_FILE"
#echo "curl -L -o dmInit-darwin-arm64-${VERSION} https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dmInit-darwin-arm64-${VERSION}
#chmod +x dmInit-darwin-arm64-${VERSION}
#./dmInit-darwin-arm64-${VERSION}" >> "$OUTPUT_FILE"
#echo '```' >> "$OUTPUT_FILE"
#echo "" >> "$OUTPUT_FILE"
#
echo "done：$OUTPUT_FILE"

git add .
#git tag -d "${VERSION}"
git commit -m "release ${VERSION}"
git tag -a "${VERSION}" -m "release ${VERSION}"
git push origin "${VERSION}"
git push origin main