# Quickly Initialize Your Project
## Recommend method
### Wget
```bash
wget https://github.com/fanqie/dcmigrate/releases/download/v0.1.3/dc_migrations_template.zip -O dc_migrations_template.zip 
unzip dc_migrations_template.zip
rm -f dc_migrations_template.zip

```
### CURL
```bash
curl -L https://github.com/fanqie/dcmigrate/releases/download/v0.1.3/dc_migrations_template.zip -o  dc_migrations_template.zip
unzip dc_migrations_template.zip
rm -f dc_migrations_template.zip

```
### POWER SHELL
```bash
curl  https://github.com/fanqie/dcmigrate/releases/download/v0.1.3/dc_migrations_template.zip -o  dc_migrations_template.zip
Expand-Archive -Path dc_migrations_template.zip -DestinationPath .
Remove-Item -Path dc_migrations_template.zip

```
### Manual mode
**You can manually download and extract it to the root directory of the project**
[Download Template Zip](https://github.com/fanqie/dcmigrate/releases/download/v0.1.3/dc_migrations_template.zip)
## Successful 
When you see the following directory structure, it means it has been successful
```shell
example/
├── dmc.go // This is the command-line tool for dcmigration
├── dc_migrations // This is the migration file directory for dcmigration
└── register.go // This is the migration file registration file for dcmigration, which is automatically generated and maintained by dcmigration. Please do not manually modify it
  └── 20230301_000000_create_users_table.go // This is the migration file for dcmigration
├── go.mod
├── go.sum
└── ... you project files
```

