package PackageC

import "github.com/ChinmayR/PackageD"

func FuncInPackageC() string {
	return "From PackageC: " + PackageD.FuncInPackageD()
}
