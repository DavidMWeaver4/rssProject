package config
import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error){
	fullPath, err := getConfigFilePath()
	if err != nil{
		return Config{}, err
	}
	file, err := os.Open(fullPath)
	if err!= nil{
		return Config{}, err
	}
	defer file.Close()
	var cfg Config
	err := json.NewDecoder(file).Decode(&cfg)
	if err != nil{
		return Config{}, err
	}

	return cfg, nil

}
func (cfg *Config) SetUser(name string) error{
	cfg.CurrentUserName = name
	err := write(*cfg)
	if err != nil{
		return err
	}
	return nil

}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil{
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil{
		return err
	}
	defer file.Close()
	err := json.NewEncoder(file).Encode(&cfg)
	if err != nil{
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"
	home, err := os.UserHomeDir()
	if err != nil{
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
