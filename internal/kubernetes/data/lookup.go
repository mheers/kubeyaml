package data

import "fmt"

type StaticFiles struct{}

// Swagger is a fairly meh function. It's poorly named and tied to the update-schemas file.
func (s *StaticFiles) Swagger(version string) []byte {
	switch version {
	case "1.8":
		return s.OneEight()
	case "1.9":
		return s.OneNine()
	case "1.10":
		return s.OneTen()
	case "1.11":
		return s.OneEleven()
	case "1.12":
		return s.OneTwelve()
	case "1.13":
		return s.OneThirteen()
	case "1.14":
		return s.OneFourteen()
	case "1.15":
		return s.OneFifteen()
	case "1.16":
		return s.OneSixteen()
	case "1.17":
		return s.OneSeventeen()
	case "1.18":
		return s.OneEighteen()
	case "1.19":
		return s.OneNineteen()
	case "1.20":
		return s.OneTwenty()
	case "1.21":
		return s.OneTwentyOne()
	case "1.22":
		return s.OneTwentyTwo()
	case "1.23":
		return s.OneTwentyThree()
	case "1.24":
		return s.OneTwentyFour()
	case "1.25":
		return s.OneTwentyFive()
	case "1.26":
		return s.OneTwentySix()
	default:
		panic(fmt.Sprintf("unknown version %v", version))
	}
}
