package tiff

import (
	"testing"
	"time"
)

func TestDuration_Term(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utc, _ := time.LoadLocation("UTC")

	type fields struct {
		From time.Time
		To   time.Time
	}
	tests := []struct {
		name        string
		fields      fields
		wantYears   int
		wantMonths  int
		wantDays    int
		wantHours   int
		wantMinutes int
		wantSeconds int
	}{
		{
			name: "to_date_is_later_than_from_date",
			fields: fields{
				From: time.Date(2019, 11, 11, 0, 0, 0, 0, jst),
				To:   time.Date(2020, 12, 13, 10, 15, 20, 0, jst),
			},
			wantYears:   1,
			wantMonths:  1,
			wantDays:    2,
			wantHours:   10,
			wantMinutes: 15,
			wantSeconds: 20,
		},
		{
			name: "to_date_is_later_than_from_date",
			fields: fields{
				From: time.Date(2020, 12, 13, 10, 15, 20, 0, jst),
				To:   time.Date(2019, 11, 11, 0, 0, 0, 0, jst),
			},
			wantYears:   1,
			wantMonths:  1,
			wantDays:    2,
			wantHours:   10,
			wantMinutes: 15,
			wantSeconds: 20,
		},
		{
			name: "same_date",
			fields: fields{
				From: time.Date(2020, 11, 10, 10, 40, 20, 0, jst),
				To:   time.Date(2020, 11, 10, 10, 40, 20, 0, jst),
			},
			wantYears:   0,
			wantMonths:  0,
			wantDays:    0,
			wantHours:   0,
			wantMinutes: 0,
			wantSeconds: 0,
		},
		{
			name: "same_date_but_time_zone_is_different",
			fields: fields{
				From: time.Date(2020, 11, 10, 10, 40, 20, 0, utc),
				To:   time.Date(2020, 11, 10, 10, 40, 20, 0, jst),
			},
			wantYears:   0,
			wantMonths:  0,
			wantDays:    0,
			wantHours:   9,
			wantMinutes: 0,
			wantSeconds: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(tt.fields.From, tt.fields.To)
			gotYears, gotMonths, gotDays, gotHours, gotMinutes, gotSeconds := d.Term()
			if gotYears != tt.wantYears {
				t.Errorf("Duration.Term() gotYears = %v, want %v", gotYears, tt.wantYears)
			}
			if gotMonths != tt.wantMonths {
				t.Errorf("Duration.Term() gotMonths = %v, want %v", gotMonths, tt.wantMonths)
			}
			if gotDays != tt.wantDays {
				t.Errorf("Duration.Term() gotDays = %v, want %v", gotDays, tt.wantDays)
			}
			if gotHours != tt.wantHours {
				t.Errorf("Duration.Term() gotHours = %v, want %v", gotHours, tt.wantHours)
			}
			if gotMinutes != tt.wantMinutes {
				t.Errorf("Duration.Term() gotMinutes = %v, want %v", gotMinutes, tt.wantMinutes)
			}
			if gotSeconds != tt.wantSeconds {
				t.Errorf("Duration.Term() gotSeconds = %v, want %v", gotSeconds, tt.wantSeconds)
			}
		})
	}
}
