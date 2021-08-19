






package marine

import "fmt"
import "time"
import "strconv"




func GetTimeStamp() string {
     loc, _ := time.LoadLocation("America/Los_Angeles")
     t := time.Now().In(loc)
     return t.Format("20060102150405")
}
func GetTodaysDate() string {
    loc, _ := time.LoadLocation("America/Los_Angeles")
    current_time := time.Now().In(loc)
    return current_time.Format("2006-01-02")
}

func GetTodaysDateTime() string {
    loc, _ := time.LoadLocation("America/Los_Angeles")
    current_time := time.Now().In(loc)
    return current_time.Format("2006-01-02 15:04:05")
}

func GetTodaysDateTimeFormatted() string {
    loc, _ := time.LoadLocation("America/Los_Angeles")
    current_time := time.Now().In(loc)
    return current_time.Format("Jan 2, 2006 at 3:04 PM")
}

func GetTimeStampFromDate(dtformat string) string {
    form := "Jan 2, 2006 at 3:04 PM"
    t2, _ := time.Parse(form, dtformat)
    return t2.Format("20060102150405")
}


func epoch2Date_() {
//Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds, respectively.

    fmt.Printf("\n-------------------------- epoch -> Date------------------------------------\n")

    //time.LoadLocation("Local")
    //loc, _ := time.LoadLocation("America/Los_Angeles")

    loc, _ := time.LoadLocation("Local")

    now := time.Now().In(loc)
    secs := now.Unix() + 32400 // type -> int64
    nanos := now.UnixNano()
    fmt.Println(now)

//Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.

    millis := nanos / 1000000
    fmt.Println(`secs : `, secs , "-----> (+32400) Korea/Seoul") 
    fmt.Println(`millis : `, millis)
    fmt.Println(`nanos : `, nanos)

//You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.

    // Seconds to Date

    dateSecs := time.Unix(secs, 0)

    fmt.Println(dateSecs, "-----> Korea/Seoul")
    //fmt.Println(time.Unix(0, nanos))


    fmt.Printf("Epoch (%v)  --> Date (%v)\n", secs, dateSecs)


}



func epoch2Date(epochSec string) string {
//Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds, respectively.

    epochSecI64, _ := strconv.ParseInt(epochSec, 10, 64)

    fmt.Printf("\n-------------------------- epoch -> Date------------------------------------\n")

    // Seconds to Date
    //fmt.Println(time.Unix(epochSecI64 + int64(32400), 0), "-----> Korea/Seoul")
    //fmt.Println(time.Unix(0, nanos))


    fmt.Printf("Epoch(sec) (%s)  --> Date (%s)\n", epochSec, time.Unix(epochSecI64, 0))

    return string(time.Unix(epochSecI64, 0).String())

}

func date2Epoch(date string) string{

    fmt.Printf("\n--------------------------- date -> Epoch-----------------------------------\n")

    referenceTime := "2006-01-02 15:04:05"  // Fixed,..no need to change this
    //myTime := "2018-12-20 06:50:01"

    myTime := date

    datetime_epoch, _ := time.Parse(referenceTime, myTime)

 

    epoch := datetime_epoch.Unix()


    fmt.Printf("Date (%s)  --> Epoch (%d)\n", myTime, epoch)

    return strconv.FormatInt(epoch, 10)

}


func ParseTime() {

    fmt.Println(`-------------------------------`)


    t, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
    trunc := []time.Duration{
        time.Nanosecond,
        time.Microsecond,
        time.Millisecond,
        time.Second,
        2 * time.Second,
        time.Minute,
        10 * time.Minute,
    }

    for _, d := range trunc {
        fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
    }
    // To round to the last midnight in the local timezone, create a new Date.
    midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
    _ = midnight

}

func testA() {

    fmt.Println(`-------------------------------`)


    start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
    oneDayLater := start.AddDate(0, 0, 1)
    oneMonthLater := start.AddDate(0, 1, 0)
    oneYearLater := start.AddDate(1, 0, 0)

    fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
    fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
    fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)

}



























