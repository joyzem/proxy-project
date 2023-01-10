package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Функция десериализирует тело запроса и возвращает ошибку, если тело
// запроса не соответствует структуре
func DecodeBody(r *http.Request, dest interface{}) error {
	// Ожидаемые поля
	expected := map[string]struct{}{}
	elem := reflect.ValueOf(dest).Elem()
	for i := 0; i < elem.NumField(); i++ {
		expected[string(elem.Type().Field(i).Tag.Get("json"))] = struct{}{}
	}

	// Создание копии тела запроса, так как r.Body может вызываться только один раз
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// Декодировать тело запроса, чтобы проверить корректность
	var jsonKeys map[string]interface{}
	if err := json.Unmarshal(body, &jsonKeys); err != nil {
		return err
	}

	// Проверить все ожидаемые поля
	for key := range expected {
		if _, ok := jsonKeys[key]; !ok {
			return fmt.Errorf("missing field: %s", key)
		}
	}

	// Проверить лишние поля
	for key := range jsonKeys {
		if _, ok := expected[key]; !ok {
			return fmt.Errorf("additional field: %s", key)
		}
	}

	// Decode the request body into the v variable.
	r.Body.Close()
	return json.Unmarshal(body, &dest)
}

type BaseResponse struct {
	Response interface{} `json:"data"`
	Success  bool        `json:"success"`
}

// Закодировать ответ, содержащий ошибку
func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	return http.StatusInternalServerError
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Ошибка бизнес-логики
		EncodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var baseResponse = BaseResponse{
		Response: response,
		Success:  true,
	}
	return json.NewEncoder(w).Encode(baseResponse)
}

// Для проверки, является ли объект ошибкой
type errorer interface {
	error() error
}
