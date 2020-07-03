module("audiocore", package.seeall)

local modname = "audiocore"
local M = {
    ["LOUDSPEAKER"]=1
}
_G[modname] = M                      -- 将模块表加入到全局变量中
package.loaded[modname] = M    -- 将模块表加入到package.loaded中，防止多次加载

function M:setchannel()

end
function M:setMicVolume()

end
function M:setvol()

end