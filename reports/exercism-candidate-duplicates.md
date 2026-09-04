# Reviewed Exercism candidate duplicates

Reviewed semantic clusters: **20**. Exact hashes, concept sets, public API shapes, and parser-only AST shapes are signals; the recorded decision is a human review and is never inferred from one signal. No upstream code was executed.

Mechanical baseline: **2** shared individual-file hash groups and **0** identical non-empty test-hash groups. Shared file hashes represented below: `56907be466c737285212da71881cd2f11803dd69465ceabb7183778a8ff961c5`, `7943d357f137a7600be90732638e9d6645283585da8016628f7584352b9a77c3`.

| Cluster | Members | Signals | Decision | Preferred | Reviewed conclusion |
|---|---|---|---|---|---|
| `base-parser-family` | `practice/all-your-base`, `practice/binary`, `practice/hexadecimal`, `practice/octal`, `practice/trinary` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `reject-redundant` | — | All members center on positional base conversion; their language mechanisms are incidental and several have broken or weak test contracts. |
| `checksum-validation` | `practice/isbn-verifier`, `practice/luhn` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `prefer-one` | `practice/luhn` | Both candidates scan normalized digits and validate a checksum; keeping both would repeat the same loop-and-validation ability. |
| `cipher-transforms` | `practice/affine-cipher`, `practice/atbash-cipher`, `practice/crypto-square`, `practice/rail-fence-cipher`, `practice/rotational-cipher`, `practice/simple-cipher` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `prefer-one` | `practice/atbash-cipher` | These candidates are dominated by reversible or layout-based cipher rules; only one may remain as idea evidence, not several cosmetic transforms. |
| `concurrent-counting` | `practice/paasio`, `practice/parallel-letter-frequency` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/paasio`, `practice/parallel-letter-frequency` | Both count work under concurrency, but one exercises transparent I/O decoration while the other exercises result ownership and aggregation. |
| `countdown-song` | `practice/beer-song`, `practice/bottle-song` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `reject-redundant` | — | Both expose the same countdown loop, boundary wording, and large literal-output burden without distinct Go mastery. |
| `grid-neighborhood` | `practice/flower-field`, `practice/game-of-life`, `practice/minesweeper` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `prefer-one` | `practice/flower-field`, `practice/game-of-life` | The candidates share neighbor scans over row slices; only the snapshot/aliasing variant contributes a distinct capability beyond boundary handling. |
| `higher-order-collection-operations` | `concept/expenses`, `practice/accumulate` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `merge-as-seed` | `concept/expenses` | Both pass a function over a collection; the concept exercise adds closure construction and offers the richer basis for separated seeds. |
| `lasagna-family` | `concept/lasagna`, `concept/lasagna-master` | `ast-shape`, `concept-set`, `file-hash`, `public-api`, `semantic` | `merge-as-seed` | `concept/lasagna-master` | Shared framing and module boilerplate do not justify separate cosmetic tasks; retain only independently observable function and slice mechanisms. |
| `list-abstraction` | `practice/linked-list`, `practice/list-ops`, `practice/simple-linked-list` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/list-ops`, `practice/simple-linked-list` | All present list-shaped APIs, but function values, singly linked pointer ownership, and bidirectional invariants are separate abilities. |
| `matrix-module-skeleton` | `practice/matrix`, `practice/saddle-points` | `ast-shape`, `concept-set`, `file-hash`, `public-api`, `semantic` | `keep-distinct` | — | The only exact similarity is shared module boilerplate; reviewed source and test semantics are not duplicates, demonstrating why hashes are only signals. |
| `number-theory-drills` | `practice/armstrong-numbers`, `practice/difference-of-squares`, `practice/nth-prime`, `practice/palindrome-products`, `practice/perfect-numbers`, `practice/prime-factors`, `practice/pythagorean-triplet`, `practice/sieve`, `practice/square-root` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `reject-redundant` | — | These candidates repeatedly use loops and arithmetic as incidental tools for mathematical search or classification rather than atomic Go practice. |
| `parser-suite` | `practice/dot-dsl`, `practice/forth`, `practice/markdown`, `practice/ocr-numbers`, `practice/sgf-parsing`, `practice/wordy` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/markdown`, `practice/sgf-parsing` | All parse structured text, but their state machines and outputs differ; only refactoring and decomposed parser evidence is retained. |
| `randomness-statistical` | `concept/animal-magic`, `practice/baffling-birthdays`, `practice/dnd-character` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `reject-redundant` | — | Each relies on repeated random sampling without a deterministic seam; test volume does not repair the same reproducibility defect. |
| `remote-control-car` | `concept/jedliks-toys`, `concept/need-for-speed` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `concept/jedliks-toys`, `concept/need-for-speed` | The state model and operation names are near-identical, but receiver mutation and struct value semantics can remain distinct only with explicit targets. |
| `repetitive-song-builder` | `practice/food-chain`, `practice/house`, `practice/proverb`, `practice/twelve-days`, `practice/two-fer` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `reject-redundant` | — | The group mostly reproduces large or tiny formatted literals and offers little distinct deliberate practice beyond string assembly. |
| `resistor-color-family` | `practice/resistor-color`, `practice/resistor-color-duo`, `practice/resistor-color-trio` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `prefer-one` | `practice/resistor-color` | The larger variants add positional arithmetic and formatting to the same lookup, so only the smallest focused member is preferred. |
| `state-machine-objects` | `practice/bowling`, `practice/hangman`, `practice/react`, `practice/save-the-cow`, `practice/split-second-stopwatch` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/react`, `practice/split-second-stopwatch` | All maintain state transitions, but callback dependency, transition errors, and lifecycle invariants are semantically distinct and require decomposition. |
| `stateful-collections` | `practice/grade-school`, `practice/high-scores`, `practice/kindergarten-garden` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/grade-school`, `practice/high-scores` | Members share method-backed collection state, while ordering, defensive copies, and instance isolation supply different review targets. |
| `time-calculation` | `concept/booking-up-for-beauty`, `practice/gigasecond`, `practice/meetup`, `practice/swift-scheduling` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/gigasecond`, `practice/meetup`, `practice/swift-scheduling` | All use time values, but deterministic clocks, fixed durations, calendar occurrence, and parsing layouts are different abilities. |
| `tree-structures` | `practice/binary-search-tree`, `practice/pov`, `practice/satellite`, `practice/tree-building`, `practice/zipper` | `ast-shape`, `concept-set`, `public-api`, `semantic` | `keep-distinct` | `practice/tree-building`, `practice/zipper` | The common recursive tree shape is not enough for duplication: construction validation, rerooting, reconstruction, and persistent navigation differ. |

