+++
date = "2017-07-26T16:57:54+09:00"
title = "slide17_2"

+++
#### JSONで書くと
```json
{
  "bastion-dev": {
    "host":"192.168.100.1",
    "port":"22",
    "user":"bastion-user",
    "cert_path":"/Users/minori.tokuda/.ssh/bastion-dev.pem",
    "cascades": {
      "app-server": {
        "host":"192.168.100.2",
        "port":"22",
        "user":"appuser",
        "cert_path":"/Users/minori.tokuda/.ssh/bastion-dev.pem",
        "tunnels": {
          "postgres": {
            "local_port":"15432",
            "remote_host":"192.168.101.10",
            "remote_port":"5432"
          }
        }
      }
    }
}
```
※ 設定ファイルフォーマットはTOMLのみです
