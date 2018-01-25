#### utopian_data

utopian_data is a CLI application which fetches public data of Utopian
into your Mongodb instance.

Supported collections:

- moderators (api/moderators)[https://api.utopian.io/api/moderators]
- sponsors (api/sponsors)[https://api.utopian.io/api/sponsors]
- **posts** (api/posts)[https://api.utopian.io/api/posts]

#### Why

I have been working on a separate *statistics* application (private) for
a while to analyze Utopian data better. However, using a REST API for
statistics is not practical.

You need a fast system where you can make complex queries. Syncing Utopian 
posts' data into your system (you can set a crontab and sync it in every 5 minutes)
is much better.


#### Example Mongodb aggreegation Queries

**Total Contribution Count**

```
db.posts.find({}).count()
```

![mongoshell output](https://i.hizliresim.com/y0RzOn.png)

**Top 3 moderators on blog category (moderation count)**

```
db.posts.aggregate(
  [
    {"$match": {"json_metadata.type": "blog"}},
    {"$group": {_id: "$moderator", count: {$sum : 1}}},
    {"$sort": {"count": -1}},
    {"$limit": 3},
  ]
)
```

![mongoshell output](https://i.hizliresim.com/JQ6O1W.png)

**Top Contributors**

```
db.posts.aggregate(
  [
    {"$group": {_id: "$author", count: {$sum : 1}}},
    {"$sort": {"count": -1}},
    {"$limit": 3},
  ]
)
```

![mongoshell output](https://i.hizliresim.com/az01VB.png)

And more... You can pretty much filter/group everything you want once
you learn how [Mongodb aggregation queries](https://docs.mongodb.com/manual/aggregation/) work.

#### Installation and Usage

```
$ go get github.com/emre/utopian_data
$ cd $GOPATH
$ go install github.com/emre/utopian_data
$ cd `$GOPATH`/bin
$  utopian_data --mongodb_uri=localhost
```

**Output**

![Console Output](https://i.hizliresim.com/D79yqO.png)