## Detailed signal notes

### base-parser-family

- Public API: Parser-only starter declarations: practice/all-your-base=func:ConvertToBase; practice/binary=func:ParseBinary; practice/hexadecimal=func:ParseHex; practice/octal=func:ParseOctal; practice/trinary=func:ParseTrinary
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/all-your-base=f2/t1/i4/l0/r1/s0/c18; practice/binary=f3/t0/i4/l1/r2/s0/c7; practice/hexadecimal=f4/t0/i7/l1/r4/s0/c21; practice/octal=f3/t0/i3/l1/r2/s0/c7; practice/trinary=f3/t0/i1/l1/r2/s1/c7
- Concept sets: reviewed cluster: base-parser-family
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### checksum-validation

- Public API: Parser-only starter declarations: practice/isbn-verifier=func:IsValidISBN; practice/luhn=func:Valid
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/isbn-verifier=f3/t1/i1/l1/r2/s0/c6; practice/luhn=f3/t1/i1/l1/r2/s0/c6
- Concept sets: reviewed cluster: checksum-validation
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### cipher-transforms

- Public API: Parser-only starter declarations: practice/affine-cipher=func:Decode,func:Encode; practice/atbash-cipher=func:Atbash; practice/crypto-square=func:Encode; practice/rail-fence-cipher=func:Decode,func:Encode; practice/rotational-cipher=func:RotationalCipher; practice/simple-cipher=func:NewCaesar,func:NewShift,func:NewVigenere,method:Decode,method:Decode,method:Encode,method:Encode
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/affine-cipher=f7/t1/i8/l1/r3/s0/c16; practice/atbash-cipher=f3/t1/i1/l1/r2/s0/c6; practice/crypto-square=f3/t1/i1/l1/r2/s0/c6; practice/rail-fence-cipher=f6/t1/i2/l2/r4/s0/c12; practice/rotational-cipher=f3/t1/i1/l1/r2/s0/c6; practice/simple-cipher=f20/t5/i11/l9/r14/s0/c68
- Concept sets: reviewed cluster: cipher-transforms
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### concurrent-counting

