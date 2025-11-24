# Console-weather-client-on-Golang
Консольный прогноз погоды на Golang

Настрой API-ключ

Windows (PowerShell)
$env:OPENWEATHER_API_KEY="ТВОЙ_КЛЮЧ"
go run main.go

Linux / macOS (bash/zsh)
export OPENWEATHER_API_KEY="ТВОЙ_КЛЮЧ"
go run main.go


После запуска:

go run main.go


И в консоли:

Введите город: Moscow
======================================
Погода в городе: Moscow
Состояние:      Clouds (переменная облачность)
Температура:    -3.2°C (ощущается как -7.1°C)
Влажность:      80%
Скорость ветра: 3.5 м/с
======================================
