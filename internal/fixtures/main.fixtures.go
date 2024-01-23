package fixtures

import "TOPO/appctr"

func MakeFixtures() {
	appctr.Log().Debug("Starting fixtures")
	MakeUsers()
	appctr.Log().Debug("Finishing fixtures")
}
