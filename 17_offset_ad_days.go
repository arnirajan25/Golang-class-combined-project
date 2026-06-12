package main

import "time"

// Concept: Given a signed number of days from the base AD date (2017-02-11),
// compute the resulting Gregorian date. This is the AD counterpart to the
// offsetBSDays function we studied in Lessons 14-16.
//
// DIFFERENCE FROM offsetBSDays:
//   offsetADDays is MUCH simpler because Go's time.AddDate() handles all
//   the month-length and leap-year logic for us. We don't need a lookup table!
//
// The Go standard library automatically handles:
//   - Different month lengths (28, 29, 30, 31 days)
//   - Leap year calculations
//   - Year/month wraparound (e.g., Dec 31 + 1 day = Jan 1 next year)
//
// STEPS:
//   1. Create base AD date: time.Date(2017, 2, 11, 0, 0, 0, 0, time.UTC)
//   2. Add the day offset:     date.AddDate(0, 0, dayCount)
//      (AddDate adds years, months, days — we only add days)
//   3. Extract and format:     year, month, dayOfWeek, etc.
//
// RESULT STRUCTURE:
//   map[string]interface{}{
//       "year":              2017,
//       "month":             3,
//       "strMonth":          "March",
//       "strShortMonth":     "Mar",
//       "day":               11,
//       "dayOfWeek":         6,         // 0=Sunday, 6=Saturday
//       "strDayOfWeek":      "Saturday",
//       "strShortDayOfWeek": "Sat",
//   }
//
// EXAMPLE:
//   offsetADDays(28) → map containing {year:2017, month:3, day:11, ...}
//   (28 days after 2017-02-11 is 2017-03-11)
// =============================================================================

func offsetADDays(dayCount int) map[string]interface{} {
        // Create the base date
        date := time.Date(baseAD.year, time.Month(baseAD.month), baseAD.day, 0, 0, 0, 0, time.UTC)

        // Add the day offset — Go handles all month/year edge cases
        date = date.AddDate(0, 0, dayCount)

        // Extract date components
        month := int(date.Month())
        dayOfWeek := int(date.Weekday())

        // Build the result map
        return map[string]interface{}{
                "year":              date.Year(),
                "month":             month,
                "strMonth":          engMonthsName[month-1],
                "strShortMonth":     engMonthsShortName[month-1],
                "day":               date.Day(),
                "dayOfWeek":         dayOfWeek,
                "strDayOfWeek":      engDaysName[dayOfWeek],
                "strShortDayOfWeek": engDaysShortName[dayOfWeek],
        }
}
