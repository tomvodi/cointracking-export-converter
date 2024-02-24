package config

import
"github.com/tomvodi/cointracking-export-converter/internal/common"

var AllTimezones = []common.TimezoneData{
	{
		Value:         "Pacific/Midway",
		Title:         "(GMT-11:00) Midway Island, Samoa",
		OffsetSeconds: -39600,
	},
	{
		Value:         "America/Adak",
		Title:         "(GMT-10:00) Hawaii-Aleutian",
		OffsetSeconds: -36000,
	},
	{
		Value:         "Etc/GMT+10",
		Title:         "(GMT-10:00) Hawaii",
		OffsetSeconds: -36000,
	},
	{
		Value:         "Pacific/Marquesas",
		Title:         "(GMT-09:30) Marquesas Islands",
		OffsetSeconds: -34200,
	},
	{
		Value:         "Pacific/Gambier",
		Title:         "(GMT-09:00) Gambier Islands",
		OffsetSeconds: -32400,
	},
	{
		Value:         "America/Anchorage",
		Title:         "(GMT-09:00) Alaska",
		OffsetSeconds: -32400,
	},
	{
		Value:         "America/Ensenada",
		Title:         "(GMT-08:00) Tijuana, Baja California",
		OffsetSeconds: -28800,
	},
	{
		Value:         "Etc/GMT+8",
		Title:         "(GMT-08:00) Pitcairn Islands",
		OffsetSeconds: -28800,
	},
	{
		Value:         "America/Los_Angeles",
		Title:         "(GMT-08:00) Pacific Time (US & Canada)",
		OffsetSeconds: -28800,
	},
	{
		Value:         "America/Denver",
		Title:         "(GMT-07:00) Mountain Time (US & Canada)",
		OffsetSeconds: -25200,
	},
	{
		Value:         "America/Chihuahua",
		Title:         "(GMT-07:00) Chihuahua, La Paz, Mazatlan",
		OffsetSeconds: -25200,
	},
	{
		Value:         "America/Dawson_Creek",
		Title:         "(GMT-07:00) Arizona",
		OffsetSeconds: -25200,
	},
	{
		Value:         "America/Belize",
		Title:         "(GMT-06:00) Saskatchewan, Central America",
		OffsetSeconds: -21600,
	},
	{
		Value:         "America/Cancun",
		Title:         "(GMT-06:00) Guadalajara, Mexico City, Monterrey",
		OffsetSeconds: -21600,
	},
	{
		Value:         "Chile/EasterIsland",
		Title:         "(GMT-06:00) Easter Island",
		OffsetSeconds: -21600,
	},
	{
		Value:         "America/Chicago",
		Title:         "(GMT-06:00) Central Time (US & Canada)",
		OffsetSeconds: -21600,
	},
	{
		Value:         "America/New_York",
		Title:         "(GMT-05:00) Eastern Time (US & Canada)",
		OffsetSeconds: -18000,
	},
	{
		Value:         "America/Havana",
		Title:         "(GMT-05:00) Cuba",
		OffsetSeconds: -18000,
	},
	{
		Value:         "America/Bogota",
		Title:         "(GMT-05:00) Bogota, Lima, Quito, Rio Branco",
		OffsetSeconds: -18000,
	},
	{
		Value:         "America/Caracas",
		Title:         "(GMT-04:30) Caracas",
		OffsetSeconds: -16200,
	},
	{
		Value:         "America/Santiago",
		Title:         "(GMT-04:00) Santiago",
		OffsetSeconds: -14400,
	},
	{
		Value:         "America/La_Paz",
		Title:         "(GMT-04:00) La Paz",
		OffsetSeconds: -14400,
	},
	{
		Value:         "Atlantic/Stanley",
		Title:         "(GMT-04:00) Faukland Islands",
		OffsetSeconds: -14400,
	},
	{
		Value:         "America/Campo_Grande",
		Title:         "(GMT-04:00) Brazil",
		OffsetSeconds: -14400,
	},
	{
		Value:         "America/Goose_Bay",
		Title:         "(GMT-04:00) Atlantic Time (Goose Bay)",
		OffsetSeconds: -14400,
	},
	{
		Value:         "America/Glace_Bay",
		Title:         "(GMT-04:00) Atlantic Time (Canada)",
		OffsetSeconds: -14400,
	},
	{
		Value:         "America/St_Johns",
		Title:         "(GMT-03:30) Newfoundland",
		OffsetSeconds: -12600,
	},
	{
		Value:         "America/Araguaina",
		Title:         "(GMT-03:00) UTC-3",
		OffsetSeconds: -10800,
	},
	{
		Value:         "America/Montevideo",
		Title:         "(GMT-03:00) Montevideo",
		OffsetSeconds: -10800,
	},
	{
		Value:         "America/Miquelon",
		Title:         "(GMT-03:00) Miquelon, St. Pierre",
		OffsetSeconds: -10800,
	},
	{
		Value:         "America/Godthab",
		Title:         "(GMT-03:00) Greenland",
		OffsetSeconds: -10800,
	},
	{
		Value:         "America/Argentina/Buenos_Aires",
		Title:         "(GMT-03:00) Buenos Aires",
		OffsetSeconds: -10800,
	},
	{
		Value:         "America/Sao_Paulo",
		Title:         "(GMT-03:00) Brasilia",
		OffsetSeconds: -10800,
	},
	{
		Value:         "America/Noronha",
		Title:         "(GMT-02:00) Mid-Atlantic",
		OffsetSeconds: -7200,
	},
	{
		Value:         "Atlantic/Cape_Verde",
		Title:         "(GMT-01:00) Cape Verde Is.",
		OffsetSeconds: -3600,
	},
	{
		Value:         "Atlantic/Azores",
		Title:         "(GMT-01:00) Azores",
		OffsetSeconds: -3600,
	},
	{
		Value:         "UTC",
		Title:         "(UTC) Coordinated Universal Time",
		OffsetSeconds: 0,
	},
	{
		Value:         "Europe/Belfast",
		Title:         "(GMT) Greenwich Mean Time : Belfast",
		OffsetSeconds: 0,
	},
	{
		Value:         "Europe/Dublin",
		Title:         "(GMT) Greenwich Mean Time : Dublin",
		OffsetSeconds: 0,
	},
	{
		Value:         "Europe/Lisbon",
		Title:         "(GMT) Greenwich Mean Time : Lisbon",
		OffsetSeconds: 0,
	},
	{
		Value:         "Europe/London",
		Title:         "(GMT) Greenwich Mean Time : London",
		OffsetSeconds: 0,
	},
	{
		Value:         "Africa/Abidjan",
		Title:         "(GMT) Monrovia, Reykjavik",
		OffsetSeconds: 0,
	},
	{
		Value:         "Europe/Amsterdam",
		Title:         "(GMT+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna",
		OffsetSeconds: 3600,
	},
	{
		Value:         "Europe/Belgrade",
		Title:         "(GMT+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague",
		OffsetSeconds: 3600,
	},
	{
		Value:         "Europe/Brussels",
		Title:         "(GMT+01:00) Brussels, Copenhagen, Madrid, Paris",
		OffsetSeconds: 3600,
	},
	{
		Value:         "Africa/Algiers",
		Title:         "(GMT+01:00) West Central Africa",
		OffsetSeconds: 3600,
	},
	{
		Value:         "Africa/Windhoek",
		Title:         "(GMT+01:00) Windhoek",
		OffsetSeconds: 3600,
	},
	{
		Value:         "Europe/Helsinki",
		Title:         "(GMT+02:00) Helsinki",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Asia/Beirut",
		Title:         "(GMT+02:00) Beirut",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Africa/Cairo",
		Title:         "(GMT+02:00) Cairo",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Asia/Gaza",
		Title:         "(GMT+02:00) Gaza",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Africa/Blantyre",
		Title:         "(GMT+02:00) Harare, Pretoria",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Asia/Jerusalem",
		Title:         "(GMT+02:00) Jerusalem",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Europe/Minsk",
		Title:         "(GMT+02:00) Minsk",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Asia/Damascus",
		Title:         "(GMT+02:00) Syria",
		OffsetSeconds: 7200,
	},
	{
		Value:         "Europe/Moscow",
		Title:         "(GMT+03:00) Moscow, St. Petersburg, Volgograd",
		OffsetSeconds: 10800,
	},
	{
		Value:         "Africa/Addis_Ababa",
		Title:         "(GMT+03:00) Nairobi",
		OffsetSeconds: 10800,
	},
	{
		Value:         "Asia/Tehran",
		Title:         "(GMT+03:30) Tehran",
		OffsetSeconds: 12600,
	},
	{
		Value:         "Asia/Dubai",
		Title:         "(GMT+04:00) Abu Dhabi, Muscat",
		OffsetSeconds: 14400,
	},
	{
		Value:         "Asia/Yerevan",
		Title:         "(GMT+04:00) Yerevan",
		OffsetSeconds: 14400,
	},
	{
		Value:         "Asia/Kabul",
		Title:         "(GMT+04:30) Kabul",
		OffsetSeconds: 16200,
	},
	{
		Value:         "Asia/Yekaterinburg",
		Title:         "(GMT+05:00) Ekaterinburg",
		OffsetSeconds: 18000,
	},
	{
		Value:         "Asia/Tashkent",
		Title:         "(GMT+05:00) Tashkent",
		OffsetSeconds: 18000,
	},
	{
		Value:         "Asia/Kolkata",
		Title:         "(GMT+05:30) Chennai, Kolkata, Mumbai, New Delhi",
		OffsetSeconds: 19800,
	},
	{
		Value:         "Asia/Katmandu",
		Title:         "(GMT+05:45) Kathmandu",
		OffsetSeconds: 20700,
	},
	{
		Value:         "Asia/Dhaka",
		Title:         "(GMT+06:00) Astana, Dhaka",
		OffsetSeconds: 21600,
	},
	{
		Value:         "Asia/Novosibirsk",
		Title:         "(GMT+06:00) Novosibirsk",
		OffsetSeconds: 21600,
	},
	{
		Value:         "Asia/Rangoon",
		Title:         "(GMT+06:30) Yangon (Rangoon)",
		OffsetSeconds: 23400,
	},
	{
		Value:         "Asia/Bangkok",
		Title:         "(GMT+07:00) Bangkok, Hanoi, Jakarta",
		OffsetSeconds: 25200,
	},
	{
		Value:         "Asia/Krasnoyarsk",
		Title:         "(GMT+07:00) Krasnoyarsk",
		OffsetSeconds: 25200,
	},
	{
		Value:         "Asia/Hong_Kong",
		Title:         "(GMT+08:00) Beijing, Chongqing, Hong Kong, Urumqi",
		OffsetSeconds: 28800,
	},
	{
		Value:         "Asia/Irkutsk",
		Title:         "(GMT+08:00) Irkutsk, Ulaan Bataar",
		OffsetSeconds: 28800,
	},
	{
		Value:         "Australia/Perth",
		Title:         "(GMT+08:00) Perth",
		OffsetSeconds: 28800,
	},
	{
		Value:         "Australia/Eucla",
		Title:         "(GMT+08:45) Eucla",
		OffsetSeconds: 31500,
	},
	{
		Value:         "Asia/Tokyo",
		Title:         "(GMT+09:00) Osaka, Sapporo, Tokyo",
		OffsetSeconds: 32400,
	},
	{
		Value:         "Asia/Seoul",
		Title:         "(GMT+09:00) Seoul",
		OffsetSeconds: 32400,
	},
	{
		Value:         "Asia/Yakutsk",
		Title:         "(GMT+09:00) Yakutsk",
		OffsetSeconds: 32400,
	},
	{
		Value:         "Australia/Adelaide",
		Title:         "(GMT+09:30) Adelaide",
		OffsetSeconds: 34200,
	},
	{
		Value:         "Australia/Darwin",
		Title:         "(GMT+09:30) Darwin",
		OffsetSeconds: 34200,
	},
	{
		Value:         "Australia/Brisbane",
		Title:         "(GMT+10:00) Brisbane",
		OffsetSeconds: 36000,
	},
	{
		Value:         "Australia/Hobart",
		Title:         "(GMT+10:00) Sydney, Melbourne, Hobart",
		OffsetSeconds: 36000,
	},
	{
		Value:         "Asia/Vladivostok",
		Title:         "(GMT+10:00) Vladivostok",
		OffsetSeconds: 36000,
	},
	{
		Value:         "Australia/Lord_Howe",
		Title:         "(GMT+10:30) Lord Howe Island",
		OffsetSeconds: 37800,
	},
	{
		Value:         "Etc/GMT-11",
		Title:         "(GMT+11:00) Solomon Is., New Caledonia",
		OffsetSeconds: 39600,
	},
	{
		Value:         "Asia/Magadan",
		Title:         "(GMT+11:00) Magadan",
		OffsetSeconds: 39600,
	},
	{
		Value:         "Pacific/Norfolk",
		Title:         "(GMT+11:30) Norfolk Island",
		OffsetSeconds: 41400,
	},
	{
		Value:         "Asia/Anadyr",
		Title:         "(GMT+12:00) Anadyr, Kamchatka",
		OffsetSeconds: 43200,
	},
	{
		Value:         "Pacific/Auckland",
		Title:         "(GMT+12:00) Auckland, Wellington",
		OffsetSeconds: 43200,
	},
	{
		Value:         "Etc/GMT-12",
		Title:         "(GMT+12:00) Fiji, Kamchatka, Marshall Is.",
		OffsetSeconds: 43200,
	},
	{
		Value:         "Pacific/Chatham",
		Title:         "(GMT+12:45) Chatham Islands",
		OffsetSeconds: 45900,
	},
	{
		Value:         "Pacific/Tongatapu",
		Title:         "(GMT+13:00) Nuku'alofa",
		OffsetSeconds: 46800,
	},
	{
		Value:         "Pacific/Kiritimati",
		Title:         "(GMT+14:00) Kiritimati",
		OffsetSeconds: 50400,
	},
}
