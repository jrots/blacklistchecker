package blacklistchecker

var BlacklistEntries []string = []string{
	"http.dnsbl.sorbs.net",
	"socks.dnsbl.sorbs.net",
	"misc.dnsbl.sorbs.net",
	"smtp.dnsbl.sorbs.net",
	"new.spam.dnsbl.sorbs.net",
	"spam.dnsbl.sorbs.net",
	"web.dnsbl.sorbs.net",
	"zombie.dnsbl.sorbs.net",
	"noptr.spamrats.com",
	"spam.spamrats.com",
	"b.barracudacentral.org",
	"dnsbl.justspam.org",
	"xbl.spamhaus.org",
	"sbl.spamhaus.org",
	"cbl.anti-spam.org.cn",
	"cdl.anti-spam.org.cn",
	"access.redhawk.org",
	"bl.spamcop.net",
	"blackholes.mail-abuse.org",
	"bogons.cymru.com",
	"cbl.abuseat.org",
	"combined.njabl.org",
	"csi.cloudmark.com",
	"db.wpbl.info",
	"dnsbl.dronebl.org",
	"dnsbl.inps.de",
	"dnsbl.njabl.org",
	"drone.abuse.ch",
	"dsn.rfc-ignorant.org",
	"httpbl.abuse.ch",
	"ips.backscatterer.org",
	"ix.dnsbl.manitu.net",
	"multi.surbl.org",
	"netblock.pedantic.org",
	"opm.tornevall.org",
	"psbl.surriel.com",
	"rf.senderbase.org",
	"rbl-plus.mail-abuse.org",
	"rbl.efnetrbl.org",
	"rbl.interserver.net",
	"rbl.spamlab.com",
	"rbl.suresupport.com",
	"relays.mail-abuse.org",
	"spamguard.leadmon.net",
	"spamrbl.imp.ch",
	"tor.dan.me.uk",
	"ubl.unsubscore.com",
	"virbl.bit.nl",
	"xbl.abuseat.org",
}

func GetBlacklistHosts() []string {
	return BlacklistEntries
}
