package treblle

var Config internalConfiguration

const defaultServerURL = "https://rocknrolla.treblle.com"

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

// Configuration sets up and customizes communication with the Treblle API
type Configuration struct {
	APIKey     string
	ProjectID  string
	KeysToMask []string
	ServerURL  string
	Logger     Logger
}

// internalConfiguration is used for communication with Treblle API and contains optimizations
type internalConfiguration struct {
	Configuration
	KeysMap      map[string]interface{}
	serverInfo   ServerInfo
	languageInfo LanguageInfo
}

func Configure(config Configuration) {
	if config.ServerURL != "" {
		Config.ServerURL = config.ServerURL
	} else {
		Config.ServerURL = defaultServerURL
	}
	if config.APIKey != "" {
		Config.APIKey = config.APIKey
	}
	if config.ProjectID != "" {
		Config.ProjectID = config.ProjectID
	}
	if len(config.KeysToMask) > 0 {
		Config.KeysToMask = config.KeysToMask

		// transform the string slice to a map for faster retrieval
		Config.KeysMap = make(map[string]interface{})
		for _, v := range config.KeysToMask {
			Config.KeysMap[v] = nil
		}
	}

	if config.Logger != nil {
		Config.Logger = config.Logger
	}

	Config.serverInfo = getServerInfo()
	Config.languageInfo = getLanguageInfo()
}
