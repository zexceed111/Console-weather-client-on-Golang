package main

import (
	"fmt"
	"os"
)

// Структура для ответа OpenWeather
type WeatherResponse struct {
	Name string `json:"name"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"temp_feels"`
		Humidity  int     `json:"temp_humidity"`
	} `json:"main"`

	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func main() {
	// Читаем API ключ из переменной окружения
	apikey := os.Getenv("a3036b726f510b14504c8901ac666e6e")
	if apikey == "" {
		fmt.Println("Ошибка: не найден API ключ. Уставновите переменную окружения OpenWeather_API_KEY")
		fmt.Println("Пример (Windows PowerShell):  $env:OPENWEATHER_API_KEY=\"ТВОЙ_КЛЮЧ\"")
		fmt.Println("Пример (Linux/macOS bash):   export OPENWEATHER_API_KEY=\"ТВОЙ_КЛЮЧ\"")
		return
	}
	// Спрашиваем у пользователя город

}
