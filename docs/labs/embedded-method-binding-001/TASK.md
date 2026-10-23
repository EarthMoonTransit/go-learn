# Embedded live bindings

Counter has Read on its value receiver and Add on its pointer receiver. ValueBox embeds Counter; PointerBox embeds *Counter. Left and Right each expose Label, and Joined embeds both. Preserve these types and methods. Implement Bind(v *ValueBox, p PointerBox) so its two Reader/Adder pairs observe and change the original counters. Neither pair may own a copied counter. Inputs and p.Counter are non-nil and refer to different counters. Implement Joined.Label by explicitly combining its left and right labels with one vertical bar; do not replace either embedded value.

Before executing the compiler probes, predict which value, pointer and promoted method sets satisfy each interface. Explain why an addressable v.Add call does not make ValueBox implement Adder. Predict the ambiguous Joined.Label outcome before your method and compare it with explicit Left.Label selection. Retain actual diagnostics, then demonstrate your repaired bindings with unchanged public and grader checks. Valid results require both useful state behavior and the method-set explanation.

