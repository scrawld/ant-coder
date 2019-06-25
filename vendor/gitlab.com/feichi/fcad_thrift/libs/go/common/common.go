// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package common

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = enums.GoUnusedProtection__
// Attributes:
//  - Requester: 调用者对自身的描述，如:ad_server#192.168.10.2:8211 *
//  - Sessid: 请求的唯一ID，必须，建议使用UUID，由调用者填写 *
//  - Timestamp: 事件发生的时间，必须，由调用者填写 *
//  - Version: 传输协议版本号，必须，由调用者填写 *
//  - Operator: 操作者UID，如web操作的操作用户，由调用者填写 *
//  - Code: RPC的响应状态，由服务端填写
// 一般约定200为正常响应，其他值为异常
// 
//  - Metadata: 交互数据 *
type Header struct {
  Requester string `thrift:"requester,1" db:"requester" json:"requester"`
  Sessid string `thrift:"sessid,2" db:"sessid" json:"sessid"`
  Timestamp int64 `thrift:"timestamp,3" db:"timestamp" json:"timestamp"`
  Version int32 `thrift:"version,4" db:"version" json:"version"`
  Operator int64 `thrift:"operator,5" db:"operator" json:"operator"`
  Code enums.ResponseCode `thrift:"code,6" db:"code" json:"code"`
  // unused fields # 7 to 9
  Metadata map[string]string `thrift:"metadata,10" db:"metadata" json:"metadata"`
}

func NewHeader() *Header {
  return &Header{}
}


func (p *Header) GetRequester() string {
  return p.Requester
}

func (p *Header) GetSessid() string {
  return p.Sessid
}

func (p *Header) GetTimestamp() int64 {
  return p.Timestamp
}

func (p *Header) GetVersion() int32 {
  return p.Version
}

func (p *Header) GetOperator() int64 {
  return p.Operator
}

func (p *Header) GetCode() enums.ResponseCode {
  return p.Code
}

func (p *Header) GetMetadata() map[string]string {
  return p.Metadata
}
func (p *Header) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 10:
      if fieldTypeId == thrift.MAP {
        if err := p.ReadField10(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Header)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Requester = v
}
  return nil
}

func (p *Header)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Sessid = v
}
  return nil
}

func (p *Header)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Timestamp = v
}
  return nil
}

func (p *Header)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Version = v
}
  return nil
}

func (p *Header)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Operator = v
}
  return nil
}

func (p *Header)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  temp := enums.ResponseCode(v)
  p.Code = temp
}
  return nil
}

func (p *Header)  ReadField10(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]string, size)
  p.Metadata =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.Metadata[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *Header) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Header"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField10(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Header) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("requester", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:requester: ", p), err) }
  if err := oprot.WriteString(string(p.Requester)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.requester (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:requester: ", p), err) }
  return err
}

func (p *Header) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("sessid", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sessid: ", p), err) }
  if err := oprot.WriteString(string(p.Sessid)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sessid (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sessid: ", p), err) }
  return err
}

func (p *Header) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err) }
  if err := oprot.WriteI64(int64(p.Timestamp)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err) }
  return err
}

