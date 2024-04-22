package firebase

import (
	"context"
	"io"
	"log"
	"mime/multipart"

	acl "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"github.com/ansxy/nagabelajar-be-go/config"
	"google.golang.org/api/option"
)

type IFaceFCM interface {
	UploudFile(ctx context.Context, file *multipart.FileHeader) error
	GetOneFile(ctx context.Context, fileName string) (io.Reader, error)
}

type FCS struct {
	Client  *firebase.App
	Storage *storage.Client
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
func (f *FCS) UploudFile(ctx context.Context, file *multipart.FileHeader) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	bucket, err := f.Storage.DefaultBucket()
	if err != nil {
		return err
	}

	object := bucket.Object(file.Filename)
	uploud := object.NewWriter(ctx)

	if _, err := io.Copy(uploud, fileReader); err != nil {
		return err
	}

	if err := uploud.Close(); err != nil {
		return err
	}

	if err := object.ACL().Set(ctx, acl.AllUsers, acl.RoleReader); err != nil {
		return err
	}

	return nil

}

func NewFCMClient(cnf config.FirebaseConfig) (IFaceFCM, error) {
	opts := option.WithCredentialsFile(cnf.ServiceAccountPath)
	config := &firebase.Config{
		StorageBucket: cnf.FirebaseStorageBucket,
	}
	app, err := firebase.NewApp(context.Background(), config, opts)
	if err != nil {
		log.Fatal("Error initializing firebase app: ", err)
		return nil, err
	}

	storage, err := app.Storage(context.Background())
	if err != nil {
		log.Fatal("Error initializing firebase storage: ", err)
		return nil, err
	}

	return &FCS{
		Client:  app,
		Storage: storage,
	}, nil
}