- Public API: Parser-only starter declarations: practice/paasio=func:NewReadCounter,func:NewReadWriteCounter,func:NewWriteCounter,method:Read,method:ReadCount,method:Write,method:WriteCount; practice/parallel-letter-frequency=func:ConcurrentFrequency,func:Frequency
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/paasio=f28/t7/i19/l0/r6/s0/c125; practice/parallel-letter-frequency=f6/t2/i1/l2/r4/s0/c20
- Concept sets: reviewed cluster: concurrent-counting
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### countdown-song

- Public API: Parser-only starter declarations: practice/beer-song=func:Song,func:Verse,func:Verses; practice/bottle-song=func:Recite
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/beer-song=f8/t0/i10/l2/r3/s0/c21; practice/bottle-song=f4/t2/i4/l0/r1/s1/c11
- Concept sets: reviewed cluster: countdown-song
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### grid-neighborhood

- Public API: Parser-only starter declarations: practice/flower-field=func:Annotate; practice/game-of-life=func:Tick; practice/minesweeper=func:Annotate
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/flower-field=f3/t1/i1/l1/r1/s0/c7; practice/game-of-life=f3/t1/i1/l1/r2/s0/c7; practice/minesweeper=f4/t1/i4/l1/r2/s0/c12
- Concept sets: reviewed cluster: grid-neighborhood
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### higher-order-collection-operations

- Public API: Parser-only starter declarations: concept/expenses=func:ByCategory,func:ByDaysPeriod,func:CategoryExpenses,func:Filter,func:TotalByPeriod; practice/accumulate=func:Accumulate
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): concept/expenses=f9/t2/i9/l0/r6/s0/c31; practice/accumulate=f7/t0/i5/l1/r3/s1/c18
- Concept sets: reviewed cluster: higher-order-collection-operations
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### lasagna-family

- Public API: Parser-only starter declarations: concept/lasagna=func:ElapsedTime,func:PreparationTime,func:RemainingOvenTime; concept/lasagna-master=no-starter-function
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): concept/lasagna=f7/t1/i4/l0/r4/s0/c14; concept/lasagna-master=f4/t4/i8/l0/r5/s0/c33
- Concept sets: reviewed cluster: lasagna-family
- Shared file hashes: 7943d357f137a7600be90732638e9d6645283585da8016628f7584352b9a77c3
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### list-abstraction

- Public API: Parser-only starter declarations: practice/linked-list=func:NewList,method:Count,method:Delete,method:First,method:Last,method:Next,method:Pop,method:Prev,method:Push,method:Reverse,method:Shift,method:Unshift; practice/list-ops=method:Append,method:Concat,method:Filter,method:Foldl,method:Foldr,method:Length,method:Map,method:Reverse; practice/simple-linked-list=func:New,method:Array,method:Peek,method:Pop,method:Push,method:Reverse,method:Size
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/linked-list=f22/t6/i14/l3/r6/s2/c133; practice/list-ops=f19/t9/i8/l0/r10/s0/c70; practice/simple-linked-list=f18/t2/i6/l7/r8/s3/c62
- Concept sets: reviewed cluster: list-abstraction
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### matrix-module-skeleton

- Public API: Parser-only starter declarations: practice/matrix=func:New,method:Cols,method:Rows,method:Set; practice/saddle-points=func:New,method:Saddle
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/matrix=f12/t1/i26/l5/r6/s0/c73; practice/saddle-points=f6/t1/i5/l1/r7/s0/c24
- Concept sets: reviewed cluster: matrix-module-skeleton
- Shared file hashes: 56907be466c737285212da71881cd2f11803dd69465ceabb7183778a8ff961c5
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### number-theory-drills

