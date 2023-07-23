package config

import (
	"ectary/skin"
	"encoding/json"
	"github.com/troian/toml"
	"image/png"
	"os"
)

type config struct {
	Port     int `toml:"port" comment:"the port of the website"`
	Database struct {
		User     string `toml:"user" comment:"the user that connect to db"`
		Password string `toml:"password" comment:"the db password"`
		NetType  string `toml:"netType" comment:"the type of net connection for the db"`
		Host     string `toml:"host" comment:"the host address of the db"`
		DbName   string `toml:"dbName" comment:"the name of the database"`
	} `toml:"database"`
}

var (
	C           config
	DefaultSkin *skin.Skin
	Tokens      []string
)

func Init() {
	f, err := os.OpenFile("config.toml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			createConfigFile()
		} else {
			panic(err)
		}
	} else {
		_, err = toml.DecodeReader(f, &C)
		if err != nil {
			panic(err)
		}
	}
	f1, err := os.OpenFile("token.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			f1, err = os.Create("token.json")
			if err != nil {
				panic(err)
			}
			_, err = f1.WriteString("[]")
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		err = json.NewDecoder(f1).Decode(&Tokens)
		if err != nil {
			panic(err)
		}
	}
	if _, err = os.Stat("resource/"); os.IsNotExist(err) {
		err = os.Mkdir("resource/", os.ModePerm)
		if err != nil {
			panic(err)
		}
		panic("creating resource folder, please add default skin image with the name of \"default_skin.png\" in resource folder")
	} else if err != nil {
		if err != nil {
			panic(err)
		}
	}
	f2, err := os.OpenFile("resource/default_skin.png", os.O_RDONLY, os.ModePerm)
	if err == os.ErrNotExist {
		panic("please add default skin image with the name of \"default_skin.png\" in resource folder")
	}
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(f2)
	if err != nil {
		panic(err)
	}
	DefaultSkin = skin.FromImage(img)
}
func createConfigFile() {
	d := config{
		Port: 8000,
		Database: struct {
			User     string `toml:"user" comment:"the user that connect to db"`
			Password string `toml:"password" comment:"the db password"`
			NetType  string `toml:"netType" comment:"the type of net connection for the db"`
			Host     string `toml:"host" comment:"the host address of the db"`
			DbName   string `toml:"dbName" comment:"the name of the database"`
		}{
			User:     "",
			Password: "",
			NetType:  "tcp",
			Host:     "",
			DbName:   "EctaryNetwork2",
		},
	}
	C = d
	f, err := os.Create("config.toml")
	if err != nil {
		panic(err)
	}
	err = toml.NewEncoder(f).Encode(d)
	if err != nil {
		panic(err)
	}
}
