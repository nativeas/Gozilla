package socket

import (
	"../packet"
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"reflect"
)

type PCodec struct {
	Name string
}

func (p *PCodec) Read(r io.Reader) (packet.IGozillaPacket, error) {
	log.Println("HELLO")
	dec := gob.NewDecoder(r)
	var obj gobPObj
	dec.Decode(&obj)
	buf := bytes.NewBuffer(obj.Binary)
	dec = gob.NewDecoder(buf)
	var pkt = (GetCommand(obj.Name))
	dec.Decode(&pkt)
	log.Print(pkt)
	var ppt packet.IGozillaPacket = pkt.(packet.IGozillaPacket)
	return ppt, nil
	// return pkt, nil
}

func (p *PCodec) Write(w io.Writer, data packet.IGozillaPacket) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	// bin := buf.Bytes()
	obj := new(gobPObj)
	obj.Name = reflect.TypeOf(data).Name()
	wrt := gob.NewEncoder(w)
	wrt.Encode(obj)
}

type gobPObj struct {
	Name   string
	Binary []byte
}
