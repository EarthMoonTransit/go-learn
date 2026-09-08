# Progressive hints

1. ScanLines can return a final token at EOF.
2. Check Err after Scan stops.
3. An io.Closer capability does not authorize the scanner to close the reader.
