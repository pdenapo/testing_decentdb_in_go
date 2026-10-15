This repository contains test for different embededd sql databases. It might be helpfull as example of how to use
them in different languages.

# Sqlite 

Is by far the most used embedded database in the world. 

# Turso

Is a reimplemntation of Sqlite in rust with additional features.

# DecentDb

Decentdb is an embedded database created by Steven Hildreth, which provides
an interesting alternative to Sqlite.

https://decentdb.org/

https://github.com/sphildreth/decentdb

The idea of this repository is to test it from Go.

To run the tests clone this repository with

git clone --recurse-submodules 

Conclusions so far: decentdb can be used with simple explicit sql queries.
But it does not work with Bun Orm.

Todo: portar el driver de turso para aceptar el dialecto postgresql.

# Duckdb

Is an embedded database for analytical work-flows (using a columnar-oriented storage). It has greater compatibility 
with Postgresql.
