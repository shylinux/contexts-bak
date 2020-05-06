[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone https://github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh
# declare -f ish_help_repos &>/dev/null || require conf.sh

require help.sh
require miss.sh

ish_miss_prepare
ish_miss_compile_prepare
ish_miss_install_prepare

ish_miss_volcanos_prepare
ish_miss_icebergs_prepare
ish_miss_intshell_prepare
