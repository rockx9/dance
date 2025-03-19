package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type PaginatorRequest struct {
	Page     int64 `json:"page" form:"page" default:"1"`
	PageSize int64 `json:"pageSize" form:"pageSize" default:"10"`
}

func GetStringPtr(s string) *string {
	return &s
}

func GetStringVal(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r BaseResponse) Verify() error {
	if r.Code == 0 {
		return nil
	}
	return errors.New(r.Msg)
}

func BuildResponse(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func ToJsonString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("to json err: %s", err.Error())
	}
	return string(b)
}

func PrintJson(data interface{}) {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(data)
		return
	}
	fmt.Println(string(b))
}

func StructToReader(p interface{}) (io.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}
