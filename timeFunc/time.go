package timeFunc

import "time"

func FormatTime(t time.Time) string {
	return t.Format("02.01.2006 15:04:05") // Аналогично: YYYY.MM.DD-hh.mm.ss
}
func FormatTimeOnlyDate(t time.Time) string {
	return t.Format("02.01.2006") // Аналогично: YYYY.MM.DD-hh.mm.ss
}
