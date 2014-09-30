// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package api

import (
	"code.google.com/p/goprotobuf/proto"
	"utils/errors"
	"utils/packet"
)

// nothing should be done for heart beat
type HeartBeatHandler struct {
}

func (p *HeartBeatHandler) ParseInput(req *packet.Packet, reqMsg proto.Message) errors.GalaxyError {
	return errors.ErrorOK()
}

func (p *HeartBeatHandler) VerifyParams(reqMsg proto.Message) errors.GalaxyError {
	return errors.ErrorOK()
}

func (p *HeartBeatHandler) ProcessLogic(reqMsg, rspMsg proto.Message) errors.GalaxyError {
	return errors.ErrorOK()
}

func (p *HeartBeatHandler) ParseOutPut(rspMsg proto.Message) (outBuffer []byte) {
	return nil
}
