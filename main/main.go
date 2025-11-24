package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
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
	//Читаем API ключ из переменной окружения
	apikey := os.Getenv("OPENWEATHER_API_KEY")
	if apikey == "" {
		fmt.Println("Ошибка: не найден API ключ. Уставновите переменную окружения OpenWeather_API_KEY")
		fmt.Println("Пример (Windows PowerShell):  $env:OPENWEATHER_API_KEY=\"ТВОЙ_КЛЮЧ\"")
		fmt.Println("Пример (Linux/macOS bash):   export OPENWEATHER_API_KEY=\"ТВОЙ_КЛЮЧ\"")
		return
	}
	// Спрашиваем у пользователя город
	var city string
	fmt.Println("Введите ваш город")
	fmt.Scan(&city)

	if city == "" {
		fmt.Println("Город не найден")
		return
	}

	//Формируем URL-запрос
	endpoint := "https://api.openweathermap.org/data/2.5/weather"

	params := url.Values{}
	params.Set("q", city)
	params.Set("appid", apikey)
	params.Set("units", "metric")
	params.Set("lang", "ru")

	reqURL := endpoint + "?" + params.Encode()

	//Делаем HTTP-запрос
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(reqURL)
	if err != nil {
		fmt.Println("Ошибка при запросе: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("API вернул ошибку. HTTP статус: ", resp.Status)
		return
	}

	//Парсим JSON
	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Println("Ошибка при разборе ответа:\", err")
		return
	}
	if weather.Name == "" {
		fmt.Println("Не удалось получить данные по городу. Проверь название.")
		return
	}

	//Выводим красивый текстовый прогноз
	fmt.Println("======================================")
	fmt.Printf("Погода в городе: %s\n", weather.Name)
	if len(weather.Weather) > 0 {
		fmt.Printf("Состояние:      %s (%s)\n", weather.Weather[0].Main, weather.Weather[0].Description)
	}
	fmt.Printf("Температура:    %.1f°C (ощущается как %.1f°C)\n", weather.Main.Temp, weather.Main.FeelsLike)
	fmt.Printf("Влажность:      %d%%\n", weather.Main.Humidity)
	fmt.Printf("Скорость ветра: %.1f м/с\n", weather.Wind.Speed)
	fmt.Println("======================================")
}
