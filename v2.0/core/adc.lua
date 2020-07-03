module("adc", package.seeall)

local modname = "adc"
local M = {}
_G[modname] = M                      -- 将模块表加入到全局变量中
package.loaded[modname] = M    -- 将模块表加入到package.loaded中，防止多次加载

function M:open()

end

function M:read()
    return 0,0
end