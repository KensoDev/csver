# csver

Export queries to CSV and upload to S3

Dynamically define queries,filenames and buckets. CSVer will take care of the rest.

At [Gogobot](http://www.gogobot.com) we have quite a few tasks that depend on a CSV.  
We want a task to run in cron and export these files then upload to S3 for processing.

Right now, this only works with mysql (our main database), but all of our microservices are working with Postgres so in the future, this will probably be extended to work with Postgres as well.

## Usage

```
usage: csver [<flags>]

Flags:
  --help               Show help (also see --help-long and --help-man).
  --configuration-file=CONFIGURATION-FILE
                       JSON configuration file
  --username=USERNAME  DB username
  --pass=PASS          DB password
  --host=HOST          DB host
  --db-name=DB-NAME    DB Name
```

### Configuration file

The way csver works is you supply configuration file.  
This configuration file is paired with a query file (can be any test file) that the query will be read from.  

The reason this is another file and not inline in the JSON is that often the queries are more complex and easier to read and format when in a separate file.

Here's what it looks like

```
[
  {
    "QueryFile": "query.sql",
    "OutFile": "query.csv"
  },
  {
    "QueryFile": "query_2.sql",
    "OutFile", "query_2.csv"
  }
]
```

The `sql` file are really just a query, here's an example:

```
select id, name from users;
```

