package main

import (
  "context"
  "log"
  "testing"

  "github.com/testcontainers/testcontainers-go"
)

func TestMain(m *testing.T) {
  var testC testcontainers.Container
  var err error
  testC, testDb, err = getTestDb()
  if err != nil {
    log.Fatal(err)
  }
  m.Run()
  testC.Terminate(context.Background())
}
