# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Complete stock/item.go and subject/describe.go. New stores Name and Weight with both private counts zero.
- Each setter changes only its matching count; getters return current counts. Inputs are nonnegative integers.
- Describe returns "Name:Weight:OnHand:Sold" with decimal integers, using stock exports and methods.
- Preserve Name exactly, including empty text and spaces; separate Item values have independent state.

Public/reference assertions (actual/expected diagnostics):

- Case01: `item:=stock.New("part",4);if item.Name!="part"||item.Weight!=4||item.OnHand()!=0||item.Sold()!=0{t.Errorf("new got%+v counts%d,%d",item,item.OnHand(),item.Sold())};item.SetOnHand(8);item.SetSold(3);if got:=subject.Describe(item);got!="part:4:8:3"{t.Errorf("summary got%q wantpart:4:8:3",got)}`
- Case02: `item:=stock.New(" x ",0);other:=stock.New("",2);item.SetOnHand(9);item.SetSold(2);item.SetSold(0);if item.OnHand()!=9||item.Sold()!=0||item.Name!=" x "||item.Weight!=0{t.Errorf("set sold changed other state %+v",item)};item.SetOnHand(0);other.SetSold(7);if got:=subject.Describe(item);got!=" x :0:0:0"{t.Errorf("first got%q",got)};if got:=subject.Describe(other);got!=":2:0:7"{t.Errorf("independent got%q",got)}`

Realistic required defect models:

- missing-fields
- reset-onhand
- swapped-accessors
- copy-only-setter
- wrong-summary

Retained source behavior: ardan:topics/go/language/exporting/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.

The source Toy public Name/Weight, private onHand/sold, New and accessors are retained as Item in a separate stock package used from subject. Both implementation files are learner editable through the existing workspace-tests profile. No struct-tag interpretation is claimed by this introduction.
