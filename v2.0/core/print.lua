oldprint=print
notprintexps={
    "%[I%]%-%[ril",
    "%[I%]%-%[net%.rsp",
}
print=function (...)
    if not DISPSYSLOG then
        for k,v in ipairs({...}) do
            if k==1 then
                for pk,pv in ipairs(notprintexps) do
                    if string.find(v,pv) then
                        return
                    end
                end
            end
        end
    end
    oldprint(...)
end