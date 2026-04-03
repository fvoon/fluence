---
name: fluence
description: |
  Create, edit, view, and search Confluence pages using the fluence CLI. Use this skill when the user asks to:
  - Create or update Confluence documentation
  - View or fetch a Confluence page (by ID or URL)
  - Search Confluence for existing pages
  - Document code, APIs, or project information in Confluence
  - Sync markdown files to/from Confluence
  
  Trigger on: "confluence", "wiki page", "document this in confluence", confluence URLs, page IDs, "update the docs", "create documentation", "search confluence for".
---

# fluence CLI - Confluence Page Management

A CLI tool for managing Confluence pages with bidirectional Markdown conversion. Content is written in Markdown and automatically converted to Confluence storage format.

## Prerequisites

Environment variables must be set:
```bash
CONFLUENCE_BASE_URL="https://your-instance.atlassian.net"
CONFLUENCE_EMAIL="your-email@example.com"
CONFLUENCE_API_TOKEN="your-api-token"
CONFLUENCE_SPACE_KEY="YOURSPACE"  # optional default
```

## Quick Reference

### View a Page
```bash
# Get page content as Markdown
fluence page view PAGE_ID

# Get raw JSON response
fluence page view PAGE_ID -j
```

Extract PAGE_ID from Confluence URLs: `https://....atlassian.net/wiki/spaces/SPACE/pages/PAGE_ID/...`

### Create a Page
```bash
# From stdin
echo "# My Page\n\nContent here" | fluence page create -t "Page Title" -s SPACE_KEY

# From file
fluence page create -t "Page Title" -f content.md -s SPACE_KEY

# With parent page
fluence page create -t "Child Page" -f content.md -p PARENT_PAGE_ID
```

### Update a Page
```bash
# From stdin
cat updated-content.md | fluence page update PAGE_ID

# From file with version message
fluence page update PAGE_ID -f content.md -m "Added new section"

# Update title too
fluence page update PAGE_ID -f content.md -t "New Title"
```

### Search Pages
```bash
# Full-text search
fluence search "deployment guide"

# Search in specific space
fluence search "API documentation" -s MYSPACE

# Search by title
fluence search --title "Migration"

# Search by label
fluence search --label "architecture"

# Raw CQL query
fluence search --cql "type=page AND space=MYSPACE AND title~'API'"
```

### List Pages
```bash
# List pages in a space
fluence page list -s SPACE_KEY

# List children of a page
fluence page list -p PARENT_PAGE_ID

# With sorting
fluence page list -s SPACE --sort modified --desc -l 50
```

### Other Operations
```bash
# Delete a page
fluence page delete PAGE_ID

# Move a page to new parent
fluence page move PAGE_ID -p NEW_PARENT_ID

# List spaces
fluence space list

# View space details
fluence space view SPACE_KEY
```

## Workflow Patterns

### Update an Existing Page
1. Fetch current content: `fluence page view PAGE_ID > current.md`
2. Edit the markdown file
3. Update: `fluence page update PAGE_ID -f current.md -m "Description of changes"`

### Create Documentation from Code
1. Generate markdown documentation
2. Pipe to fluence: `cat docs.md | fluence page create -t "API Docs" -s DOCS -p PARENT_ID`

### Search and Update
1. Search: `fluence search "topic" -s SPACE -j` (JSON for page IDs)
2. View: `fluence page view PAGE_ID`
3. Update: `fluence page update PAGE_ID -f updated.md`

## Markdown Support

Supported:
- Headings, paragraphs, lists (ordered/unordered)
- Code blocks with syntax highlighting
- Tables (GFM format)
- Bold, italic, strikethrough
- Links and images
- Task lists (checkboxes)
- Blockquotes

The CLI handles conversion between Markdown and Confluence storage format automatically.

## Tips

- Use `-j` flag to get JSON output for scripting
- Page IDs are numeric (e.g., `6552387870`)
- Space keys are uppercase (e.g., `DOCS`, `ENG`)
- The `--verbose` flag shows detailed operation info
- Content can be piped via stdin or passed with `-f file.md`
