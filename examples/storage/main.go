package main

// Use `dev_appserver.py --default_gcs_bucket_name GCS_BUCKET_NAME`
// when running locally.

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	//	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/log"
)

const URL = "http://localhost:8080/nmg/image/team"

func serveError(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	log.Errorf(ctx, "%v", err)
}

var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const rootTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form></body></html>
`

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	bucket, err := file.DefaultBucketName(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
	}

	fmt.Println()
	fmt.Println("bucket", bucket)
	fmt.Println()

	//	uploadURL, err := blobstore.UploadURL(ctx, "/upload", nil)
	//	if err != nil {
	//		serveError(ctx, w, err)
	//		return
	//	}

	w.Header().Set("Content-Type", "text/html")
	err = rootTemplate.Execute(w, URL)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

func handleServe(w http.ResponseWriter, r *http.Request) {
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	//	ctx := appengine.NewContext(r)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/serve/", handleServe)
	http.HandleFunc("/upload", handleUpload)
	appengine.Main()
}
