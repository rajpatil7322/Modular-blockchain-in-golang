package main

// "github.com/rajpatil7322/Modular-blockchain-in-golang/network"

func main() {
	// trlocal := network.NewLocalTransport("LOCAL")
	// trRemote := network.NewLocalTransport("REMOTE")

	// trlocal.Connect(trRemote)
	// trRemote.Connect(trlocal)
	// // fmt.Println(trRemote.Addr())
	// // fmt.Println(trlocal)

	// go func() {
	// 	for {
	// 		trRemote.SendMessage(trlocal.Addr(), []byte("Hii Appa"))
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// opts := network.ServerOpts{
	// 	Transports: []network.Transport{trlocal},
	// }

	// server := network.NewServer(opts)

	// server.Start()
	// private_key := crypto.GenratePrivateKey()
	// fmt.Println(private_key)
	// public_key := private_key.GenratePublicKey()
	// fmt.Println(public_key)
	// address := public_key.Address().String()
	// fmt.Println(address)

	// msg := []byte("Hello World")
	// sig, err := private_key.Sign(msg)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// res := sig.Verify(public_key, msg)
	// if res {
	// 	fmt.Println("Verfied")
	// }

}