- Public API: Parser-only starter declarations: practice/armstrong-numbers=func:IsNumber; practice/difference-of-squares=func:Difference,func:SquareOfSum,func:SumOfSquares; practice/nth-prime=func:Nth; practice/palindrome-products=func:Products; practice/perfect-numbers=func:Classify; practice/prime-factors=func:Factors; practice/pythagorean-triplet=func:Range,func:Sum; practice/sieve=func:Sieve; practice/square-root=func:SquareRoot
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/armstrong-numbers=f3/t1/i1/l1/r2/s0/c6; practice/difference-of-squares=f7/t1/i1/l3/r1/s0/c12; practice/nth-prime=f3/t1/i1/l1/r1/s1/c8; practice/palindrome-products=f4/t3/i8/l1/r2/s0/c22; practice/perfect-numbers=f5/t1/i3/l2/r3/s1/c12; practice/prime-factors=f3/t1/i1/l1/r2/s0/c9; practice/pythagorean-triplet=f6/t3/i2/l2/r2/s0/c14; practice/sieve=f3/t1/i2/l1/r2/s0/c9; practice/square-root=f3/t1/i2/l1/r2/s0/c7
- Concept sets: reviewed cluster: number-theory-drills
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### parser-suite

- Public API: Parser-only starter declarations: practice/dot-dsl=func:Parse; practice/forth=func:Forth; practice/markdown=func:Render; practice/ocr-numbers=func:Recognize; practice/sgf-parsing=func:Parse; practice/wordy=func:Answer
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/dot-dsl=f4/t3/i8/l0/r1/s0/c21; practice/forth=f3/t1/i4/l1/r2/s0/c27; practice/markdown=f3/t1/i14/l3/r2/s1/c22; practice/ocr-numbers=f2/t1/i4/l0/r1/s0/c8; practice/sgf-parsing=f3/t2/i4/l1/r2/s0/c9; practice/wordy=f3/t1/i1/l1/r2/s1/c8
- Concept sets: reviewed cluster: parser-suite
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### randomness-statistical

- Public API: Parser-only starter declarations: concept/animal-magic=func:GenerateWandEnergy,func:RollADie,func:ShuffleAnimals; practice/baffling-birthdays=func:EstimatedProbability,func:RandomBirthdates,func:SharedBirthday; practice/dnd-character=func:Ability,func:GenerateCharacter,func:Modifier
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): concept/animal-magic=f7/t0/i14/l0/r5/s0/c28; practice/baffling-birthdays=f12/t2/i9/l2/r11/s0/c41; practice/dnd-character=f11/t3/i4/l2/r4/s0/c26
- Concept sets: reviewed cluster: randomness-statistical
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### remote-control-car

- Public API: Parser-only starter declarations: concept/jedliks-toys=func:NewCar; concept/need-for-speed=func:CanFinish,func:Drive,func:NewCar,func:NewTrack
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): concept/jedliks-toys=f5/t1/i4/l0/r4/s0/c12; concept/need-for-speed=f8/t0/i4/l0/r4/s0/c16
- Concept sets: reviewed cluster: remote-control-car
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### repetitive-song-builder

- Public API: Parser-only starter declarations: practice/food-chain=func:Song,func:Verse,func:Verses; practice/house=func:Song,func:Verse; practice/proverb=func:Proverb; practice/twelve-days=func:Song,func:Verse; practice/two-fer=func:ShareWith
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/food-chain=f8/t2/i4/l2/r2/s1/c31; practice/house=f6/t1/i5/l2/r3/s0/c20; practice/proverb=f3/t1/i1/l1/r2/s0/c8; practice/twelve-days=f7/t1/i3/l3/r2/s1/c25; practice/two-fer=f3/t1/i1/l1/r2/s0/c5
- Concept sets: reviewed cluster: repetitive-song-builder
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### resistor-color-family

