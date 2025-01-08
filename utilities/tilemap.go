package utilities

import (
	"encoding/json"
	"os"
)

type TilesetLayersJSON struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type TilesetMapJSON struct {
	Layers []*TilesetLayersJSON `json:"layers"`
}

func NewTileMapJSON(filePath string) (*TilesetMapJSON, error) {
	tileFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tileMap TilesetMapJSON
	err = json.Unmarshal(tileFile, &tileMap)
	if err != nil {
		return nil, err
	}
	return &tileMap, nil
}
