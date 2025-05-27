package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/router"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"gorm.io/gorm/logger"
)

func HttpGet(path string, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	resp, err = sendHttp("GET", path, "", headers)
	return
}

func HttpPost(path string, body string, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	resp, err = sendHttp("POST", path, body, headers)
	return
}

func HttpPostAndMarshalBody(path string, body interface{}, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	bodyInByte, err := json.Marshal(body)
	if err != nil {
		return
	}

	resp, err = sendHttp("POST", path, string(bodyInByte), headers)
	return
}

func HttpPut(path string, body string, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	resp, err = sendHttp("PUT", path, body, headers)
	return
}

func HttpPatch(path string, body string, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	resp, err = sendHttp("PATCH", path, body, headers)
	return
}

func HttpDelete(path string, body string, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	resp, err = sendHttp("DELETE", path, body, headers)
	return
}

func sendHttp(method string, path string, body string, headers map[string]string) (resp *httptest.ResponseRecorder, err error) {
	resp = httptest.NewRecorder()

	req, err := http.NewRequest(method, path, bytes.NewBufferString(body))
	if err != nil {
		return
	}

	setHeader(req, headers)
	router.Router.ServeHTTP(resp, req)

	return
}

func setHeader(req *http.Request, headers map[string]string) {
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func getRespVersion(w *httptest.ResponseRecorder) string {
	resp := model.CommonErrorRes{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	return resp.Version
}

func getRespCustomErrorcode(w *httptest.ResponseRecorder) string {
	resp := model.CommonErrorRes{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	return resp.Error
}

func getRespCustomErrorMsg(w *httptest.ResponseRecorder) string {
	resp := model.CommonErrorRes{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	return resp.Msg
}

func resetDbForTesting() {
	database.GetSqlDb().Orm.Logger = logger.Default.LogMode(logger.Warn)

	if err := database.DropTables(database.GetSqlDb().Orm); err != nil {
		log.L.Fatal(err)
	}

	if err := database.AutoMigrate(database.GetSqlDb().Orm); err != nil {
		log.L.Fatal(err)
	}

	log.L.Info("----- DB DATA HAS BEEN RESET -----")
}
