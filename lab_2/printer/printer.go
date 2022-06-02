package printer

import (
	"errors"
	"math"
	"strconv"
)

var InvalidVariables = errors.New("invalid variables")

func Printer(pageCount, count int) (pages []string, err error) {
	if pageCount < 0 || count < 0 {
		err = InvalidVariables
		return
	}
	if pageCount == 0 {
		return
	}

	var (
		listCount = int(math.Ceil(float64(pageCount) / float64(4)))
		blocks    = 1
		maxPages  = listCount * 4
	)
	if count != 0 {
		blocks = int(math.Ceil(float64(pageCount) / float64(4*count)))
	}

	for j := 1; j < blocks+1; j++ {
		if count != 0 {
			maxPages = 4 * count * j
			listCount = count
		}

		for i := 1; i < (2*listCount+1)+1; i++ {
			if i <= listCount {
				first := 0
				if pageCount >= maxPages-2*(i-1) {
					first = maxPages - 2*(i-1)
				}
				second := 0
				if pageCount >= 4*count*(j-1)+1+2*(i-1) {
					second = 4*count*(j-1) + 1 + 2*(i-1)
				}
				pages = append(pages, "("+strconv.Itoa(first)+", "+strconv.Itoa(second)+")")
				continue
			}

			if i == listCount+1 && pageCount > 1 {
				pages = append(pages, "@")
				continue
			}

			if pageCount > 1 {
				index := i - listCount - 1

				first := 0
				if pageCount >= 4*count*(j-1)+2+2*(index-1) {
					first = 4*count*(j-1) + 2 + 2*(index-1)
				}
				second := 0
				if pageCount >= maxPages-1-2*(index-1) {
					second = maxPages - 1 - 2*(index-1)
				}
				pages = append(pages, "("+strconv.Itoa(first)+", "+strconv.Itoa(second)+")")
			}
		}

		if blocks > 1 && j != blocks {
			pages = append(pages, "|")
		}
	}

	return
}
