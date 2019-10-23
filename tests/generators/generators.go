package generators

import (
	"math/rand"

	"phoenixnap.com/pnap-cli/pnapctl/commands/get/servers"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateServers(n int) []servers.LongServer {
	var serverlist []servers.LongServer
	for i := 0; i < n; i++ {
		serverlist = append(serverlist, GenerateServer())
	}
	return serverlist
}

func GenerateServer() servers.LongServer {
	return servers.LongServer{
		ID:          randSeq(10),
		Status:      randSeq(10),
		Name:        randSeq(10),
		Description: randSeq(10),
		Os:          randSeq(10),
		Type:        randSeq(10),
		Location:    randSeq(10),
		CPU:         randSeq(10),
		RAM:         randSeq(10),
		Storage:     randSeq(10),
	}
}

// ConvertLongToShortServers will convert a slice of long servers to a short servers
func ConvertLongToShortServers(long_servers []servers.LongServer) []servers.ShortServer {
	var serverlist []servers.ShortServer
	for i := 0; i < len(long_servers); i++ {
		serverlist = append(serverlist, ConvertLongToShortServer(long_servers[i]))
	}
	return serverlist
}

// ConvertLongToShortServer will convert a single long server to a short server
func ConvertLongToShortServer(long_server servers.LongServer) servers.ShortServer {
	return servers.ShortServer{
		ID:          long_server.ID,
		Status:      long_server.Status,
		Name:        long_server.Name,
		Description: long_server.Description,
	}
}
