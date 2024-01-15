## ABOUT

Use makefiles and sqitch to managing the database design.

(Evolutionary Database Design)[https://martinfowler.com/articles/evodb.html]

## INSTALL

* cp env.sqitch.example env.sqitch
* update the property values in env.sqitch
* env.sqitch is included in .gitignore

## SETUP DATABASE

* run make to see command line menu
    * make
* create database
    * make create

## SETUP SCHEMA

* create schema
    * make create-schema SCHEMA_NAME
* work with scheam
    * cd SCEHAM_NAME
* run make to see command line menu
    * make

## VERSION

When adding a new migration (make add) a versioned file will added to the deploy, revert,
and verify directories.  The makefile scripts will check for the last migration with the
same name, if found the previous file will be copied to the new version.

```
make add xyx

001-XYZ.sql # content will be copied to 002-XYZ.sql
002-ZZZ.sql # filename does not mach xyz
003-XYZ.sql # new migration added
```
