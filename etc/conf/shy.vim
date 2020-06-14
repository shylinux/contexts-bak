
set foldmethod=indent
syn match Comment	    "#.*$"
syn match Comment	    "\"[^\"]*\""

highlight shyString     ctermfg=magenta
syn match shyString	    "\`[^\`]*\`"
syn region shyString	start="`" end="`"

highlight shyContext    ctermfg=red
syntax match shyContext "\~[a-z\.]\+"

highlight shyCommand    ctermfg=green
syntax match shyCommand "^    [a-zA-Z0-9:._]\+"
syntax match shyCommand "^[a-zA-Z0-9:._]\+"

highlight shyConfig    ctermfg=yellow
syntax match shyConfig "^    config"
syntax match shyConfig "^source"


highlight shyTitle    ctermbg=darkred ctermfg=white
syntax match shyTitle "^title"
syntax match shyTitle "^chapter"
syntax match shyTitle "^section"

highlight shySpark    ctermbg=darkgreen ctermfg=white
syntax match shySpark "^brief"
syntax match shySpark "^spark"

