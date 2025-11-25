package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/smartystreets/go-aws-auth"
	"github.com/smartystreets/go-aws-loggly"
)

func ValidateRequest(req *http.Request) error {
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		return fmt.Errorf("unsupported request method: %s", req.Method)
	}
	return nil
}

func GetStatsGrouping(req *http.Request) (string, error) {
	grouping := req.URL.Query().Get("grouping")
	if grouping == "" {
		return "", fmt.Errorf("missing grouping query parameter")
	}
	return grouping, nil
}

func GetStatsAggregation(req *http.Request) (string, error) {
	aggregation := req.URL.Query().Get("aggregation")
	if aggregation == "" {
		return "", fmt.Errorf("missing aggregation query parameter")
	}
	return aggregation, nil
}

func GetStatsFields(req *http.Request) ([]string, error) {
	fields := req.URL.Query().Get("fields")
	if fields == "" {
		return nil, fmt.Errorf("missing fields query parameter")
	}
	var fieldList []string
	for _, field := range strings.Split(fields, ",") {
		fieldList = append(fieldList, strings.TrimSpace(field))
	}
	if len(fieldList) == 0 {
		return nil, fmt.Errorf("fields query parameter is empty")
	}
	return fieldList, nil
}

func GetStatsPeriod(req *http.Request) (string, error) {
	period := req.URL.Query().Get("period")
	if period == "" {
		return "", fmt.Errorf("missing period query parameter")
	}
	return period, nil
}

func GetStatsFilter(req *http.Request) (map[string]interface{}, error) {
	filter := req.URL.Query().Get("filter")
	if filter == "" {
		return nil, fmt.Errorf("missing filter query parameter")
	}
	var filterMap map[string]interface{}
	err := json.Unmarshal([]byte(filter), &filterMap)
	if err != nil {
		return nil, fmt.Errorf("invalid filter query parameter: %s", err)
	}
	return filterMap, nil
}

func GetStatsOrderBy(req *http.Request) (string, error) {
	order := req.URL.Query().Get("order")
	if order == "" {
		return "", fmt.Errorf("missing order query parameter")
	}
	return order, nil
}

func GetStatsLimit(req *http.Request) (int64, error) {
	limit := req.URL.Query().Get("limit")
	if limit == "" {
		return 0, fmt.Errorf("missing limit query parameter")
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid limit query parameter: %s", err)
	}
	return limitInt, nil
}

func GetAwsCredentials() (*aws.Credentials, error) {
	creds, err := aws.GetCredentials()
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func GetLogglyConfig() (*loggly.Config, error) {
	creds, err := GetAwsCredentials()
	if err != nil {
		return nil, err
	}
	return &loggly.Config{
		Region:      creds.Region,
		AccessKey:   creds.AccessKey,
		SecretKey:   creds.SecretKey,
		SessionToken: creds.SessionToken,
	}, nil
}