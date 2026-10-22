# Format and diagnose a small package

Prerequisite: modules-build-001 (modules/workspaces, then build/run). This manual first encounter assesses toolchain.formatting-vetting.

Copy the supplied subject and test into the approved disposable module. Use gofmt to inspect and apply formatting, then demonstrate an empty second formatting diff. Run go vet, interpret its actual finding, repair the formatting verb without changing the intended `item:<label>` result and rerun vet and tests. Labels are arbitrary strings; preserve them exactly.

Submit before/after files, commands and outputs. Explain which changes only format the source, which repair the API behavior, and why clean vet is not a proof that a program satisfies its tests. In a separate wrong-result specimen demonstrate clean vet plus a failed assertion. Follow PROCEDURE.md exactly; edit only subject.go. Supplied tests/code earn no implementation or testing credit. No personal course progress is created by the reference run.
