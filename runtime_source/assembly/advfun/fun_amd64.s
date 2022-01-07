#include "textflag.h"

TEXT ·printnl_nosplit(SB), NOSPLIT, $8
CALL runtime·printnl(SB)
RET
