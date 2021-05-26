package mqpro

import (
  "github.com/caarlos0/env/v6"
)

// configuration for ibm mq
type envCfg struct {
  MQ0Host     string `env:"ENV_MQ_0_HOST"`
  MQ0Port     int    `env:"ENV_MQ_0_PORT"`
  MQ0Mgr      string `env:"ENV_MQ_0_MGR"`
  MQ0Channel  string `env:"ENV_MQ_0_CHANNEL"`
  MQ0PutQueue string `env:"ENV_MQ_0_PUT_QUEUE"`
  MQ0GetQueue string `env:"ENV_MQ_0_GET_QUEUE"`
  MQ0BrowseQ  string `env:"ENV_MQ_0_BROWSE_QUEUE"`
  MQ0App      string `env:"ENV_MQ_0_APP"`
  MQ0User     string `env:"ENV_MQ_0_USER"`
  MQ0Pass     string `env:"ENV_MQ_0_PASS"`
  MQ0Priority string `env:"ENV_MQ_0_PRIORITY"`

  MQ1Host     string `env:"ENV_MQ_1_HOST"`
  MQ1Port     int    `env:"ENV_MQ_1_PORT"`
  MQ1Mgr      string `env:"ENV_MQ_1_MGR"`
  MQ1Channel  string `env:"ENV_MQ_1_CHANNEL"`
  MQ1PutQueue string `env:"ENV_MQ_1_PUT_QUEUE"`
  MQ1GetQueue string `env:"ENV_MQ_1_GET_QUEUE"`
  MQ1BrowseQ  string `env:"ENV_MQ_1_BROWSE_QUEUE"`
  MQ1App      string `env:"ENV_MQ_1_APP"`
  MQ1User     string `env:"ENV_MQ_1_USER"`
  MQ1Pass     string `env:"ENV_MQ_1_PASS"`
  MQ1Priority string `env:"ENV_MQ_1_PRIORITY"`
}

func (p *Mqpro) UseDefEnv() {
  p.SetConn(getConnFromEnv()...)
}

func getConnFromEnv() []*Mqconn {
  var cfg = &envCfg{}

  if err := env.Parse(cfg); err != nil {
    Log.Error(err)
  }

  Log.Debugf("Application configuration: %+v", *cfg)

  connLi := make([]*Mqconn, 0)

  if cfg.MQ0Host != "" {
    if cfg.MQ0PutQueue != "" {
      conn := MqconnNew(TypePut, Log, &Cfg{
        Host:        cfg.MQ0Host,
        Port:        cfg.MQ0Port,
        MgrName:     cfg.MQ0Mgr,
        ChannelName: cfg.MQ0Channel,
        QueueName:   cfg.MQ0PutQueue,
        AppName:     cfg.MQ0App,
        User:        cfg.MQ0User,
        Pass:        cfg.MQ0Pass,
        Priority:    cfg.MQ0Priority,
      })
      connLi = append(connLi, conn)
    }

    if cfg.MQ0GetQueue != "" {
      conn := MqconnNew(TypeGet, Log, &Cfg{
        Host:        cfg.MQ0Host,
        Port:        cfg.MQ0Port,
        MgrName:     cfg.MQ0Mgr,
        ChannelName: cfg.MQ0Channel,
        QueueName:   cfg.MQ0GetQueue,
        AppName:     cfg.MQ0App,
        User:        cfg.MQ0User,
        Pass:        cfg.MQ0Pass,
        Priority:    cfg.MQ0Priority,
      })
      connLi = append(connLi, conn)
    }

    if cfg.MQ0BrowseQ != "" {
      conn := MqconnNew(TypeBrowse, Log, &Cfg{
        Host:        cfg.MQ0Host,
        Port:        cfg.MQ0Port,
        MgrName:     cfg.MQ0Mgr,
        ChannelName: cfg.MQ0Channel,
        QueueName:   cfg.MQ0BrowseQ,
        AppName:     cfg.MQ0App,
        User:        cfg.MQ0User,
        Pass:        cfg.MQ0Pass,
        Priority:    cfg.MQ0Priority,
      })
      connLi = append(connLi, conn)
    }
  }

  if cfg.MQ1Host != "" {
    if cfg.MQ1PutQueue != "" {
      conn := MqconnNew(TypePut, Log, &Cfg{
        Host:        cfg.MQ1Host,
        Port:        cfg.MQ1Port,
        MgrName:     cfg.MQ1Mgr,
        ChannelName: cfg.MQ1Channel,
        QueueName:   cfg.MQ1PutQueue,
        AppName:     cfg.MQ1App,
        User:        cfg.MQ1User,
        Pass:        cfg.MQ1Pass,
        Priority:    cfg.MQ1Priority,
      })
      connLi = append(connLi, conn)
    }

    if cfg.MQ1GetQueue != "" {
      conn := MqconnNew(TypeGet, Log, &Cfg{
        Host:        cfg.MQ1Host,
        Port:        cfg.MQ1Port,
        MgrName:     cfg.MQ1Mgr,
        ChannelName: cfg.MQ1Channel,
        QueueName:   cfg.MQ1GetQueue,
        AppName:     cfg.MQ1App,
        User:        cfg.MQ1User,
        Pass:        cfg.MQ1Pass,
        Priority:    cfg.MQ1Priority,
      })
      connLi = append(connLi, conn)
    }

    if cfg.MQ1BrowseQ != "" {
      conn := MqconnNew(TypeBrowse, Log, &Cfg{
        Host:        cfg.MQ1Host,
        Port:        cfg.MQ1Port,
        MgrName:     cfg.MQ1Mgr,
        ChannelName: cfg.MQ1Channel,
        QueueName:   cfg.MQ1BrowseQ,
        AppName:     cfg.MQ1App,
        User:        cfg.MQ1User,
        Pass:        cfg.MQ1Pass,
        Priority:    cfg.MQ1Priority,
      })
      connLi = append(connLi, conn)
    }
  }

  return connLi
}