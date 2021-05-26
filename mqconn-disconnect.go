package mqpro

func (c *Mqconn) Disconnect() {
  c.mx.Lock()
  defer c.mx.Unlock()
  if c.stateConn == stateDisconnect {
    return
  }
  c.log.Trace("Request to disconnect...")
  c.reqDisconnect()
}

func (c *Mqconn) disconnect() {
  if c.stateConn == stateConnect {
    c.log.Trace("Disconnecting...")
    c._disconnect()
    c.log.Info("Disconnected")
  }
}

func (c *Mqconn) _disconnect() {
  c._unregisterInMsg()

  // TODO обработка ситуации, когда идет обмен данных - подождать
  // MQRC_CALL_IN_PROGRESS

  if c.que != nil {
    c.isWarnConn(c.que.Close(0))
    c.que = nil
  }

  if c.mgr != nil {
    c.isWarnConn(c.mgr.Disc())
    c.mgr = nil
  }
}