"安装plug-vim {{{
"$ curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
"
"安装vim各种插件
":PlugInstall
"
"}}}
"加载插件"{{{
call plug#begin()
Plug 'vim-scripts/tComment'
Plug 'tpope/vim-fugitive'
Plug 'airblade/vim-gitgutter'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'ntpeters/vim-better-whitespace'
Plug 'easymotion/vim-easymotion'

Plug 'gcmt/taboo.vim'
set sessionoptions+=tabpages,globals

Plug 'vim-scripts/taglist.vim'
let g:Tlist_WinWidth=45
let g:Tlist_Exit_OnlyWindow=1
let g:Tlist_Enable_Fold_Column=0
nnoremap <F2> :TlistToggle<CR>

Plug 'scrooloose/nerdtree'
let g:NERDTreeWinPos="right"
nnoremap <F4> :NERDTreeToggle<CR>

Plug 'kien/ctrlp.vim'
let g:ctrlp_cmd='CtrlPBuffer'

Plug 'rking/ag.vim'
nnoremap <C-G> :Ag <C-R>=expand("<cword>")<CR><CR>

Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
nnoremap <C-N> :FZF -q <C-R>=expand("<cword>")<CR><CR>

Plug 'benmills/vimux'
let mapleader=";"
nnoremap <Leader>; :VimuxPromptCommand<CR>
" nnoremap <Leader>j :VimuxZoomRunner<CR>
" nnoremap <Leader>l :VimuxRunLastCommand<CR>
" nnoremap <Leader>vx :VimuxInterruptRunner<CR>

Plug 'fatih/vim-go'
Plug 'chr4/nginx.vim'
Plug 'othree/html5.vim'
Plug 'godlygeek/tabular'
Plug 'plasticboy/vim-markdown'
Plug 'vim-scripts/python.vim'

Plug 'mbbill/echofunc'
Plug 'vim-syntastic/syntastic'
let g:syntastic_quiet_messages = { "regex": [
            \ "Missing module docstring",
            \ "Missing class docstring",
            \ "Missing method docstring",
            \ "Missing function docstring",
            \ "Invalid class name",
            \ "Invalid method name",
            \ "Invalid function name",
            \ "Invalid constant name",
            \ "Invalid variable name",
            \ "Method could be a function",
            \ "Too many instance attributes",
            \ "Wrong continued indentation",
            \ "Too many lines in module",
            \ "Too many arguments",
            \ "Too many local variables",
            \ "Too many branches",
            \ "Too many statements",
            \ "Too many return statements",
            \ "Line Too Long",
            \ "defined outside __init__",
            \ "Catching too general exception Exception",
         \ ] }

Plug 'Valloric/YouCompleteMe'
let g:ycm_confirm_extra_conf=1
let g:syntastic_enable_signs=1
let g:ycm_python_binary_path='/usr/bin/python'
nnoremap gd :YcmCompleter GoToDeclaration<CR>
nnoremap gD :YcmCompleter GoToReferences<CR>

Plug 'vim-scripts/matrix.vim--Yang'
call plug#end()
"}}}
" 基本配置"{{{
set number
set relativenumber
set cursorline
set cursorcolumn
set ruler
set showcmd
set showmode
set cc=80
set nowrap
set scrolloff=3

set tabstop=4
set shiftwidth=4
set cindent
set expandtab
set backspace=indent,eol,start

set showmatch
set matchtime=2
set foldenable
set foldmethod=marker

set hlsearch
set incsearch
set nowrapscan
set smartcase
set ignorecase

set hidden
set autowrite
set encoding=utf-8
set mouse=a

" colorscheme darkblue
" colorscheme default
set t_Co=256
"}}}
"快捷键映射"{{{
nnoremap <C-H> <C-W>h
nnoremap <C-J> <C-W>j
nnoremap <C-K> <C-W>k
nnoremap <C-L> <C-W>l
" nnoremap <C-M> :make<CR>
nnoremap <Space> :

nnoremap j gj
nnoremap k gk

nnoremap <C-M> :make<CR>
inoremap <C-M> <ESC>:make<CR>i

nnoremap df :FZF<CR>
inoremap df _
inoremap jk <Esc>
cnoremap jk <CR>
"}}}
" 编程配置{{{
set keywordprg=man\ -a
set splitbelow
set splitright

function! Config(type)
    if a:type == "go"
        set foldmethod=syntax
    elseif a:type == "shy"
        set filetype=shy
        set commentstring=#%s
    elseif a:type == "json"
        set foldmethod=syntax
    elseif a:type == "conf"
        set filetype=nginx
    elseif a:type == "xml"
        set filetype=xml
    elseif a:type == "css"
        set filetype=css
    elseif a:type == "txt"
        set noexpandtab
    endif
endfunction

autocmd BufReadPost * normal `"

autocmd BufNewFile,BufReadPost *.go call Config("go")
autocmd BufNewFile,BufReadPost *.shy call Config("shy")
autocmd BufNewFile,BufReadPost *.conf call Config("conf")
autocmd BufNewFile,BufReadPost *.json call Config("json")

autocmd BufNewFile,BufReadPost *.wxml call Config("xml")
autocmd BufNewFile,BufReadPost *.wxss call Config("css")
autocmd BufNewFile,BufReadPost *.txt call Config("txt")

command! RR wa | source ~/.vimrc |e
command! SS mksession! etc/session.vim
source ~/.vim_local

let g:colorscheme=1
let g:colorlist = [ "ron", "torte", "darkblue", "peachpuff" ]
function! ColorNext()
    if g:colorscheme >= len(g:colorlist)
        let g:colorscheme = 0
    endif
    let g:scheme = g:colorlist[g:colorscheme]
    exec "colorscheme " . g:scheme
    let g:colorscheme = g:colorscheme+1
endfunction
call ColorNext()
command! NN call ColorNext()<CR>

function! BenchCode(cmd, arg)
    let l:line = "web.code." . a:cmd . " " . join(a:arg, " ")
    exe "silent !bench " l:line
endfunction

" autocmd BufReadPost * call BenchCode("counter", ["nopen", 1])
" autocmd BufWritePost * call BenchCode("counter", ["nsave", 1])
"}}}
