+++
date = "2017-07-26T16:06:39+09:00"
title = "slide06"

+++

#### 実装

```go
func tomlToJson(data []byte) ([]byte, error) {
	tree, err := toml.LoadBytes(data)
	if err != nil {
		return []byte{}, err
	}
	json, err := json.Marshal(tree.ToMap())
	if err != nil {
		return []byte{}, err
	}
	return json, nil
}
```
