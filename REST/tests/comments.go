package tests

import (
	"encoding/json"
	"fmt"

	"github.com/iLopezosa/api-wars/rest/db"
	"github.com/iLopezosa/api-wars/rest/models"
)

func TestCommentsComplete() {
	fmt.Println("\n---TestCommentsComplete---")
	fmt.Println("\n#1 > Creating new comment...")

	c := &models.Comment{
		PostID:  1,
		UserID:  1,
		Content: "I am a comment",
	}

	if err := db.CommentUpsert(c); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	c2, err := db.CommentRead(c.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err := json.MarshalIndent(c2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	fmt.Printf("%d: %s\n", c.ID, string(b))
	fmt.Println("\n#2 > Updating comment...")

	c.Content = "I am an updated comment"

	if err = db.CommentUpsert(c); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	c2, err = db.CommentRead(c.ID)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(c2, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#3 > List all comments:")

	comments, err := db.CommentList()
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(comments, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#4 > List all comments from a user:")

	comments, err = db.CommentListByUserID(1)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(comments, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#5 > List all comments from a post:")

	comments, err = db.CommentListByPostID(1)
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	b, err = json.MarshalIndent(comments, "", "  ")
	if err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("\n#6 > Deleting comment...")

	if err = db.CommentDelete(c.ID); err != nil {
		fmt.Println("TEST ERROR:", err)
		return
	}

	if _, err = db.CommentRead(c.ID); err != nil {
		fmt.Printf("\nComment %d deleted successfully", c.ID)
	} else {
		fmt.Printf("\nTEST ERROR: Comment %d not deleted", c.ID)
		return
	}

	fmt.Println("DONE!")
}
