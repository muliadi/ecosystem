package ecosystem

//Config sets up a blank configuration map which can then be accessed by all ECOSystem modules
var Config = make(map[string]string)

//SetConfig provides a hook to populate the configuration map from anywhere
func SetConfig(settings map[string]string) {
	for k, v := range settings {
		Config[k] = v
	}
}
