#### utopian_data

utopian_data is a CLI application which fetches public data of Utopian
into your Mongodb instance.

Supported collections:

- moderators (api/moderators)[https://api.utopian.io/api/moderators]
- sponsors (api/sponsors)[https://api.utopian.io/api/sponsors]
- **posts** (api/posts)[https://api.utopian.io/api/posts]

#### Why

I have been working on a seperate *statistics* application (private) for
a while to analyze Utopian data better. However, using a REST API for
statistics is not practical.

You need a fast system where you can make complex queries. Syncing Utopian 
posts' data into your system (you can set a crontab and sync it in every 5 minutes)
is much better.


#### Example Mongodb aggreegation Queries

TODO

#### Installation and Usage
```
$ go get github.com/emre/utopian_data
$ utopian_data MONGODB_URI
```







