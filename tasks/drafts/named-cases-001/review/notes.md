# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- For each Case, create one real t.Run subtest using its Name and invoke Check once with that child *testing.T.
- Run cases serially in input order. Empty input creates no subtests. Names are unique nonempty ASCII letters/digits/hyphens; callbacks are non-nil.
- Preserve the input table. Do not invoke callbacks with the parent, change names, duplicate calls or mark subtests parallel.
- The provided public callbacks inspect their real test names and identities. Callback failures use normal testing package propagation.

Public/reference assertions (actual/expected diagnostics):

- Case01: `var calls []string;t.Run("group",func(parent *testing.T){cases:=[]subject.Case{{Name:"first",Check:func(child *testing.T){if child==parent||child.Name()!=parent.Name()+"/first"{t.Errorf("child first name%q",child.Name())};calls=append(calls,"first")}},{Name:"second",Check:func(child *testing.T){if child==parent||child.Name()!=parent.Name()+"/second"{t.Errorf("child second name%q",child.Name())};calls=append(calls,"second")}}};subject.RunNamed(parent,cases);if !slices.Equal(calls,[]string{"first","second"})||cases[0].Name!="first"||cases[1].Name!="second"{t.Errorf("serial cases calls%v names%s,%s",calls,cases[0].Name,cases[1].Name)}})`
- Case02: `subject.RunNamed(t,nil);subject.RunNamed(t,[]subject.Case{});calls:=0;t.Run("single",func(parent *testing.T){subject.RunNamed(parent,[]subject.Case{{Name:"case-1",Check:func(child *testing.T){calls++;if child.Name()!=parent.Name()+"/case-1"{t.Errorf("single name%q",child.Name())}}}})});if calls!=1{t.Errorf("single calls%d want1",calls)}`

Realistic required defect models:

- parent-callback
- constant-name
- skip-last
- reverse-order
- double-check

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.

Slot 119 changes from planned test-writing to implementation of a reusable test runner. Real child *testing.T identities/names and serial callback ordering are observable through the existing fixed-tests profile; no unsupported syntax-only mutation requirement is added. C13 deliberately contains four implementation, four test-writing and two debugging tasks.
