package mysql

import (
	"bytes"
	"compress/zlib"
	"database/sql"
	error2 "ectary/utils/error"
	_ "embed"
	"encoding/base64"
	"github.com/qustavo/dotsql"
	"io/ioutil"
	"strings"
)

var (
	//go:embed query.sql
	d           string
	prepares    = map[string]*sql.Stmt{}
	Ranks       []*RankData
	DefaultRank *RankData
)

func loadPrepares() {
	dot, err := dotsql.LoadFromString(d)
	if err != nil {
		panic(err)
	}
	for k, v := range dot.QueryMap() {
		p, err := DB.Prepare(v)
		if err != nil {
			panic(err)
		}
		prepares[k] = p
	}
}

func decompressSkin(b []byte) ([]uint8, error) {
	bytesReader := bytes.NewReader(b)
	zipReader, err := zlib.NewReader(bytesReader)
	if err != nil {
		return nil, &error2.ApiError{
			Code: error2.CodeErrDecompress,
			AdditionalData: map[string]interface{}{
				"bytes": b,
			},
			Err: err,
		}
	}
	result, err := ioutil.ReadAll(zipReader)
	if err != nil {
		return nil, &error2.ApiError{
			Code: error2.CodeErrDecompress,
			AdditionalData: map[string]interface{}{
				"bytes": b,
			},
			Err: err,
		}
	}
	data, err := base64.StdEncoding.DecodeString(string(result))
	if err != nil {
		return nil, &error2.ApiError{
			Code: error2.CodeErrDecompress,
			AdditionalData: map[string]interface{}{
				"bytes": b,
			},
			Err: err,
		}
	}
	return data, nil
}

func decompressCape(b []byte) ([]uint8, error) {
	bytesReader := bytes.NewReader(b)
	zipReader, err := zlib.NewReader(bytesReader)
	if err != nil {
		return nil, &error2.ApiError{
			Code: error2.CodeErrDecompress,
			AdditionalData: map[string]interface{}{
				"bytes": b,
			},
			Err: err,
		}
	}
	r, err := ioutil.ReadAll(zipReader)
	if err != nil {
		return nil, &error2.ApiError{
			Code: error2.CodeErrDecompress,
			AdditionalData: map[string]interface{}{
				"bytes": b,
			},
			Err: err,
		}
	}
	by, err := base64.StdEncoding.DecodeString(strings.Split(string(r), "\"")[1])
	if err != nil {
		return nil, &error2.ApiError{
			Code: error2.CodeErrDecompress,
			AdditionalData: map[string]interface{}{
				"bytes": b,
			},
			Err: err,
		}
	}
	return by, nil
}

func LoadRanks() {
	Ranks = []*RankData{}
	Ranks = ranks()
	for _, r := range Ranks {
		if r.RankDefault {
			DefaultRank = r
		}
	}
}
