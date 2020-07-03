oldopen=io.open
io.open=function (p1,p2)
    if #p1>1 and string.sub(p1,1,1)=="/" then
        p1=string.sub(p1,2,-1)
    end
    file,r2,r3= oldopen(p1,p2)
    if file~=nil then
        return file
    else
        return nil,r2,r3
    end
end

oldrename=os.rename
io.rename=function (p1,p2)
    if #p1>1 and string.sub(p1,1,1)=="/" then
        p1=string.sub(p1,2,-1)
    end
    if #p2>2 and string.sub(p2,1,1)=="/" then
        p2=string.sub(p2,2,-1)
    end
    return oldrename(p1,p2)
end