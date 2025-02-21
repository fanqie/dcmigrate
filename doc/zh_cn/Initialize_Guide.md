# 快速初始化您的项目
## 推荐方法
### Wget
```bash
wget https://github.com/fanqie/dcmigrate/releases/download/v0.1.7/dc_migrations_template.zip -O dc_migrations_template.zip 
unzip dc_migrations_template.zip
rm -f dc_migrations_template.zip

```
### CURL
```bash
curl -L  https://github.com/fanqie/dcmigrate/releases/download/v0.1.7/dc_migrations_template.zip -o  dc_migrations_template.zip
unzip dc_migrations_template.zip
rm -f dc_migrations_template.zip

```
### POWER SHELL
```bash
curl  https://github.com/fanqie/dcmigrate/releases/download/v0.1.7/dc_migrations_template.zip -o  dc_migrations_template.zip
Expand-Archive -Path dc_migrations_template.zip -DestinationPath .
Remove-Item -Path dc_migrations_template.zip

```
### 手动模式
**您可以手动下载并将其解压到项目的根目录**
[下载模板压缩包](https://github.com/fanqie/dcmigrate/releases/download/v0.1.7/dc_migrations_template.zip)
## 成功
当您看到以下目录结构时，说明已成功
```shell
example/
├── dmc.go // 这是用于 dcmigration 的命令行工具
├── dc_migrations // 这是 dcmigration 的迁移文件目录
└── register.go // 这是 dcmigration 的迁移文件注册文件，由 dcmigration 自动生成和维护。请勿手动修改
  └── 20230301_000000_create_users_table.go // 这是 dcmigration 的迁移文件
├── go.mod
├── go.sum
└── ... 您的项目文件
```

