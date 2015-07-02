package gopher

var defaultConfig = map[string]map[string]interface{}{
	LOGGER: {
		"FullTimestamp": true,
	},
	ROUTER: {
		"name": "Router",
	},
	RENDERER: {
		"provider": "Renderer",
	},
}
