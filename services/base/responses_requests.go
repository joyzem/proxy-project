package base

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/levigross/grequests"
)

// Функция десериализирует тело запроса и возвращает ошибку, если тело
// запроса не соответствует структуре
func DecodeBody(r *http.Request, dest interface{}) (interface{}, error) {
	// Ожидаемые поля
	expected := map[string]struct{}{}
	elem := reflect.ValueOf(dest).Elem()
	for i := 0; i < elem.NumField(); i++ {
		expected[string(elem.Type().Field(i).Tag.Get("json"))] = struct{}{}
	}

	// Создание копии тела запроса, так как r.Body может вызываться только один раз
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// Декодировать тело запроса, чтобы проверить корректность
	var jsonKeys map[string]interface{}
	if err := json.Unmarshal(body, &jsonKeys); err != nil {
		return nil, err
	}

	// Проверить все ожидаемые поля
	for key := range expected {
		if _, ok := jsonKeys[key]; !ok {
			return nil, fmt.Errorf("missing field: %s", key)
		}
	}

	// Проверить лишние поля
	for key := range jsonKeys {
		if _, ok := expected[key]; !ok {
			return nil, fmt.Errorf("additional field: %s", key)
		}
	}

	// Decode the request body into the v variable.
	r.Body.Close()
	err = json.Unmarshal(body, &dest)
	return dest, err
}

// Закодировать ответ, содержащий ошибку
func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func CreateJsonRequestOption(body interface{}) *grequests.RequestOptions {
	return &grequests.RequestOptions{
		JSON: body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
