package gorbkrates

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {

	type args struct {
		number string
	}

	type test struct {
		name     string
		args     args
		wantRate float64
		wantErr  bool
	}

	var tests []test

	add := func(name string, arg args, wantRate float64, wantErr bool) {
		tests = append(tests, test{
			name:     name,
			args:     arg,
			wantRate: wantRate,
			wantErr:  wantErr,
		})
	}

	add("TestNow 1 on 2020-09-23 (USD)", args{"840"}, 76.2711, false)
	add("TestNow 2 false rate", args{"840"}, 78.2711, false)
	add("TestNow 3 not found rate", args{"9"}, 0, true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotRate, err := Now(tt.args.number)

			if gotRate != tt.wantRate && (err != nil) != tt.wantErr {
				t.Errorf("Now() = %v, want %v", gotRate, tt.wantRate)
			}
		})
	}
}

func TestOnDate(t *testing.T) {
	type args struct {
		number string
		date   time.Time
	}
	type test struct {
		name     string
		args     args
		wantRate float64
		wantErr  bool
	}

	var tests []test

	add := func(name string, arg args, wantRate float64, wantErr bool) {
		tests = append(tests, test{
			name:     name,
			args:     arg,
			wantRate: wantRate,
			wantErr:  wantErr,
		})
	}

	t1, err := time.Parse("2006/01/02", "2020/09/23")
	if err != nil {
		t.Errorf("Error Parse() date %v", err)
	}
	add("TestOnDate 1 2020-09-23 (USD)", args{"840", t1}, 76.2711, false)
	add("TestOnDate 2 2020-09-23 (EUR)", args{"978", t1}, 89.4813, false)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRate, err := OnDate(tt.args.number, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRate != tt.wantRate {
				t.Errorf("OnDate() = %v, want %v", gotRate, tt.wantRate)
			}
		})
	}
}

func TestForPeriod(t *testing.T) {
	type args struct {
		number string
		start  time.Time
		end    time.Time
	}

	type test struct {
		name         string
		args         args
		wantRatesLen int
		wantErr      bool
	}

	var tests []test

	add := func(name string, arg args, wantRatesLen int, wantErr bool) {
		tests = append(tests, test{
			name:         name,
			args:         arg,
			wantRatesLen: wantRatesLen,
			wantErr:      wantErr,
		})
	}

	t1, err := time.Parse("2006/01/02", "2020/09/01")
	if err != nil {
		t.Errorf("Error Parse() date %v", err)
	}

	t2, err := time.Parse("2006/01/02", "2020/09/23")
	if err != nil {
		t.Errorf("Error Parse() date %v", err)
	}

	add("TestForPeriod 1 2020-09-01 - 2020-09-23 (USD)", args{"840", t1, t2}, 23, false)
	add("TestForPeriod 2 2020-09-01 - 2020-09-23 (EUR)", args{"978", t1, t2}, 23, false)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRates, err := ForPeriod(tt.args.number, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForPeriod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotRates) == tt.wantRatesLen {
				t.Errorf("ForPeriod() = %v, want %v", len(gotRates), tt.wantRatesLen)
			}
		})
	}
}
