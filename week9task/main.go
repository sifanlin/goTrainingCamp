package week9task

import (
	"encoding/binary"
	"errors"
	"github.com/Terry-Mao/goim/api/protocol"
)

const (
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	MaxBodySize    = int32(1 << 12)
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	_packOffset    = 0
	_headerOffset  = _packOffset + _packSize
	_verOffset     = _headerOffset + _headerSize
	_opOffset      = _verOffset + _verSize
	_seqOffset     = _opOffset + _opSize
	_seqLen        = 4
	_bodyOffset    = _seqOffset + _seqLen
)

// decode 解码器
func decode(buf []byte) (protocol.Proto, error) {
	var (
		headerLen uint16
		packLen   uint32
		proto     protocol.Proto
	)

	packLen = binary.BigEndian.Uint32(buf[_packOffset:_headerOffset])
	headerLen = binary.BigEndian.Uint16(buf[_headerOffset:_verOffset])
	proto.Ver = int32(binary.BigEndian.Uint32(buf[_verOffset:_opOffset]))
	proto.Op = int32(binary.BigEndian.Uint32(buf[_opOffset:_seqOffset]))
	proto.Seq = int32(binary.BigEndian.Uint32(buf[_seqOffset : _seqOffset+_seqLen]))
	if int32(packLen) > _maxPackSize {
		return protocol.Proto{}, errors.New("超过最大包长度")
	}
	if int32(headerLen) != _rawHeaderSize {
		return protocol.Proto{}, errors.New("header err")
	}
	proto.Body = buf[16:]
	return proto, nil
}

func encode(proto protocol.Proto) (err error) {
	var (
		packLen int32
	)
	packLen = _rawHeaderSize + int32(len(proto.Body))
	buf := make([]byte, packLen)
	binary.BigEndian.PutUint32(buf[_packOffset:_headerOffset], uint32(packLen))
	binary.BigEndian.PutUint16(buf[_headerOffset:_verOffset], uint16(_rawHeaderSize))
	binary.BigEndian.PutUint16(buf[_verOffset:_opOffset], uint16(proto.Ver))
	binary.BigEndian.PutUint32(buf[_opOffset:_seqOffset], uint32(proto.Op))
	binary.BigEndian.PutUint32(buf[_seqOffset:_bodyOffset], uint32(proto.Seq))
	copy(buf[_bodyOffset:], proto.Body)
	return nil
}
