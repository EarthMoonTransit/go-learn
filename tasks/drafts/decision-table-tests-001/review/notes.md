# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Weights: ace=11; two through nine=2 through 9; ten, jack, queen, king=10; every other exact token=0.
- Sum the first two weights. At 22 return P. At 21 return S if the visible weight is at least 10, otherwise W.
- For sums 17..20 return S. For sums 12..16 return S when the visible weight is below 7, otherwise H. For sums below 12 return H.
- Codes are literal uppercase strings. Token matching is case-sensitive and preserves whitespace; unknown and empty tokens are allowed.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.SelectAction("ace","ace","jack");got!="P"{t.Errorf("pair got%q wantP",got)}`
- Case02: `for _,c:=range []struct{a,b,v,want string}{{"king","ace","queen","S"},{"ace","ten","five","W"},{"ace","king","nine","W"},{"ten","king","ace","S"},{"seven","king","ace","S"},{"ten","six","six","S"},{"ten","six","seven","H"},{"ten","two","six","S"},{"ten","two","seven","H"},{"nine","two","two","H"},{"two","two","five","H"}}{if got:=subject.SelectAction(c.a,c.b,c.v);got!=c.want{t.Errorf("decision %s/%s/%s got%q want%q",c.a,c.b,c.v,got,c.want)}}`
- Case03: `for _,c:=range []struct{a,b,v,want string}{{"Ace","king","two","H"},{"ace ","ten","two","H"},{"","ace","","H"},{"ace","ten","king ","W"},{"ten","two","unknown","S"},{"ten","two","Seven","S"}}{if got:=subject.SelectAction(c.a,c.b,c.v);got!=c.want{t.Errorf("exact token %q/%q/%q got%q want%q",c.a,c.b,c.v,got,c.want)}}`

Realistic required defect models:

- pair-after-range
- win-against-ten
- stand-at-sixteen
- exclude-twelve
- visible-seven-low
- normalize-token
- unknown-is-ten

Retained source behavior: blackjack selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
