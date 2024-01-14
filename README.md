## ABOUT

Uses makefiles and sqitch to managing a database

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
