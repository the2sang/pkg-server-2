package main

import (
  "context"
  "database/sql"
  "fmt"
  "net/url"
  "os"
  "path/filepath"
  "time"

  "github.com/docker/go-connections/nat"
  "github.com/testcontainers/testcontainers-go"
  "github.com/testcontainers/testcontainers-go/wait"
  "gocloud.dev/blob"
  "gocloud.dev/blob/fileblob"
)

func getTestBucket(tmpDir string) (*blob.BUcket, error) {
  myDir, err := os.MkdirTemp(tmpDir, "test-bucket")
  if err != nil {
    return nil, err
  }
  u, err := url.Parse(fmt.Sprintf("file:///%s", myDir))

}
