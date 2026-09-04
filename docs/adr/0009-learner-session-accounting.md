# ADR 0009: Learner work sessions across verification attempts

- Status: accepted
- Date: 2026-09-04

## Defect and decision

The first human-trial audit reproduced a ten-minute retry sequence recorded as
one minute: the first `check` ended the session, and later edits were not timed.
`resume` also discarded time already accumulated in an open session.

A work session now starts with `start` or `resume`, continues across public
checks and failed completion attempts, and ends with `pause`, successful
completion, or selection of another task. Resuming the current active task is
idempotent. Verification time is excluded; a check made while paused leaves the
session paused. `progress` includes the current elapsed interval in its display
without persisting or ending that interval.

This is an explicit wall-clock work-session approximation, not an idle detector.
The learner must pause before a break. Human calibration still records active
time independently so interruptions and missed pauses can be identified.

## Compatibility and migration

The progress schema remains version 1. `activity_seconds` still means accumulated
work-session seconds and `active_session_started_at` still identifies the open
interval. No task schema or taxonomy fields change. Existing totals are retained;
previously unrecorded time cannot be reconstructed and is not invented. Existing
paused records remain paused until `resume`. Switching tasks closes any other
open intervals left by the earlier implementation.

Old and new CLI binaries must not be used concurrently against one learner home.
Tests cover retries, idempotent resume, pause, task switching, successful
completion, and exclusion of verification time.
