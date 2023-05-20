#!/bin/sh

echo "Setting up bbdb for testing.."

if [ -d "pb_data" ]; then
  echo "Removing old pb_data directory"
  rm -rf pb_data
fi

echo "Setting up test data"
tar xfz pb_data.tar.gz

echo "Building UI"
npm --prefix=./ui run build

echo "Running app"
go run bbdb.go serve
