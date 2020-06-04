// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista

package apiserver

import (
	"github.com/cloud-barista/cb-tumblebug/src/common"
	"github.com/cloud-barista/cb-tumblebug/src/mcir"
	"github.com/cloud-barista/cb-tumblebug/src/mcis"

	//"os"

	"fmt"

	// REST API (echo)
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// CB-Store
)

/*
// CB-Store
var cblog *logrus.Logger
var store icbs.Store

func init() {
	cblog = config.Cblogger
	store = cbstore.GetStore()
}

type KeyValue struct {
	Key   string
	Value string
}
*/

//var masterConfigInfos confighandler.MASTERCONFIGTYPE

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	Version = " Version: Cappuccino"
	website = " Repository: https://github.com/cloud-barista/cb-tumblebug"
	banner  = `

  ██████╗██╗      ██████╗ ██╗   ██╗██████╗       ██████╗  █████╗ ██████╗ ██╗███████╗████████╗ █████╗
 ██╔════╝██║     ██╔═══██╗██║   ██║██╔══██╗      ██╔══██╗██╔══██╗██╔══██╗██║██╔════╝╚══██╔══╝██╔══██╗
 ██║     ██║     ██║   ██║██║   ██║██║  ██║█████╗██████╔╝███████║██████╔╝██║███████╗   ██║   ███████║
 ██║     ██║     ██║   ██║██║   ██║██║  ██║╚════╝██╔══██╗██╔══██║██╔══██╗██║╚════██║   ██║   ██╔══██║
 ╚██████╗███████╗╚██████╔╝╚██████╔╝██████╔╝      ██████╔╝██║  ██║██║  ██║██║███████║   ██║   ██║  ██║
  ╚═════╝╚══════╝ ╚═════╝  ╚═════╝ ╚═════╝       ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚══════╝   ╚═╝   ╚═╝  ╚═╝

 ████████╗██╗   ██╗███╗   ███╗██████╗ ██╗     ███████╗██████╗ ██╗   ██╗ ██████╗
 ╚══██╔══╝██║   ██║████╗ ████║██╔══██╗██║     ██╔════╝██╔══██╗██║   ██║██╔════╝
    ██║   ██║   ██║██╔████╔██║██████╔╝██║     █████╗  ██████╔╝██║   ██║██║  ███╗
    ██║   ██║   ██║██║╚██╔╝██║██╔══██╗██║     ██╔══╝  ██╔══██╗██║   ██║██║   ██║
    ██║   ╚██████╔╝██║ ╚═╝ ██║██████╔╝███████╗███████╗██████╔╝╚██████╔╝╚██████╔╝
    ╚═╝    ╚═════╝ ╚═╝     ╚═╝╚═════╝ ╚══════╝╚══════╝╚═════╝  ╚═════╝  ╚═════╝              

 Multi-cloud infra service managemenet framework
 ________________________________________________`
)

// Main Body

