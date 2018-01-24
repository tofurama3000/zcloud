package cmd

import (
	"testing"
)

func TestConvertURL (t *testing.T) {
	url := "cloud://bucket/file.txt"
	if res := convertURL(url, "s3"); res != "s3://bucket/file.txt" {
		t.Fatal(res)
	}
	url = "s3://bucket/dir/file.txt"
	if res := convertURL(url, "s3"); res != url {
		t.Fatal(res)
	}
}

func TestIsCloudURL (t *testing.T) {
	url := "cloud://bucket"
	if !isCloudURL(url) {
		t.Fatal(url)
	}
	url = "gs://bucket/file.txt"
	if isCloudURL(url) {
		t.Fatal(url)
	}
	url = "zcloud://bucket/dir/"
	if isCloudURL(url) {
		t.Fatal(url)
	}
}