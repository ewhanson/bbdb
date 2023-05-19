 #!/bin/sh

echo "Setting up bbdb for testing.."

if [ -d "pb_data" ]; then
  echo "Removing old pb_data directory"
  rm -rf pb_data
fi

echo "Setting up test data"
tar xfz pb_data.tar.gz

echo "Installing UI dependencies"
npm --prefix=./ui ci && npm --prefix=./ui run build

echo "Building bbdb"
go build bbdb.go