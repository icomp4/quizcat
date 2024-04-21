FROM node:latest as build-stage
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend/ .
RUN npm run build

FROM golang:latest
WORKDIR /go/src/app
COPY . .
COPY --from=build-stage /app/dist /go/src/app/frontend/dist
RUN apt-get update && apt-get install -y make  # Install make utility
RUN go get -d -v ./...
RUN make dbup 
RUN go build -o main .
CMD ["sh", "-c", "serve -s frontend/dist & ./main"]