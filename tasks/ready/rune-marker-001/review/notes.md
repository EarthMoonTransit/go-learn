# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Input is valid UTF-8; inspect Unicode code points from left to right.
- Map ❗ to recommendation, 🔍 to search and ☀ to weather.
- When multiple markers occur, return the result for the earliest marker.
- Return default for empty input or text without a recognized marker.

Public/reference assertions (each reports actual/expected values):

- Case01: `if got := subject.Application("❗ recommended product"); got != "recommendation" { 		t.Errorf("Application: got %q, want %q", got, "recommendation") 	}`
- Case02: `if got := subject.Application("executed search 🔍"); got != "search" { 		t.Errorf("Application: got %q, want %q", got, "search") 	}`
- Case03: `if got := subject.Application("forecast: ☀ sunny"); got != "weather" { 		t.Errorf("Application: got %q, want %q", got, "weather") 	}`
- Case04: `if got := subject.Application("☀x❗🔍"); got != "weather" { 		t.Errorf("Application: got %q, want %q", got, "weather") 	}`
- Case05: `if got := subject.Application(""); got != "default" { 		t.Errorf("Application: got %q, want %q", got, "default") 	}`
- Case06: `if got := subject.Application("plain ⚠"); got != "default" { 		t.Errorf("Application: got %q, want %q", got, "default") 	}`

Realistic required defect models:

- fixed-priority
- last-marker
- byte-loop
- empty-default
- wrong-weather

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| rune interpretation and marker mapping | 01–03 | byte-loop; wrong-weather |
| earliest marker wins | 04 | fixed-priority; last-marker |
| empty/default text | 05–06 | empty-default |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.
