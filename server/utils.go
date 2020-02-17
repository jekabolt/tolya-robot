package server

import (
	"fmt"
	"io"

	"encoding/json"
	"io/ioutil"

	"github.com/jekabolt/tolya-robot/schemas"
)

func UnmarshalConsumer(body io.ReadCloser) (*schemas.Consumer, error) {
	consumer := &schemas.Consumer{}
	rawBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalConsumer:ioutil.ReadAll:err [%v]", err.Error())
	}
	err = json.Unmarshal(rawBody, consumer)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalConsumer:json.Unmarshal: [%v]", err.Error())
	}
	return consumer, nil
}
