'use strict';
const byId = id => document.getElementById(id);
const state = { selected: '', course: null, task: null, job: null, pending: false, busy: false, refreshing: false, authorized: true, taskLoad: 0, routeSignature: '', topics: [], topic: '', theoryLoad: 0 };
let token = location.hash.slice(1);
try {
  if (token) sessionStorage.setItem('go-learn-token', token);
  else token = sessionStorage.getItem('go-learn-token') || '';
} catch (_) { /* The printed URL still works if storage is disabled. */ }
if (location.hash) history.replaceState(null, '', '/');
// Reopening a new session URL on the same host/port can be a same-document
// fragment navigation. Reload so the new token is consumed and stale jobs reset.
window.addEventListener('hashchange', () => { if (location.hash) location.reload(); });
function alertMessage(text) { byId('alert').textContent = text; byId('alert').hidden = !text; }
async function api(path, body) {
  let response;
  try {
    response = await fetch(path, { method: body ? 'POST' : 'GET', credentials: 'omit', cache: 'no-store', headers: { Authorization: 'Bearer ' + token, ...(body ? { 'Content-Type': 'application/json' } : {}) }, ...(body ? { body: JSON.stringify(body) } : {}) });
  } catch (_) { throw new Error('Cannot reach the local UI. Start learn ui and reopen the full printed URL.'); }
  let data;
  try { data = await response.json(); } catch (_) { throw new Error('The local UI returned an unreadable response. Refresh or reopen the printed URL.'); }
  if (!response.ok) {
    if (response.status === 401) { state.authorized = false; byId('connection').textContent = 'Reopen printed URL'; actionsEnabled(); }
    throw new Error(data.error || 'The request could not finish.');
  }
  return data;
}
function element(tag, text, className) { const node = document.createElement(tag); if (text !== undefined) node.textContent = text; if (className) node.className = className; return node; }
function safeLink(raw) {
  if (/^#[a-zA-Z][a-zA-Z0-9_-]*$/.test(raw)) return true;
  if (/[\\\s]/.test(raw)) return false;
  try { const url = new URL(raw); return url.protocol === 'https:' && !!url.hostname && !url.username && !url.password && !url.port; } catch (_) { return false; }
}
function readingLink(label, href, target) {
  const link = element('a', label); link.href = href;
  if (href.startsWith('#')) link.addEventListener('click', event => { event.preventDefault(); document.getElementById(target.id + '-' + href.slice(1))?.scrollIntoView(); });
  else { link.target = '_blank'; link.rel = 'noopener noreferrer'; link.referrerPolicy = 'no-referrer'; }
  return link;
}
function inline(node, text, target = node) {
  for (const piece of text.split(/(`[^`]+`|\*\*[^*]+\*\*|\[[^\]\n]+\]\([^)\n]*\))/g)) {
    if (piece.startsWith('`') && piece.endsWith('`')) node.append(element('code', piece.slice(1, -1)));
    else if (piece.startsWith('**') && piece.endsWith('**')) node.append(element('strong', piece.slice(2, -2)));
    else { const match = piece.match(/^\[([^\]]+)\]\(([^)]*)\)$/); if (match && safeLink(match[2])) node.append(readingLink(match[1], match[2], target)); else node.append(document.createTextNode(piece)); }
  }
}
// Render the small Markdown subset used by tasks using text nodes only. Raw
// HTML, script, URLs and diagnostics are never passed to innerHTML or eval.
function markdown(target, text) {
  target.replaceChildren();
  const lines = (text || '').replace(/\r/g, '').split('\n');
  for (let i = 0; i < lines.length;) {
    const line = lines[i];
    if (!line.trim()) { i++; continue; }
    if (/^\s*```/.test(line)) {
      const code = []; i++;
      while (i < lines.length && !/^\s*```/.test(lines[i])) code.push(lines[i++]);
      i++; const pre = element('pre'); pre.append(element('code', code.join('\n'))); target.append(pre); continue;
    }
    const heading = line.match(/^\s*(#{1,6})\s+(.+)$/);
    if (heading) { const node = element('h' + Math.min(heading[1].length, 4)); node.id = target.id + '-' + heading[2].toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-|-$/g, ''); inline(node, heading[2], target); target.append(node); i++; continue; }
    const list = line.match(/^\s*(?:[-*]|\d+\.)\s+(.+)$/);
    if (list) {
      const node = element(/^\s*\d+\./.test(line) ? 'ol' : 'ul');
      while (i < lines.length) { const item = lines[i].match(/^\s*(?:[-*]|\d+\.)\s+(.+)$/); if (!item) break; const li = element('li'); inline(li, item[1], target); node.append(li); i++; }
      target.append(node); continue;
    }
    const paragraph = [line.trim()]; i++;
    while (i < lines.length && lines[i].trim() && !/^\s*(?:#|```|[-*]\s|\d+\.\s)/.test(lines[i])) paragraph.push(lines[i++].trim());
    const node = element('p'); inline(node, paragraph.join(' '), target); target.append(node);
  }
}
function renderCourse() {
  const c = state.course; if (!c) return;
  byId('counts').textContent = c.completed + ' completed · ' + c.ready + ' ready lessons · ' + c.learning.mastered + ' skills confirmed · ' + c.learning.due + ' reviews due';
  if (c.learning.upgrade_required) byId('counts').textContent += ' · Enable learning history: learn organize';
  byId('next').disabled = !c.next_task_id;
  byId('next-message').textContent = c.next_message;
  byId('next-message').hidden = !c.next_message;
  byId('connection').textContent = state.busy ? 'Operation running' : 'Local · connected';
  const signature = JSON.stringify([c.blocks, state.selected]);
  if (signature === state.routeSignature) return;
  state.routeSignature = signature;
  const open = new Set([...byId('route').querySelectorAll('details[open]')].map(node => node.dataset.block));
  byId('route').replaceChildren();
  for (const block of c.blocks) {
    const details = element('details', undefined, 'route-block'); details.dataset.block = block.id;
    details.open = open.has(block.id) || block.status === 'current';
    const summary = element('summary', block.title); summary.append(element('span', block.status, 'block-status')); details.append(summary);
    if (!block.lessons.length) details.append(element('p', 'Content planned', 'planned'));
    for (const lesson of block.lessons) {
      const button = element('button', undefined, 'lesson-link' + (lesson.id === state.selected ? ' selected' : ''));
      button.disabled = !lesson.available; button.title = lesson.id + ' · ' + lesson.status;
      button.append(element('span', lesson.status === 'completed' ? '✓' : (lesson.id === c.next_task_id ? '→' : '·'), 'lesson-marker'), element('span', (lesson.number ? lesson.number + ' — ' : '') + lesson.title));
      button.addEventListener('click', () => selectTask(lesson.id).catch(error => alertMessage(error.message)));
      details.append(button);
    }
    byId('route').append(details);
  }
}
function actionsEnabled() {
  const t = state.task; const busy = state.busy || state.pending || !state.authorized;
  for (const button of document.querySelectorAll('[data-action]')) {
    const action = button.dataset.action;
    button.hidden = (action === 'start' && !!t?.started) || (action === 'resume' && (!t?.started || (t.current && (t.session_active || t.completed))));
    let enabled = !!t && !busy && !t.workspace_unavailable;
    if (['check', 'complete', 'hint', 'pause'].includes(action)) enabled = enabled && t.started && t.current;
    if (action === 'complete') enabled = enabled && !t.completed;
    if (action === 'hint') enabled = enabled && t.hints_remaining > 0;
    if (action === 'pause') enabled = enabled && t.session_active;
    if (action === 'open') enabled = enabled && !!t.workspace && !!state.course?.ide_available;
    button.disabled = !enabled;
  }
}
function renderTask() {
  const t = state.task; if (!t) return;
  byId('welcome').hidden = true; byId('lesson').hidden = !!state.topic;
  byId('title').textContent = (t.number ? t.number + ' — ' : '') + t.info.Title;
  byId('task-meta').textContent = t.info.ID + ' · ' + t.info.EstimatedTimeMinutes.min + '–' + t.info.EstimatedTimeMinutes.max + ' min';
  byId('task-kind').textContent = t.info.Activity;
  byId('task-status').textContent = (t.completed ? 'Completed' : t.started ? 'In progress' : 'Ready to start') + ' · ' + t.attempts + ' attempts' + (t.started && !t.completed ? (t.session_active ? ' · Work session active' : ' · Work session paused') : '');
  markdown(byId('statement'), t.info.Prompt);
  markdown(byId('preparation'), state.course?.preparation || 'Use learn next to see the preparation guide.');
  byId('workspace-path').value = t.workspace;
  byId('ide-project-path').value = t.ide_project;
  byId('copy-project').disabled = !t.ide_project;
  byId('copy').disabled = !t.workspace;
  byId('workspace-guidance').textContent = t.workspace_unavailable ? 'Workspace needs attention. Use the CLI to inspect it; your files have not been replaced.' : !t.started ? 'Start this lesson to create its private workspace.' : !t.current ? 'Resume this lesson before checking. Viewing it has not changed your active task.' : 'Edit the files named in the task, then check your work here.';
  byId('launcher-note').textContent = state.course?.ide_available ? '' : 'GoLand launcher unavailable. Copy the project path and open that folder in your IDE.';
  byId('hints').replaceChildren();
  if (!t.used_hints.length) byId('hints').append(element('p', 'No hints used for this lesson.', 'muted'));
  for (const hint of t.used_hints) { const section = element('div', undefined, 'hint'); section.append(element('h3', 'Hint ' + hint.number)); const text = element('p'); inline(text, hint.text); section.append(text); byId('hints').append(section); }
  byId('review-card').hidden = !t.review;
  if (t.review) markdown(byId('review'), t.review);
  renderTaskTheory(); actionsEnabled(); renderJob();
}
function renderJob() {
  const job = state.job;
  if (!job || job.task_id !== state.selected) { byId('result').hidden = true; return; }
  byId('result').hidden = false; byId('diagnostic').hidden = true;
  let title = 'Operation finished'; let summary = job.result?.message || 'Your progress is up to date.'; let outcome = 'passed';
  if (job.state === 'running') { title = job.action === 'complete' ? 'Running full verification…' : job.action === 'check' ? 'Running public tests…' : 'Working…'; summary = 'You can keep viewing the route. Another action can start when this one finishes.'; outcome = 'running'; }
  else if (job.state === 'failed') { title = 'Action could not finish'; summary = job.error; outcome = 'failed'; }
  else if (job.result?.verification) {
    const v = job.result.verification; title = v.Passed ? (job.action === 'complete' ? 'Lesson completed' : 'Public tests passed') : 'Verification did not pass'; summary = v.Summary; outcome = v.Passed ? 'passed' : 'failed';
    if (job.action === 'check' && v.Diagnostic) { byId('diagnostic').textContent = v.Diagnostic; byId('diagnostic').hidden = false; }
  } else if (job.action === 'start' || job.action === 'resume') { title = 'Workspace ready'; summary = 'This lesson is in your shared project. Open its folder in GoLand, or use Open lesson in GoLand to jump to the code.'; }
  else if (job.action === 'hint') { title = 'Next hint revealed'; summary = 'It is saved in your lesson progress and shown below.'; }
  byId('result-title').textContent = title; byId('result-summary').textContent = summary; byId('result').dataset.state = outcome;
}
async function loadTask() {
  const selected = state.selected; const serial = ++state.taskLoad;
  if (!selected) return;
  const task = await api('/api/task?id=' + encodeURIComponent(selected));
  if (serial !== state.taskLoad || selected !== state.selected) return;
  state.task = task; renderTask();
}
async function selectTask(id) {
  state.topic = ''; state.theoryLoad++; byId('theory').hidden = true; state.selected = id; state.task = null; state.taskLoad++;
  byId('lesson').hidden = true; byId('welcome').hidden = false; byId('result').hidden = true;
  renderCourse(); actionsEnabled(); await loadTask();
}
async function refresh() {
  if (state.refreshing || !state.authorized) return;
  state.refreshing = true;
  try {
    const data = await api('/api/state'); state.course = data.course;
    state.busy = !!data.job || data.stopping;
    if (data.job) state.job = data.job;
    else if (state.job?.state === 'running') state.job = await api('/api/job?id=' + encodeURIComponent(state.job.id));
    renderCourse(); actionsEnabled(); renderJob();
    if (!state.selected && !state.topic) { const id = data.course.current_task_id || data.course.next_task_id; if (id) await selectTask(id); }
    else if (!state.task || state.task.revision !== data.course.revision) await loadTask();
    if (data.stopping) alertMessage('The UI is stopping. The active operation will finish before the process exits.');
  } finally { state.refreshing = false; }
}
async function action(name) {
  if (!state.task || state.pending || state.busy || !state.authorized) return;
  const request = { action: name, task_id: state.selected, revision: state.task.revision, request_id: crypto.randomUUID() };
  state.pending = true; actionsEnabled(); alertMessage('');
  try { state.job = await api('/api/action', request); state.busy = state.job.state === 'running'; renderJob(); }
  catch (error) { alertMessage(error.message); }
  finally { state.pending = false; actionsEnabled(); await refresh(); }
}
for (const button of document.querySelectorAll('[data-action]')) button.addEventListener('click', () => action(button.dataset.action).catch(error => alertMessage(error.message)));
byId('next').addEventListener('click', () => { if (state.course?.next_task_id) selectTask(state.course.next_task_id).catch(error => alertMessage(error.message)); });
byId('refresh').addEventListener('click', () => refresh().catch(error => alertMessage(error.message)));
for (const [button, field, message] of [['copy', 'workspace-path', 'Lesson path copied.'], ['copy-project', 'ide-project-path', 'Project path copied. Open this folder once in GoLand.']]) {
  byId(button).addEventListener('click', async () => {
    const input = byId(field); if (!input.value) return;
    try { await navigator.clipboard.writeText(input.value); alertMessage(message); }
    catch (_) { input.focus(); input.select(); alertMessage('Path selected. Copy it and open it in GoLand.'); }
  });
}
if (!token) { state.authorized = false; alertMessage('Open the full URL printed by learn ui, including its session fragment.'); byId('connection').textContent = 'Session URL required'; }
else refresh().then(loadTopics).catch(error => alertMessage(error.message));
setInterval(() => refresh().catch(error => alertMessage(error.message)), 1800);

function topicButton(topic) {
  const button = element('button', 'Theory: ' + topic.title + (topic.status === 'planned' ? ' (planned)' : ''), 'lesson-link');
  button.addEventListener('click', () => selectTheory(topic.id).catch(error => alertMessage(error.message))); return button;
}
function renderTaskTheory() {
  const target = byId('task-theory'); target.replaceChildren();
  for (const topic of state.topics) if (topic.task_ids.includes(state.selected)) target.append(topicButton(topic));
}
async function loadTopics() {
  state.topics = await api('/api/topics'); const target = byId('topics'); target.replaceChildren();
  if (state.topics.length) target.append(element('h2', 'Theory'));
  for (const block of state.course?.blocks || []) {
    const topics = state.topics.filter(topic => topic.section_id === block.id); if (!topics.length) continue;
    const group = element('details'); group.append(element('summary', block.title));
    for (const topic of topics) group.append(topicButton(topic)); target.append(group);
  }
  renderTaskTheory();
}
async function selectTheory(id) {
  const serial = ++state.theoryLoad;
  const view = await api('/api/theory?id=' + encodeURIComponent(id)); if (serial !== state.theoryLoad) return;
  state.topic = id; byId('theory').hidden = false; byId('lesson').hidden = true; byId('welcome').hidden = true;
  markdown(byId('theory-content'), view.topic.status === 'planned' ? '# ' + view.topic.title + '\n\nThis explanation is planned. Existing exercises remain available.' : view.markdown);
  byId('theory-back').hidden = !state.selected;
  const prerequisites = byId('theory-prerequisites'); prerequisites.replaceChildren();
  if (view.topic.prerequisites.length) prerequisites.append(element('h2', 'Before reading'));
  for (const prior of view.topic.prerequisites) prerequisites.append(topicButton(state.topics.find(topic => topic.id === prior)));
  const lessons = byId('theory-lessons'); lessons.replaceChildren();
  for (const lesson of view.lessons) { const button = element('button', (lesson.number ? lesson.number + ' — ' : '') + lesson.title, 'lesson-link'); button.addEventListener('click', () => selectTask(lesson.id).catch(error => alertMessage(error.message))); lessons.append(button); }
  const links = byId('theory-links'); links.replaceChildren();
  for (const link of view.topic.official_links) { const p = element('p'); p.append(readingLink(link.description, link.url, links)); links.append(p); }
}
byId('theory-back').addEventListener('click', () => selectTask(state.selected).catch(error => alertMessage(error.message)));
