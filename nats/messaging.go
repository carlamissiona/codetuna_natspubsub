package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var nc *nats.Conn
var ns server.Server
var subject string  = "whatsmyurl"
var URLstr = "server"

func Serve() {
	opts := &server.Options{}
 
	ns, err := server.NewServer(opts)

	if err != nil {
		panic(err)
	}

	// Start the server via goroutine
	go ns.Start()

     
	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	URLstr = ns.ClientURL()
    
    Connect(ns.ClientURL())
	 
	 
    Subs(subject)
 
	// Wait for server shutdown
	ns.WaitForShutdown()
}

func Pubs( str_pub []byte){
    fmt.Println("@ publish")
	nc.Publish(subject,str_pub)

}
func Subs(topic string ){
	fmt.Println("@ subs waiting") 
	nc.Subscribe(topic, func(msg *nats.Msg) {
		// Print message data
		fmt.Println("@ subscribe nc")
		data := string(msg.Data)
		fmt.Println(data)

		// Shutdown the server (optional)
		// ns.Shutdown()
	})
 
}
func Connect(nats_url string){
    var err error
	nc, err = nats.Connect(nats_url )

	if err != nil {
		panic(err)
	}

    
}