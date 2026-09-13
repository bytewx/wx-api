package models

type Response struct {
	Latitude             float32 `json:"latitude"`
	Longitude            float32 `json:"longitude"`
	GenerationTimeMS     float32 `json:"generation_time_ms"`
	UTCOffsetSeconds     uint    `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float32 `json:"elevation"`

	CurrentUnits   CurrentUnitsResponse   `json:"current_units"`
	CurrentWeather CurrentWeatherResponse `json:"current"`

	HourlyUnits   HourlyUnitsResponse   `json:"hourly_units"`
	HourlyWeather HourlyWeatherResponse `json:"hourly"`

	Minutely15Units   Minutely15UnitsResponse   `json:"minutely_15_units"`
	Minutely15Weather Minutely15WeatherResponse `json:"minutely_15"`

	DailyUnits   DailyUnitsResponse   `json:"daily_units"`
	DailyWeather DailyWeatherResponse `json:"daily"`

	PressureLevelUnits   PressureLevelUnitsResponse   `json:"pressure_level_units"`
	PressureLevelWeather PressureLevelWeatherResponse `json:"pressure_level"`
}

type CurrentUnitsResponse struct {
	Time                *string `json:"time,omitempty"`
	Interval            *string `json:"seconds,omitempty"`
	Temperature2M       *string `json:"temperature_2m,omitempty"`
	RelativeHumidity2M  *string `json:"relative_humidity_2m,omitempty"`
	IsDay               *string `json:"is_day,omitempty"`
	ApparentTemperature *string `json:"apparent_temperature,omitempty"`
	WindSpeed10M        *string `json:"wind_speed_10m,omitempty"`
	WindDirection10M    *string `json:"wind_direction_10m,omitempty"`
	WindGusts10M        *string `json:"wind_gusts_10m,omitempty"`
	Precipitation       *string `json:"precipitation,omitempty"`
	Showers             *string `json:"showers,omitempty"`
	Rain                *string `json:"rain,omitempty"`
	Snowfall            *string `json:"snowfall,omitempty"`
	WeatherCode         *string `json:"weather_code,omitempty"`
	CloudCover          *string `json:"cloud_cover,omitempty"`
	PressureMSL         *string `json:"pressure_msl,omitempty"`
	SurfacePressure     *string `json:"surface_pressure,omitempty"`
}

type CurrentWeatherResponse struct {
	Time                *string  `json:"time,omitempty"`
	Interval            *uint    `json:"interval,omitempty"`
	Temperature2M       *float32 `json:"temperature_2m,omitempty"`
	RelativeHumidity2M  *float32 `json:"relative_humidity_2m,omitempty"`
	IsDay               *uint    `json:"is_day,omitempty"`
	ApparentTemperature *float32 `json:"apparent_temperature,omitempty"`
	WindSpeed10M        *float32 `json:"wind_speed_10m,omitempty"`
	WindDirection10M    *float32 `json:"wind_direction_10m,omitempty"`
	WindGusts10M        *float32 `json:"wind_gusts_10m,omitempty"`
	Precipitation       *float32 `json:"precipitation,omitempty"`
	Showers             *float32 `json:"showers,omitempty"`
	Rain                *float32 `json:"rain,omitempty"`
	Snowfall            *float32 `json:"snowfall,omitempty"`
	WeatherCode         *float32 `json:"weather_code,omitempty"`
	CloudCover          *float32 `json:"cloud_cover,omitempty"`
	PressureMSL         *float32 `json:"pressure_msl,omitempty"`
	SurfacePressure     *float32 `json:"surface_pressure,omitempty"`
}
