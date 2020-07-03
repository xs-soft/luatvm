
--module(..., package.seeall)
local modname = "pio"
local pin={

}
function pin:close(pin)
end
function pin:setpull()
end
function pin:getval()
end
function pin:setdir()
end
function pin:setval(val, pin)
end
local M = {
    pin=pin,
    P2_0=0,
}
_G[modname] = M
package.loaded[modname] = M