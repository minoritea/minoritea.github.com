基本的に `key  = value` 形式で記述

```toml
id = 10
name = "snowcrush"
birthday = 1980-01-01T00:00:00Z # RFC3339
speed = 1.5e-3 # or 0.0015

#=> {"id":10,
#    "name":"snowcrush",
#    "birthday": "1980-01-01T00:00:00Z",
#    "speed": 0.0015}
```
