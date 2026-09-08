package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alibiamanzhol/open-kazakhstan/internal/model"
)

var fixture = []model.Dataset{
	{ID: "demo-001", Title: "Road traffic intensity", Organization: "Transport authority", Category: "Transport", Format: "CSV", Updated: "2026-08-30", Records: 12480},
	{ID: "demo-002", Title: "Public schools registry", Organization: "Education authority", Category: "Education", Format: "JSON", Updated: "2026-09-01", Records: 7812},
	{ID: "demo-003", Title: "Air monitoring stations", Organization: "Environmental authority", Category: "Environment", Format: "JSON", Updated: "2026-09-05", Records: 96},
	{ID: "demo-004", Title: "Regional budget indicators", Organization: "Finance authority", Category: "Finance", Format: "XLSX", Updated: "2026-08-28", Records: 440},
}

type envelope struct {
	Datasets []model.Dataset `json:"datasets"`
}

func Load() ([]model.Dataset, error) {
	url := strings.TrimSpace(os.Getenv("KZDATA_URL"))
	if url == "" {
		return fixture, nil
	}
	hc := &http.Client{Timeout: 10 * time.Second}
	resp, err := hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("upstream returned %s", resp.Status)
	}
	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}
	var list []model.Dataset
	if err := json.Unmarshal(raw, &list); err == nil {
		return list, nil
	}
	var env envelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, err
	}
	return env.Datasets, nil
}
