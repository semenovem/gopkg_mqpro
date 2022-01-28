package main

import (
  "bytes"
  "fmt"
  "github.com/semenovem/mqm/v2/queue"
  "time"
)

func logMsgIn(m *queue.Msg) {
  logMsg(m, "\n--------------------------------\nПолучили сообщение:")
}

func logMsgOut(m *queue.Msg) {
  logMsg(m, "\n--------------------------------\nОтправили сообщение:")
}

func logMsgBrowse(m *queue.Msg) {
  logMsg(m, "\n--------------------------------\nСообщение:")
}

func logMsgDel(m *queue.Msg) {
  logMsg(m, "\n--------------------------------\nУдалено сообщение:")
}

func logMsg(msg *queue.Msg, s string) {
  if !cfg.logInfo {
    return
  }

  var buf = bytes.NewBufferString(s)
  f := func(s string, i ...interface{}) {
    buf.WriteString(fmt.Sprintf(s, i...))
  }

  if len(msg.Payload) < 300 {
    f(">>>>> msg.Payload  = %s\n", string(msg.Payload))
  } else {
    f(">>>>> len msg.Payload  = %d\n", len(msg.Payload))
  }
  f(">>>>> msg.Props    = %+v\n", msg.Props)
  f(">>>>> msg.CorrelId = %x\n", msg.CorrelId)
  f(">>>>> msg.MsgId    = %x\n", msg.MsgId)
  f(">>>>> msg.MQRFH2   = %s\n", msg.MQRFH2)
  f(">>>>> msg.Time     = %s\n", msg.Time.Format(time.RFC822))

  fmt.Println(buf.String())
}
