// Package Users comment
// This file was generated by tars2go 1.1.6
// Generated from ManageUsers.tars
package Users

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// User struct implement
type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Idcard   string `json:"idcard"`
	Level    int32  `json:"level"`
	Score    int32  `json:"score"`
	Phone    string `json:"phone"`
	State    int32  `json:"state"`
	Promise  bool   `json:"promise"`
	Username string `json:"username"`
	Money    int32  `json:"money"`
}

func (st *User) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *User) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Name, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Idcard, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Level, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Score, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Phone, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.State, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadBool(&st.Promise, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Username, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Money, 9, false)
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
func (st *User) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require User, but not exist. tag %d", tag)
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
func (st *User) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Name, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Idcard, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Level, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Score, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Phone, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.State, 6)
	if err != nil {
		return err
	}

	err = buf.WriteBool(st.Promise, 7)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Username, 8)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Money, 9)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *User) WriteBlock(buf *codec.Buffer, tag byte) error {
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
