package manager

import (
  "time"
)

func (m *Mqpro) workerState() {
  var (
    err error
    st  state
    l   = m.log.WithField("mod", "workerState")
  )

worker:
  for st = range m.chState {
    l.Debug(stateMapByKey[m.state], " >>> ", stateMapByKey[st])

    if st == m.state {
      continue
    }

    switch st {
    case stateConn:
      if m.state == stateConnecting {
        continue
      }
      m.state = stateConnecting

      for {
        err = m.connect()
        if err == nil {
          m.state = stateConn
          m.fireConn()
          continue worker
        }
        l.WithField("oper", "conn").Warn(err)

        select {
        case <-m.ctx.Done():
          continue worker
        case <-time.After(m.reconnDelay):
        }
      }

    case stateDisconn:
      m.state = stateDisconn
      m.disconnect()

    case stateErr:
      m.state = stateErr
      m.disconnect()
      go m.stateConn()
    }
  }
}

func (m *Mqpro) stateConn() {
  m.chState <- stateConn
}

func (m *Mqpro) stateDisconn() {
  m.chState <- stateDisconn
}

func (m *Mqpro) stateErr() {
  if m.state == stateConn {
    m.chState <- stateErr
  }
}

func (m *Mqpro) IsConn() bool {
  return m.state == stateConn
}

func (m *Mqpro) IsDisconn() bool {
  return m.state == stateDisconn
}