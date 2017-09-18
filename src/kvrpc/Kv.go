package main

import (
	"fmt"
	"log"
	"sync"
)

const (
	OK = "OK"
	ErrNoKey = "ErrNoKey"
)

type Err string

type PutArgs struct {
	Key string
	Value string
}

type GetArgs struct {
	Key string
}

type PutReply struct {
	Err Err
}

type GetReply struct {
	Err Err
	Value string
}

type KV struct {
	mu sync.Mutex
	keyvalue map[string]string
}

func (kv *KV)Get(args *GetArgs, reply *GetReply)  {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Err = "OK"
	val, ok := kv.keyvalue[args.Key]
	if ok {
		reply.Err = OK
		reply.Value = val
	} else {
		reply.Err = OK
		reply.Value = ""
	}
	return nil
}

func (kv *KV)Put(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Lock()

	kv.keyvalue[args.Key] = args.Value
	reply.Err = OK
	return nil
}

func server()  {
	kv := new(KV)
	
}
