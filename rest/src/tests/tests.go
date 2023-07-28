package tests

import (
	"fmt"
	"strings"
)

// runTests runs the tests passed as arguments
func Run(testsPtr *string) {

	toDo := strings.Split(*testsPtr, ",")

	for _, t := range toDo {
		switch t {
		case "users":
			TestUsersComplete()
		case "posts":
			TestPostsComplete()
		case "comments":
			TestCommentsComplete()
		case "messages":
			TestMessagesComplete()
		case "chats":
			TestChatsComplete()
		default:
			fmt.Printf("\nUnknown test: %s\n", t)
		}
	}
}
