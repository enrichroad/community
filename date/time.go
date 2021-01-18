package date

import "time"

// direction of time
const (
	Prev = "previous"
	Next = "next"
	Cur  = "current"
)

// time period
const (
	Minute = "minute"
	Hour   = "hour"
	Day    = "day"
	Week   = "week"
)

// unit
const (
	SecUnit      int64 = 1
	MilliSecUnit int64 = 1000
)

// interval of time by second unit
var (
	SecPerMin  int64 = 60
	SecPerHour       = 60 * SecPerMin
	SecPerDay        = 24 * SecPerHour
)

func fetchValueByDirection(s string) int64 {
	switch s {
	case Prev:
		return -1
	case Next:
		return 1
	default:
		return 0
	}
}

func fetchValueByUnit(period string, unit int64) int64 {
	switch period {
	case Minute:
		return SecPerMin * unit
	case Hour:
		return SecPerHour * unit
	case Day:
		return SecPerDay * unit
	default:
		return 1 * unit
	}
}

// IsSameDay if is some day return true
func IsSameDay(s, d, unit int64) bool {
	return s/SecPerDay*unit == d/SecPerDay*unit
}

// Period period struct
type Period struct {
	start time.Time
	end   time.Time
}

// Start get period start time
func (p *Period) Start() time.Time {
	return p.start
}

// End get period end time
func (p *Period) End() time.Time {
	return p.end
}

// CalcPeriodStartEnd giving time and period, calc period start time and end time
// default period is Day
func CalcPeriodStartEnd(t time.Time, period string) *Period {
	var interval int64
	switch period {
	case Minute:
		interval = SecPerMin
	case Hour:
		interval = SecPerHour
	case Day:
		interval = SecPerDay

	case Week:
		tsSec := time.Unix(t.Unix()/SecPerDay*SecPerDay, 0)
		return &Period{
			calcTimeWithDirByWeekInterval(tsSec, fetchValueByDirection(Cur)),
			time.Unix(calcTimeWithDirByWeekInterval(tsSec, fetchValueByDirection(Next)).Unix()-1, 0),
		}
	default:
		return nil
	}

	return &Period{
		time.Unix(t.Unix()/interval*interval, 0),
		time.Unix((t.Unix()/interval+1)*interval-1, 0),
	}
}

// CalcPeriodStartByDirection calc period start time
func CalcPeriodStartByDirection(ts time.Time, unit int64, period, direction string) int64 {
	var intervalSecond int64 = 0
	switch period {
	case Minute, Hour, Day:
		intervalSecond = fetchValueByUnit(period, unit)
	case Week:
		sec := fetchValueByUnit(Day, unit)
		tsSec := time.Unix(ts.Unix()*unit/sec*sec, 0)
		return calcTimeWithDirByWeekInterval(tsSec, fetchValueByDirection(direction)).Unix() * unit
	default:
		return 0
	}

	return calcNonWeekPeriod(ts.Unix(), intervalSecond, fetchValueByDirection(direction))
}

func calcNonWeekPeriod(ts, interval, direction int64) int64 {
	return ((ts / interval) + direction) * interval
}

func calcTimeWithDirByWeekInterval(t time.Time, direction int64) time.Time {
	day := int(direction) * 7
	return t.AddDate(0, 0, -int(t.Weekday())).AddDate(0, 0, day)
}
