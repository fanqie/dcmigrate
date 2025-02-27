#!/bin/bash
VERSION="v0.3.1"
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
OUTPUT_FILE="doc/en/Initialize_Guide.md"

echo "# Quickly Initialize Your Project" > "$OUTPUT_FILE"

echo "## Recommend method" >> "$OUTPUT_FILE"
echo "### Wget" >> "$OUTPUT_FILE"
echo '```bash' >> "$OUTPUT_FILE"
echo "wget https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -O dc_migrations_template.zip " >> "$OUTPUT_FILE"
echo "unzip dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo "rm -f dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo '' >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"

echo "### CURL" >> "$OUTPUT_FILE"
echo '```bash' >> "$OUTPUT_FILE"
echo "curl -L https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -o  dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo "unzip dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo "rm -f dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo '' >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"

echo "### POWER SHELL" >> "$OUTPUT_FILE"
echo '```bash' >> "$OUTPUT_FILE"
echo "curl  https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -o  dc_migrations_template.zip" >> "$OUTPUT_FILE"
echo 'Expand-Archive -Path dc_migrations_template.zip -DestinationPath .' >> "$OUTPUT_FILE"
echo 'Remove-Item -Path dc_migrations_template.zip' >> "$OUTPUT_FILE"
echo '' >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"




echo "### Manual mode" >> "$OUTPUT_FILE"
echo "**You can manually download and extract it to the root directory of the project**" >> "$OUTPUT_FILE"
echo "[Download Template Zip](https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip)" >> "$OUTPUT_FILE"

echo "## Successful " >> "$OUTPUT_FILE"
echo "When you see the following directory structure, it means it has been successful" >> "$OUTPUT_FILE"
echo '```shell' >> "$OUTPUT_FILE"
echo "example/" >> "$OUTPUT_FILE"
echo "├── dmc.go // This is the command-line tool for dcmigration" >> "$OUTPUT_FILE"
echo "├── dc_migrations // This is the migration file directory for dcmigration" >> "$OUTPUT_FILE"
echo "└── register.go // This is the migration file registration file for dcmigration, which is automatically generated and maintained by dcmigration. Please do not manually modify it" >> "$OUTPUT_FILE"
echo "  └── 20230301_000000_create_users_table.go // This is the migration file for dcmigration" >> "$OUTPUT_FILE"
echo "├── go.mod" >> "$OUTPUT_FILE"
echo "├── go.sum" >> "$OUTPUT_FILE"
echo "└── ... you project files" >> "$OUTPUT_FILE"
echo '```' >> "$OUTPUT_FILE"

echo "" >> "$OUTPUT_FILE"

OUTPUT_FILE_CN="doc/zh_cn/Initialize_Guide.md"
echo "# Quickly Initialize Your Project" > "$OUTPUT_FILE_CN"
echo "# 快速初始化您的项目" > "$OUTPUT_FILE_CN"

echo "## 推荐方法" >> "$OUTPUT_FILE_CN"
echo "### Wget" >> "$OUTPUT_FILE_CN"
echo '```bash' >> "$OUTPUT_FILE_CN"
echo "wget https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -O dc_migrations_template.zip " >> "$OUTPUT_FILE_CN"
echo "unzip dc_migrations_template.zip" >> "$OUTPUT_FILE_CN"
echo "rm -f dc_migrations_template.zip" >> "$OUTPUT_FILE_CN"
echo '' >> "$OUTPUT_FILE_CN"
echo '```' >> "$OUTPUT_FILE_CN"

echo "### CURL" >> "$OUTPUT_FILE_CN"
echo '```bash' >> "$OUTPUT_FILE_CN"
echo "curl -L  https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -o  dc_migrations_template.zip" >> "$OUTPUT_FILE_CN"
echo "unzip dc_migrations_template.zip" >> "$OUTPUT_FILE_CN"
echo "rm -f dc_migrations_template.zip" >> "$OUTPUT_FILE_CN"
echo '' >> "$OUTPUT_FILE_CN"
echo '```' >> "$OUTPUT_FILE_CN"

echo "### POWER SHELL" >> "$OUTPUT_FILE_CN"
echo '```bash' >> "$OUTPUT_FILE_CN"
echo "curl  https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip -o  dc_migrations_template.zip" >> "$OUTPUT_FILE_CN"
echo 'Expand-Archive -Path dc_migrations_template.zip -DestinationPath .' >> "$OUTPUT_FILE_CN"
echo 'Remove-Item -Path dc_migrations_template.zip' >> "$OUTPUT_FILE_CN"
echo '' >> "$OUTPUT_FILE_CN"
echo '```' >> "$OUTPUT_FILE_CN"


echo "### 手动模式" >> "$OUTPUT_FILE_CN"
echo "**您可以手动下载并将其解压到项目的根目录**" >> "$OUTPUT_FILE_CN"
echo "[下载模板压缩包](https://github.com/fanqie/dcmigrate/releases/download/${VERSION}/dc_migrations_template.zip)" >> "$OUTPUT_FILE_CN"

echo "## 成功" >> "$OUTPUT_FILE_CN"
echo "当您看到以下目录结构时，说明已成功" >> "$OUTPUT_FILE_CN"
echo '```shell' >> "$OUTPUT_FILE_CN"
echo "example/" >> "$OUTPUT_FILE_CN"
echo "├── dmc.go // 这是用于 dcmigration 的命令行工具" >> "$OUTPUT_FILE_CN"
echo "├── dc_migrations // 这是 dcmigration 的迁移文件目录" >> "$OUTPUT_FILE_CN"
echo "└── register.go // 这是 dcmigration 的迁移文件注册文件，由 dcmigration 自动生成和维护。请勿手动修改" >> "$OUTPUT_FILE_CN"
echo "  └── 20230301_000000_create_users_table.go // 这是 dcmigration 的迁移文件" >> "$OUTPUT_FILE_CN"
echo "├── go.mod" >> "$OUTPUT_FILE_CN"
echo "├── go.sum" >> "$OUTPUT_FILE_CN"
echo "└── ... 您的项目文件" >> "$OUTPUT_FILE_CN"
echo '```' >> "$OUTPUT_FILE_CN"

echo "" >> "$OUTPUT_FILE_CN"

echo "done：$OUTPUT_FILE_CN"

git add .
git tag -d "${VERSION}"
git commit -m "release ${VERSION}"
git tag -a "${VERSION}" -m "release ${VERSION}"
git push origin "${VERSION}" -f
git push origin main