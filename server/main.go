package server

import (
	"database/sql"
	"ectary/cape"
	"ectary/handlers/config"
	"ectary/handlers/mysql"
	"ectary/skin"
	error2 "ectary/utils/error"
	"github.com/gofiber/fiber/v2"
	cache2 "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"image/png"
	"log"
	"strconv"
	"strings"
	"time"
)

type requestError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func Load(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Set("ARC-DefaultContent", "false")
		return ctx.Next()
	})
	app.Use(logger.New(logger.Config{Format: "${ip} [${time}] ${status} - ${latency}(${bytesSent}b) ${method} ${url}\n"}))
	app.Use(cache2.New(cache2.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.ToLower(c.Query("renew", "false")) == "true"
		},
		ExpirationGenerator: func(ctx *fiber.Ctx, c *cache2.Config) time.Duration {
			s, _ := strconv.Atoi(ctx.GetRespHeader("ARC-CacheDelay", "0"))
			return time.Second * time.Duration(s)
		},
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return utils.CopyString(ctx.OriginalURL()) + utils.CopyString(ctx.Get("Authorization"))
		},
	}))
	app.Get("/players/:name/skin/head", headSkin)
	app.Get("/players/:name/skin/full", fullSkin)
	app.Get("/players/:name/info/full", auth, playerFullInfos)
	app.Get("/players/:name/info/bedwars", auth, playerBedwarsInfos)
	app.Get("/players/:name/info/practice", auth, playerPracticeInfos)
	app.Get("/players/:name/info", auth, playerInfos)
	app.Get("/players/:name/cape", playerCape)
	app.Get("/servers/:server", query)
}

