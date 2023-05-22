## Development Guide
### Directory and Files
<pre style="font-size: 15px; font-weigt: bold;">
|-- <a target="_blank" href="api/">api</a>: developed set of APIs
    |-- <a target="_blank" href="api/auth">auth</a>
        |-- <a target="_blank" href="api/auth/controllers">controllers</a>: controller package for api auth
            |- <a target="_blank" href="api/auth/controllers/auth.go">auth.go</a>
        |-- <a target="_blank" href="api/auth/domain">domain</a>: interface and entity package for api auth domain
            |- <a target="_blank" href="api/auth/domain/entity.go">entity.go</a>
            |- <a target="_blank" href="api/auth/domain/interface.go">interface.go</a>
        |-- <a target="_blank" href="api/auth/mapping">mapping</a>: response data mapping package for api auth
            |- <a target="_blank" href="api/auth/mapping/permission_policy_user.go">permission_policy_user.go</a>
        |-- <a target="_blank" href="api/auth/repositories">repositories</a>: repository package that can be used for api auth
            |- <a target="_blank" href="api/auth/repositories/auth.go">auth.go</a>
            |- <a target="_blank" href="api/auth/repositories/mongodb.go">mongodb.go</a>
            |- <a target="_blank" href="api/auth/repositories/mysql.go">mysql.go</a>
            |- <a target="_blank" href="api/auth/repositories/postgresql.go">postgresql.go</a>
        |-- <a target="_blank" href="api/auth/services">services</a>: service package that can be used for the auth api
            |- <a target="_blank" href="api/auth/services/auth.go">auth.go</a>
    |-- <a target="_blank" href="api/bank">bank</a>
        |-- <a target="_blank" href="api/bank/controllers">controllers</a>: controller package for api bank
            |- <a target="_blank" href="api/bank/controllers/bank.go">bank.go</a>
        |-- <a target="_blank" href="api/bank/domain">domain</a>: interface and entity package for api bank domain
            |- <a target="_blank" href="api/bank/domain/entity.go">entity.go</a>
            |- <a target="_blank" href="api/bank/domain/interface.go">interface.go</a>
        |-- <a target="_blank" href="api/bank/mapping">mapping</a>: response data mapping package for api bank
            |- <a target="_blank" href="api/bank/mapping/bank.go">bank.go</a>
        |-- <a target="_blank" href="api/bank/repositories">repositories</a>: repository package that can be used for api bank
            |- <a target="_blank" href="api/bank/repositories/bank.go">bank.go</a>
            |- <a target="_blank" href="api/bank/repositories/mongodb.go">mongodb.go</a>
            |- <a target="_blank" href="api/bank/repositories/mysql.go">mysql.go</a>
            |- <a target="_blank" href="api/bank/repositories/postgresql.go">postgresql.go</a>
        |-- <a target="_blank" href="api/bank/services">services</a>: service package that can be used for the bank api
            |- <a target="_blank" href="api/bank/services/bank.go">bank.go</a>
    |-- <a target="_blank" href="api/file">file</a>
        |-- <a target="_blank" href="api/file/controllers">controllers</a>: controller package for api file
            |- <a target="_blank" href="api/file/controllers/file.go">file.go</a>
        |-- <a target="_blank" href="api/file/domain">domain</a>: interface and entity package for api file domain
            |- <a target="_blank" href="api/file/domain/entity.go">entity.go</a>
            |- <a target="_blank" href="api/file/domain/interface.go">interface.go</a>
        |-- <a target="_blank" href="api/file/repositories">repositories</a>: repository package that can be used for api file
            |- <a target="_blank" href="api/file/repositories/file.go">file.go</a>
            |- <a target="_blank" href="api/file/repositories/mongodb.go">mongodb.go</a>
            |- <a target="_blank" href="api/file/repositories/mysql.go">mysql.go</a>
            |- <a target="_blank" href="api/file/repositories/postgresql.go">postgresql.go</a>
        |-- <a target="_blank" href="api/file/services">services</a>: service package that can be used for the file api
            |- <a target="_blank" href="api/file/services/file.go">file.go</a>
    |-- <a target="_blank" href="api/mail">mail</a>
        |-- <a target="_blank" href="api/mail/controllers">controllers</a>: controller package for api mail
            |- <a target="_blank" href="api/mail/controllers/mail.go">mail.go</a>
        |-- <a target="_blank" href="api/mail/domain">domain</a>: interface and entity package for api mail domain
            |- <a target="_blank" href="api/mail/domain/entity.go">entity.go</a>
            |- <a target="_blank" href="api/mail/domain/interface.go">interface.go</a>
        |-- <a target="_blank" href="api/mail/repositories">repositories</a>: repository package that can be used for api mail
            |- <a target="_blank" href="api/mail/repositories/mail.go">mail.go</a>
            |- <a target="_blank" href="api/mail/repositories/mongodb.go">mongodb.go</a>
            |- <a target="_blank" href="api/mail/repositories/mysql.go">mysql.go</a>
            |- <a target="_blank" href="api/mail/repositories/postgresql.go">postgresql.go</a>
        |-- <a target="_blank" href="api/mail/services">services</a>: service package that can be used for the mail api
            |- <a target="_blank" href="api/mail/services/mail.go">mail.go</a>
    |-- <a target="_blank" href="api/rajaongkir">rajaongkir</a>
        |-- <a target="_blank" href="api/rajaongkir/controllers">controllers</a>: controller package for api rajaongkir
            |- <a target="_blank" href="api/rajaongkir/controllers/rajaongkir.go">rajaongkir.go</a>
        |-- <a target="_blank" href="api/rajaongkir/domain">domain</a>: interface and entity package for api rajaongkir domain
            |- <a target="_blank" href="api/rajaongkir/domain/entity.go">entity.go</a>
            |- <a target="_blank" href="api/rajaongkir/domain/interface.go">interface.go</a>
        |-- <a target="_blank" href="api/rajaongkir/repositories">repositories</a>: repository package that can be used for api rajaongkir
            |- <a target="_blank" href="api/rajaongkir/repositories/rajaongkir.go">rajaongkir.go</a>
            |- <a target="_blank" href="api/rajaongkir/repositories/mongodb.go">mongodb.go</a>
            |- <a target="_blank" href="api/rajaongkir/repositories/mysql.go">mysql.go</a>
            |- <a target="_blank" href="api/rajaongkir/repositories/postgresql.go">postgresql.go</a>
        |-- <a target="_blank" href="api/rajaongkir/services">services</a>: service package that can be used for the rajaongkir api
            |- <a target="_blank" href="api/rajaongkir/services/rajaongkir.go">rajaongkir.go</a>
