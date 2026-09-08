# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Sum Amount for exact matching Category and Day within inclusive From..To.
- If the category occurs anywhere in input, return the sum and nil even when no matching day lies in the period or From exceeds To.
- If no input record has that category, return zero and the original non-nil ErrCategory sentinel; preserve the sentinel and every input record.
- Empty categories, negative and zero amounts are valid; finite amounts have magnitude <=1000 and at most 100 records. Allow absolute sum error 0.000000001.

Public/reference assertions (actual/expected diagnostics):

- Case01: `sentinel:=subject.ErrCategory;if sentinel==nil{t.Fatal("nil sentinel")};defer func(){if subject.ErrCategory!=sentinel{t.Error("sentinel changed")}}();in:=[]subject.Record{{Day:1,Amount:2.5,Category:"a"},{Day:2,Amount:-1,Category:"a"},{Day:3,Amount:9,Category:"b"},{Day:4,Amount:3,Category:"a"}};before:=slices.Clone(in);n,e:=subject.CategoryTotal(in,subject.Period{From:1,To:2},"a");if e!=nil||math.IsNaN(n)||math.Abs(n-1.5)>1e-9{t.Errorf("total got %g,%v want1.5,nil",n,e)};if !slices.Equal(in,before){t.Errorf("input changed %v",in)}`
- Case02: `for _,p:=range []subject.Period{{From:9,To:10},{From:3,To:1}}{n,e:=subject.CategoryTotal([]subject.Record{{Day:2,Amount:8,Category:"a"}},p,"a");if n!=0||e!=nil{t.Errorf("known outside got %g,%v want0,nil",n,e)}}`
- Case03: `for _,in:=range [][]subject.Record{nil,{{Category:"A",Amount:5}}}{sentinel:=subject.ErrCategory;n,e:=subject.CategoryTotal(in,subject.Period{},"a");if sentinel==nil||e!=sentinel||n!=0||subject.ErrCategory!=sentinel{t.Errorf("unknown got %g,%v want original sentinel",n,e)}}`
- Case04: `n,e:=subject.CategoryTotal([]subject.Record{{Day:0,Amount:0,Category:""}},subject.Period{},"");if n!=0||e!=nil{t.Errorf("empty category got %g,%v want0,nil",n,e)}`

Realistic required defect models:

- known-only-in-period
- exclusive-boundaries
- no-category-filter
- fresh-error
- absolute-amount
- empty-category-rejected
- mutate-input

Retained source behavior: expenses selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