func headSkin(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "300")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	border, _ := strconv.Atoi(ctx.Query("border", "0"))
	size, _ := strconv.Atoi(ctx.Query("size", "0"))
	withUUID := ctx.Query("xuid", "false")
	if border > 100 || border < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(requestError{
			Code:       error2.CodeErrInvalidBorderSize,
			Message:    error2.ErrInvalidBorderSize.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	if !validSize(size) {
		return ctx.Status(fiber.StatusBadRequest).JSON(requestError{
			Code:       error2.CodeErrInvalidImageSize,
			Message:    error2.ErrInvalidImageSize.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	var d []uint8
	var err error
	if withUUID == "true" {
		d, err = mysql.PlayerSkinByXUID(name)
	} else {
		d, err = mysql.PlayerSkinByUsername(name)
	}
	if err != nil {
		if err, ok := err.(*error2.ApiError); ok {
			if err.Code == error2.CodeErrNoSkin {
				img := config.DefaultSkin.Head(border, size)
				ctx.Type("png")
				err2 := png.Encode(ctx, img)
				if err2 != nil {
					log.Println(err2)
					return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
						Code:       error2.CodeErrPngEncoding,
						Message:    error2.ErrPngEncoding.Error(),
						StatusCode: fiber.StatusInternalServerError,
					})
				}
				ctx.Set("ARC-DefaultContent", "true")
				ctx.Status(fiber.StatusOK)
				return nil
			}
			// other possibility are decompressing error
			log.Println(err.ErrorWithDebug())
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrDecompress,
				Message:    error2.ErrDecompress.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})

		}
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	s, err2 := skin.FromData(d)
	if s == nil && err2 == nil {
		img := config.DefaultSkin.Head(border, size)
		ctx.Type("png")
		err2 := png.Encode(ctx, img)
		if err2 != nil {
			log.Println(err2)
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrPngEncoding,
				Message:    error2.ErrPngEncoding.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
		ctx.Set("ARC-DefaultContent", "true")
		ctx.Status(fiber.StatusOK)
		return nil
	}
	if err2 != nil {
		log.Println(err2.ErrorWithDebug())
		if err2.Code == error2.CodeErrDecompress {
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrDecompress,
				Message:    error2.ErrDecompress.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		} else {
			// other possibility : invalid skin size
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrInvalidSkinSize,
				Message:    error2.ErrInvalidSkinSize.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
	}
	img := s.Head(border, size)
	ctx.Type("png")
	err = png.Encode(ctx, img)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrPngEncoding,
			Message:    error2.ErrPngEncoding.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	ctx.Status(fiber.StatusOK)
	return nil
}
func fullSkin(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "300")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	withUUID := ctx.Query("xuid", "false")
	var d []uint8
	var err error
	if withUUID == "true" {
		d, err = mysql.PlayerSkinByXUID(name)
	} else {
		d, err = mysql.PlayerSkinByUsername(name)
	}
	if err != nil {
		if err, ok := err.(*error2.ApiError); ok {
			if err.Code == error2.CodeErrNoSkin {
				skin2 := config.DefaultSkin
				img := skin2.FullSkin()
				ctx.Type("png")
				err2 := png.Encode(ctx, img)
				if err2 != nil {
					log.Println(err2)
					return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
						Code:       error2.CodeErrPngEncoding,
						Message:    error2.ErrPngEncoding.Error(),
						StatusCode: fiber.StatusInternalServerError,
					})
				}
				ctx.Set("ARC-DefaultContent", "true")
				ctx.Status(fiber.StatusOK)
				return nil
			}
			// other possibility are decompressing error
			log.Println(err.ErrorWithDebug())
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrDecompress,
				Message:    error2.ErrDecompress.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})

		}
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	s, err2 := skin.FromData(d)
	if s == nil && err2 == nil {
		img := config.DefaultSkin.FullSkin()
		ctx.Type("png")
		err2 := png.Encode(ctx, img)
		if err2 != nil {
			log.Println(err2)
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrPngEncoding,
				Message:    error2.ErrPngEncoding.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
		ctx.Set("ARC-DefaultContent", "true")
		ctx.Status(fiber.StatusOK)
		return nil
	}
	if err2 != nil {
		log.Println(err2.ErrorWithDebug())
		if err2.Code == error2.CodeErrDecompress {
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrDecompress,
				Message:    error2.ErrDecompress.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		} else {
			// other possibility : invalid skin size
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrInvalidSkinSize,
				Message:    error2.ErrInvalidSkinSize.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
	}
	img := s.FullSkin()
	ctx.Type("png")
	err = png.Encode(ctx, img)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrPngEncoding,
			Message:    error2.ErrPngEncoding.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	ctx.Set("ARC-DefaultContent", "false")
	ctx.Status(fiber.StatusOK)
	return nil
}

func playerFullInfos(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "30")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	withUUID := ctx.Query("xuid", "false")
	var d mysql.PlayerDataFull
	var err error
	if withUUID == "true" {
		d, err = mysql.PlayerFullInfosByXUID(name)
	} else {
		d, err = mysql.PlayerFullInfosByUsername(name)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(PlayerFullInfosFromMysql(d))
}

func playerInfos(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "30")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%", " ")
	withUUID := ctx.Query("xuid", "false")
	var d mysql.PlayerData
	var err error
	if withUUID == "true" {
		d, err = mysql.PlayerInfosByXUID(name)
	} else {
		d, err = mysql.PlayerInfosByUsername(name)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(PlayerInfosFromMysql(d))
}

func playerBedwarsInfos(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "30")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	withUUID := ctx.Query("xuid", "false")
	var d mysql.BedwarsStats
	var err error
	if withUUID == "true" {
		d, err = mysql.PlayerBedwarsInfosByXUID(name)
	} else {
		d, err = mysql.PlayerBedwarsInfosByUsername(name)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(BedWarsInfosFromMysql(d))
}
func playerPracticeInfos(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "30")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	withUUID := ctx.Query("xuid", "false")
	var d mysql.PracticeStats
	var err error
	if withUUID == "true" {
		d, err = mysql.PlayerPracticeInfosByXUID(name)
	} else {
		d, err = mysql.PlayerPracticeInfosByUsername(name)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(PracticeInfosFromMysql(d))
}

func playerCape(ctx *fiber.Ctx) error {
	ctx.Set("ARC-CacheDelay", "300")
	name := ctx.Params("name")
	name = strings.ReplaceAll(name, "%20", " ")
	withUUID := ctx.Query("xuid", "false")
	var (
		d   []uint8
		err error
	)
	if withUUID == "true" {
		d, err = mysql.PlayerCapeByXUID(name)
	} else {
		d, err = mysql.PlayerCapeByUsername(name)
	}
	if err != nil {
		if err2, ok := err.(*error2.ApiError); ok {
			if err2.Code == error2.CodeErrNoCape {
				return ctx.Status(fiber.StatusNotFound).JSON(requestError{
					Code:       error2.CodeErrNoCape,
					Message:    error2.ErrNoCape.Error(),
					StatusCode: fiber.StatusNotFound,
				})
			}
			// other possibility : decompress error
			log.Println(err2.ErrorWithDebug())
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrDecompress,
				Message:    error2.ErrDecompress.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(requestError{
				Code:       error2.CodeErrPlayerNotFound,
				Message:    error2.ErrPlayerNotFound.Error(),
				StatusCode: fiber.StatusNotFound,
			})
		}
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDatabase,
			Message:    error2.ErrDatabase.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	if len(d) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(requestError{
			Code:       error2.CodeErrNoCape,
			Message:    error2.ErrNoCape.Error(),
			StatusCode: fiber.StatusNotFound,
		})
	}
	c, err := cape.FromData(d)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrDecompress,
			Message:    error2.ErrDecompress.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	img := c.Image()
	ctx.Type("png")
	err = png.Encode(ctx, img)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
			Code:       error2.CodeErrPngEncoding,
			Message:    error2.ErrPngEncoding.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	ctx.Status(fiber.StatusOK)
	return nil
}
