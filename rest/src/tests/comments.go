package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/models"
)

func TestCommentsComplete() {
	fmt.Println("\n---TestCommentsComplete---")
	fmt.Println("\n#1 > Creating new comment...")

	c := &models.Comment{
		PostID:  1,
		UserID:  1,
		Content: "I am a comment",
	}
	fmt.Printf("%+v\n", c)
	if err := db.CommentUpsert(c); err != nil {
		log.Fatal(err)
	}

	c2, err := db.CommentRead(c.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(c2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", c.ID, string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#2 > Updating comment...")

	c.Content = "I am an updated comment"

	if err = db.CommentUpsert(c); err != nil {
		log.Fatal(err)
	}

	c2, err = db.CommentRead(c.ID)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(c2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#3 > List all comments:")

	comments, err := db.CommentList()
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(comments, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#4 > List all comments from a user:")

	comments, err = db.CommentListByUserID(1)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(comments, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#5 > List all comments from a post:")

	comments, err = db.CommentListByPostID(1)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(comments, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	time.Sleep(5 * time.Second)
	fmt.Println("\n#6 > Deleting comment...")

	if err = db.CommentDelete(c.ID); err != nil {
		log.Fatal(err)
	}

	if _, err = db.CommentRead(c.ID); err != nil {
		fmt.Println("Comment deleted")
	} else {
		log.Fatal("Comment not deleted")
	}

	fmt.Println("DONE!")
}
