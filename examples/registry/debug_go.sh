# go build -o goapp -gcflags "all=-N -l" .
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./go/goapp