// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package api

import (
	"code.google.com/p/goprotobuf/proto"
	"gls/proto/pb"
	"reflect"
	"utils/errors"
	"utils/packet"
)

const (
	_heart_beat_req = int16(iota)
	_user_login_req
)

type ProtoHandler interface {
	//parse input into reqMsg
	ParseInput(*packet.Packet, proto.Message) errors.GalaxyError

	VerifyParams(proto.Message) errors.GalaxyError

	ProcessLogic(proto.Message, proto.Message) errors.GalaxyError

	//parse rspMsg into output
	ParseOutPut(proto.Message) []byte
}

var (
	ServiceList   map[int16]string // used for check cliet protocol num, if not bind, can't be in api package
	protoHandlers map[int16]reflect.Type
	protoReqs     map[int16]reflect.Type
	protoRsps     map[int16]reflect.Type
)

func init() {

	ServiceList = map[int16]string{
		_heart_beat_req: "heartBeat",
	}

	protoHandlers = map[int16]reflect.Type{
		_heart_beat_req: reflect.TypeOf(HeartBeatHandler{}),
	}

	protoReqs = map[int16]reflect.Type{
		_heart_beat_req: reflect.TypeOf(pb.ReqHeartBeat{}),
	}

	protoRsps = map[int16]reflect.Type{
		_heart_beat_req: reflect.TypeOf(pb.RspHeartBeat{}),
	}
}

func SafeProtoHandler(svr int16, reader *packet.Packet) []byte {

	handleType := protoHandlers[svr]
	reqType := protoReqs[svr]
	rspType := protoRsps[svr]

	handler := reflect.New(handleType).Interface().(ProtoHandler)
	reqMsg := reflect.New(reqType).Interface().(proto.Message)
	rspMsg := reflect.New(rspType).Interface().(proto.Message)

	e := handler.ParseInput(reader, reqMsg)
	if e.IsError() {
		return handler.ParseOutPut(rspMsg)
	}

	e = handler.VerifyParams(reqMsg)
	if e.IsError() {
		return handler.ParseOutPut(rspMsg)
	}

	// game logic
	e = handler.ProcessLogic(reqMsg, rspMsg)

	return handler.ParseOutPut(rspMsg)
}
