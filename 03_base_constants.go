package main

// Concept: var declarations at the package level create variables that are
// accessible from ALL files in the same package. These are our "global" data.
//
// Think of these as the fixed starting points from which all conversions
// are calculated. When converting BS ↔ AD, we always measure the distance
// in days from these reference dates.
//
// Why 2017-02-11 AD = 2073-10-29 BS?
// Because on that specific day, both calendars aligned at Saturday (dayOfWeek=6).
// This known correspondence allows us to count days in either direction.
//
// daysInYear = 365 is the standard non-leap-year day count used as a default.

var (
        // baseAD: The reference point in the English (Gregorian) calendar.
        // This is the anchor date — February 11, 2017, which was a Saturday.
        baseAD = dateAD{year: 2017, month: 2, day: 11, dayOfWeek: 6}

        // baseBS: The reference point in the Bikram Sambat (Nepali) calendar.
        // This is the equivalent anchor date — 2073-10-29, also a Saturday.
        baseBS = dateBS{year: 2073, month: 10, day: 29, dayOfWeek: 6}

        // daysInYear: Used as a fallback when a BS year is not found in calendarData.
        // Also serves as the basis for leap year detection:
        //   if daysInYear != countDaysInYear(year) → it's a leap year!
        daysInYear = 365
)
