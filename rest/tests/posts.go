package tests

import (
	"encoding/json"
	"fmt"

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
		fmt.Println("TEST ERROR:", err)
		return
	}

	p2, err := db.PostRead(p.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err := json.MarshalIndent(p2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Printf("%d: %s\n", p.ID, string(b))
	fmt.Println("\n#2 > Updating post...")

	p.Content = "It is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy. Evading the dreaded Imperial Starfleet, a group of freedom fighters led by Luke Skywalker has established a new secret base on the remote ice world of Hoth."

	if err = db.PostUpsert(p); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	p2, err = db.PostRead(p.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(p2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#3 > List all posts:")

	posts, err := db.PostList()
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	for _, p := range posts {
		b, err = json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Println("TEST ERROR:", err)
			return
		}
		fmt.Println(string(b))
	}
	fmt.Println("\n#4 > List all posts by user:")

	posts, err = db.PostListByUserID(1)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(posts, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#5 > Deleting post...")
	if err = db.PostDelete(p.ID); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	posts, err = db.PostList()

	_, err = db.PostRead(p.ID)
	if err != nil {
		fmt.Printf("\nPost %d deleted successfully\n", p.ID)
	} else {
		fmt.Printf("\nTEST ERROR: Post %d not deleted\n", p.ID)
		return
	}

	fmt.Println("DONE!")
}
