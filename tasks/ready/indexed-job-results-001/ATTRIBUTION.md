# Attribution

This task substantively adapts the finite fan-out/all-results operation from
Ardan Labs Go Training, registry `ardanlabs-gotraining`, commit
`4ae5a46bfc665909a7cb29e11c18775330c63401`, path
`topics/go/concurrency/channels/exercises/template2/template2.go`.
Source selection: PR85, merge `b20ca9feacaa30b7a2ad20ca2761c4ee14c8c0f8`,
`catalog/intake/2026-10-18/p17-p19-source-review.json`.

William Kennedy introduced the fan-out behavior; Jacob Walker added slice
collection and refined its comments. The selected template has no completed
implementation. The inherited operation starts a finite group and collects all
results. Local modifications replace random generation/printing with supplied
jobs and indexed records carrying exact values/errors. The function API, source
and test cases, implementations, controlled-gate checks and defects are local.
No upstream answer implementation or test fixture is copied. No other template
or reference-only SHAD material is used as task content.

Preserved source notice:

> All material is licensed under the Apache License Version 2.0, January 2004
> http://www.apache.org/licenses/LICENSE-2.0

Local additions are Apache-2.0 under project policy. The selected ancestors have
no additional NOTICE or license override. The full upstream LICENSE follows.

Copyright 2014 Ardan Studios

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
