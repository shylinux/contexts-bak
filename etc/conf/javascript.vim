syntax match Comment "\<Volcanos(\"[^\"]*\""
syntax match Comment "\<shy(\"[^\"]*\""
syntax match Comment "\<name: \"[^\"]*\""
syntax match Comment "\<help: \"[^\"]*\""

highlight canConst    ctermfg=yellow
syntax match canConst "\<can\>"
syntax match canConst "\<sub\>"
syntax match canConst "\<sup\>"
syntax match canConst "\<you\>"
syntax match canConst "\<msg\>"

highlight msgConst    ctermfg=cyan
syntax match msgConst "\<m\.[a-z0-9A-Z_.]*"
syntax match msgConst "\<msg\.[a-z0-9A-Z_.]*"
syntax match msgConst "\<can\.[a-z0-9A-Z_]*"
syntax match msgConst "\<sub\.[a-z0-9A-Z_.]*"
syntax match msgConst "\<sup\.[a-z0-9A-Z_.]*"
syntax match msgConst "\<you\.[a-z0-9A-Z_]*"

syntax match canConst "\<can\.base"
syntax match canConst "\<can\.core"
syntax match canConst "\<can\.misc"
syntax match canConst "\<can\.page"
syntax match canConst "\<can\.user"


