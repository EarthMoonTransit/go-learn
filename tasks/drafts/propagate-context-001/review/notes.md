# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Call the non-nil operation exactly once with the exact supplied non-nil context. Return its integer and error unchanged, including a nonzero result with an error.
- Pass through cancellation, deadlines and values by preserving context identity. Do not replace, derive or cancel the context.
- Invoke operation even if the context is already canceled; that operation owns its cancellation response.
- Use synchronous composition: return only after the operation returns, and do not launch a background callback.

Public/reference assertions (actual/expected diagnostics):

- Case01: `type key struct{};ctx:=context.WithValue(context.Background(),key{},"v");cause:=errors.New("operation failed");calls:=0;n,e:=subject.CallWithContext(ctx,func(got context.Context)(int,error){calls++;if got!=ctx||got.Value(key{})!="v"{t.Error("context identity/value lost")};return 7,cause});if n!=7||e!=cause||calls!=1{t.Errorf("result got%d,%v calls%d",n,e,calls)}`
- Case02: `ctx,cancel:=context.WithCancel(context.Background());cancel();calls:=0;n,e:=subject.CallWithContext(ctx,func(got context.Context)(int,error){calls++;if got!=ctx||got.Err()!=context.Canceled{t.Error("canceled context lost")};return -2,nil});if n!= -2||e!=nil||calls!=1{t.Errorf("canceled operation got%d,%v calls%d",n,e,calls)}`
- Case03: `ctx,cancel:=context.WithCancel(context.Background());defer cancel();if n,e:=subject.CallWithContext(ctx,func(context.Context)(int,error){return 0,nil});n!=0||e!=nil||ctx.Err()!=nil{t.Error("successful composition changed context") }`

Realistic required defect models:

- background-context
- skip-canceled
- clear-error-result
- discard-error
- double-call

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
