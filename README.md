# tiff

calculates the difference between two times

## install

```bash
go get -u github.com/usk81/tiff
```

## usage

```go
d := New(time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2001, 12, 10, 11, 10, 20, 0, time.Local))

years, months, days, hours, mins, sec := d.Term()
// years:  1
// months: 11
// days:   9
// hours:  11
// mins:   10
// sec:    20

years = d.Years()
// 1

months = d.Months()
// 23
```
