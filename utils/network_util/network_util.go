package network_util

import (
	"errors"
	"github.com/liyuhuana/go-common/common_logger"
	"github.com/liyuhuana/go-common/network/network_tcp"
	"google.golang.org/protobuf/proto"
)

func PushMessage(s *network_tcp.Session, msgId int32, msg proto.Message) error {
	if s == nil {
		err := errors.New("PushMessage fail, session is nil")
		common_logger.LogError(err)
		return err
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		common_logger.LogError("proto.Marshal fail, msgId:", msgId, "error:", err)
		return err
	}

	err = s.Push(data)
	if err != nil {
		common_logger.LogError(err)
	}
	return err
}
