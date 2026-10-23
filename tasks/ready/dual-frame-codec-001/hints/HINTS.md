# Hints

- A matching version and a decodable prefix do not establish a complete frame.
- Keep candidate state separate until every byte has been accepted.
- Test exact literal formats as well as round trips, then mutate storage on each side.
