package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math"
	"strings"
)

type PaginationParam struct {
	Count int64       `json:"count"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  interface{} `json:"data"`
}

type PaginationResult struct {
	TotalPage    int         `json:"totalPage"`
	TotalData    int64       `json:"totalData"`
	NextPage     *int        `json:"nextPage"`
	PreviousPage *int        `json:"previousPage"`
	Page         int         `json:"page"`
	Limit        int         `json:"limit"`
	Data         interface{} `json:"data"`
}

func GeneratePagination(params PaginationParam) PaginationResult {
	totalPage := int(math.Ceil(float64(params.Count) / float64(params.Limit)))

	var (
		nextPage     int
		previousPage int
	)
	if params.Page < totalPage {
		nextPage = params.Page + 1
	}

	if params.Page > 1 {
		previousPage = params.Page - 1
	}

	result := PaginationResult{
		TotalPage:    totalPage,
		TotalData:    params.Count,
		NextPage:     &nextPage,
		PreviousPage: &previousPage,
		Page:         params.Page,
		Limit:        params.Limit,
		Data:         params.Data,
	}
	return result
}

func GenerateSHA256(inputString string) string {
	hash := sha256.New()
	hash.Write([]byte(inputString))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func RupiahFormat(amount *float64) string {
	stringValue := "0"
	if amount != nil {
		humanizeValue := humanize.CommafWithDigits(*amount, 0)
		stringValue = strings.ReplaceAll(humanizeValue, ",", ".")
	}
	return fmt.Sprintf("Rp. %s", stringValue)
}

// membaca file config.json
func BindFromJSON(dest any, filename, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(filename)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("failed to unmarshal: %v", err)
		return err
	}

	return nil
}

// Membaca konfigurasi dari .env
func ReadFromEnv() map[string]interface{} {
	v := viper.New()
	v.AutomaticEnv() // Aktifkan pembacaan dari environment variables

	configMap := map[string]interface{}{
		"port":                  v.GetInt("PORT"),
		"appName":               v.GetString("APP_NAME"),
		"appEnv":                v.GetString("APP_ENV"),
		"signatureKey":          v.GetString("SIGNATURE_KEY"),
		"rateLimiterMaxRequest": v.GetInt("RATE_LIMITER_MAX_REQUEST"),
		"rateLimiterTimeSecond": v.GetInt("RATE_LIMITER_TIME_SECOND"),

		// Database
		"database.host":                  v.GetString("DB_HOST"),
		"database.port":                  v.GetInt("DB_PORT"),
		"database.name":                  v.GetString("DB_NAME"),
		"database.username":              v.GetString("DB_USERNAME"),
		"database.password":              v.GetString("DB_PASSWORD"),
		"database.maxOpenConnections":    v.GetInt("DB_MAX_OPEN_CONNECTIONS"),
		"database.maxLifeTimeConnection": v.GetInt("DB_MAX_LIFETIME_CONNECTION"),
		"database.maxIdleConnections":    v.GetInt("DB_MAX_IDLE_CONNECTIONS"),
		"database.maxIdleTime":           v.GetInt("DB_MAX_IDLE_TIME"),

		// Clients
		"internalService.user.host":         v.GetString("INTERNAL_SERVICE_USER_HOST"),
		"internalService.user.signatureKey": v.GetString("INTERNAL_SERVICE_USER_SIGNATURE_KEY"),

		// Minio
		"minio.address":    v.GetString("MINIO_ADDRESS"),
		"minio.accessKey":  v.GetString("MINIO_ACCESS_KEY"),
		"minio.secret":     v.GetString("MINIO_SECRET"),
		"minio.useSsl":     v.GetBool("MINIO_USE_SSL"),
		"minio.bucketName": v.GetString("MINIO_BUCKET_NAME"),
	}

	return configMap
}

func BindFromEnv(dest any) error {
	v := viper.New()
	v.AutomaticEnv()

	for key, value := range ReadFromEnv() {
		v.Set(key, value)
	}

	err := v.Unmarshal(dest)
	if err != nil {
		return err
	}

	return nil
}
