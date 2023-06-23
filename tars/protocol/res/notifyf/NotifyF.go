// Code generated by tars2go 1.2.3, DO NOT EDIT.
// This file was generated from NotifyF.tars
// Package notifyf comment
package notifyf

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

//go:generate stringer -type NOTIFYLEVEL -trimprefix NOTIFYLEVEL_ -output notifylevel_string.go
type NOTIFYLEVEL int32

const (
	NOTIFYLEVEL_NOTIFYNORMAL NOTIFYLEVEL = 0
	NOTIFYLEVEL_NOTIFYWARN   NOTIFYLEVEL = 1
	NOTIFYLEVEL_NOTIFYERROR  NOTIFYLEVEL = 2
)

//go:generate stringer -type ReportType -trimprefix ReportType_ -output reporttype_string.go
type ReportType int32

const (
	ReportType_REPORT ReportType = 0
	ReportType_NOTIFY ReportType = 1
)

// NotifyKey struct implement
type NotifyKey struct {
	Name string `json:"name" tars:"name,tag:1,require:true"`
	Ip   string `json:"ip" tars:"ip,tag:2,require:true"`
	Page int32  `json:"page" tars:"page,tag:3,require:true"`
}

func (st *NotifyKey) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NotifyKey) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Name, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Ip, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Page, 3, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NotifyKey) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NotifyKey, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NotifyKey) WriteTo(buf *codec.Buffer) (err error) {
	err = buf.WriteString(st.Name, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Ip, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Page, 3)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *NotifyKey) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// NotifyItem struct implement
type NotifyItem struct {
	STimeStamp string `json:"sTimeStamp" tars:"sTimeStamp,tag:1,require:true"`
	SServerId  string `json:"sServerId" tars:"sServerId,tag:2,require:true"`
	ILevel     int32  `json:"iLevel" tars:"iLevel,tag:3,require:true"`
	SMessage   string `json:"sMessage" tars:"sMessage,tag:4,require:true"`
}

func (st *NotifyItem) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NotifyItem) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.STimeStamp, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SServerId, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.ILevel, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SMessage, 4, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NotifyItem) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NotifyItem, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NotifyItem) WriteTo(buf *codec.Buffer) (err error) {
	err = buf.WriteString(st.STimeStamp, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SServerId, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.ILevel, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SMessage, 4)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *NotifyItem) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// NotifyInfo struct implement
type NotifyInfo struct {
	Nextpage    int32        `json:"nextpage" tars:"nextpage,tag:1,require:true"`
	NotifyItems []NotifyItem `json:"notifyItems" tars:"notifyItems,tag:2,require:true"`
}

func (st *NotifyInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *NotifyInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Nextpage, 1, true)
	if err != nil {
		return err
	}

	_, ty, err = readBuf.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}
	if ty == codec.LIST {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}
		st.NotifyItems = make([]NotifyItem, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {
			err = st.NotifyItems[i0].ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SimpleList {
		err = fmt.Errorf("not support SimpleList type")
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *NotifyInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NotifyInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *NotifyInfo) WriteTo(buf *codec.Buffer) (err error) {
	err = buf.WriteInt32(st.Nextpage, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = buf.WriteInt32(int32(len(st.NotifyItems)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.NotifyItems {
		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}
	}

	return err
}

// WriteBlock encode struct
func (st *NotifyInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// ReportInfo struct implement
type ReportInfo struct {
	EType      ReportType  `json:"eType" tars:"eType,tag:1,require:true"`
	SApp       string      `json:"sApp" tars:"sApp,tag:2,require:true"`
	SSet       string      `json:"sSet" tars:"sSet,tag:3,require:true"`
	SContainer string      `json:"sContainer" tars:"sContainer,tag:4,require:true"`
	SServer    string      `json:"sServer" tars:"sServer,tag:5,require:true"`
	SMessage   string      `json:"sMessage" tars:"sMessage,tag:6,require:true"`
	SThreadId  string      `json:"sThreadId" tars:"sThreadId,tag:7,require:false"`
	ELevel     NOTIFYLEVEL `json:"eLevel" tars:"eLevel,tag:8,require:false"`
	SNodeName  string      `json:"sNodeName" tars:"sNodeName,tag:9,require:false"`
}

func (st *ReportInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ReportInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32((*int32)(&st.EType), 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SApp, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SSet, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SContainer, 4, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SServer, 5, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SMessage, 6, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SThreadId, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32((*int32)(&st.ELevel), 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SNodeName, 9, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ReportInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ReportInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ReportInfo) WriteTo(buf *codec.Buffer) (err error) {
	err = buf.WriteInt32(int32(st.EType), 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SApp, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SSet, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SContainer, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SServer, 5)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SMessage, 6)
	if err != nil {
		return err
	}

	if st.SThreadId != "" {
		err = buf.WriteString(st.SThreadId, 7)
		if err != nil {
			return err
		}
	}

	err = buf.WriteInt32(int32(st.ELevel), 8)
	if err != nil {
		return err
	}

	if st.SNodeName != "" {
		err = buf.WriteString(st.SNodeName, 9)
		if err != nil {
			return err
		}
	}

	return err
}

// WriteBlock encode struct
func (st *ReportInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}
