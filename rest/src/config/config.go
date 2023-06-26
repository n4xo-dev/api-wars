package config

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	NumOfUsers           int
	NumOfChats           int
	NumOfPosts           int
	MaxNumOfMessages     int
	MaxNumOfParticipants int
}

var config Config

var configOnce sync.Once

func GetConfig() Config {
	configOnce.Do(func() {
		numOfUsers, err := strconv.Atoi(os.Getenv("NUM_OF_USERS"))
		if err != nil {
			numOfUsers = 5
		}
		numOfChats, err := strconv.Atoi(os.Getenv("NUM_OF_CHATS"))
		if err != nil {
			numOfChats = 3
		}
		numOfPosts, err := strconv.Atoi(os.Getenv("NUM_OF_POSTS"))
		if err != nil {
			numOfPosts = 3
		}
		maxNumOfMessages, err := strconv.Atoi(os.Getenv("MAX_NUM_OF_MESSAGES"))
		if err != nil {
			maxNumOfMessages = 5
		}
		maxNumOfParticipants, err := strconv.Atoi(os.Getenv("MAX_NUM_OF_PARTICIPANTS"))
		if err != nil {
			maxNumOfParticipants = 4
		}

		config = Config{
			NumOfUsers:           numOfUsers,
			NumOfChats:           numOfChats,
			NumOfPosts:           numOfPosts,
			MaxNumOfMessages:     maxNumOfMessages,
			MaxNumOfParticipants: maxNumOfParticipants,
		}
	})

	return config
}
