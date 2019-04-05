
# Must enable kms api: cloudkms.googleapis.com

## Creating KeyRing

### What is a KeyRing??

A key ring is a grouping of keys for organizational purposes. A key ring belongs to a GCP Project and resides in a specific location. Keys inherit permissions from the key ring that contains them. Grouping keys with related permissions together in a key ring allows you to grant, revoke, or modify permissions to those keys at the key ring level, without needing to act on each one individually.

$ gcloud kms keyrings create <keyring> --location global

## Creating Key

### What is a Key??

A key is a named object representing a cryptographic key used for a specific purpose. The key material, the actual bits used for encryption, can change over time as new key versions are created.

A key is used to protect some corpus of data. You could encrypt a collection of files with the same key, and people with decrypt permissions on that key would be able to decrypt those files.

### What is the purpose??

$ gcloud kms keys create <key> --location global --keyring <keyring> --purpose encryption

## Listing keys

```
$ gcloud kms keys list --location global --keyring <keyring>

NAME                                                                        PURPOSE          LABELS  PRIMARY_ID  PRIMARY_STATE
projects/vishen-admin/locations/global/keyRings/test/cryptoKeys/quickstart  ENCRYPT_DECRYPT          1           ENABLED
```

## Listing version of keys

### What is a key version

A key version represents the key material associated with a key at some point in time. Each key can have arbitrarily many versions, but must have at least one. Versions are numbered sequentially, starting with 1.

Returning to the example of a set of encrypted files, files encrypted with the same key may be encrypted with different key versions. Some of the files may be encrypted with version 1 and others with version 2. When you ask Cloud KMS to decrypt any one of these files, you specify the name of the key (not a specific version) which encrypted it. Cloud KMS automatically identifies which version was used for encryption, and uses it to decrypt the file if the version is still enabled.

```
$ gcloud kms keys versions list --location global --keyring <keyring> --key <key>

NAME                                                                                            STATE
projects/vishen-admin/locations/global/keyRings/test/cryptoKeys/quickstart/cryptoKeyVersions/1  ENABLED

```

## Key rotation

In Cloud KMS, a key rotation is represented by generating a new key version of a key, and marking that version as the primary version.

Creating a new key version generates the new cryptographic key material, and marking that key version as primary causes it to be used to encrypt any new data. Each key has a designated primary version at any point in time, which Cloud KMS uses to encrypt data, by default.

Rotating a key doesn't disable or destroy previous key versions. The previous key versions will no longer be primary, but they remain available for decrypting data.

```
$ gcloud kms keys versions create --location global --keyring test --key quickstart
$ gcloud kms keys versions list --location global --keyring test --key quickstart

NAME                                                                                            STATE
projects/vishen-admin/locations/global/keyRings/test/cryptoKeys/quickstart/cryptoKeyVersions/2  ENABLED
projects/vishen-admin/locations/global/keyRings/test/cryptoKeys/quickstart/cryptoKeyVersions/1  ENABLED
```

## Encrypt data

When you have a Key, you can then use that Key to encrypt data.

```
$ echo "Hello, KMS!" > mysecret.txt
$ gcloud kms encrypt --location global --keyring <keyring> --key <key> --plaintext-file mysecret.txt --ciphertext-file mysecret.txt.encrypted
$ ls
mysecret.txt  mysecret.txt.encrypted 
```

## Decrypt data

```
$ gcloud kms decrypt --location global --keyring <keyring> --key <key> --ciphertext-file mysecret.txt.encrypted --plaintext-file mysecret.txt.decrypted
$ ls
mysecret.txt  mysecret.txt.decrypted  mysecret.txt.encrypted
$ cat mysecret.txt.decrypted
Hello, KMS!
```

