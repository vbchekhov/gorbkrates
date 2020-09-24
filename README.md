# 📈 rbc.ru rates

Simple library for read exchange rates rbc.ru on Go.

## Examples

Rate now:
```golang
rate, err := gorbkrates.Now("840") // check code number by ISO 4217
if err != nil { ... }

// return float64
```

Rate on date:
```golang
date, err := time.Parse("2006/01/02", "2020/09/23")
if err != nil { ... }

rate, err := gorbkrates.OnDate("840") // check code number by ISO 4217
if err != nil { ... }

// return float64
```

Rate for period:
```golang
start, err := time.Parse("2006/01/02", "2020/09/23")
if err != nil { ... }
end, err := time.Parse("2006/01/02", "2020/09/23")
if err != nil { ... }

rates, err := gorbkrates.ForPeriod("840", start, end) // check code number by ISO 4217

// return map[time.Time]float64
```