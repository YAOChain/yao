package mock

import (
	"github.com/YAOChain/yao/core/libs/log"
	"github.com/YAOChain/yao/core/p2p"
	"github.com/YAOChain/yao/core/p2p/conn"
)

type Reactor struct {
	p2p.BaseReactor
}

func NewReactor() *Reactor {
	r := &Reactor{}
	r.BaseReactor = *p2p.NewBaseReactor("Reactor", r)
	r.SetLogger(log.TestingLogger())
	return r
}

func (r *Reactor) GetChannels() []*conn.ChannelDescriptor            { return []*conn.ChannelDescriptor{} }
func (r *Reactor) AddPeer(peer p2p.Peer)                             {}
func (r *Reactor) RemovePeer(peer p2p.Peer, reason interface{})      {}
func (r *Reactor) Receive(chID byte, peer p2p.Peer, msgBytes []byte) {}
