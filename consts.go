package goValidin

const dnsInfra = "https://pilot.validin.com/api/axon/domain/dns/history/%s"
const dnsInfraA = "https://pilot.validin.com/api/axon/domain/dns/history/%s/A"
const dnsInfraAAAA = "https://pilot.validin.com/api/axon/domain/dns/history/%s/AAAA"
const dnsInfraNS = "https://pilot.validin.com/api/axon/domain/dns/history/%s/NS"
const dnsInfraNSFOR = "https://pilot.validin.com/api/axon/domain/dns/history/%s/NS_FOR"
const dnsPtr = "https://pilot.validin.com/api/axon/domain/dns/hostname/%s"
const dnsExtra = "https://pilot.validin.com/api/axon/domain/dns/extra/%s"
const dnsSubDomain = "https://pilot.validin.com/api/axon/domain/subdomains/%s"
const dnsOsintHistory = "https://pilot.validin.com/api/axon/domain/osint/history/%s"
const dnsOsintContext = "https://pilot.validin.com/api/axon/domain/osint/context/%s"
const dnsFastRep = "https://pilot.validin.com/api/axon/domain/reputation/quick/%s"
const HostResponsePivots = "https://pilot.validin.com/api/axon/domain/pivots/%s"
const dnsCtStreamLogs = "https://pilot.validin.com/api/axon/domain/certificates/%s"
const dnsRegStreamLogs = "https://pilot.validin.com/api/axon/domain/registration/history/%s"
const dnsLiveReg = "https://pilot.validin.com/api/axon/domain/registration/live/%s"

const ipDnsHistory = "https://pilot.validin.com/api/axon/ip/dns/history/%s"
const ipDnsHistoryCidr = "https://pilot.validin.com/api/axon/ip/dns/history/%s/%s" // IP/CIDR
const ipPtrRecords = "https://pilot.validin.com/api/axon/ip/dns/hostname/%s"
const ipPtrRecordsCidr = "https://pilot.validin.com/api/axon/ip/dns/hostname/%s/%s" // IP/CIDR
const ipExtra = "https://pilot.validin.com/api/axon/ip/dns/extra/%s"
const ipExtraCidr = "https://pilot.validin.com/api/axon/ip/dns/extra/%s/%s" // IP/CIDR
const ipOsintHistory = "https://pilot.validin.com/api/axon/ip/osint/history/%s"
const ipOsintHistoryCidr = "https://pilot.validin.com/api/axon/ip/osint/history/%s/%s" // IP/CIDR
const ipFastRep = "https://pilot.validin.com/api/axon/ip/reputation/quick/%s"
const ipOsintContext = "https://pilot.validin.com/api/axon/ip/osint/context/%s"
const ipHostResponsePivots = "https://pilot.validin.com/api/axon/ip/pivots/%s"
const ipHostResponsePivotsCidr = "https://pilot.validin.com/api/axon/ip/pivots/%s/%s" // IP/CIDR

const hashHostResponsePivots = "https://pilot.validin.com/api/axon/hash/pivots/%s"

const stringExtraDnsRecords = "https://pilot.validin.com/api/axon/string/dns/extra/%s"
const stringHostResponsePivots = "https://pilot.validin.com/api/axon/string/pivots/%s"
const stringRegHistory = "https://pilot.validin.com/api/axon/string/registration/history/%s"

const utilApiUsage = "https://pilot.validin.com/api/profile/usage"
const utilAllPaths = "https://pilot.validin.com/api/paths"
const utilPing = "https://pilot.validin.com/api/ping"

const projectAll = "https://pilot.validin.com/api/projects"
const projectInfo = "https://pilot.validin.com/api/project/%s"
const projectAllIndicators = "https://pilot.validin.com/api/projects/%s/indicators"
const projectAlertDates = "https://pilot.validin.com/api/projects/%s/alerts/dates"
const RecentAlerts = "https://pilot.validin.com/api/projects/%s/alerts/latest"
const projectAlertsOnDate = "https://pilot.validin.com/api/projects/%s/alerts/%s"
