go build -ldflags "-X main.GitTag=$(git describe --tags) -X main.GitHash=$(git rev-parse --short HEAD) -s -w"
