syntax match Comment "#.*$"
syntax match Comment "Name: \"[^\"]*\""
syntax match Comment "Help: \"[^\"]*\""

highlight kitConst    ctermfg=yellow
syntax match kitConst "kit\.[a-z0-9A-Z_.]*"
syntax match kitConst "app\.[a-z0-9A-Z_.]*"

highlight msgConst    ctermfg=cyan
syntax match msgConst "\Am\.[a-z0-9A-Z_.]*"
syntax match msgConst "\Amsg\.[a-z0-9A-Z_.]*"
syntax match msgConst "\Asub\.[a-z0-9A-Z_.]*"

