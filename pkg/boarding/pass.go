package boarding

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type BoardingPass struct {
	SpacePartition string
}

// GetSeatNumber converts seat partition data
// to a binary representation and then decodes the seat num
func (p *BoardingPass) GetSeatNumber() int {
	row := p.SpacePartition[0:7]
	col := p.SpacePartition[7:]
	row = strings.ReplaceAll(row, "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	col = strings.ReplaceAll(col, "L", "0")
	col = strings.ReplaceAll(col, "R", "1")
	colNum, _ := strconv.ParseInt(col, 2, 64)
	rowNum, _ := strconv.ParseInt(row, 2, 64)

	log.WithFields(logrus.Fields{"col": col, "row": row}).Debug("Getting seat number")
	log.Debugf("Col num: %d, Row num: %d", colNum, rowNum)
	return int((rowNum * 8) + colNum)
}
