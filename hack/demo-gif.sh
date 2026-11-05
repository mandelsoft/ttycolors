#!/usr/bin/env bash
#
# Generates the file demo.gif for an example in the folder examples.
#
# The example is executed with a pseudo terminal (using script), so that
# colored output is generated. The output is rendered with terminalizer
# based on the configuration hack/terminalizer.json. Lines appear one
# after the other, the final output is shown for some seconds.
#
# Requirements: bash, jq, script (util-linux), terminalizer.
#
# Usage: hack/demo-gif.sh <example>
#        for example: hack/demo-gif.sh ttycolors/strings
#
# The gifs for all examples are generated with make gifs.

set -euo pipefail

declare -A TITLES=(
  [ttycolors/strings]="Cascaded Colorized Text"
  [ttycolors/formats]="Composed Formats"
  [ttycolors/context]="Output Contexts"
  [ttycolors/embedded]="Embedded Formatted Strings"
  [ttycolors/underline]="Underline Styles and Colors"
  [plain]="Plain Formatted Strings"
  [ansi/codes]="ANSI Escape Sequences"
  [ansi/scanner]="Scanning Escape Sequences"
  [ansi/strings]="Strings with Escape Sequences"
)

LINE_DELAY=400     # ms between lines
FINAL_DELAY=3000   # ms the final output is shown
RENDER_TIMEOUT=180 # s for rendering a single gif

if [ $# -ne 1 ]; then
  echo "usage: $0 <example>" >&2
  exit 1
fi

example="${1%/}"
repo="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
if [ ! -f "$repo/examples/$example/main.go" ]; then
  echo "unknown example $example" >&2
  exit 1
fi
title="${TITLES[$example]:-$example}"
gif="$repo/examples/$example/demo.gif"

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

# execute the example with a pseudo terminal to get colored output
(cd "$repo" && go build -o "$tmp/example" "./examples/$example")
script -qec "$tmp/example" /dev/null >"$tmp/output"

# Output lines using output formats not supported by the terminal
# emulation of terminalizer (xterm.js 3.x) are omitted: blink (5, 6),
# strike through (9), double underline (21), overline (53), underline
# color (58) and all sequences with sub parameters (for example 4:3),
# which are ignored completely. The arguments of extended colors
# (38;5;n and 38;2;r;g;b) are no output formats.
FILTER='
  def unsupported:
    if test(":") then true
    else (split(";") | map(if . == "" then 0 else tonumber end)) as $codes
    | reduce range(0; $codes | length) as $i ({skip: 0, bad: false};
        if .skip > 0 then .skip -= 1
        else $codes[$i] as $c
        | if $c == 38 or $c == 48 then
            .skip = (if $codes[$i+1] == 5 then 2 elif $codes[$i+1] == 2 then 4 else 0 end)
          elif any(5, 6, 9, 21, 53, 58; . == $c) then .bad = true
          else . end
        end)
    | .bad
    end;
  def supported: [scan("\u001b\\[([0-9;:]*)m")[0] | unsupported] | any | not;
  def lines: $output | gsub("\r\n"; "\n") | rtrimstr("\n") | split("\n");
'

# report omitted lines
jq -nr --rawfile output "$tmp/output" "$FILTER"' lines[] | select(supported | not)
  | "omitting line with unsupported formats: " + gsub("\u001b\\[[0-9;:?]*[A-Za-z]"; "")' >&2

# generate the terminalizer configuration (JSON is valid YAML,
# but terminalizer requires the file extension .yml).
# Notes:
# - the cursor is hidden by the first record
# - terminalizer hangs on records with empty content, therefore
#   the final record contains an invisible escape sequence
jq --rawfile output "$tmp/output" \
   --arg example "$example" \
   --arg title "$title" \
   --argjson lineDelay "$LINE_DELAY" \
   --argjson finalDelay "$FINAL_DELAY" "$FILTER"' (lines | map(select(supported))) as $lines
  | ($lines | map(gsub("\u001b\\[[0-9;:?]*[A-Za-z]"; "") | length) | max) as $width
  | .config.command = "go run ./examples/" + $example
  | .config.cols = $width + 2
  | .config.rows = ($lines | length) + 1
  | .config.frameBox.title = $title
  | .records = ([$lines | to_entries[]
                 | {delay: (if .key == 0 then 0 else $lineDelay end),
                    content: ((if .key == 0 then "\u001b[?25l" else "" end) + .value + "\r\n")}]
                + [{delay: $finalDelay, content: "\u001b[0m"}])
' "$repo/hack/terminalizer.json" >"$tmp/demo.yml"

if ! timeout "$RENDER_TIMEOUT" terminalizer render "$tmp/demo.yml" -o "$gif" >"$tmp/render.log" 2>&1; then
  echo "rendering $example failed:" >&2
  cat "$tmp/render.log" >&2
  exit 1
fi
echo "rendered examples/$example/demo.gif"
