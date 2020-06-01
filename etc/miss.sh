[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone https://github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh
# declare -f ish_help_repos &>/dev/null || require conf.sh

require show.sh
require help.sh
require miss.sh
# cd $ISH_CONF_TASK

ish_miss_prepare_compile
ish_miss_prepare_install
ish_miss_prepare_session miss

ish_miss_prepare_volcanos
ish_miss_prepare_icebergs
ish_miss_prepare_intshell
# ish_miss_prepare toolkits
# ish_miss_prepare learning

# ish_miss_prepare wubi-dict
# ish_miss_prepare word-dict
