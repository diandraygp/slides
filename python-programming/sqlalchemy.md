# SQLAlchemy
{id: sqlalchemy}

## SQLAlchemy hierarchy
{id: sqlalchemy-hierarchy}

* ORM
* Table, Metadata, Reflection, DDL - standardized language
* Engine - standardize low-level access (placeholders)



## SQLAlchemy engine
{id: sqlalchemy-engine}
{i: create_engine}

```
engine = create_engine('sqlite:///test.db')               # relative path
engine = create_engine('sqlite:////full/path/to/test.db') # full path
engine = create_engine('sqlite://')                       # in memory database
```


**PostgreSQL**


```
engine = create_engine('postgresql://user:password@hostname/dbname')
engine = create_engine('postgresql+psycopg2://user:password@hostname/dbname')
```


**MySQL**


```
engine = create_engine("mysql://user:password@hostname/dbname", encoding='latin1') # defaults to utf-8
```


## SQLAlchemy autocommit
{id: sqlalchemy-autocommit}
{i: begin}
{i: commit}
{i: rollback}


Unlike the underlying database engines, SQLAlchemy uses autocommit.
That is, usually we don't need to call <b>commit()</b>, but if we would like to have a transaction we need to
start it using <b>begin()</b> and end it either with <b>commit()</b> or with <b>rollback()</b>.




## SQLAlchemy engine CREATE TABLE
{id: sqlalchemy-engine-create}
{i: execute}
{i: Engine}
![](examples/sqla/sqlite_engine_create.py)


## SQLAlchemy engine INSERT
{id: sqlalchemy-engine-insert}
{i: INSERT}
![](examples/sqla/sqlite_engine_insert.py)


## SQLAlchemy engine SELECT
{id: sqlalchemy-engine-select}
{i: fetchone}
{i: close}
{i: SELECT}
![](examples/sqla/sqlite_engine_select.py)


## SQLAlchemy engine SELECT all
{id: sqlalchemy-engine-select-all}
![](examples/sqla/sqlite_engine_select_all.py)


## SQLAlchemy engine SELECT fetchall
{id: sqlalchemy-engine-select-fetchall}
![](examples/sqla/sqlite_engine_select_fetchall.py)



## SQLAlchemy engine SELECT aggregate
{id: sqlalchemy-engine-select-aggregate}
![](examples/sqla/sqlite_engine_select_aggregate.py)


## SQLAlchemy engine SELECT IN
{id: sqlalchemy-engine-select-in}
![](examples/sqla/sqlite_engine_select_in.py)


## SQLAlchemy engine SELECT IN with placeholders
{id: sqlalchemy-engine-select-in-placeholder}
![](examples/sqla/sqlite_engine_select_in_placeholders.py)


## SQLAlchemy engine connection
{id: sqlalchemy-engine-connection}
![](examples/sqla/sqlite_engine_connection.py)



## SQLAlchemy engine transaction
{id: sqlalchemy-engine-transaction}
![](examples/sqla/sqlite_engine_transaction.py)




## SQLAlchemy engine using context managers
{id: sqlalchemy-engine-context-manager}

```
with engine.begin() as trans:
    conn.execute(...)
    conn.execute(...)
    raise Exception()  # for rollback
```


## Exercise: Create table
{id: exercise-create-table}

Create the following schema

![](examples/sqla/exercise2.sql)

Insert a few data items. Write a few select statements.




## SQLAlchemy Metada
{id: sqlalchemy-metadata}


    Describe the Schema, the structure of the database (tables, columns, constraints, etc.) in Python.



* SQL generation from the metadata, generate to a schema.
* Reflection (Introspection) - Create the metadata from an existing database, from an existing schema.

![](examples/sqla/metadata.py)



## SQLAlchemy types
{id: sqlalchemy-types}

* Integer() - INT
* String()  - ASCII strings - VARCHAR
* Unicode() - Unicode string - VARCHAR or NVARCHAR depending on database
* Boolean()  - BOOLEAN, INT, TINYINT depending on db support for boolean type
* DateTime()  - DATETIME or TIMESTAMP returns Python datetime() objects.
* Float()     - floating point values
* Numeric()   - precision numbers using Python Decimal()



## SQLAlchemy ORM - Object Relational Mapping
{id: sqlalchemy-orm}

* Domain model
* Mapping between Domain Object - Table Row




## SQLAlchemy ORM create
{id: sqlalchemy-orm-create}
![](examples/sqla/orm_create_db.py)


## SQLAlchemy ORM schema
{id: orm-schema}

```
    echo .schema | sqlite3 imdb.db
```
![](examples/sqla/orm_schema.sql)



## SQLAlchemy ORM reflection
{id: orm-reflection}
![](examples/sqla/orm_reflection.py)


## SQLAlchemy ORM INSERT after automap
{id: orm-automap-insert}
![](examples/sqla/orm_automap_insert.py)


## SQLAlchemy ORM INSERT
{id: orm-insert}
![](examples/sqla/orm_insert.py)


## SQLAlchemy ORM SELECT
{id: orm-select}
![](examples/sqla/orm_select.py)


## SQLAlchemy ORM SELECT cross tables
{id: orm-select-cross-table}
![](examples/sqla/orm_select_cross_table.py)


## SQLAlchemy ORM SELECT and INSERT
{id: orm-select-insert}
![](examples/sqla/orm_select_insert.py)


## SQLAlchemy ORM UPDATE
{id: orm-update}
![](examples/sqla/orm_update.py)


## SQLAlchemy ORM logging
{id: orm-logging}
![](examples/sqla/orm_logging.py)


