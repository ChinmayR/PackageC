package PackageC

import "github.com/ChinmayR/PackageD"

func FuncInPackageC() (string, error) {
	str, err := PackageD.FuncInPackageD()
	return "From PackageC: " + str, err
}