- Public API: Parser-only starter declarations: practice/resistor-color=func:ColorCode,func:Colors; practice/resistor-color-duo=func:Value; practice/resistor-color-trio=func:Label
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/resistor-color=f8/t2/i5/l2/r5/s0/c21; practice/resistor-color-duo=f3/t1/i1/l1/r2/s0/c6; practice/resistor-color-trio=f3/t1/i1/l1/r2/s0/c6
- Concept sets: reviewed cluster: resistor-color-family
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### state-machine-objects

- Public API: Parser-only starter declarations: practice/bowling=func:NewGame,method:Roll,method:Score; practice/hangman=func:NewGame,method:Guess,method:MaskedWord,method:RemainingGuesses,method:State; practice/react=func:New,method:AddCallback,method:Cancel,method:CreateCompute1,method:CreateCompute2,method:CreateInput,method:SetValue,method:Value; practice/save-the-cow=func:NewGame,method:Guess,method:MaskedWord,method:RemainingGuesses,method:State; practice/split-second-stopwatch=func:NewSplitSecondStopwatch,method:AdvanceTime,method:CurrentLap,method:Lap,method:PreviousLaps,method:Reset,method:Start,method:State,method:Stop,method:Total
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/bowling=f6/t2/i6/l0/r3/s1/c19; practice/hangman=f7/t2/i8/l1/r4/s0/c27; practice/react=f16/t9/i6/l1/r10/s4/c51; practice/save-the-cow=f7/t2/i8/l1/r4/s0/c27; practice/split-second-stopwatch=f14/t4/i8/l1/r4/s1/c45
- Concept sets: reviewed cluster: state-machine-objects
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### stateful-collections

- Public API: Parser-only starter declarations: practice/grade-school=func:New,method:Add,method:Enrollment,method:Grade; practice/high-scores=func:NewHighScores,method:Latest,method:PersonalBest,method:Scores,method:TopThree; practice/kindergarten-garden=func:NewGarden,method:Plants
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/grade-school=f12/t5/i3/l7/r16/s0/c64; practice/high-scores=f8/t3/i2/l1/r4/s0/c35; practice/kindergarten-garden=f9/t2/i5/l2/r5/s3/c37
- Concept sets: reviewed cluster: stateful-collections
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### time-calculation

- Public API: Parser-only starter declarations: concept/booking-up-for-beauty=func:AnniversaryDate,func:Description,func:HasPassed,func:IsAfternoonAppointment,func:Schedule; practice/gigasecond=func:AddGigasecond; practice/meetup=func:Day; practice/swift-scheduling=func:DeliveryDate
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): concept/booking-up-for-beauty=f10/t5/i5/l0/r5/s0/c28; practice/gigasecond=f4/t1/i3/l1/r1/s0/c11; practice/meetup=f3/t1/i1/l1/r2/s0/c6; practice/swift-scheduling=f3/t1/i1/l1/r2/s0/c6
- Concept sets: reviewed cluster: time-calculation
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.

### tree-structures

- Public API: Parser-only starter declarations: practice/binary-search-tree=func:NewBst,method:Insert,method:SortedData; practice/pov=func:New,method:Children,method:FromPov,method:PathTo,method:String,method:Value; practice/satellite=func:TreeFromTraversals; practice/tree-building=func:Build; practice/zipper=func:NewZipper,method:Left,method:Right,method:SetLeft,method:SetRight,method:SetValue,method:ToTree,method:Up,method:Value
- AST shape: Parser-only starter/public-test node counts (func/type/if/for/range/switch/call): practice/binary-search-tree=f11/t3/i15/l2/r4/s1/c55; practice/pov=f15/t3/i10/l2/r7/s1/c187; practice/satellite=f5/t2/i7/l1/r2/s0/c19; practice/tree-building=f11/t3/i5/l3/r5/s1/c29; practice/zipper=f17/t5/i15/l1/r3/s1/c70
- Concept sets: reviewed cluster: tree-structures
- Shared file hashes: none
- Shared test hashes: none
- Rationale: Independent reviewers compared API purpose, concept assignments, edge partitions, and likely solution mechanisms before recording this versioned decision.
