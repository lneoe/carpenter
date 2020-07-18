![Go](https://github.com/lneoe/carpenter/workflows/Go/badge.svg)
# carpenter
carpenter is a user management tool for [trojan](https://github.com/trojan-gfw/trojan)

# usage
```bash
$ export DB_DSN="user:password@(host:port)/dbname"
$ carpenter user --dsn=$DB_DSN create -u test1 -p raw_pwd_for_test1 -q 10
$ carpenter user --dsn=$DB_DSN list
id      username        download        upload  quota
13      test1   0.00GB  0.00GB  -0.00GB
14      test2   0.00GB  0.00GB  -0.00GB
15      test3   0.00GB  0.00GB  -0.00GB
$ carpenter user --dsn=$DB_DSN delete --id=15
```

# build
```bash
$ make build
$ ls ./bin
carpenter

# build cross platform
$ GOOS=linux GOARCH=amd64 make build
```