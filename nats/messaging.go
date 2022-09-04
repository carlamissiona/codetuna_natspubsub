package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var nc *nats.Conn
var subject string  = "whatsmyurl"

func Serve() {
	opts := &server.Options{}
 
	ns, err := server.NewServer(opts)

	if err != nil {
		panic(err)
	}

	// Start the server via goroutine
	go ns.Start()

    
	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

    
    connect(ns.ClientURL())
	 
	 
 Subs(subject)
 
	// Wait for server shutdown
	ns.WaitForShutdown()
}

func Pubs( str_pub []byte){
     
	nc.Publish(subject,str_pub)

}
func Subs(subject string ){
	 
	nc.Subscribe(subject, func(msg *nats.Msg) {
		// Print message data
		data := string(msg.Data)
		fmt.Println(data)

		// Shutdown the server (optional)
		// ns.Shutdown()
	})
    
}
func connect(nats_url string){
    var err error
	nc, err = nats.Connect(nats_url )

	if err != nil {
		panic(err)
	}

    
}