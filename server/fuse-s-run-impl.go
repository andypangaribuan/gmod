/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"log"
	"maps"

	"github.com/andypangaribuan/gmod/gm"
)

func (slf *stuFuseSRouter) run(path string, registerCondition *func(ctx FuseSContext) bool, userUidCondition *func(ctx FuseSContext) (ok bool, userUid string)) *stuFuseSRun {
	sock := &stuFuseSRunWebsocket{
		register:   make(chan *stuFuseSRunClient),
		unregister: make(chan *stuFuseSRunClient),
		broadcast:  make(chan *stuFuseSRunBroadcastMessage),
		clients:    make(map[string]*stuFuseSRunClient, 0),
	}

	go sock.handler()

	slf.Register(path, func(ctx FuseSContext) {
		clientObj := &stuFuseSRunClient{
			ctx: ctx,
			uid: gm.Util.UID(),
		}

		if registerCondition != nil && !(*registerCondition)(ctx) {
			ctx.Close()
			return
		}

		if userUidCondition != nil {
			ok, userUid := (*userUidCondition)(ctx)
			if !ok {
				ctx.Close()
				return
			}

			clientObj.userUid = userUid
		}

		sock.register <- clientObj
		defer func() {
			sock.unregister <- clientObj
		}()

		for {
			// message, err := ctx.ReadMessage()
			_, err := ctx.ReadMessage()
			if err != nil {
				return
			}

			// if message != "" {
			// 	sock.broadcast <- &stuFuseSRunBroadcastMessage{
			// 		message: message,
			// 	}
			// }
		}
	})

	return &stuFuseSRun{
		sock: sock,
	}
}

func (slf *stuFuseSRunWebsocket) handler() {
	for {
		select {
		case client := <-slf.register:
			go slf.addClient(client)

		case client := <-slf.unregister:
			go slf.removeClient(client)

		case message := <-slf.broadcast:
			go slf.broadcastMessage(message)
		}
	}
}

func (slf *stuFuseSRunWebsocket) addClient(client *stuFuseSRunClient) {
	slf.mx.Lock()
	defer slf.mx.Unlock()

	slf.clients[client.uid] = client
	log.Printf("add-client   : %v\n", client.uid)
}

func (slf *stuFuseSRunWebsocket) removeClient(client *stuFuseSRunClient) {
	slf.mx.Lock()
	defer slf.mx.Unlock()

	delete(slf.clients, client.uid)
	client.ctx.Close()
	log.Printf("remove-client: %v\n", client.uid)
}

func (slf *stuFuseSRunWebsocket) broadcastMessage(msg *stuFuseSRunBroadcastMessage) {
	var (
		clients    map[string]*stuFuseSRunClient
		deleteList = make([]*stuFuseSRunClient, 0)
		delivered  = false
	)

	clients = slf.cloneClients()

	if len(clients) == 0 {
		return
	}

	if msg.userUid != "" {
		filtered := make(map[string]*stuFuseSRunClient, 0)
		for k, c := range clients {
			if c.userUid == msg.userUid {
				filtered[k] = c
			}
		}

		if len(filtered) == 0 {
			return
		}

		clients = filtered
	}

	for _, client := range clients {
		err := client.ctx.WriteMessage(msg.message)
		if err != nil {
			deleteList = append(deleteList, client)
		} else {
			delivered = true
		}
	}

	if delivered {
		log.Printf("broadcast    : %v\n", msg.message)
	}

	if len(deleteList) > 0 {
		slf.deleteClients(deleteList)
		for _, client := range deleteList {
			client.ctx.Close()
		}
	}
}

func (slf *stuFuseSRunWebsocket) cloneClients() map[string]*stuFuseSRunClient {
	slf.mx.Lock()
	defer slf.mx.Unlock()
	return maps.Clone(slf.clients)
}

func (slf *stuFuseSRunWebsocket) deleteClients(deletedClients []*stuFuseSRunClient) {
	slf.mx.Lock()
	defer slf.mx.Unlock()

	for _, client := range deletedClients {
		delete(slf.clients, client.uid)
		log.Printf("remove-client: %v\n", client.uid)
	}
}

func (slf *stuFuseSRun) Broadcast(message string) {
	if slf == nil || slf.sock == nil {
		return
	}

	slf.sock.broadcast <- &stuFuseSRunBroadcastMessage{
		message: message,
	}
}

func (slf *stuFuseSRun) BroadcastUser(userUid string, message string) {
	if slf == nil || slf.sock == nil {
		return
	}

	slf.sock.broadcast <- &stuFuseSRunBroadcastMessage{
		userUid: userUid,
		message: message,
	}
}
