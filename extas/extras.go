package extas

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var mesesEnEspanol = map[string]string{
	"ene": "Jan",
	"feb": "Feb",
	"mar": "Mar",
	"abr": "Apr",
	"may": "May",
	"jun": "Jun",
	"jul": "Jul",
	"ago": "Aug",
	"sep": "Sep",
	"oct": "Oct",
	"nov": "Nov",
	"dic": "Dec",
}

func ParseFechaEspanol(fechaStr string) (time.Time, error) {
	fechaStr = strings.TrimSpace(fechaStr)
	for esp, eng := range mesesEnEspanol {
		if strings.Contains(strings.ToLower(fechaStr), esp) {
			fechaStr = strings.Replace(strings.ToLower(fechaStr), esp, eng, 1)
			break
		}
	}
	fechaParseada, err := time.Parse("02 Jan 2006", fechaStr)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, errors.New("error parseando")
	}
	return fechaParseada, nil
}
