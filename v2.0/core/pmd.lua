module("pmd", package.seeall)

local modname = "pmd"
local M = {}
_G[modname] = M                      -- 将模块表加入到全局变量中
package.loaded[modname] = M    -- 将模块表加入到package.loaded中，防止多次加载
M["LDO_VLCD"]=55

function M:sleep()

end
function M:ldoset()

end
function M:param_get()
    return 5,5,5,5,5
end
function M:sys32k_clk_out()

end