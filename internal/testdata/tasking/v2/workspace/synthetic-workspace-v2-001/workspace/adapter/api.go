package adapter

import "task.local/synthetic-workspace-v2-001/core"

func Transform(value int) int { return core.Apply(value) }
