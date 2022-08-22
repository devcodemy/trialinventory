# trialinventory
Go Lang implementation of Big Inventory - Trial Inventory

# Some specific Go language commands
```shell
# create work directory
mkdir -p WORK_DIR_NAME

# go into work directory and initiate module
cd WORK_DIR_NAME
go mod init PROJECT_NAME

# also may be the reason reset go env
go env -w GO111MODULE=

go get github.com/gin-gonic/gin

# git flow (main branch for github)
git init -b main
git remote add origin git@github.com:devcodemy/trialinventory.git

# git pull if it needs

git remote -v
git pull origin main

# add and commit a code
git add .
git commit -m "initial commit"

git push --set-upstream otigin main
```
