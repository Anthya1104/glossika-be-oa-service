## when you firstly clone this repo

`go mod tidy`

## Run Service

## Run Test

## TODO

- 目前 user 只用 email 作為 unique key ，並且不存在 user 的其他詳細 info，實際實作或許可以把 id 改為用 uuid, 並擴充 username 等等欄位
- 專案內有些 naming 的 consistency 有時間的話可以再整理過
- 應該還能再整理插一些 log ，有些 log 的 pattern 可以再清楚明確一點 e.g.在前面加上 functio name 或插 traceID
- API doc 有需要擴充的話可以掛 swagger
- 目前 DB schema 基本上沒有考慮軟刪除的情況，如果業務上有需要可能可以載做一些調整
