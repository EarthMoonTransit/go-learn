# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- NewCache uses the supplied non-nil Feed; Lookup is non-nil. Promoted Count returns the current Feed.Total.
- First Fetch calls Feed.Lookup through Feed.Fetch, increments Feed.Calls, and caches a successful Document under the exact requested key.
- Repeated successful keys return the cached document without another lookup. Different keys and cache instances are independent.
- Errors return a zero Document and the original error; failures are not cached. Successful empty documents are valid.

Public/reference assertions (actual/expected diagnostics):

- Case01: `f:=&subject.Feed{Total:42,Lookup:func(k string)(subject.Document,error){return subject.Document{Key:k,Title:"Title for "+k},nil}};c:=subject.NewCache(f);if c==nil{t.Fatal("nil cache")};d,e:=c.Fetch("a");if e!=nil||d!=(subject.Document{Key:"a",Title:"Title for a"})||f.Calls!=1{t.Errorf("fetch got%v,%v calls%d",d,e,f.Calls)}`
- Case02: `f:=&subject.Feed{Total:42,Lookup:func(k string)(subject.Document,error){return subject.Document{Key:"different",Title:k},nil}};c:=subject.NewCache(f);if c==nil{t.Fatal("nil cache")};for _,k:=range []string{"a","a","A","A"}{d,e:=c.Fetch(k);if e!=nil||d.Title!=k||d.Key!="different"{t.Errorf("key%q got%v,%v",k,d,e)}};if f.Calls!=2{t.Errorf("lookup calls got%d want2",f.Calls)};if c.Count()!=42{t.Errorf("promoted count got%d want42",c.Count())};f.Total=9;if c.Count()!=9{t.Errorf("live count got%d want9",c.Count())};other:=subject.NewCache(f);_,_=other.Fetch("a");if f.Calls!=3{t.Errorf("independent cache calls%d want3",f.Calls)}`
- Case03: `sentinel:=errors.New("missing");attempts:=0;f:=&subject.Feed{Lookup:func(string)(subject.Document,error){attempts++;if attempts<3{return subject.Document{Key:"ignored"},sentinel};return subject.Document{},nil}};c:=subject.NewCache(f);if c==nil{t.Fatal("nil cache")};for i:=0;i<2;i++{d,e:=c.Fetch("");if d!=(subject.Document{})||e!=sentinel{t.Errorf("failure got%v,%v wantzero,original",d,e)}};for i:=0;i<2;i++{d,e:=c.Fetch("");if d!=(subject.Document{})||e!=nil{t.Errorf("empty success got%v,%v",d,e)}};if attempts!=3||f.Calls!=3{t.Errorf("retry/empty cache attempts%d calls%d want3",attempts,f.Calls)}`

Realistic required defect models:

- uncached
- cache-by-document-key
- fixed-count
- error-cached
- partial-error-document
- empty-not-cached
- shared-cache

Retained source behavior: ardan:topics/go/language/embedding/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.

The source Feed.Fetch/CachingFeed.Fetch cache and promoted Count form the retained component. Sleep is replaced by a controlled lookup and call count. Errors and empty-success retry boundaries are explicit extensions. This introduction assesses promotion and override, not ambiguity resolution.
