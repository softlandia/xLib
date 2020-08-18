package xlib

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type tSortBytes = struct {
	i []byte
	o []byte
}

var dSortBytes = []tSortBytes{
	tSortBytes{
		[]byte{0, 'u', 'c', 'k'},
		[]byte{0, 'c', 'k', 'u'}},
	tSortBytes{
		[]byte{0, 'c', 'k', 0},
		[]byte{0, 0, 'c', 'k'}},
	tSortBytes{
		[]byte{'f', 'u', 'c', 'k', 'e', 'm', 'a', 'l', 'l'},
		[]byte{'a', 'c', 'e', 'f', 'k', 'l', 'l', 'm', 'u'}},
}

func TestSortBytes(t *testing.T) {
	for _, d := range dSortBytes {
		c := SortBytes(d.i)
		assert.Equal(t, true, bytes.Equal(c, d.o), fmt.Sprintf("%v\n", c))
	}
}

func TestWeekDayR(t *testing.T) {
	tmp, _ := time.Parse("2006-01-02", "2020-01-01")
	assert.Equal(t, "Ср", WeekdayR(tmp))
	tmp, _ = time.Parse("2006-01-02", "2020-12-31")
	assert.Equal(t, "Чт", WeekdayR(tmp))
	tmp, _ = time.Parse("2006-01-02", "2020-08-30")
	assert.Equal(t, "Вс", WeekdayR(tmp))
}

func TestDay(t *testing.T) {
	tmp, _ := time.Parse("2006-01-02", "2020-01-01")
	assert.Equal(t, tmp, Day(0))
	tmp, _ = time.Parse("2006-01-02", "2020-12-31")
	assert.Equal(t, tmp, Day(365))
}

func TestFirstDay(t *testing.T) {
	fd, _ := time.Parse("2006-01-02", "2020-01-01")
	assert.Equal(t, fd, FirstDay())
}

func TestLastDay(t *testing.T) {
	fd, _ := time.Parse("2006-01-02", "2020-12-31")
	assert.Equal(t, fd, LastDay())
}

func TestIndex(t *testing.T) {
	fd, _ := time.Parse("2006-01-02", "2020-01-01")
	assert.Equal(t, 0, Index(fd))

	fd, _ = time.Parse("2006-01-02", "2020-01-02")
	assert.Equal(t, 1, Index(fd))

	fd, _ = time.Parse("2006-01-02", "2020-12-31")
	assert.Equal(t, 365, Index(fd))
}

func TestDayCount(t *testing.T) {
	assert.Equal(t, 366, DayCount())
}
