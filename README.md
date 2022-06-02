# go-inmem-store



Go inmem store is simple implementation of redis memstore
## Dependency
- Go
this repo only using builtin package from go.

## Test
```
$ go test ./... --race
```


## How to use 
this repo containing client - server implementation, in order to use this you need to spin up server first
```
$ cd /server
$ go build . && ./server [PORT]
```

then you can connect from multiple client by running this command
```
$ cd /client
$ go build . && ./client [HOST] [PORT]
```

## Client command features
```
- SET [Key] [Value]
- DUMP [Key]
- RENAME [oldKey] [newKey]
- DELETE [key]
- HELP
- QUIT
```

## Error Definition
```
errCmdNotFound     = "command not found"
errInvalidCommand  = "invalid command"
errKeyAlreadyExist = "key already exist"
errNotFound        = "key value not found"
```