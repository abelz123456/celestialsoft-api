## Development Guide

### Run Development Server
##### Using native go run
```sh
go run main.go
```

##### With Air Live reload development
Install Air Love reload <a target="_blank" href="https://github.com/cosmtrek/air">https://github.com/cosmtrek/air</a>
```sh
go install github.com/cosmtrek/air@latest
```
Run
```sh
air
```


### Run Database Migration
database configuration in [base.env](base.env) file. database migration will only succeed for value **APP_ENV=development**

##### Using native go run
```sh
go run cmd/db_migration.go
```

##### Using makefile
```sh
make migration
```


### Running Go test
###### Using shell script
Run testing oly
```sh
test/run.sh
```

Run test and get coverage test result
```sh
test/coverage.sh
```

##### Using makefile
Run test and get coverage test result for all directories
```sh
make runtest
```

Run test and get coverage test result with selected directories
```sh
make runtest DIR=<directory>
```
Example: run test for directory api
```sh
make runtest DIR=./api/...
```