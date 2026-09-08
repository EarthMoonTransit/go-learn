package learner

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"text/tabwriter"

	"go-practice-bank/internal/tasking"
)

const cliUsage = `usage: learn [-root DIR] [-home DIR] <command> [arguments]

commands:
  ui [-port PORT]
  list
  show <task-id>
  start <task-id>
  resume <task-id>
  pause
  check
  complete
  hint
  review
  solution --reveal
  next
  roadmap
  organize [--dry-run]
  progress
  feedback <text>
  reset-task [task-id] --yes
  export-review-context
`

func RunCLI(config Config, arguments []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("learn", flag.ContinueOnError)
	flags.SetOutput(stderr)
	root := config.Root
	home := config.Home
	flags.StringVar(&root, "root", root, "repository root")
	flags.StringVar(&home, "home", home, "local learner state directory")
	flags.Usage = func() { _, _ = io.WriteString(stderr, cliUsage) }
	if err := flags.Parse(arguments); err != nil {
		return 2
	}
	arguments = flags.Args()
	if len(arguments) == 0 {
		_, _ = io.WriteString(stderr, cliUsage)
		return 2
	}
	service, err := New(Config{Root: root, Home: home, Now: config.Now})
	if err != nil {
		return cliError(stderr, err)
	}
	command, commandArguments := arguments[0], arguments[1:]
	switch command {
	case "ui":
		uiFlags := flag.NewFlagSet("learn ui", flag.ContinueOnError)
		uiFlags.SetOutput(stderr)
		port := uiFlags.Int("port", 0, "local UI port (0 selects an available port)")
		if err := uiFlags.Parse(commandArguments); err != nil {
			return 2
		}
		if err := noArguments("ui", uiFlags.Args()); err != nil {
			return cliError(stderr, err)
		}
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		if err := ServeUI(ctx, service, *port, stdout); err != nil {
			return cliError(stderr, err)
		}
	case "list":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		tasks, progress, err := service.ListTasks()
		if err != nil {
			return cliError(stderr, err)
		}
		route, err := loadRoute(service.root)
		if err != nil {
			return cliError(stderr, err)
		}
		if route != nil && route.SchemaVersion == 2 {
			writeNumberedTaskList(stdout, *route, tasks, progress)
		} else {
			writeTaskList(stdout, tasks, progress)
		}
	case "show":
		id, err := oneArgument(command, commandArguments)
		if err != nil {
			return cliError(stderr, err)
		}
		info, err := service.Show(id)
		if err != nil {
			return cliError(stderr, err)
		}
		number, err := service.LessonNumber(id)
		if err != nil {
			return cliError(stderr, err)
		}
		if number != "" {
			fmt.Fprintf(stdout, "Lesson: %s\n", number)
		}
		writeTask(stdout, info)
	case "start":
		id, err := oneArgument(command, commandArguments)
		if err != nil {
			return cliError(stderr, err)
		}
		workspace, err := service.Start(id)
		if err != nil {
			return cliError(stderr, err)
		}
		fmt.Fprintf(stdout, "Started %s\nWorkspace: %s\n", id, workspace)
	case "resume":
		id, err := oneArgument(command, commandArguments)
		if err != nil {
			return cliError(stderr, err)
		}
		workspace, err := service.Resume(id)
		if err != nil {
			return cliError(stderr, err)
		}
		fmt.Fprintf(stdout, "Resumed %s\nWorkspace: %s\n", id, workspace)
	case "pause":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		if err := service.Pause(); err != nil {
			return cliError(stderr, err)
		}
		fmt.Fprintln(stdout, "Work session paused. Use resume <task-id> to continue.")
	case "check":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		result, err := service.Check()
		if err != nil {
			return cliError(stderr, err)
		}
		writeVerification(stdout, result, true)
		if !result.Passed {
			return 1
		}
	case "complete":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		result, err := service.Complete()
		if err != nil {
			return cliError(stderr, err)
		}
		writeVerification(stdout, result, false)
		if !result.Passed {
			return 1
		}
	case "hint":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		index, hint, err := service.Hint()
		if err != nil {
			return cliError(stderr, err)
		}
		fmt.Fprintf(stdout, "Hint %d: %s\n", index, hint)
	case "review":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		review, err := service.Review()
		if err != nil {
			return cliError(stderr, err)
		}
		_, _ = io.WriteString(stdout, review)
	case "solution":
		reveal, err := explicitFlag(command, commandArguments, "--reveal")
		if err != nil {
			return cliError(stderr, err)
		}
		files, err := service.Solution(reveal)
		if err != nil {
			return cliError(stderr, err)
		}
		writeSolution(stdout, files)
	case "organize":
		organizeFlags := flag.NewFlagSet("learn organize", flag.ContinueOnError)
		organizeFlags.SetOutput(stderr)
		dry := organizeFlags.Bool("dry-run", false, "preview organization without changing files")
		if err := organizeFlags.Parse(commandArguments); err != nil {
			return 2
		}
		if err := noArguments(command, organizeFlags.Args()); err != nil {
			return cliError(stderr, err)
		}
		plan, err := service.Organize(*dry)
		if err != nil {
			return cliError(stderr, err)
		}
		fmt.Fprintf(stdout, "Workspace organization: progress v%d → v%d; %d moves\n", plan.FromVersion, plan.ToVersion, len(plan.Moves))
		for _, move := range plan.Moves {
			fmt.Fprintf(stdout, "%s: %s → %s\n", move.TaskID, move.From, move.To)
		}
		if plan.Backup != "" {
			fmt.Fprintf(stdout, "Backup: %s\n", plan.Backup)
		}
		if *dry {
			fmt.Fprintln(stdout, "Preview only; no files changed.")
		} else {
			fmt.Fprintln(stdout, "Organization complete. Open the new workspace path in GoLand.")
		}
	case "roadmap":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		roadmap, err := service.Roadmap()
		if err != nil {
			return cliError(stderr, err)
		}
		writeRoadmap(stdout, roadmap)
	case "next":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		choice, err := service.Next()
		if err != nil {
			return cliError(stderr, err)
		}
		writeNext(stdout, choice)
		if choice.ReviewSkill != "" {
			fmt.Fprintf(stdout, "Delayed review: %s\n", choice.ReviewSkill)
		}
	case "progress":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		progress, err := service.Progress()
		if err != nil {
			return cliError(stderr, err)
		}
		writeProgress(stdout, progress)
		writeLearning(stdout, progress)
	case "feedback":
		if len(commandArguments) == 0 {
			return cliError(stderr, errors.New("feedback requires text"))
		}
		if err := service.Feedback(strings.Join(commandArguments, " ")); err != nil {
			return cliError(stderr, err)
		}
		_, _ = io.WriteString(stdout, "Feedback stored locally.\n")
	case "reset-task":
		id, confirmed, err := resetArguments(commandArguments)
		if err != nil {
			return cliError(stderr, err)
		}
		if err := service.ResetTask(id, confirmed); err != nil {
			return cliError(stderr, err)
		}
		_, _ = io.WriteString(stdout, "Local task workspace and progress removed.\n")
	case "export-review-context":
		if err := noArguments(command, commandArguments); err != nil {
			return cliError(stderr, err)
		}
		path, err := service.ExportReviewContext()
		if err != nil {
			return cliError(stderr, err)
		}
		fmt.Fprintf(stdout, "Review context: %s\n", path)
	case "help", "-h", "--help":
		_, _ = io.WriteString(stdout, cliUsage)
	default:
		fmt.Fprintf(stderr, "learn: unknown command %q\n", command)
		_, _ = io.WriteString(stderr, cliUsage)
		return 2
	}
	return 0
}

