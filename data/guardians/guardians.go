//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/DharitriME/protobuf/protobuf --gogoslick_out=. guardians.proto
package guardians
