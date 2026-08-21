package state

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	LastMessageID int64 `json:"last_message_id"`
}

type Store struct {
	path string
}

func NewStore() (*Store, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(
		home,
		".imsgbridge",
	)

	err = os.MkdirAll(
		dir,
		0755,
	)

	if err != nil {
		return nil, err
	}


	return &Store{
		path: filepath.Join(
			dir,
			"state.json",
		),
	}, nil
}


func (s *Store) Load() (*State, error) {

	data, err := os.ReadFile(
		s.path,
	)

	if os.IsNotExist(err) {
		return &State{
			LastMessageID: 0,
		}, nil
	}

	if err != nil {
		return nil, err
	}


	var state State

	err = json.Unmarshal(
		data,
		&state,
	)

	if err != nil {
		return nil, err
	}


	return &state, nil
}


func (s *Store) Save(state *State) error {

	data, err := json.MarshalIndent(
		state,
		"",
		"  ",
	)

	if err != nil {
		return err
	}


	return os.WriteFile(
		s.path,
		data,
		0644,
	)
}
