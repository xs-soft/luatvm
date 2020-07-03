module("pio", package.seeall)

local modname = "pio"
local M = {}
local pin = {}
M["pin"]=pin-- 定义用于返回的模块表
_G[modname] = M                      -- 将模块表加入到全局变量中
package.loaded[modname] = M    -- 将模块表加入到package.loaded中，防止多次加载
for i=0,3 do
    for x=0,31 do
        M["P"..i.."_"..x]=i*32+x
    end
end
function M:setup()

end
function M:on()

end


function pin:setup()

end
function pin:on()

end
function pin:pinfunc()
    return 1
end
function pin:setdir()
    return 1
end
function pin:setpull()
    return 1
end
function pin:close()

end
function pin:index()

end
function pin:getval(val,...)

end
function pin:setval(val,...)
    --ui.PioSet(L.CheckInt(i),val)
end
