# Metafile DFN

To fool around with the mysql database of Metafile, particularly tables related with DFN.


## Techs:
- mysql.
- [mssql](https://github.com/denisenkom/go-mssqldb).
- all-drivers: https://github.com/golang/go/wiki/SQLDrivers


### mysql
#### connection string
```
username:password@tcp(ip:port)/database
```

### mssql

#### connection string:

```
Server=myServerAddress;Database=myDatabase;User Id=myUsername;Password=myPassword;Failover Partner=myMirror;Max Pool Size=200;Compatibility Mode=Sybase
```
