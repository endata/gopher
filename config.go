package gopher

var defaultConfig = map[string]map[string]interface{}{
	LOGGER: {
		"FullTimestamp": true,
	},
	RENDERER: {
		"ViewsDir": "templates",
	},
}
