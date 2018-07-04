# testgin

[![Build Status](https://www.travis-ci.org/xujintao/testgin.svg?branch=master)](https://www.travis-ci.org/xujintao/testgin)

testgin is a gin based server.
## Quick start
1. Prepare a json config file  
```json
{
    "db":{
        "name": "mysql",
        "user": "root",
        "password": "1234",
        "ip": "127.0.0.1",
        "port": 3306,
        "table":"test"
    }
}
```

2. Run it  
For example, your config file is $GOPATH/src/github.com/xujintao/config/config.json  

* Run from source directly  
```sh
$ go get -u github.com/xujintao/testgin
$ cd $GOPATH/src/github.com/xujintao/testgin
$ go build
$ ./testgin $GOPATH/src/github.com/xujintao/config/config.json

```
* or run from docker  
```sh
docker run --rm \
           -it \
           -p 8080:8080 \
           -v $GOPATH/src/github.com/xujintao/config:/etc/testgin \
           xujintao/testgin:1.0.0.21 \
           /etc/testgin/config.json
```
more tags: https://hub.docker.com/r/xujintao/testgin/tags/  

3. Test  
Open chrome, press F12, then visit http://172.0.0.1:8080 
