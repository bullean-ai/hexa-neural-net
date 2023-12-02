package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/bullean-ai/hexa-neural-net/config"
	"google.golang.org/api/option"
	"log"
)

var (
	connStr string
)

// NewFireStoreDB Return new FirestoreDB client
func NewFireStoreDB(cfg *config.Config) (db *firestore.Client, err error) {
	println("Driver Firestore Initialized")

	connStr = fmt.Sprintf("%s", cfg.Firestore.CREDENTIALS_PATH)
	opt := option.WithCredentialsFile(connStr)

	db, err = firestore.NewClient(context.Background(), cfg.Firestore.PROJECT_ID, opt)
	if err != nil {
		log.Fatalf("firestore new error:%s \n", err)
	}
	//defer db.Close()

	return
}
