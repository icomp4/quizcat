#!/bin/bash

set -e

# Build the Vue.js frontend
echo "Building Vue.js frontend..."
cd frontend
npm install
npm run build
cd ..

# Compile the main Go application
echo "Compiling main Go application..."
cd migrate/up
GOOS=linux GOARCH=amd64 go build -o migrate
cd ../..
GOOS=linux GOARCH=amd64 go build -o mainapp

echo "Build and compilation process completed successfully."
