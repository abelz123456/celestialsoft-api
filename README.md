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
##### Using shell script
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

<hr>
## Project Architecture

<pre style="font-size: 15px; font-weigt: bold;">
-- <a target="_blank" href="api/">api</a>
-- <a target="_blank" href="cmd/">cmd</a>
-- <a target="_blank" href="entity/">entity</a>
-- <a target="_blank" href="package/">package</a>
-- <a target="_blank" href="routes/">routes</a>
-- <a target="_blank" href="test/">test</a>
-- <a target="_blank" href="utils/">utils</a>
- <a target="_blank" href="main.go">main.go</a>
- <a target="_blank" href="base.env.go">base.env.go</a>
</pre>

## Folder 1

Deskripsi Folder 1.

### Subfolder 1.1

Deskripsi Subfolder 1.1.

### Subfolder 1.2

Deskripsi Subfolder 1.2.

## Folder 2

Deskripsi Folder 2.

### Subfolder 2.1

Deskripsi Subfolder 2.1.

### Subfolder 2.2

Deskripsi Subfolder 2.2.
