package main

import (
	"flag"
	"github.com/emre/utopian_data/client"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/fatih/set.v0"
	"fmt"
)

func main() {
	mongodbURI := flag.String("mongodb_uri", "", "Mongodb connection uri")
	flag.Parse()

	log.Println("Connecting to", *mongodbURI)

	session, err := mgo.Dial(*mongodbURI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	utopianClient := client.Client{BaseUrl: "https://api.utopian.io/api"}

	categoryCollection := session.DB("utopiandata").C("categories")
	categories := set.New()

	// sync moderators
	log.Println("Sync moderators")
	moderatorCollection := session.DB("utopiandata").C("moderators")
	moderators := utopianClient.GetModerators()
	for _, moderator := range moderators.Moderators {
		moderatorCollection.Upsert(bson.M{"account": moderator.Account}, moderator)
	}

	// sync sponsors
	log.Println("Sync sponsors")
	sponsorCollection := session.DB("utopiandata").C("sponsors")
	sponsors := utopianClient.GetSponsors()
	for _, sponsor := range sponsors.Sponsors {
		sponsorCollection.Upsert(bson.M{"account": sponsor.Account}, sponsor)
	}

	// sync approved posts
	log.Println("Sync approved posts")
	postCollection := session.DB("utopiandata").C("posts")
	postList := []client.Post{}
	posts := utopianClient.GetPosts(nil, nil, postList, false)
	for _, post := range posts {
		categories.Add(post.JSONMetadata.Type)
		postCollection.Upsert(bson.M{"_id": post.MongoId}, post)
	}

	// sync hidden posts
	log.Println("Sync hidden posts")
	postList = []client.Post{}
	posts = utopianClient.GetPosts(nil, nil, postList, true)
	for _, post := range posts {
		postCollection.Upsert(bson.M{"_id": post.MongoId}, post)
		categories.Add(post.JSONMetadata.Type)
	}

	// sync categories
	log.Println("Sync categories")
	for _, category := range categories.List() {
		if len(fmt.Sprintf("%s", category)) < 3 {
			log.Println("Broken category name. Skipping.")
			continue
		}
		categoryCollection.Upsert(bson.M{"name": category}, bson.M{"name": category})
	}


}
