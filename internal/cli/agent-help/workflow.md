# Confluence CLI - fluence

Workflow Patterns:

```
fluence page create -t "Page Title" -f content.md -s SPACE
fluence page create -t "Child Page" -f content.md -s SPACE --parent PAGE_ID
echo "# Heading\n\nContent here" | fluence page create -t "Title" -s SPACE
echo "# Heading\n\nContent here" | fluence page create -t "Title" -s SPACE -f -
cat document.md | fluence page create -t "Title" -s SPACE
URL=$(fluence page create -t "Title" -f content.md -s SPACE)
ID=$(fluence page create -t "Title" -f content.md -s SPACE --json | jq -r '.id')

fluence page update PAGE_ID -f updated.md
fluence page update PAGE_ID -f content.md -m "Fixed typos"
fluence page update PAGE_ID -t "New Title" -f content.md
echo "# Heading\n\nContent here" | fluence page update PAGE_ID -f -

fluence space list
fluence space list --json | jq '.[].key'

fluence page list -s SPACE
fluence page list -s SPACE --limit 100
fluence page list -s SPACE --sort title
fluence page list -s SPACE --sort modified --desc

fluence page list --parent PAGE_ID
fluence page list --parent PAGE_ID --sort title

fluence search "error handling"
fluence search "API documentation" -s SPACE
fluence search --title "README"
fluence search --title "Architecture" -s SPACE
fluence search --label documentation
fluence search --label api-reference -s SPACE
fluence search "query" --label docs --creator me
fluence search --cql "type=page AND space=SPACE AND label=important"
fluence search --cql "creator=currentUser() AND lastModified > now('-7d')"

fluence search "query" --limit 50
CURSOR=$(fluence search "query" --json | jq -r '.nextCursor // empty')
fluence search "query" --cursor "$CURSOR"

fluence debug md < document.md
fluence page view PAGE_ID --json | jq -r '.body.storage.value' | fluence debug storage

for id in PAGE_ID1 PAGE_ID2 PAGE_ID3; do
  fluence page view "$id" > "page-$id.md"
done

fluence search --label outdated --json | jq -r '.results[].content.id' | while read id; do
  echo "Processing $id"
  fluence page view "$id"
done

fluence page move PAGE_ID --parent NEW_PARENT_ID
fluence page delete PAGE_ID
```
