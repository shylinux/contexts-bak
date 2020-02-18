
set foldmarker={{,}}
highlight shyContext    ctermfg=yellow
syntax match shyContext '^\.[a-z\.]*'

highlight shyCommand    ctermfg=yellow
syntax match shyCommand "^\a*"
syntax match shyCommand "^    \a*"

highlight shyTitle    ctermbg=darkred ctermfg=white
syntax match shyTitle "^title"
syntax match shyTitle "^chapter"
syntax match shyTitle "^section"

highlight shySpark    ctermbg=darkgreen ctermfg=white
syntax match shySpark "^brief"
syntax match shySpark "^spark"

highlight shyString     ctermfg=magenta
syn match shyString	    "\`[^\`]*\`"
syn region shyString	start="`" end="`"

syn match Comment	    "\"[^\"]*\""

