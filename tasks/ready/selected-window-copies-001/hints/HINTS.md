# Hints

Count retained requests before choosing the outer slice length and capacity. A capped header can still point into a large allocation. Give every selected interval its own storage, including repeated and overlapping intervals.
