package rtc

import (
	"flag"
    "fmt"
	"github.com/pion/webrtc/v4"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/intervalpli"
)

func Rtc(){		
   port:=flag.Int("port",8080,"http server port")
   flag.Parse()
	fmt.Println(*port)
   sdpChan :=HttpSDPServer(*port)
   offer:=webrtc.SessionDescription{}
   Decode(<-sdpChan,&offer)
   
   peerConnectionConfig := webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs: []string{"stun:stun.l.google.com:19302"},
		},
	},
	}

	m:=&webrtc.MediaEngine{}
    if err := m.RegisterDefaultCodecs(); err != nil {
		panic(err)
	}


    //interceptor creation
	i := &interceptor.Registry{}

	if err := webrtc.RegisterDefaultInterceptors(m, i); err != nil {
		panic(err)
	}

	intervalPliFactory, err := intervalpli.NewReceiverInterceptor()
	if err != nil {
		panic(err)
	}
	i.Add(intervalPliFactory)

    
	
	//creating new Api
	peerConnection, err := webrtc.NewAPI(webrtc.WithMediaEngine(m), webrtc.WithInterceptorRegistry(i)).NewPeerConnection(peerConnectionConfig)
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()


	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			fmt.Println("Peer Connection has gone to failed exiting")
			return
		}

		if s == webrtc.PeerConnectionStateClosed {
			// PeerConnection was explicitly closed. This usually happens from a DTLS CloseNotify
			fmt.Println("Peer Connection has gone to closed exiting")
		    return
		}
	})


    peerConnection.OnDataChannel(func(d *webrtc.DataChannel){
		fmt.Printf("new data channel is created at %d",d.ID())



		d.OnOpen(func(){
			message := "connected first time"
				
			fmt.Printf("Sending '%s'\n", message)
			if sendErr := d.SendText(message); sendErr != nil {
				panic(sendErr)
			}
		   })

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	
	})


	


	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		panic(err)
	}

	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
	}
	<-gatherComplete

	fmt.Println(Encode(peerConnection.LocalDescription()))



	for {
		fmt.Println("")
		fmt.Println("Curl an base64 SDP to start sendonly peer connection")

		recvOnlyOffer := webrtc.SessionDescription{}
		Decode(<-sdpChan, &recvOnlyOffer)

		// Create a new PeerConnection
		peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfig)
		if err != nil {
			panic(err)
		}
 
		peerConnection.OnDataChannel(func(d *webrtc.DataChannel){
			fmt.Printf("new data channel is created at %d",d.ID())
	
	
	       d.OnOpen(func(){
			message := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				
			fmt.Printf("Sending '%s'\n", message)
			if sendErr := d.SendText(message); sendErr != nil {
				panic(sendErr)
			}
		   })
			
	
			d.OnMessage(func(msg webrtc.DataChannelMessage) {
				fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
			})
		
		})


		// Set the remote SessionDescription
		err = peerConnection.SetRemoteDescription(recvOnlyOffer)
		if err != nil {
			panic(err)
		}

		// Create answer
		answer, err := peerConnection.CreateAnswer(nil)
		if err != nil {
			panic(err)
		}

		// Create channel that is blocked until ICE Gathering is complete
		gatherComplete = webrtc.GatheringCompletePromise(peerConnection)

		// Sets the LocalDescription, and starts our UDP listeners
		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			panic(err)
		}

		// Block until ICE Gathering is complete, disabling trickle ICE
		// we do this because we only can exchange one signaling message
		// in a production application you should exchange ICE Candidates via OnICECandidate
		<-gatherComplete

		// Get the LocalDescription and take it to base64 so we can paste in browser
		fmt.Println(Encode(peerConnection.LocalDescription()))
	}

}

