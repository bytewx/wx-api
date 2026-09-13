package config

type WeatherType string

const (
	Current WeatherType = "current"
	Daily   WeatherType = "daily"
	Hourly  WeatherType = "hourly"
)

type WeatherTypeConfig struct {
	ParamName     string
	DefaultFields string
}

var WeatherTypes = map[WeatherType]WeatherTypeConfig{
	"current": {
		ParamName: "current",
		DefaultFields: "temperature_2m,relative_humidity_2m,is_day," +
			"apparent_temperature,wind_speed_10m,wind_direction_10m,wind_gusts_10m," +
			"precipitation,showers,rain,snowfall,weather_code,cloud_cover,pressure_msl,surface_pressure",
	},
	"daily": {
		ParamName: "daily",
		DefaultFields: "weather_code,temperature_2m_max,temperature_2m_min,apparent_temperature_max," +
			"apparent_temperature_min,uv_index_max,uv_index_clear_sky_max,wind_speed_10m_max," +
			"wind_direction_10m_dominant,wind_gusts_10m_max,shortwave_radiation_sum," +
			"et0_fao_evapotranspiration,sunrise,sunset,daylight_duration,sunshine_duration," +
			"moonrise,moon_phase,moonset,rain_sum,showers_sum,snowfall_sum,precipitation_sum," +
			"precipitation_hours,precipitation_probability_max,temperature_2m_mean,apparent_temperature_mean," +
			"cape_mean,cape_max,cloud_cover_mean,cape_min,cloud_cover_max,cloud_cover_min,dew_point_2m_mean," +
			"dew_point_2m_max,dew_point_2m_min,wet_bulb_temperature_2m_mean,wet_bulb_temperature_2m_min," +
			"wet_bulb_temperature_2m_max,vapour_pressure_deficit_max,et0_fao_evapotranspiration_sum," +
			"growing_degree_days_base_0_limit_50,leaf_wetness_probability_mean,precipitation_probability_mean," +
			"precipitation_probability_min,relative_humidity_2m_mean,relative_humidity_2m_max," +
			"relative_humidity_2m_min,snowfall_water_equivalent_sum,pressure_msl_mean,pressure_msl_max," +
			"pressure_msl_min,surface_pressure_mean,surface_pressure_max,surface_pressure_min,updraft_max," +
			"visibility_mean,visibility_min,visibility_max,wind_gusts_10m_mean,wind_speed_10m_mean," +
			"wind_speed_10m_min,wind_gusts_10m_min",
	},
	"hourly": {
		ParamName: "hourly",
		DefaultFields: "temperature_2m,relative_humidity_2m,dew_point_2m,apparent_temperature," +
			"precipitation_probability,precipitation,rain,showers,snowfall,weather_code," +
			"pressure_msl,surface_pressure,cloud_cover,visibility,wind_speed_10m," +
			"wind_direction_10m,wind_gusts_10m,uv_index,is_day",
	},
}
