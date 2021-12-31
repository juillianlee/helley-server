## Golang hot reload
``` 
npm i -g npm 
```

### Execute
```
nodemon --exec go run main.go --signal SIGTERM
```

### Run Test
nodemon --exec "go test -v ./..." --signal SIGTERM
