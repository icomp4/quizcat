#!/bin/bash

# Navigating to 'frontend' directory and building the project
cd frontend || exit
npm run build

# Returning to the root directory
cd ..

# Navigating to 'migrate' and executing commands
cd migrate || exit
cd up || exit

# Assuming 'cd up' is a directory change and not a command. If 'up' is a command, adjust accordingly.
cd .. || exit

# Running a Go program, assuming 'main.go' is in the 'migrate' directory
go run main.go

# Navigating back to the root directory
cd ../..

# Running the main Go program from the root directory
go run main.go
