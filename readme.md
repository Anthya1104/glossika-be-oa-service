## How to use?

1. Required version:

```
go: 1.24.1
mysql: 8.0
redis: 7.2

```

and make sure your docker environment and go IDE tools have been settled down.

2. clone the repo and install the dependencies

`go mod tidy`

## Run Service

1. Run docker-compose to start required DB/redis containers

```
docker-compose up -d
```

2. Run and Debug -> RUN to start server, or you can use the following command line:

```
go run cmd/main.go
```

3. the db schema should be auto-migrated, if it's not work, try the SQL commands inside `./sql`, and make sure the `user` table and `product` table are migrated bofere `user_recommendation` table

4. insert the mock data by running commands in `./sql/mock_data.sql`

5. the server would be run on port `:9999`, while the mysql db and redis would be on `:3306` and `:6379` correspondingly. If there is any conflict on your machine, please modify the `./docker-compose.yaml`, `./configs/config.yaml` files

## Run Test

1.  Run and Debug -> TEST to start test

## API docs

1. you can find the API docs in `./doc`
2. or you can import the postman collection by using `./doc/Glossika_BE_OA_service.postman_collection.json`

## About Test

- 目前只實作了 User Registration API 的 test suite ，包含一個 positive case 和兩個 negative case
- 其他 TODO 的 test case 可以在 `.doc/TODO_test_cases.md` 內找到

## TODO List

- 目前為簡化設計用於 demo ， user 只用 email 作為 unique key ，並且不存在其他詳細 info，實際實作或許可以把 id 改為用 uuid, 並擴充 username 等等欄位
- 專案內有些 naming 以及 redundant methods/models 還沒整理過，有時間的話可以再整理一次
- 這次比較無法仔細考慮 log ，應該還有些地方能補，且相關 pattern 可以再清楚明確一點 e.g. 加上當前 function name 或是掛好 traceID
- API doc 有需要擴充的話可以掛 swagger
- 目前 DB schema 基本上沒有考慮軟刪除的情況，如果業務上有需要可能可以再做一些調整
- 目前 user activation 的部分是採寄送 activation link 並讓 user 點擊驗證的流程，因為不考慮 FE 所以直接回了 json response ，如果不想讓這些資料暴露給 user ，可以在這裡做一個 redirect 回 FE page 的機制並讓 FE 做 handling
- user activation 目前沒有考慮到註冊過程中寄送失敗，需要重新寄信的 case ，如果需要實作，可能可以另開 API ，確認 user 存在並重新 gen auth token 寄出
- user login 目前會擋還沒做完 activation 的 user ，如果沒有要擋這麼嚴，想讓 user 還是能登入，只是可能被限制使用部分功能的話，可能可以在 gen auth token 的時候再做一些調整，並調整 login API 的相關 error handling
- doc: API doc 目前只寫了 200 OK response，時間充足的情況下會想再補上幾個常見的 bad response 範例
- 時間上考量再加上這次 DEMO 的功能較簡單，所以 auth 是直接做在 server 內部，如果未來可能擴充多平台或支援不同產品，或許可以考慮把 auth 相關 feature 另開 account server ，搭配 API Gateway 改用 OAuth2.0 架構
- Test 目前是直接 run 一個 DB container 來做 feature test ，但其實 sql DB 的部分有把 db repo 抽成抽象 interface ，未來可以替換用 gomock+mockgen 直接 mock DB response
- Known issues :
  1. get recommendation API paging slice 數量會變成 1.5 倍 (e.g. page_size 指定 50 會一次給出 75 筆資料) ，問題應該是發生在最後 slice data 時，還需要再做 debug
  2. test container 在跑起來的時候會有一段時間 server 無法連入 DB ，推測可能是 run test 時 sql container 還沒完全起來，這裡需要再思考怎麼改寫法排除
- Known improvement
  1. redis 取得 cache 還沒有做 error handling ，雖然拿不到就會回 DB 撈，但考量到題目情境下，這個 DB operation 蠻耗時的，這裡如果能補上相關處置(如 logging, notification 機制)可能會更好
