テーブルの配列という記法もある
```toml
[[array_of_table]]
id = 10
[[array_of_table]]
id = 20

# => {"array_of_table": [{"id":10},{"id":20}]}
```