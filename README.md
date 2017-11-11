# govendoring
Go Vendor Sandbox 

## Prepare your linux box
If you are running Ubuntu do following steps
```
sudo apt-get install golang-1.7-go
echo 'export PATH=/usr/lib/go-1.7/bin/:$PATH' >> ~/.bashrc
cd govendoring
export GOPATH=$PWD
go get -u github.com/kardianos/govendor
export PATH=$GOPATH/bin:$PATH
```

## Start to use it
init
```
govendor init
```
pull all dependencies from network remotes
```
govendor fetch +out
```
