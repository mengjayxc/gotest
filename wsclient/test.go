package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Command struct {
	Req     string      `json:"req"`
	Rid     string      `json:"rid"`
	Args    interface{} `json:"args"`
	Expires int64       `json:"expires"`
}

type UnSubOption []string

type Topic struct {
	Subj   string `json:"subj"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
}

func (t *Topic) Decode(topic string) {
	items := strings.SplitN(topic, "_", 3)
	t.Subj = items[0]
	if len(items) == 2 {
		t.Symbol = items[1]
	} else if len(items) == 3 {
		t.Type = items[1]
		t.Symbol = items[2]
	}
}

func main() {
	var cmd = Command{
		Req:     "UnSub",
		Rid:     "12",
		Args:    []string{"index_BTC", "index_ETH"},
		Expires: 0,
	}

	onCmdUnSub(&cmd)

}

// 处理取消订阅
func onCmdUnSub(cmd *Command) {
	var option UnSubOption
	jsonArgs, err := json.Marshal(cmd.Args)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonArgs, &option)
	if err != nil {
		return
	}

	for _, v := range option {
		var topic Topic
		topic.Decode(v)
		//if !topic.IsValid() {
		//	log.ZapLogger.Warn(fmt.Sprintf("UnSub Invalid topic: %+v", topic))
		//	continue
		//}
		//if topic.Subj == "index" {
		//	s.topics.Delete(topic.EncodeIndex())
		//} else if topic.Subj == "kline" {
		//	s.topics.Delete(topic.EncodeKline())
		//}

		fmt.Printf("topic: ++++ %+v", topic)
	}

	//s.sendResponse(cmd.Rid, types.StatusOK, types.StatusText(types.StatusOK))
}
