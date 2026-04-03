# Confluence CLI - fluence

stdout is Markdown, token efficient

Use --json only for structured data

Quick Start:

```
fluence space list
fluence space view SPACE_KEY
fluence page list -s SPACE_KEY
fluence page list --parent PAGE_ID
fluence page view PAGE_ID
fluence page view PAGE_ID --json
fluence search "query text"
fluence search --title "page name"
fluence search --label documentation
fluence search --cql "type=page AND space=SPACE"
fluence page create -t "Title" -f content.md -s SPACE --parent PAGE_ID
echo "# Title" | fluence page create -t "Page Title" -s SPACE
echo "# Heading\n\nContent here" | fluence page update PAGE_ID -f -
fluence page update PAGE_ID -f updated.md
fluence page update PAGE_ID -f content.md -m "Update message"
fluence page move PAGE_ID --parent NEW_PARENT_ID
fluence page delete PAGE_ID
fluence debug md < input.md
fluence debug storage < storage.html
```

Global Flags:

```
--verbose   Show detailed warnings and debug information
--json, -j  Output in JSON format (most commands)
--help, -h  Help for command
--version   Print version
```

Command Flags:

```
page create:
  -t, --title <title>   Page title (required)
  -f, --file <path>     Markdown file, or - for stdin
  -s, --space <key>     Space key (uses CONFLUENCE_SPACE_KEY if not set)
  -p, --parent <id>     Parent page ID
  -j, --json            Output as JSON
page view:
  -j, --json            Output as JSON (returns full API response)
page update:
  -t, --title <title>   New page title (optional, keeps existing)
  -f, --file <path>     Markdown file, or - for stdin
  -m, --message <msg>   Version update message
  -j, --json            Output as JSON
page list:
  -s, --space <key>     Space key (uses CONFLUENCE_SPACE_KEY if not set)
  -p, --parent <id>     Parent page ID (list children)
  -l, --limit <n>       Maximum results (default: 25)
  --sort <field>        Sort: web, title, created, modified, id
  --desc                Sort descending
  -j, --json            Output as JSON
page move:
  -p, --parent <id>     Target parent page ID (required)
  -j, --json            Output as JSON
page delete:
  (no additional flags)
space list:
  -l, --limit <n>       Maximum results (default: 25)
  -j, --json            Output as JSON
space view:
  -j, --json            Output as JSON
search:
  --title <text>        Search in page titles
  --label <label>       Search by label (exact match)
  --creator <email>     Filter by creator (email or 'me')
  -s, --space <key>     Filter by space key
  --type <type>         Content type (page, blogpost, attachment)
  -l, --limit <n>       Maximum results (default: 25)
  --cursor <cursor>     Pagination cursor from previous search
  --cql <query>         Raw CQL query (overrides other search flags)
  -j, --json            Output as JSON
debug md:
  (reads markdown from stdin, outputs storage format)
debug storage:
  (reads storage format from stdin, outputs markdown)
```

More Help:

```
fluence help agents workflow
fluence help agents all
```

Bug Reports:

Repository and issue tracker available via `fluence --version`.
