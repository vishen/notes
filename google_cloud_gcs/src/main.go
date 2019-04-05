package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	// projectID := "vishen-admin"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "vishen-admin-test"
	bucket := client.Bucket(bucketName)

	objName := "some-obj-key"
	obj := bucket.Object(objName)

	wc := obj.NewWriter(ctx)
	message := "Hello, some-obj-key!"
	if _, err := wc.Write([]byte(message)); err != nil {
		log.Fatalf("unable to write to %s: %v", objName, err)
	}
	wc.Close()

	rc, err := bucket.Object(objName).NewReader(ctx)
	if err != nil {
		log.Fatalf("unable to create new reader: %v", err)
	}
	slurp, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		log.Fatalf("unable to read %s: %v", objName, err)
	}
	fmt.Printf("file contents: %s\n", slurp)
}