func writeTaskList(output io.Writer, tasks []tasking.LearnerTaskInfo, progress Progress) {
	writer := tabwriter.NewWriter(output, 0, 4, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tSTATUS\tACTIVITY\tDIFFICULTY\tPRIMARY SKILL")
	for _, task := range tasks {
		status := "not-started"
		if item := progress.Tasks[task.ID]; item != nil {
			status = item.Status
		}
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", task.ID, status, task.Activity, task.Difficulty, task.PrimarySkill)
	}
	_ = writer.Flush()
}

func writeTask(output io.Writer, info tasking.LearnerTaskInfo) {
	fmt.Fprintf(output, "ID: %s\nTitle: %s\nActivity: %s\nSubmission: %s / %s\nDifficulty: %s\nEstimated time: %d-%d minutes\nPrimary skill: %s\n", info.ID, info.Title, info.Activity, info.SubmissionKind, info.VerificationProfile, info.Difficulty, info.EstimatedTimeMinutes.Min, info.EstimatedTimeMinutes.Max, info.PrimarySkill)
	if len(info.SecondarySkills) != 0 {
		fmt.Fprintf(output, "Secondary skills: %s\n", strings.Join(info.SecondarySkills, ", "))
	}
	if len(info.Prerequisites) != 0 {
		fmt.Fprintf(output, "Prerequisites: %s\n", strings.Join(info.Prerequisites, ", "))
	}
	fmt.Fprintf(output, "\n%s\n", strings.TrimSpace(info.Prompt))
}

func writeVerification(output io.Writer, result tasking.LearnerVerificationResult, public bool) {
	fmt.Fprintln(output, result.Summary)
	if public && result.Diagnostic != "" {
		fmt.Fprintln(output)
		fmt.Fprintln(output, strings.TrimSpace(result.Diagnostic))
	}
}

func writeSolution(output io.Writer, files []tasking.LearnerWorkspaceFile) {
	_, _ = io.WriteString(output, "Solution revealed. This does not complete the task.\n\n")
	for _, file := range files {
		fence := markdownFence(file.Content)
		fmt.Fprintf(output, "## %s\n\n%sgo\n%s", file.Path, fence, file.Content)
		if len(file.Content) == 0 || file.Content[len(file.Content)-1] != '\n' {
			fmt.Fprintln(output)
		}
		fmt.Fprintf(output, "%s\n\n", fence)
	}
}

func writeNext(output io.Writer, choice SchedulerChoice) {
	if choice.RouteBlock != "" {
		fmt.Fprintf(output, "Route: %s (%s)\n", choice.RouteBlock, choice.RouteRole)
	}
	if choice.LessonNumber != "" {
		fmt.Fprintf(output, "%s — %s\nID: %s\n", choice.LessonNumber, choice.Task.Title, choice.Task.ID)
	} else {
		fmt.Fprintf(output, "%s — %s\n", choice.Task.ID, choice.Task.Title)
	}
	fmt.Fprintf(output, "Activity: %s; difficulty: %s; primary skill: %s\n", choice.Task.Activity, choice.Task.Difficulty, choice.Task.PrimarySkill)
	if len(choice.MissingPrerequisites) == 0 {
		fmt.Fprintln(output, "Prerequisites: satisfied by recorded completions.")
	} else {
		fmt.Fprintf(output, "Prerequisites not yet encountered: %s\n", strings.Join(choice.MissingPrerequisites, ", "))
	}
	if len(choice.MissingPrerequisites) > 0 && choice.Preparation != "" {
		fmt.Fprintf(output, "Preparation guide: %s\n", choice.Preparation)
	}
	if choice.Resume {
		fmt.Fprintf(output, "Continue: learn resume %s\n", choice.Task.ID)
	} else {
		fmt.Fprintf(output, "Start: learn start %s\n", choice.Task.ID)
	}
	if choice.RecentTypeCount > 0 {
		fmt.Fprintf(output, "Recent activity-type count: %d\n", choice.RecentTypeCount)
	}
}

func writeProgress(output io.Writer, progress Progress) {
	if len(progress.Tasks) == 0 {
		_, _ = io.WriteString(output, "No local progress.\n")
		return
	}
	ids := make([]string, 0, len(progress.Tasks))
	for id := range progress.Tasks {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		task := progress.Tasks[id]
		current := ""
		if id == progress.CurrentTaskID {
			current = " (current)"
		}
		fmt.Fprintf(output, "%s%s\n", id, current)
		fmt.Fprintf(output, "  status: %s\n  attempts: %d\n  hints: %v\n  solution revealed: %t\n  completed: %t\n  active duration: %s\n  feedback entries: %d\n", task.Status, task.Attempts, task.UsedHints, task.SolutionRevealed, task.Completed, formatDuration(task.ActivitySeconds), len(task.Feedback))
		if len(task.SkillEncounters) != 0 {
			values := make([]string, 0, len(task.SkillEncounters))
			for _, encounter := range task.SkillEncounters {
				values = append(values, fmt.Sprintf("%s:%s=%d", encounter.SkillID, encounter.Role, encounter.Count))
			}
			fmt.Fprintf(output, "  skill encounters: %s\n", strings.Join(values, ", "))
		}
	}
}

func noArguments(command string, arguments []string) error {
	if len(arguments) != 0 {
		return fmt.Errorf("%s does not accept arguments", command)
	}
	return nil
}

func oneArgument(command string, arguments []string) (string, error) {
	if len(arguments) != 1 {
		return "", fmt.Errorf("%s requires exactly one task id", command)
	}
	return arguments[0], nil
}

func explicitFlag(command string, arguments []string, expected string) (bool, error) {
	if len(arguments) != 1 || arguments[0] != expected {
		return false, fmt.Errorf("%s requires the explicit %s flag", command, expected)
	}
	return true, nil
}

func resetArguments(arguments []string) (string, bool, error) {
	confirmed := false
	id := ""
	for _, argument := range arguments {
		switch argument {
		case "--yes":
			if confirmed {
				return "", false, errors.New("reset-task received --yes more than once")
			}
			confirmed = true
		default:
			if strings.HasPrefix(argument, "-") || id != "" {
				return "", false, errors.New("reset-task accepts at most one task id and --yes")
			}
			id = argument
		}
	}
	return id, confirmed, nil
}

func cliError(output io.Writer, err error) int {
	fmt.Fprintf(output, "learn: %v\n", err)
	return 2
}
