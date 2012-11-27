package socket

import (
	"../packet"
	"bytes"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"reflect"
)

type RemoteCodec interface {
	Read(r io.Reader) (packet.IGozillaPacket, error)
	Write(w io.Writer, data packet.IGozillaPacket)
}

type GobCodec2 struct {
	Name string
}

func (g *GobCodec2) Read(r io.Reader) (packet.IGozillaPacket, error) {
	dec := gob.NewDecoder(r)
	// var tobj transObj
	// dec.Decode(tobj)
	// log.Println("read", tobj)
	// data, error := tobj.decode()
	// return data, error

	var cmdName string
	dec.Decode(&cmdName)
	log.Println("READ", cmdName)
	obj := GetCommand(cmdName)
	if obj == nil {
		return nil, errors.New("GetCommand error!")
	}
	err := dec.DecodeValue(reflect.ValueOf(obj))
	return obj.(packet.IGozillaPacket), err
}

func (g *GobCodec2) Write(w io.Writer, data packet.IGozillaPacket) {
	// tobj := new(transObj)
	// err := tobj.encode(data)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("write", tobj)
	// enc := gob.NewEncoder(w)
	// enc.Encode(tobj)
	enc := gob.NewEncoder(w)
	cmdName := GetCommandMark(data)
	log.Println("dowrite", cmdName)
	enc.Encode(cmdName)
	enc.Encode(data)
}

type transObj struct {
	Name   string
	Binary []byte
}

func (t *transObj) encode(data packet.IGozillaPacket) error {
	t.Name = GetCommandMark(data)
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return err
	}
	t.Binary = buf.Bytes()
	return nil
}

func (t *transObj) decode() (packet.IGozillaPacket, error) {
	objRef := GetCommand(t.Name)
	buf := bytes.NewBuffer(t.Binary)
	dec := gob.NewDecoder(buf)
	err := dec.DecodeValue(reflect.ValueOf(objRef))
	if err != nil {
		return nil, err
	}
	return objRef.(packet.IGozillaPacket), nil
}
