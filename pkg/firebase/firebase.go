package firebase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	acl "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"github.com/ansxy/nagabelajar-be-go/config"
	"google.golang.org/api/option"
)

type IFaceFCM interface {
	UploudFile(ctx context.Context, file *multipart.FileHeader, fileReader io.Reader) (string, error)
	GetOneFile(ctx context.Context, fileName string) (io.Reader, error)
	GetMd5Hash(ctx context.Context, fileName string) (string, error)
	VerifiyToken(ctx context.Context, token string) (*auth.Token, error)
}

type FCS struct {
	Client  *firebase.App
	Auth    *auth.Client
	Storage *storage.Client
}

// VerifiyToken implements IFaceFCM.
func (f *FCS) VerifiyToken(ctx context.Context, token string) (*auth.Token, error) {
	return f.Auth.VerifyIDToken(ctx, token)
}

// GetMd5Hash implements IFaceFCM.
func (f *FCS) GetMd5Hash(ctx context.Context, fileName string) (string, error) {
	bucket, err := f.Storage.DefaultBucket()
	if err != nil {
		return "", err
	}

	object := bucket.Object(fileName)
	reader, err := object.NewReader(ctx)
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)
	md5Hash := hex.EncodeToString(hash)

	return md5Hash, nil
}

// GetOneFile implements IFaceFCM.
func (f *FCS) GetOneFile(ctx context.Context, fileName string) (io.Reader, error) {
	bucket, err := f.Storage.DefaultBucket()
	if err != nil {
		return nil, err
	}

	object := bucket.Object(fileName)
	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

// UploudFile implements IFaceFCM.
func (f *FCS) UploudFile(ctx context.Context, file *multipart.FileHeader, fileReader io.Reader) (string, error) {
	bucket, err := f.Storage.DefaultBucket()
	if err != nil {
		return "", err
	}

	object := bucket.Object(file.Filename)
	uploud := object.NewWriter(ctx)

	if _, err := io.Copy(uploud, fileReader); err != nil {
		return "", err
	}
	if err := uploud.Close(); err != nil {
		return "", err
	}

	if err := object.ACL().Set(ctx, acl.AllUsers, acl.RoleReader); err != nil {
		return "", err
	}

	// Generate and return the URL of the uploaded file
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s%s", "alpro-module.appspot.com", file.Filename, "?alt=media")
	return url, nil
}
func NewFCMClient(cnf config.FirebaseConfig) (IFaceFCM, error) {
	opts := option.WithCredentialsFile(cnf.ServiceAccountPath)
	config := &firebase.Config{
		StorageBucket: cnf.FirebaseStorageBucket,
	}
	app, err := firebase.NewApp(context.Background(), config, opts)
	if err != nil {
		log.Fatalf("Error initializing firebase app: %v", err)
		return nil, err
	}
	log.Println("Firebase app initialized successfully")

	storage, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf("Error initializing firebase storage: %v", err)
		return nil, err
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error initializing firebase auth: %v", err)
		return nil, err
	}

	return &FCS{
		Client:  app,
		Auth:    authClient,
		Storage: storage,
	}, nil
}
