// idealy we would keep state files in a store where multiple devices can access like s3
// but for simplicity we'll keep it as a local file ./state/app.json

package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"slices"
)

type environment struct {
	Name string   `json:"name"`
	Keys []string `json:"keys"`
}

type appState struct {
	Environments []environment `json:"environments"`
}

func (s *appState) findIndex(env string) int {
	removeIndex := -1

	for i, s := range s.Environments {
		if s.Name == env {
			removeIndex = i
			break
		}
	}

	return removeIndex

}

func (s *appState) Load() error {
	filePath, _ := filepath.Abs("./state/app.json")
	contents, err := os.ReadFile(filePath)
	// right now I'm gonna assume that if there's an error it's because the file
	// doesn't exist but obviusly you would wanna check
	if err != nil {
		s.Environments = make([]environment, 0)
		s.Write()
	}

	err = json.Unmarshal([]byte(contents), &s)

	if err != nil {
		return err
	}

	return nil

}
func (s *appState) Write() {
	filePath, _ := filepath.Abs("./state/app.json")
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatalf("Failed to create/truncate file: %v", err)
	}
	defer f.Close()

	// Write the new content to the file

	json, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Failed to create json from current state: %v", err)
	}

	_, err = f.Write(json)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

}
func (s *appState) AddEnv(env string) {
	envExists := s.findIndex(env)
	if envExists != -1 {
		return
	}

	newEnv := environment{Name: env, Keys: make([]string, 0)}
	s.Environments = append(s.Environments, newEnv)

	s.Write()
}
func (s *appState) RemoveEnv(env string) {
	removeIndex := s.findIndex(env)

	if removeIndex != -1 {
		//  might run into issues if the env to remove is the last element
		s.Environments = slices.Delete(s.Environments, removeIndex, removeIndex+1)
	}
}
func (s *appState) AddSecret(env string, key string) {
	envIndex := s.findIndex(env)
	defer s.Write()
	if envIndex == -1 {
		s.Environments = append(s.Environments, environment{
			Name: env,
			Keys: []string{key},
		})

		return
	}

	s.Environments[envIndex] = environment{
		Name: s.Environments[envIndex].Name,
		Keys: append(s.Environments[envIndex].Keys, key),
	}

}
func (s *appState) RemoveSecret(env string, key string) {
	envIndex := s.findIndex(env)
	if envIndex == -1 {
		return
	}

	keyIndex := -1

	for i, k := range s.Environments[envIndex].Keys {
		if k == key {
			keyIndex = i
			break
		}
	}

	if keyIndex == -1 {
		return
	}

	s.Environments[envIndex] = environment{
		Name: s.Environments[envIndex].Name,
		Keys: slices.Delete(s.Environments[envIndex].Keys, keyIndex, keyIndex+1),
	}

	s.Write()

}
