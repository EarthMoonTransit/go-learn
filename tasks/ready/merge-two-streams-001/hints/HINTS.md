# Hints

1. Closing one input does not mean the other input has finished.
2. A receive reports both a value and whether it came from an open channel. A nil channel case cannot proceed.
3. Keep consuming the remaining input after disabling the closed one; finish only after both are disabled.
