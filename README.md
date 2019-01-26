# byebyebyte
byebyte in go (https://github.com/wayspurrchen/byebyte)
Command line tool written in Go, replaces random bytes with random values.

## Usage
### Global parameters
parameter | meaning | default
----- | ----- | -----
`-i, --input` | input file path | none, required
`-o, --output` | output file path | none, required
`--min` | the lower bound of the file to alter bytes in - use percentage 0 to 1 (ex: 0.15 = 15%, 1 = 100%). If specified, you cannot use --start or --stop | none, required
`--max` | the upper bound of the file to alter bytes in - use percentage from 0 to 1 (ex: 0.15 = 15%, 1 = 100%). If specified, you cannot use --start or --stop | none, required
`--start` | a specific point at the file, in bytes, at which to begin altering bytes. If specified, you cannot use --min or --max | none, required
`--stop` | a specific point at the file, in bytes, at which to stop altering bytes. If specified, you cannot use --min or --max | none, required

They're global because they are used by all commands in order to tell byebyte what file to alter as well as what portion of the file to alter.

- Use `-i` or `--input` to specify the file to change, and `-o` or `--output` to specify the target path for the resulting modified file.
- Use `--min` and `--max` to alter a percentage range of the file (for instance, `--min 0.2` and `--max 0.8` on a 200 byte file will cause changes within only bytes 40 to 160).
- Use `--start` and `--stop` to specify a specific byte range.
- You can only use `--min` and `--max` **or** `--start` and `--stop`.

### destroy mode
----- | ----- | -----
`-p, --probability` | probability | none, required
