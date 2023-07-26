package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
)

func TestPostsComplete() {
	fmt.Println("\n---TestPostsComplete---")
	fmt.Println("\n#1 > Creating new post...")

	p := &models.Post{
		Title:   "The Empire Strikes Back",
		Content: "It is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy.",
		UserID:  1,
	}

	if err := db.PostUpsert(p); err != nil {
		log.Fatal(err)
	}

	p2, err := db.PostRead(p.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(p2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", p.ID, string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#2 > Updating post...")

	p.Content = "It is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy. Evading the dreaded Imperial Starfleet, a group of freedom fighters led by Luke Skywalker has established a new secret base on the remote ice world of Hoth."

	if err = db.PostUpsert(p); err != nil {
		log.Fatal(err)
	}

	p2, err = db.PostRead(p.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(p2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#3 > List all posts:")

	posts, err := db.PostList()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range posts {
		b, err = json.MarshalIndent(p, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
	time.Sleep(5 * time.Second)
	fmt.Println("\n#4 > List all posts by user:")

	posts, err = db.PostListByUserID(1)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(posts, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#5 > Deleting post...")
	if err = db.PostDelete(p.ID); err != nil {
		log.Fatal(err)
	}

	posts, err = db.PostList()

	_, err = db.PostRead(p.ID)
	if err != nil {
		fmt.Println("Post deleted")
	} else {
		log.Fatal("Post not deleted")
	}

	fmt.Println("DONE!")
}
