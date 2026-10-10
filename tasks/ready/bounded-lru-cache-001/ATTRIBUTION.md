# Attribution

Substantive adaptation of SHAD lrucache, registered as `shad-go` in
`sources/registry.yaml`, at commit `6718601587ec77cfaa632bbd29a8ca0716fe81ec`:
`lrucache/README.md`, `lrucache/cache.go`, `lrucache/lru.go`,
`lrucache/cache_test.go`. Exact file hashes are in provenance.yaml and the P3
artifact review. Repository: https://github.com/slon/shad-go.
The source module identifies `gitlab.com/slon/shad-go`; this historical identity
is retained without treating unrelated GitLab content as an additional source.

The source integer cache's observable cross-operation contracts are adapted.
The generic K/V API and usable zero Cache are local changes (ADR0024); the
source did not supply a generic implementation.
The statement, deterministic test data, defect implementations, hints and linear
slice reference were authored for Go Learn. A missing lookup's typed zero value,
nonnegative capacities, sequential calls and non-mutating callbacks are explicit
local boundaries. No complexity threshold is imposed. No source dependencies,
random workloads, GC settings, benchmarks, Wikipedia text or canonical Exercism
material were imported. The source's Wikipedia link is background reading only;
no article content was used. Root LICENSE is copied solely to preserve the
required MIT notice. No additional notices were found in the four admitted files.
Go Learn changes are Apache-2.0 subject to the retained upstream notice below.

MIT License

Copyright (c) 2023 Fedor Korotkiy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
