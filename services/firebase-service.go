package services

import (
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Storage interface {
	FirebaseInit() *firebase.App
	UploadToStorage(context.Context, string)
}

type storage struct {
	bucket string
	json   string
}

func NewStorage() Storage {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	bucket := os.Getenv("BUCKET")
	json := os.Getenv("CREDENTIAL")
	return &storage{
		bucket: bucket,
		json:   json,
	}
}

func (s *storage) FirebaseInit() *firebase.App {
	config := &firebase.Config{
		StorageBucket: s.bucket,
	}
	opt := option.WithCredentialsFile(s.json)

	fb, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Println(err)
	}

	return fb
}

func (s *storage) UploadToStorage(ctx context.Context, fileName string) {
	fb := s.FirebaseInit()
	id := uuid.New()
	client, err := fb.Storage(context.Background())
	if err != nil {
		log.Println(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Println(err)
	}

	wc := bucket.Object(fileName).NewWriter(ctx)
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer wc.Close()
	// var re io.Reader
	// re, _ = os.Open(fileName)

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)
	i := strings.Index(base64Encoding, ",")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Encoding[i+1:]))

	_, err = io.Copy(wc, reader)
	if err != nil {
		log.Println(err.Error())
	}

	if err := wc.Close(); err != nil {
		log.Println(err.Error())
	}
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
