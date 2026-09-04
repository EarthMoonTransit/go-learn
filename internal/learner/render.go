package learner

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"go-practice-bank/internal/tasking"
)

func renderReviewContext(info tasking.LearnerTaskInfo, progress *TaskProgress, files []tasking.LearnerWorkspaceFile, skills map[string]SkillMetadata, hints []string) string {
	var output strings.Builder
	output.WriteString("# Learner review context\n\n")
	output.WriteString("This package contains learner-visible material only. It does not contain grader tests, reference solutions, mutants, or hidden expectations.\n\n")
	output.WriteString("## Task\n\n")
	output.WriteString(strings.TrimSpace(info.Prompt))
	output.WriteString("\n\n")
	output.WriteString("## Learner code\n\n")
	for _, file := range files {
		fmt.Fprintf(&output, "### `%s`\n\n", file.Path)
		fence := markdownFence(file.Content)
		fmt.Fprintf(&output, "%sgo\n%s", fence, file.Content)
		if len(file.Content) == 0 || file.Content[len(file.Content)-1] != '\n' {
			output.WriteByte('\n')
		}
		fmt.Fprintf(&output, "%s\n\n", fence)
	}
	output.WriteString("## Last public check\n\n")
	if progress.LastPublicCheck == nil {
		output.WriteString("No public check has been run.\n\n")
	} else {
		fmt.Fprintf(&output, "- Time: %s\n- Passed: %t\n- Summary: %s\n", progress.LastPublicCheck.CheckedAt, progress.LastPublicCheck.Passed, progress.LastPublicCheck.Summary)
		if progress.LastPublicCheck.Diagnostic != "" {
			diagnostic := strings.TrimSpace(progress.LastPublicCheck.Diagnostic)
			fence := markdownFence([]byte(diagnostic))
			fmt.Fprintf(&output, "- Public diagnostic:\n\n%stext\n%s\n%s\n", fence, diagnostic, fence)
		}
		output.WriteByte('\n')
	}
	output.WriteString("## Skill metadata\n\n")
	writeExportSkill(&output, "Primary", skills[info.PrimarySkill])
	for _, id := range info.SecondarySkills {
		writeExportSkill(&output, "Secondary", skills[id])
	}
	output.WriteString("## Used hints\n\n")
	if len(progress.UsedHints) == 0 {
		output.WriteString("None.\n")
	} else {
		indexes := append([]int(nil), progress.UsedHints...)
		sort.Ints(indexes)
		for _, index := range indexes {
			if index >= 1 && index <= len(hints) {
				fmt.Fprintf(&output, "%d. %s\n", index, hints[index-1])
			}
		}
	}
	return output.String()
}

func writeExportSkill(output *strings.Builder, role string, skill SkillMetadata) {
	fmt.Fprintf(output, "### %s: `%s`\n\n", role, skill.ID)
	if skill.Definition != "" {
		fmt.Fprintf(output, "%s\n\n", skill.Definition)
	}
	if len(skill.Prerequisites) != 0 {
		fmt.Fprintf(output, "Prerequisites: %s.\n\n", strings.Join(skill.Prerequisites, ", "))
	}
}

func writeSkillReview(output *strings.Builder, label string, skill SkillMetadata) {
	fmt.Fprintf(output, "## %s: `%s`\n\n", label, skill.ID)
	if skill.Definition != "" {
		fmt.Fprintf(output, "%s\n\n", skill.Definition)
	}
	if len(skill.CommonMisconceptions) != 0 {
		output.WriteString("Common misconceptions to revisit:\n\n")
		for _, misconception := range skill.CommonMisconceptions {
			fmt.Fprintf(output, "- %s\n", misconception)
		}
		output.WriteByte('\n')
	}
	if len(skill.Mastery) != 0 {
		output.WriteString("Mastery criteria:\n\n")
		for _, criterion := range skill.Mastery {
			fmt.Fprintf(output, "- %s\n", criterion)
		}
		output.WriteByte('\n')
	}
}

func markdownFence(content []byte) string {
	longest := 2
	current := 0
	for _, character := range content {
		if character == '`' {
			current++
			if current > longest {
				longest = current
			}
		} else {
			current = 0
		}
	}
	return strings.Repeat("`", longest+1)
}

func formatDuration(seconds int64) string {
	if seconds < 0 {
		seconds = 0
	}
	return (time.Duration(seconds) * time.Second).String()
}
