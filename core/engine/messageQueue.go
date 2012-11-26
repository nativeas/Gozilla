package engine

import (
	"../packet"
	"container/list"
	"log"
)

//
//实现一个消息队列
//用于实现消息传送
type messageQueue struct {
	clientMapping           map[int]int
	outputPacketQueue       *list.List
	outputPacketClientQueue *list.List
	inputPacketQueue        *list.List
	inputPacketClientQueue  *list.List
}

func (m *messageQueue) Init() {
	m.clientMapping = make(map[int]int)
	m.outputPacketQueue = list.New()
	m.outputPacketClientQueue = list.New()
	m.inputPacketQueue = list.New()
	m.inputPacketClientQueue = list.New()
}

// messageQueue -> Engine
func (m *messageQueue) Pump() (nid int, pump_pkt packet.IGozillaPacket) {
	if m.outputPacketQueue.Len() == 0 {
		return -1, nil
	}
	cid_v := m.outputPacketClientQueue.Front()
	cid := cid_v.Value
	m.outputPacketClientQueue.Remove(cid_v)
	pkt_v := m.outputPacketQueue.Front()
	pkt := pkt_v.Value
	m.outputPacketQueue.Remove(pkt_v)
	return cid.(int), pkt.(packet.IGozillaPacket)
}

// codec -> messageQueue
func (m *messageQueue) PushPacket(NclientId int, remoteId int, packet packet.IGozillaPacket) {
	newClientId := NclientId*1000 + remoteId
	m.clientMapping[newClientId] = remoteId
	m.outputPacketQueue.PushBack(packet)
	m.outputPacketClientQueue.PushBack(newClientId)
}

//codec <- messageQueue
func (m *messageQueue) RepyPacket() (NclientId int, remoteId int, reply_pkt packet.IGozillaPacket) {
	tmp := m.inputPacketClientQueue.Back()
	nid := tmp.Value
	m.inputPacketClientQueue.Remove(tmp)
	tmp2 := m.inputPacketQueue.Back()
	pkt := tmp2.Value
	m.inputPacketQueue.Remove(tmp2)
	if val, ok := m.clientMapping[nid.(int)]; ok {
		cid := (nid.(int) - val/1000)
		return cid, val, pkt.(packet.IGozillaPacket)
	}
	return -1, -1, nil
}

//messageQueue <- engine
func (m *messageQueue) Push(nid int, packet packet.IGozillaPacket) {
	if _, ok := m.clientMapping[nid]; ok {
		m.inputPacketClientQueue.PushBack(nid)
		m.inputPacketQueue.PushBack(packet)
	} else {
		log.Println("Push() nid %d not exist in clientMapping", nid)
	}
}
