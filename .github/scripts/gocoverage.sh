#!/bin/bash

unit_coverage_test() {
  go mod download
  go test $(go list ./...) -race -covermode atomic -coverprofile=coverage.out ./...
}

unit_coverage_test