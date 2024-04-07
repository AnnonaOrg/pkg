package param

import (
	"strconv"
	"strings"
	"time"
)

var EmptyTime = time.Time{}

type String string

func (p String) String() string {
	return strings.TrimSpace(string(p))
}

func (p String) Raw() string {
	return string(p)
}

func (p String) Split(sep string, limit ...int) StringSlice {
	s := p.String()
	if len(s) == 0 {
		return StringSlice{}
	}
	if len(limit) > 0 {
		return strings.SplitN(s, sep, limit[0])
	}
	return strings.Split(s, sep)
}

func (p String) SplitAny(sep string, limit ...int) StringSlice {
	s := p.String()
	if len(limit) > 0 {
		return strings.SplitN(s, sep, limit[0])
	}
	return strings.Split(s, sep)
}

func (p String) Trim() String {
	return String(strings.TrimSpace(string(p)))
}

func (p String) Interface() interface{} {
	return interface{}(p)
}

func (p String) Int() int {
	if len(p) > 0 {
		r, _ := strconv.Atoi(p.String())
		return r
	}
	return 0
}

func (p String) Int64() int64 {
	if len(p) > 0 {
		r, _ := strconv.ParseInt(p.String(), 10, 64)
		return r
	}
	return 0
}

func (p String) Int32() int32 {
	if len(p) > 0 {
		r, _ := strconv.ParseInt(p.String(), 10, 32)
		return int32(r)
	}
	return 0
}

func (p String) Uint() uint {
	if len(p) > 0 {
		r, _ := strconv.ParseUint(p.String(), 10, 0)
		return uint(r)
	}
	return 0
}

func (p String) Uint64() uint64 {
	if len(p) > 0 {
		r, _ := strconv.ParseUint(p.String(), 10, 64)
		return r
	}
	return 0
}

func (p String) Uint32() uint32 {
	if len(p) > 0 {
		r, _ := strconv.ParseUint(p.String(), 10, 32)
		return uint32(r)
	}
	return 0
}

func (p String) Float32() float32 {
	if len(p) > 0 {
		r, _ := strconv.ParseFloat(p.String(), 32)
		return float32(r)
	}
	return 0
}

func (p String) Float64() float64 {
	if len(p) > 0 {
		r, _ := strconv.ParseFloat(p.String(), 64)
		return r
	}
	return 0
}

func (p String) Bool() bool {
	if len(p) > 0 {
		r, _ := strconv.ParseBool(p.String())
		return r
	}
	return false
}

func (p String) Timestamp() time.Time {
	if len(p) > 0 {
		s := strings.SplitN(p.String(), `.`, 2)
		var sec int64
		var nsec int64
		switch len(s) {
		case 2:
			nsec = String(s[1]).Int64()
			fallthrough
		case 1:
			sec = String(s[0]).Int64()
		}
		return time.Unix(sec, nsec)
	}
	return EmptyTime
}

func (p String) Duration(defaults ...time.Duration) time.Duration {
	if len(p) > 0 {
		t, err := time.ParseDuration(p.String())
		if err == nil {
			return t
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return 0
}

func (p String) DateTime(layouts ...string) time.Time {
	if len(p) > 0 {
		layout := `2006-01-02 15:04:05`
		if len(layouts) > 0 {
			layout = layouts[0]
		}
		t, _ := time.ParseInLocation(layout, p.String(), time.Local)
		return t
	}
	return EmptyTime
}
