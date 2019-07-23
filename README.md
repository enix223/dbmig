# dbmig 

A simple database verion tracking program learning from django migraion mechanism, currently support `postgresql` database only.

## Install

```sh
go get -u github.com/enix223/dbmig
```

## How to use

Command arguments:

```
Usage of dbmig:
  -dir string
        Migration dir (default ".")
  -host string
        Database host (default "localhost")
  -name string
        Database name (default "postgres")
  -password string
        Database password (default "password")
  -port string
        Database port (default "5432")
  -table string
        Migration table (default "ishare_migrations")
  -user string
        Database user (default "postgres")
  -version
        Print version
```

### 1. Create migration table

```
psql -d <your_database_name> < scripts/init_db.sql
```

### 2. Create migrations scripts in `migrtions` dir

```sh
ls -l migrations

-rw-r--r--  1 enix  staff  3597 Jun 27 17:16 20190220_001_init_db.sql
-rw-r--r--  1 enix  staff   318 Jun 26 16:20 20190626_001_add_fk.sql
-rw-r--r--  1 enix  staff    68 Jun 27 17:17 20190627_001_add_name.sql
```

### 3. Make migrations

```
# dbmig will check if the sql has been run before or not, and run it
# if it haven't run before
dbmig -dir migrations/ \
    -host database \
    -port 5432 \
    -name postgres \
    -username postgres \
    -pass postgres
```
