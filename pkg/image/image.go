package image

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const searchImagePath = "https://api.iconify.design/search?query="
const searchImageFillPaths = "https://api.iconify.design/search?query=fill%20"
const iconify = "https://api.iconify.design/"

type IconResponse struct {
	Icons []string `json:"icons"`
	Total int      `json:"total"`
}

func SearchImage(query string) ([]string, error) {
	images := []string{}

	queries := []string{}

	splitQuery := strings.Split(query, "_")
	for _, s := range splitQuery {
		queries = append(queries, searchImageFillPaths+s)
	}
	for _, s := range splitQuery {
		queries = append(queries, searchImagePath+s)
	}
	queries = append(queries, searchImagePath+query)

	for _, q := range queries {
		req, err := http.Get(q)
		if err != nil {
			fmt.Errorf("Error getting image: %s", err)
			continue
		}

		var iconResponse IconResponse
		if err := json.NewDecoder(req.Body).Decode(&iconResponse); err != nil {
			fmt.Errorf("Error decoding image: %s", err)
			req.Body.Close()
			continue
		}

		if iconResponse.Total == 0 {
			req.Body.Close()
			continue
		}

		for _, icon := range iconResponse.Icons {
			s := strings.Split(icon, ":")
			imageUrl := iconify + s[0] + "/" + s[1] + ".svg"
			images = append(images, imageUrl)
		}

		req.Body.Close()

		if len(images) > 0 {
			break
		}
	}

	return images, nil
}

func GetFirstImage(query string) (string, error) {
	images, err := SearchImage(query)
	if err != nil {
		return "", err
	}
	if len(images) == 0 {
		return "", fmt.Errorf("No images found")
	}
	return images[0], nil
}
