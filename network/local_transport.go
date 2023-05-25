package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr       NetAddr
	consumeChn chan RPC
	lock       sync.RWMutex
	peers      map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:       addr,
		consumeChn: make(chan RPC, 1024),
		peers:      make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) GetPeers() map[NetAddr]*LocalTransport {
	return t.peers
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeChn
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("could not send message")
	}

	peer.consumeChn <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

// func TestConnect() {
// 	tra := NewLocalTransport("A")
// 	trb := NewLocalTransport("B")

// 	tra.Connect(trb)
// 	trb.Connect(tra)
// 	// assert.Equal(t, tra.peers[trb.addr], trb)
// 	// assert.Equal(t, trb.peers[tra.addr], tra)

// 	fmt.Println(tra.peers[trb.addr])
// }
