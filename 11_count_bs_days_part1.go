package main
import (
        "strconv"
        "strings"
)

// Concept: This function converts a BS date string (e.g., "2074/01/15") into
// a signed number of days relative to our base date (2073/10/29).
//
// This file covers the FIRST HALF: parsing the input and determining direction.
// The second half (the counting loop) is in Lesson 12.
//
// INPUT FORMAT: "YYYY/MM/DD"  (e.g., "2074/01/15")
//
// Step 1 — Parse the string using strings.Split and strconv.Atoi:
//   "2074/01/15" → ["2074", "01", "15"] → year=2074, month=1, day=15
//
// Step 2 — Determine direction (is the target before or after the base?):
//   Compare (year, month, day) against baseBS = {2073, 10, 29}
//   If target > base  →  "inc" (increment) = true,  counting FORWARD  (positive)
//   If target < base  →  "inc" = false,              counting BACKWARD (negative)
//
// This comparison is done one field at a time — a common pattern for comparing
// composite values:
//   1. Compare years first
//   2. If years are equal, compare months
//   3. If months are also equal, compare days
// =============================================================================

func countBSDays(date string) int {
        // Step 1: Parse the date string
        dateParts := strings.Split(date, "/")
        year, _ := strconv.Atoi(dateParts[0])
        month, _ := strconv.Atoi(dateParts[1])
        day, _ := strconv.Atoi(dateParts[2])

        dateObj := dateBS{year: year, month: month, day: day}

        // Step 2: Determine if the target date is after or before the base date
        inc := false
        if year > baseBS.year {
                inc = true // Target year is greater → COUNT FORWARD
        } else if year == baseBS.year && month > baseBS.month {
                inc = true // Same year, target month is greater → COUNT FORWARD
        } else if year == baseBS.year && month == baseBS.month && day > baseBS.day {
                inc = true // Same year & month, target day is greater → COUNT FORWARD
        }

        // Step 3: Set up the range (start, end) and direction (factor)
        var start, end dateBS
        factor := 1 // 1 = forward (positive result), -1 = backward (negative)

        if inc {
                // Count days FROM baseBS TO the target date
                start = baseBS
                end = dateObj
        } else {
                // Count days FROM target TO baseBS, then negate at the end
                start = dateObj
                end = baseBS
                factor = -1
        }

        // (CONTINUED IN LESSON 12: the actual day-counting loop)
        // The variables start, end, and factor are used in the next part.

        dayCount := 0
        for i := start.year; i <= end.year; i++ {
                dayCount += calendarData[i][12]
        }

        for i := 0; i < start.month-1; i++ {
                dayCount -= calendarData[start.year][i]
        }

        for i := end.month - 1; i < 12; i++ {
                dayCount -= calendarData[end.year][i]
        }

        dayCount -= start.day
        dayCount += end.day
        return dayCount * factor
}
