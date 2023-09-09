# Challenge 01 - short.io
Create a URL shortner like the bit.ly


## Running the application inside docker - Make Version (Recommened) 
```
make build@docker
make run@docker
```
## Running the application inside docker 
```sh
docker build -t ismael3s/short.io .
docker run --rm -it -p 4000:4000 ismael3s/short.io
```

## Running tests - Make version (Recommend)
```sh
make test@unit
```

## Running tests
**Ensure that you bash is configured correct to use go packages globally**
**Install gomock uber**

```sh
go install go.uber.org/mock/mockgen@latest
go generate ./...
go test -v ./...
```


### Topics Applied
- DAO Pattern
- Dependency injection & Dependecy Inversion
- Decorator Pattern
- Simplified Hexagonal Architeture
- Docker

#### Topics to apply
- Add redis