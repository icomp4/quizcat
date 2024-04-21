#!/bin/bash

cd frontend || exit
npm run build

cd ..

cd migrate/up || exit

go run main.go

cd ..
cd ..

go run main.go
