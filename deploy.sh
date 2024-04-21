#!/bin/bash

# Navigating to 'frontend' directory and building the project
cd frontend || exit
npm run build

# Returning to the root directory
cd ..

# Navigating to 'migrate/up' and assuming 'up' is a directory
cd migrate/up || exit

# Running a Go program, assuming 'main.go' is in the 'migrate/up' directory
go run main.go

# Navigating back to the root directory
cd ../../..

# Running the main Go program from the root directory
go run main.go
