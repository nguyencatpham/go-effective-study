package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	// "reflect"
)

// TokenType type of token
type TokenType int8

// const of token type
const (
	Basic = iota
	Bearer
)

func (s TokenType) String() string {
	switch s {
	case Basic:
		return "Basic"
	case Bearer:
	default:
		return "Bearer"
	}
	return "Bearer"
}

// RequestModel model
type RequestModel struct {
	URL       string
	TokenType TokenType
	Token     string
	Username  string
	Password  string
	Body      string
}

// Post request
func Post(requestModel RequestModel, result *interface{}) error {

	req, err := http.NewRequest("POST", requestModel.URL, bytes.NewBufferString(requestModel.Body))
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("request err:%v\n", err)
		return err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		message := "Post error"
		if result != nil {
			fmt.Printf("resp.result:%v\n", result)
			temp, ok := (*result).(map[string]interface{})
			if ok && temp != nil && temp["data"] != nil {
				data, ok := (temp["data"]).(map[string]interface{})
				if ok && data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}

					}
				}

			}

		}

		return GetError(resp.StatusCode, message)
	}

	return nil
}

// Put request
func Put(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("PUT", requestModel.URL, bytes.NewBufferString(requestModel.Body))
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	// fmt.Printf("err:%v\n", err)
	// fmt.Printf("resp.StatusCode:%v\n", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return echo.NewHTTPError(http.StatusBadRequest, "Put error")
	}

	return nil
}

// Get request
func Get(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("GET", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		log.Println("JSON Decode error", err)
		// TODO: parse wrong json
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := "Post error"
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp := (*result).(map[string]interface{})
			if temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}

					}
				}

			}

		}
		return GetError(resp.StatusCode, message)
	}

	return nil
}

// Delete request
func Delete(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("DELETE", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		log.Println("JSON Decode error", err)
		// TODO: parse wrong json
		return nil
	}
	if resp.StatusCode == http.StatusNotFound {
		log.Println("not found ", requestModel)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		log.Println("request code ", resp.StatusCode)
		log.Println("request result ", result)
		log.Println("model request", requestModel)
		message := "Delete error"
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp := (*result).(map[string]interface{})
			if temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}

					}
				}

			}

		}
		return GetError(resp.StatusCode, message)
	}

	return nil
}
