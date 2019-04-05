package main

import (
	"context"
	"fmt"
	"log"

	cloudkms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/iterator"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func main() {
	projectID := "vishen-admin"
	// Location of the key rings.
	locationID := "global"

	// Create the KMS client.
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// The resource name of the key rings.
	parentName := fmt.Sprintf("projects/%s/locations/%s", projectID, locationID)

	// Build the request.
	req := &kmspb.ListKeyRingsRequest{
		Parent: parentName,
	}
	// Query the API.
	it := client.ListKeyRings(ctx, req)

	// Iterate and print results.
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to list key rings: %v", err)
		}
		fmt.Printf("KeyRing: %q\n", resp.Name)
	}

	keyRing := "test"
	key := "quickstart"
	name := fmt.Sprintf("%s/keyRings/%s/cryptoKeys/%s", parentName, keyRing, key)
	message := "Hello, KMS!"
	encryptReq := &kmspb.EncryptRequest{
		Name:      name,
		Plaintext: []byte(message),
	}

	encryptResp, err := client.Encrypt(ctx, encryptReq)
	if err != nil {
		log.Fatalf("failed to encrypt data: %v", err)
	}
	fmt.Printf("Encrypted with %s: %v\n", encryptResp.GetName(), encryptResp.GetCiphertext())

	decryptReq := &kmspb.DecryptRequest{
		Name:       name,
		Ciphertext: encryptResp.GetCiphertext(),
	}
	decryptResp, err := client.Decrypt(ctx, decryptReq)
	if err != nil {
		log.Fatalf("failed to decrypt data: %v", err)
	}

	fmt.Printf("Decrypted: %s\n", decryptResp.GetPlaintext())

}
