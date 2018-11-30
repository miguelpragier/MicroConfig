package microconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const defaultConfigFile = "./conf.json"

type microConfigPairs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	OSEnv bool   `json:"osEnv"`
}

type microConfig struct {
	Pairs []microConfigPairs
}

var mc microConfig

func init() {
	if err := Load(""); err != nil {
		log.Printf("microconfig.init %s", err.Error())
	}
}

// Load/Reload configuration pairs from config file.
// If configFile=="", load from defaultConfigFile constant.
func Load(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	raw, err := ioutil.ReadFile(configFile)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &mc); err != nil {
		return err
	}

	for _, p := range mc.Pairs {
		if p.OSEnv {
			_ = os.Setenv(p.Key, p.Value)
		}
	}

	return nil
}

// Exists checks if the config structure was initialized, and if the given key exists on it.
func Exists(key string, caseInsensitive bool) bool {
	if len(mc.Pairs) < 1 {
		return false
	}

	for _, p := range mc.Pairs {
		if p.Key == key || (caseInsensitive && strings.ToLower(key) == strings.ToLower(p.Key)) {
			return true
		}
	}

	return false
}

// GetString returns the text value from the given key
func GetString(key string) (string, error) {
	if !Exists(key, false) {
		return "", errors.New("key is undefined (key-sensitive)")
	}

	for _, p := range mc.Pairs {
		if p.Key == key {
			return p.Value, nil
		}
	}

	return "", errors.New(fmt.Sprintf("couldn't process the given key: %s (key-sensitive)", key))
}

// GetInt returns the integer value from the given key
func GetInt(key string) (int, error) {
	s, err := GetString(key)

	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		return i, err
	}

	return i, nil
}

// GetFloat returns the float64 value from the given key
func GetFloat(key string) (float64, error) {
	s, err := GetString(key)

	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return f, err
	}

	return f, nil
}
