package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WriteJSON(req *http.Request, v any) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	return nil
}

func ReadJSON(resp *http.Response, v any) error {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
