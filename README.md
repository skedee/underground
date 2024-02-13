# Underground

Use the Underground to create the project scaffolding.  

## Setup Golang

Run make target to install or upgrade golang.

`
make setup-go
`
## Setup Environment

* Add GOPATH to environment
* Update PATH to include GOPATH

Example (Bash)
`
export GOPATH="$HOME/go"
export PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"
`

## Features Supported

* Golang
* Sqitch
* Sqlc
* Docsify
