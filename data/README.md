# Database for Legislator DB

Method used to create database for this application.

```
mkdir data && cd data
wget https://theunitedstates.io/congress-legislators/legislators-current.csv
sqlite3 legislator.db
sqlite> .mode csv
sqlite> .import legislators-current.csv legislators
sqlite> .schema legislators
sqlite> select first_name, last_name from legislators where state = "TX";
sqlite> .quit
```
