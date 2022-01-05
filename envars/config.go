package envars

import "os"

//GetWithDefault fetches an environment variable.
// If it has not been set, it defaults to the given value.
//
//Note that if the envar is set to an empty string the default will
// not be used unless acceptEmpty is set to false
func GetWithDefault(envar, defaultStr string, acceptEmpty bool) string {
	out, ok := os.LookupEnv(envar)
	if out != "" {
		return out
	}
	if out == "" && ok && acceptEmpty {
		return out
	}
	return defaultStr
}