|-- <a target="_blank" href="cmd/">cmd</a>: go command which can be used to support development
    |- <a target="_blank" href="cmd/db_migration.go">db_migration.go</a>
|-- <a target="_blank" href="entity/">entity</a>: usable models and entities
    |- <a target="_blank" href="entity/bank.go">bank.go</a>
    |- <a target="_blank" href="entity/permission_policy_user.go">permission_policy_user.go</a>
|-- <a target="_blank" href="package/">package</a>: main package to run this application
    |-- <a target="_blank" href="package/config">config</a>: application configuration loader
        |- <a target="_blank" href="package/config/config.go">config.go</a>
        |- <a target="_blank" href="package/config/entity.go">entity.go</a>
    |-- <a target="_blank" href="package/database">database</a>: database manager for used in application
        |- <a target="_blank" href="package/database/database.go">database.go</a>
        |- <a target="_blank" href="package/database/entity.go">entity.go</a>
        |- <a target="_blank" href="package/database/vars.go">vars.go</a>
    |-- <a target="_blank" href="package/log">log</a>: logging package for the application
        |- <a target="_blank" href="package/log/interface.go">interface.go</a>
        |- <a target="_blank" href="package/log/log.go">log.go</a>
        |- <a target="_blank" href="package/log/trace.go">trace.go</a>
    |-- <a target="_blank" href="package/manager">manager</a>: package manager to prepare all application needs before the server started
        |- <a target="_blank" href="package/manager/manager.go">manager.go</a>
    |-- <a target="_blank" href="package/middleware">middleware</a>: global middleware that can be used for api routers
        |- <a target="_blank" href="package/middleware/authenticated.go">authenticated.go</a>
    |-- <a target="_blank" href="package/server">server</a>: package to prepare the server that will be used to run the application
        |-- <a target="_blank" href="package/server/middleware">middleware</a>: middleware packages that servers can use only
            |- <a target="_blank" href="package/server/middleware/cors.go">cors.go</a>
            |- <a target="_blank" href="package/server/middleware/recovery.go">recovery.go</a>
        |- <a target="_blank" href="package/server/server.go">server.go</a>
|-- <a target="_blank" href="routes/">routes</a>: route definitions for api and other handlers
    |-- <a target="_blank" href="routes/api">api</a>: router package to define the route for each api
        |- <a target="_blank" href="routes/api/auth.go">auth.go</a>
        |- <a target="_blank" href="routes/api/bank.go">bank.go</a>
        |- <a target="_blank" href="routes/api/file.go">file.go</a>
        |- <a target="_blank" href="routes/api/mail.go">mail.go</a>
        |- <a target="_blank" href="routes/api/rajaongkir.go">rajaongkir.go</a>
    |- <a target="_blank" href="routes/routes.go">routes.go</a>
|-- <a target="_blank" href="test/">test</a>: golang testing needs
    |-- <a target="_blank" href="test/mockdata">mockdata</a>: package and dummy data for testing needs
        |- <a target="_blank" href="test/mockdata/bank.go">bank.go</a>
        |- <a target="_blank" href="test/mockdata/database.go">database.go</a>
        |- <a target="_blank" href="test/mockdata/local_file.go">local_file.go</a>
        |- <a target="_blank" href="test/mockdata/manager.go">manager.go</a>
        |- <a target="_blank" href="test/mockdata/permission_policy_user.go">permission_policy_user.go</a>
    |- <a target="_blank" href="test/coverage.sh">coverage.sh</a>
    |- <a target="_blank" href="test/run.sh">run.sh</a>
|-- <a target="_blank" href="utils/">utils</a>: utility to use this app
    |-- <a target="_blank" href="utils/api/">api</a>
        |-- <a target="_blank" href="utils/api/error_handler/">error_handler</a>
        |-- <a target="_blank" href="utils/api/response/">response</a>: packages to make it easy to deliver api response content
            |- <a target="_blank" href="utils/api/response/response_failed.go">response_failed.go</a>
            |- <a target="_blank" href="utils/api/response/response_success.go">response_success.go</a>
            |- <a target="_blank" href="utils/api/response/new_response.go">new_response.go</a>
    |-- <a target="_blank" href="utils/helpers/">helpers</a>: helper package for applications
        |- <a target="_blank" href="utils/helpers/formatter.go">formatter.go</a>
        |- <a target="_blank" href="utils/helpers/jwt.go">jwt.go</a>
        |- <a target="_blank" href="utils/helpers/password.go">password.go</a>
|- <a target="_blank" href="main.go">main.go</a>: main script to run the application
|- <a target="_blank" href="base.env">base.env</a>: local environment variable
</pre>

### Run Development Server
##### Using native go run
```sh
go run main.go
```

##### With Air Live reload development
Install Air Live reload <a target="_blank" href="https://github.com/cosmtrek/air">https://github.com/cosmtrek/air</a>
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
