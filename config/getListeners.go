package config

func GetListeners() []*Listener {

	listeners := Config.Listeners

	for _, eachListener := range listeners {
		if eachListener.Filter != nil {
			for _, eachFilterRule := range eachListener.Filter.Rules {
				eachFilterRule.TargetCluster = nil
			}
		}
	}

	return listeners
}