func (p *Header) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("version", thrift.I32, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:version: ", p), err) }
  if err := oprot.WriteI32(int32(p.Version)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.version (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:version: ", p), err) }
  return err
}

func (p *Header) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("operator", thrift.I64, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:operator: ", p), err) }
  if err := oprot.WriteI64(int64(p.Operator)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.operator (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:operator: ", p), err) }
  return err
}

func (p *Header) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("code", thrift.I32, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:code: ", p), err) }
  if err := oprot.WriteI32(int32(p.Code)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.code (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:code: ", p), err) }
  return err
}

func (p *Header) writeField10(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("metadata", thrift.MAP, 10); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:metadata: ", p), err) }
  if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Metadata)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.Metadata {
    if err := oprot.WriteString(string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 10:metadata: ", p), err) }
  return err
}

func (p *Header) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Header(%+v)", *p)
}

// Attributes:
//  - Topic: topic *
//  - Requester: 调用者对自身的描述，如:ad_server#192.168.10.2:8211 *
//  - Sessid: 请求的唯一ID，必须，建议使用UUID，由调用者填写 *
//  - Timestamp: 事件发生的时间，必须，由调用者填写 *
//  - Version: 传输协议版本号，必须，由调用者填写 *
//  - Operator: 操作者UID，如web操作的操作用户，由调用者填写 *
//  - Metadata: 交互数据 *
type EventHeader struct {
  Topic enums.EventTopic `thrift:"topic,1" db:"topic" json:"topic"`
  Requester string `thrift:"requester,2" db:"requester" json:"requester"`
  Sessid string `thrift:"sessid,3" db:"sessid" json:"sessid"`
  Timestamp int64 `thrift:"timestamp,4" db:"timestamp" json:"timestamp"`
  Version int32 `thrift:"version,5" db:"version" json:"version"`
  Operator int64 `thrift:"operator,6" db:"operator" json:"operator"`
  // unused fields # 7 to 9
  Metadata map[string]string `thrift:"metadata,10" db:"metadata" json:"metadata"`
}

func NewEventHeader() *EventHeader {
  return &EventHeader{}
}


func (p *EventHeader) GetTopic() enums.EventTopic {
  return p.Topic
}

func (p *EventHeader) GetRequester() string {
  return p.Requester
}

func (p *EventHeader) GetSessid() string {
  return p.Sessid
}

func (p *EventHeader) GetTimestamp() int64 {
  return p.Timestamp
}

func (p *EventHeader) GetVersion() int32 {
  return p.Version
}

func (p *EventHeader) GetOperator() int64 {
  return p.Operator
}

func (p *EventHeader) GetMetadata() map[string]string {
  return p.Metadata
}
func (p *EventHeader) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 10:
      if fieldTypeId == thrift.MAP {
        if err := p.ReadField10(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EventHeader)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := enums.EventTopic(v)
  p.Topic = temp
}
  return nil
}

func (p *EventHeader)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Requester = v
}
  return nil
}

func (p *EventHeader)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Sessid = v
}
  return nil
}

func (p *EventHeader)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Timestamp = v
}
  return nil
}

func (p *EventHeader)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Version = v
}
  return nil
}

func (p *EventHeader)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.Operator = v
}
  return nil
}

func (p *EventHeader)  ReadField10(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]string, size)
  p.Metadata =  tMap
  for i := 0; i < size; i ++ {
var _key2 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key2 = v
}
var _val3 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val3 = v
}
    p.Metadata[_key2] = _val3
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *EventHeader) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("EventHeader"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField10(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EventHeader) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("topic", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:topic: ", p), err) }
  if err := oprot.WriteI32(int32(p.Topic)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.topic (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:topic: ", p), err) }
  return err
}

func (p *EventHeader) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("requester", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:requester: ", p), err) }
  if err := oprot.WriteString(string(p.Requester)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.requester (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:requester: ", p), err) }
  return err
}

func (p *EventHeader) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("sessid", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:sessid: ", p), err) }
  if err := oprot.WriteString(string(p.Sessid)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sessid (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:sessid: ", p), err) }
  return err
}

func (p *EventHeader) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:timestamp: ", p), err) }
  if err := oprot.WriteI64(int64(p.Timestamp)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.timestamp (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:timestamp: ", p), err) }
  return err
}

func (p *EventHeader) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("version", thrift.I32, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:version: ", p), err) }
  if err := oprot.WriteI32(int32(p.Version)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.version (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:version: ", p), err) }
  return err
}

func (p *EventHeader) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("operator", thrift.I64, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:operator: ", p), err) }
  if err := oprot.WriteI64(int64(p.Operator)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.operator (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:operator: ", p), err) }
  return err
}

func (p *EventHeader) writeField10(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("metadata", thrift.MAP, 10); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:metadata: ", p), err) }
  if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Metadata)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.Metadata {
    if err := oprot.WriteString(string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 10:metadata: ", p), err) }
  return err
}

func (p *EventHeader) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EventHeader(%+v)", *p)
}
