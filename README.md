# 201702-go-db-orm
ORM (object relational mapping) with GORM framework to Golang backend language


## Installation (linux debian Ubuntu)

### install postgres
- sudo apt-get update
- sudo apt-get install postgresql postgresql-contrib
- sudo apt-get install pgadmin3

### Switching Over to the postgres Account
- sudo -i -u postgres
- createdb user
- psql

### change password
- alter user postgres with password 'yourpassword';

### list db
- /list

### install golang
- sudo apt install golang-go

#### Add to the ~/.bashrc
- export GOROOT="/usr/local/go"
- export GOPATH="/home/americo/go"
- export PATH="/home/americo/go/bin:$PATH"

### install GORM (ORM: object relactional mapping)
http://jinzhu.me/gorm/
- go get -u github.com/jinzhu/gorm

### install postgres driver
https://github.com/lib/pq
- go get github.com/lib/pq

### install a deep pretty printer for Go data structures
https://github.com/davecgh/go-spew
- go get -u github.com/davecgh/go-spew/spew


### install bash git prompt
https://github.com/magicmonty/bash-git-prompt
- cd ~
- git clone https://github.com/magicmonty/bash-git-prompt.git .bash-git-prompt --depth=1

#### Add to the ~/.bashrc:
- GIT_PROMPT_ONLY_IN_REPO=1
- source ~/.bash-git-prompt/gitprompt.sh
