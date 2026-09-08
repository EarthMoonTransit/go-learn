# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- In values.go declare doubled before base. Initialize base with record("base",3) and doubled with record("double",base*2).
- Add an init function calling record("ready",0) after variable initialization. Do not call record from Values.
- Values returns base and doubled: 3 and 6. Supplied Trace must observe base, double, ready exactly once.
- The supplied record helper returns its value unchanged. Calls to Values and Trace do not change initialization state.

Public/reference assertions (actual/expected diagnostics):

- Case01: `a,b:=subject.Values();if a!=3||b!=6{t.Errorf("values got%d,%d want3,6",a,b)};if got:=subject.Trace();!slices.Equal(got,[]string{"base","double","ready"}){t.Errorf("init order got%v want[base double ready]",got)}`
- Case02: `subject.Values();subject.Values();got:=subject.Trace();if !slices.Equal(got,[]string{"base","double","ready"}){t.Errorf("repeated access changed init trace%v",got)};got[0]="changed";if !slices.Equal(subject.Trace(),[]string{"base","double","ready"}){t.Error("trace accessor shares storage")}`

Realistic required defect models:

- no-dependency
- wrong-base
- missing-init
- repeat-record
- reverse-values

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
