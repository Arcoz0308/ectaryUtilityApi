package server

import (
	error2 "ectary/utils/error"
	"github.com/gofiber/fiber/v2"
	query2 "github.com/sandertv/gophertunnel/query"
	"log"
	"strings"
	"time"
)

type serverInfos struct {
	Infos interface{}
	Time  time.Time
}

var cache = map[string]serverInfos{}

func query(ctx *fiber.Ctx) error {
	serv := ctx.Params("server", "ectary.club:19132")
	if strings.Contains(serv, "%") {
		log.Println(serv, strings.Split(serv, "%"))
		var servs = map[string]interface{}{}
		var passed = false
		for _, s := range strings.Split(serv, "%") {
			if len(strings.Split(s, ":")) == 1 {
				s = s + ":19132"
			}
			if i, ok := cache[s]; ok {
				if !i.Time.Before(time.Now()) {
					servs[s] = i.Infos
					continue
				}
			}
			infos, err := query2.Do(s)
			if err != nil {
				servs[s] = requestError{
					Code:       error2.CodeErrServerQuery,
					Message:    error2.ErrServerQuery(s, err).Error(),
					StatusCode: fiber.StatusInternalServerError,
				}
				continue
			}
			cache[serv] = serverInfos{
				infos,
				time.Now().Add(time.Second * 5),
			}
			servs[s] = infos
			passed = true
		}
		if !passed {
			return ctx.Status(fiber.StatusInternalServerError).JSON(servs)
		}
		return ctx.JSON(servs)
	} else {
		if len(strings.Split(serv, ":")) == 1 {
			serv = serv + ":19132"
		}
		if i, ok := cache[serv]; ok {
			if !i.Time.Before(time.Now()) {
				return ctx.JSON(i.Infos)
			}
		}
		infos, err := query2.Do(serv)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(requestError{
				Code:       error2.CodeErrServerQuery,
				Message:    error2.ErrServerQuery(serv, err).Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
		cache[serv] = serverInfos{
			infos,
			time.Now().Add(time.Second * 5),
		}
		return ctx.JSON(infos)
	}
}
