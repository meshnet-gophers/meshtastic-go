//go:generate protoc  -I ./protobufs/ --go_out=. --go_opt=paths=source_relative --go-vtproto_out=paths=source_relative:. --go-vtproto_opt=features=marshal+unmarshal+size ./protobufs/meshtastic/*.proto
package meshtastic
