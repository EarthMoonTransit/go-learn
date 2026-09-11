// Run with node internal/learner/ui/app.test.cjs; no browser/network dependencies.
const {readFileSync} = require('node:fs');
const vm = require('node:vm');
const assert = require('node:assert/strict');
class Node {
  constructor(tag, text='') { this.tag=tag; this.textContent=text; this.children=[]; this.events={}; this.id=''; }
  append(...nodes) { this.children.push(...nodes); }
  replaceChildren(...nodes) { this.children=nodes; }
  addEventListener(type, callback) { this.events[type]=callback; }
}
const context = {URL, document:{createElement:tag=>new Node(tag),createTextNode:text=>new Node('#text',text),getElementById:id=>({scrollIntoView(){context.scrolled=id;}})}};
vm.createContext(context);
const source=readFileSync(__dirname+'/app.js','utf8');
vm.runInContext(source.slice(source.indexOf('function element('),source.indexOf('function renderCourse(')),context);
for(const url of ['javascript:alert(1)','data:text/html,x','//evil.example','https://u:p@go.dev','/api/action','#token/evil','https:\\evil.example']) assert.equal(context.safeLink(url),false,url);
for(const url of ['https://go.dev/ref/spec#Slice_types','#slice-types']) assert.equal(context.safeLink(url),true,url);
const target=new Node('div');target.id='theory';
context.markdown(target,'# Slice types\n\n[Read](https://go.dev/ref/spec#Slice_types) and [return](#slice-types).\n\n<script>alert(1)</script>\n\n[bad](javascript:alert)\n\n```go\nx := "<script>"\n```');
const descendants=node=>[node,...node.children.flatMap(descendants)];
const nodes=descendants(target), links=nodes.filter(n=>n.tag==='a');
assert.equal(links.length,2);
assert.equal(links[0].rel,'noopener noreferrer');assert.equal(links[0].referrerPolicy,'no-referrer');assert.equal(links[0].target,'_blank');
let prevented=false;links[1].events.click({preventDefault(){prevented=true;}});assert(prevented);assert.equal(context.scrolled,'theory-slice-types');
assert(!nodes.some(n=>n.tag==='script'||n.tag==='img'));assert(nodes.some(n=>n.tag==='code'&&n.textContent.includes('<script>')));
console.log('Markdown links, local anchors, raw HTML and credential isolation passed');
