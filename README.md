# go-exercicio

`Install dep`

```sh
$ go get -v github.com/golang/dep/cmd/dep
```

`Install the project's dependencies`

```sh
$ dep ensure
```

`Build`

```sh
$ go build -o exercicio
```

`Run`

```sh
$ ./exercicio ./csv/contas.csv ./csv/transacoes.csv 
```

`Run unit tests`

```sh
$ go test ./... -v
```
