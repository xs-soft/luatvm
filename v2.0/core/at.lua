--直接返回OK指令的AT命令   判定前会去掉AT+或者AT*开头

resendok={"ATE0","CMEE","CREG","CGREG","CEREG","MRD_CDF","EEMOPT","STON","CMGF","CIPMUX","CIPHEAD","CIPQSEND","CENG",
                 	"CSMP","CSCS","CPMS","TCPUSERPARAM","CIICRMODE","MEDCR","CSTT","CMIC","CNMI","CCLK","MRD_CDF","QTTS"}
--[[
resend={


	["AT+CIND=1"]={"OK",'+CREG: 11,"3704","3703532"'},
	["AT+VER"]={"Luat_V0025_ASR1802_FLOAT_720H","OK"},
	["AT+ATWMFT=99"]={"SUCC","OK"},
	["AT+WISN?"]={"*CME ERROR: Missing SN","OK"},
	["AT+CGSN"]={ATCGSN},
	["AT+CALIBINFO?"]={"1","OK"},
	["AT+MUID?"]={'+MUID: "20190925033009A802082A3644464284"',"OK"},
	["AT+ICCID"]={"+ICCID: 898602b2151840000759","OK"},
	["AT+CIMI"]={"460040214210759","OK",
		"^MODE: 17,17",
		'+CREG: 1,"6498","585a702"',
		"^MODE: 17,17",
		'+CEREG: 1,"6498","0585a702",7',
		'+CGREG: 1,"6498","0585a702"',
		'+CGEV: EPS PDN ACT 5',
		"+NITZ: 19/12/06,06:16:37+32,0",
	},
	["AT+CGDCONT?"]={'+CGDCONT: 5,"IP","cmiot.MNC004.MCC460.GPRS","100.104.232.182",,,802110030100108106d38814cb8306d388116b000d04d3',"OK"},
	["AT+CENG?"]={"+CENG:0,\"573,24,99,460,0,13,49234,10,0,6311,255\"","OK"},
	["AT+CCID"]={"898602b2151840000759","OK"},
	["AT+CGATT?"]={"+CGATT: 1","OK"},
	["AT+CIPSTATUS"]={"OK","STATE: IP STATUS"},
	["AT+CIICR"]={"OK","STATE: IP GPRSACT"},
	["AT+CIFSR"]={"010.099.214.051"},
	["AT+CDNSCFG?"]={"PrimaryDns: 211.136.17.107","SecondaryDns: 211.136.20.203","OK"}
}
]]--
resend={}

--resend["AT+CSQ"]={"+CSQ: 23,99","OK"}
resend = {
    {"AT+CSQ","+CSQ: 23,99","OK"},
    {"AT+CREG?",'+CREG: 2,0,"3704","3703532"',"OK"},
    {"AT+CGREG?","+CGREG: 2,0","OK"},
    {"AT+CEREG?",'+CEREG: 2,11,"3704","03703532",7','OK'},
    {"AT*BAND?","*BAND:15,74,129,482,149,0,2,0","OK"},
    {"AT+WIMEI","OK"},
    {"AT+EEMGINFO?",
    		"+EEMLTESVC:1120, 2, 17, 14084, 38, 1825, 19825, 3, 4, 57685298, 0, 59, 19, 11, 59, 29, 24, 127, 73, 0, 0, 0, 1, 10, 0, 0, 65535, 779, 146, -419014374",
    		"+EEMGINFO : 3, 2",
    		"OK", "+NITZ: 19/12/06,06:16:37+32,0"},

	{"AT+CIND=1","OK",'+CREG: 11,"3704","3703532"'},
	{"AT+VER","Luat_V0025_ASR1802_FLOAT_720H","OK"},
	{"AT+ATWMFT=99","SUCC","OK"},
	{"AT+WISN?","*CME ERROR: Missing SN","OK"},
	{"AT+CALIBINFO?","1","OK"},
	{"AT+MUID?",'+MUID: "20190925033009A802082A3644464284"',"OK"},
	{"AT+ICCID","+ICCID: 898602b2151840000759","OK"},
	{"AT+CIMI","460040214210759","OK",
    		"^MODE: 17,17",
    		'+CREG: 1,"6498","585a702"',
    		"^MODE: 17,17",
    		'+CEREG: 1,"6498","0585a702",7',
    		'+CGREG: 1,"6498","0585a702"',
    		'+CGEV: EPS PDN ACT 5',
    		"+NITZ: 19/12/06,06:16:37+32,0",
    	},
	{"AT+CGDCONT?",'+CGDCONT: 5,"IP","cmiot.MNC004.MCC460.GPRS","100.104.232.182",,,802110030100108106d38814cb8306d388116b000d04d3',"OK"},
	{"AT+CENG?","+CENG:0,\"573,24,99,460,0,13,49234,10,0,6311,255\"","OK"},
	{"AT+CCID","898602b2151840000759","OK"},
	{"AT+CGATT?","+CGATT: 1","OK"},
	{"AT+CIPSTATUS","OK","STATE: IP STATUS"},
	{"AT+CIICR","OK","STATE: IP GPRSACT"},
	{"AT+CIFSR","010.099.214.051"},
	{"AT+CDNSCFG?","PrimaryDns: 211.136.17.107","SecondaryDns: 211.136.20.203","OK"},
}
local function IsInTable(value, tbl)
for k,v in ipairs(tbl) do
  if v == value then
  return true;
  end
end
return false;
end

local function IsInTableKey(value, tbl)
for k,v in ipairs(tbl) do
    print("k",k)
  if k == value then
  return true;
  end
end
return false;
end


local sendcsq=false
uart.rilAT(function (s)
    tmp_s = string.gsub(s,"^AT%p([%a%d_]+)=*.*","%1")
    if IsInTable(tmp_s,resendok) then
        uart.rilATSend("OK")
        return
    end

    for k,v in ipairs(resend) do
        if v[1]==s then
            for  i =2,#v do
                uart.rilATSend(v[i])
            end
        end
    end
    if s=="AT+CSQ" and not sendcsq then
        sendcsq=true
    	uart.rilATSend("898602b2151840000759")
    	uart.rilATSend("^SIMST: 0")
    	uart.rilATSend("*ADMINDATA: 0, 2, 0")
    	uart.rilATSend("+CPIN: READY")
    	uart.rilATSend("^SIMST: 1")
    end
end
)