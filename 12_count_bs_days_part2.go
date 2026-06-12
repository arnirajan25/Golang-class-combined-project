package main

// Concept: This lesson explains the core algorithm that counts how many days
// separate two BS dates. It's a "whole years minus excess months ± days" approach.
//
// THE ALGORITHM (running after Lesson 11's setup):
//
//   INPUT: start (earlier date), end (later date), factor (1 or -1)
//
 //┌─ Step A: Add ALL days in ALL years from start.year through end.year
│   for i := start.year; i <= end.year; i++ {
 │       dayCount += calendarData[i][12]  // [12] = total days per year
// │   }
// │   This counts TOO MUCH — we included full years beyond what we need.
// │
// ├─ Step B: Subtract the months BEFORE start.month
 │   for i := 0; i < start.month-1; i++ {
 │       dayCount -= calendarData[start.year][i]
 │   }
// │   e.g., if start is month 4, subtract months 1,2,3 from the total.
// │
// ├─ Step C: Subtract the months AFTER end.month
 │   for i := end.month - 1; i < 12; i++ {
 │       dayCount -= calendarData[end.year][i]
 │   }
// │   e.g., if end is month 8, subtract months 8,9,10,11,12 from the total.
// │
// ├─ Step D: Subtract the days before start.day
 │   dayCount -= start.day
// │
// ├─ Step E: Add the days up to end.day
 │   dayCount += end.day
// │
// └─ Final: Apply the direction factor
     return dayCount * factor
//
// VISUAL EXAMPLE: start=2073/10/29, end=2074/01/15
//
//   Whole years added:      2073 (366) + 2074 (365) = 731
//   Subtract pre-start months of 2073:    months 1-9 = (31+32+31+32+31+30+30+30+29) = 276
//   Subtract post-end months of 2074:     months 1-12 = 365 (all of 2074!)
//   Wait... that doesn't seem right. Let me re-examine.
//
//   Actually, the algorithm works because of how the subtraction interacts.
//   The loop subtracts months from index 0 to start.month-1 in the START year,
//   and months from end.month-1 to 11 in the END year.
//
//   For start=2073/10/29, start.month-1=9 (months 0-8 are subtracted)
//   For end=2074/01/15, end.month-1=0 (months 0-11 are subtracted — all of end year)
//   Then day adjustments: -29 + 15 = -14
//
//   Result: (366+365) - (sum of months 1-9 of 2073) - 365 - 14
//   = 731 - 276 - 365 - 14 = 76
//
//   So there are 76 days between 2073/10/29 and 2074/01/15. ✓
// =============================================================================

// This file explains the algorithm conceptually.
// The actual implementation is merged into Lesson 11 for compilation.
// Students should read Lessons 11 and 12 together.
