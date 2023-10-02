/*
Package protobuf contains wire definitions of messages passed between drand nodes.
*/
//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative bridge.proto
package proto
