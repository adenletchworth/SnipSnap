package embed

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func GetEmbedding(text string) ([]float32, error) {
	payload := map[string]string{"text": text}
	body, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		"http://localhost:8000/embed",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("embedding server non-200 status: " + resp.Status)
	}

	var result struct {
		Embedding []float32 `json:"embedding"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result.Embedding, nil
}
