package socket

import (
	"../packet"
	"container/list"
)

type gateMessageQueue struct {
	nidList    *list.List
	packetList *list.List
}

func newGateMessageQueue() *gateMessageQueue {
	queue := new(gateMessageQueue)
	queue.nidList = list.New()
	queue.packetList = list.New()
	return queue
}

func (g *gateMessageQueue) PushPacket(NclientId int, packet packet.IGozillaPacket) {
	g.nidList.PushBack(NclientId)
	g.packetList.PushBack(packet)
}

func (g *gateMessageQueue) PumpPacket() (int, packet.IGozillaPacket) {
	if g.nidList.Len() == 0 {
		return -1, nil
	}
	nid_itf := g.nidList.Front()
	nid := nid_itf.Value
	g.nidList.Remove(nid_itf)

	p_itf := g.packetList.Front()
	pac := p_itf.Value
	g.packetList.Remove(p_itf)
	return nid.(int), pac.(packet.IGozillaPacket)
}