func ApiServer() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! This is cloud-barista cb-tumblebug")
	})
	e.HideBanner = true
	//e.colorer.Printf(banner, e.colorer.Red("v"+Version), e.colorer.Blue(website))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Printf(banner)
	fmt.Println("")
	fmt.Printf(ErrorColor, Version)
	fmt.Println("")
	fmt.Printf(InfoColor, website)
	fmt.Println("")
	fmt.Println("")

	// Route
	g := e.Group("/tumblebug/ns", common.NsValidation())

	g.POST("", common.RestPostNs)
	g.GET("/:nsId", common.RestGetNs)
	g.GET("", common.RestGetAllNs)
	g.PUT("/:nsId", common.RestPutNs)
	g.DELETE("/:nsId", common.RestDelNs)
	g.DELETE("", common.RestDelAllNs)

	g.POST("/:nsId/mcis", mcis.RestPostMcis)
	g.GET("/:nsId/mcis/:mcisId", mcis.RestGetMcis)
	g.GET("/:nsId/mcis", mcis.RestGetAllMcis)
	g.PUT("/:nsId/mcis/:mcisId", mcis.RestPutMcis)
	g.DELETE("/:nsId/mcis/:mcisId", mcis.RestDelMcis)
	g.DELETE("/:nsId/mcis", mcis.RestDelAllMcis)

	g.POST("/:nsId/mcis/:mcisId/vm", mcis.RestPostMcisVm)
	g.GET("/:nsId/mcis/:mcisId/vm/:vmId", mcis.RestGetMcisVm)
	//g.GET("/:nsId/mcis", mcis.RestGetAllMcis)
	//g.PUT("/:nsId/mcis/:mcisId", mcis.RestPutMcis)
	g.DELETE("/:nsId/mcis/:mcisId/vm/:vmId", mcis.RestDelMcisVm)
	//g.DELETE("/:nsId/mcis", mcis.RestDelAllMcis)

	g.POST("/:nsId/mcis/recommend", mcis.RestPostMcisRecommand)
	g.POST("/:nsId/cmd/mcis/:mcisId", mcis.RestPostCmdMcis)
	g.POST("/:nsId/cmd/mcis/:mcisId/vm/:vmId", mcis.RestPostCmdMcisVm)
	g.POST("/:nsId/install/mcis/:mcisId", mcis.RestPostInstallAgentToMcis)
	g.GET("/:nsId/benchmark/mcis/:mcisId", mcis.RestGetBenchmark)
	g.GET("/:nsId/benchmarkall/mcis/:mcisId", mcis.RestGetAllBenchmark)

	

	g.POST("/:nsId/resources/image", mcir.RestPostImage)
	g.GET("/:nsId/resources/image/:imageId", mcir.RestGetImage)
	g.GET("/:nsId/resources/image", mcir.RestGetAllImage)
	g.PUT("/:nsId/resources/image/:imageId", mcir.RestPutImage)
	g.DELETE("/:nsId/resources/image/:imageId", mcir.RestDelImage)
	g.DELETE("/:nsId/resources/image", mcir.RestDelAllImage)

	g.POST("/:nsId/resources/sshKey", mcir.RestPostSshKey)
	g.GET("/:nsId/resources/sshKey/:sshKeyId", mcir.RestGetSshKey)
	g.GET("/:nsId/resources/sshKey", mcir.RestGetAllSshKey)
	g.PUT("/:nsId/resources/sshKey/:sshKeyId", mcir.RestPutSshKey)
	g.DELETE("/:nsId/resources/sshKey/:sshKeyId", mcir.RestDelSshKey)
	g.DELETE("/:nsId/resources/sshKey", mcir.RestDelAllSshKey)

	g.POST("/:nsId/resources/spec", mcir.RestPostSpec)
	g.GET("/:nsId/resources/spec/:specId", mcir.RestGetSpec)
	g.GET("/:nsId/resources/spec", mcir.RestGetAllSpec)
	g.PUT("/:nsId/resources/spec/:specId", mcir.RestPutSpec)
	g.DELETE("/:nsId/resources/spec/:specId", mcir.RestDelSpec)
	g.DELETE("/:nsId/resources/spec", mcir.RestDelAllSpec)

	g.POST("/:nsId/resources/fetchSpecs", mcir.RestFetchSpecs)

	g.POST("/:nsId/resources/securityGroup", mcir.RestPostSecurityGroup)
	g.GET("/:nsId/resources/securityGroup/:securityGroupId", mcir.RestGetSecurityGroup)
	g.GET("/:nsId/resources/securityGroup", mcir.RestGetAllSecurityGroup)
	g.PUT("/:nsId/resources/securityGroup/:securityGroupId", mcir.RestPutSecurityGroup)
	g.DELETE("/:nsId/resources/securityGroup/:securityGroupId", mcir.RestDelSecurityGroup)
	g.DELETE("/:nsId/resources/securityGroup", mcir.RestDelAllSecurityGroup)

	g.POST("/:nsId/resources/vNet", mcir.RestPostVNet)
	g.GET("/:nsId/resources/vNet/:vNetId", mcir.RestGetVNet)
	g.GET("/:nsId/resources/vNet", mcir.RestGetAllVNet)
	g.PUT("/:nsId/resources/vNet/:vNetId", mcir.RestPutVNet)
	g.DELETE("/:nsId/resources/vNet/:vNetId", mcir.RestDelVNet)
	g.DELETE("/:nsId/resources/vNet", mcir.RestDelAllVNet)

	/*
		g.POST("/:nsId/resources/subnet", mcir.RestPostSubnet)
		g.GET("/:nsId/resources/subnet/:subnetId", mcir.RestGetSubnet)
		g.GET("/:nsId/resources/subnet", mcir.RestGetAllSubnet)
		g.PUT("/:nsId/resources/subnet/:subnetId", mcir.RestPutSubnet)
		g.DELETE("/:nsId/resources/subnet/:subnetId", mcir.RestDelSubnet)
		g.DELETE("/:nsId/resources/subnet", mcir.RestDelAllSubnet)

		g.POST("/:nsId/resources/publicIp", mcir.RestPostPublicIp)
		g.GET("/:nsId/resources/publicIp/:publicIpId", mcir.RestGetPublicIp)
		g.GET("/:nsId/resources/publicIp", mcir.RestGetAllPublicIp)
		g.PUT("/:nsId/resources/publicIp/:publicIpId", mcir.RestPutPublicIp)
		g.DELETE("/:nsId/resources/publicIp/:publicIpId", mcir.RestDelPublicIp)
		g.DELETE("/:nsId/resources/publicIp", mcir.RestDelAllPublicIp)

		g.POST("/:nsId/resources/vNic", mcir.RestPostVNic)
		g.GET("/:nsId/resources/vNic/:vNicId", mcir.RestGetVNic)
		g.GET("/:nsId/resources/vNic", mcir.RestGetAllVNic)
		g.PUT("/:nsId/resources/vNic/:vNicId", mcir.RestPutVNic)
		g.DELETE("/:nsId/resources/vNic/:vNicId", mcir.RestDelVNic)
		g.DELETE("/:nsId/resources/vNic", mcir.RestDelAllVNic)
	*/

	// We cannot use these wildcard method below.
	// https://github.com/labstack/echo/issues/382
	//g.DELETE("/:nsId/resources/:resourceType/:resourceId", mcir.RestDelResource)
	//g.DELETE("/:nsId/resources/:resourceType", mcir.RestDelAllResources)

	g.GET("/:nsId/checkResource/:resourceType/:resourceId", mcir.RestCheckResource)
	g.GET("/:nsId/checkMcis/:mcisId", mcis.RestCheckMcis)
	g.GET("/:nsId/mcis/:mcisId/checkVm/:vmId", mcis.RestCheckVm)

	e.GET("/tumblebug/checkNs/:nsId", common.RestCheckNs)

	e.GET("/tumblebug/connConfig", common.RestGetConnConfigList)
	e.GET("/tumblebug/connConfig/:connConfigName", common.RestGetConnConfig)
	e.GET("/tumblebug/region", common.RestGetRegionList)
	e.GET("/tumblebug/region/:regionName", common.RestGetRegion)

	e.GET("/tumblebug/lookupSpec", mcir.RestLookupSpecList)
	e.GET("/tumblebug/lookupSpec/:specName", mcir.RestLookupSpec)

	e.Logger.Fatal(e.Start(":1323"))

}

var SPIDER_URL string

/*
func main() {

	//fmt.Println("\n[cb-tumblebug (Multi-Cloud Infra Service Management Framework)]")
	//fmt.Println("\nInitiating REST API Server ...")
	//fmt.Println("\n[REST API call examples]")

	SPIDER_URL = os.Getenv("SPIDER_URL")

	// load config
	//masterConfigInfos = confighandler.GetMasterConfigInfos()

	// Run API Server
	apiServer()

}
*/
