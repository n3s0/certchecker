package main

import (
    "fmt"
    "os"
    "time"
    "crypto/tls"
    "net/http"

    "github.com/spf13/cobra"
    "github.com/jedib0t/go-pretty/v6/list"
)

type CertInfo struct {
    ServerName string
    TlsVersion string
    CipherSuite string
    Subject string
    Signature string
    NotBefore string
    NotAfter string
}

var certInfo CertInfo
var host string
var port string

var rootCmd = &cobra.Command{
    Use:   "certchecker",
    Short: "Cert Checker checks that TLS certificates are valid.",
    Long: `Check that TLS certificates are valid.`,
    Run: func(cmd *cobra.Command, args []string) {
        if host != "" && port == "" {
            err := getTlsInfo(host, "")
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }

            certInfo.generateTlsInfoList()
        } else if host != "" && port != "" {
            err := getTlsInfo(host, port)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }

            certInfo.generateTlsInfoList()
        }
    },
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print version for certchecker",
    Long:  `Print version for certchecker.`,
    Run: func(cmd *cobra.Command, args []string) {
        const appVersion = "1.0.0"
        fmt.Printf("certmon v%s", appVersion)
    },
}

func main() {
    Execute()
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(versionCmd)

    rootCmd.Flags().StringVarP(&host, "server", "s", "", "hostname for site being tested (i.g. www.n3s0.tech)")
    rootCmd.Flags().StringVarP(&port, "port", "p", "", "port for site being tested, needs to be run with --server")
}

func getTlsInfo(host, port string) (e error) {
    var url string

    if port != "" {
        url = fmt.Sprintf("https://%s:%s", host, port)
    } else {
        url = fmt.Sprintf("https://%s", host)
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }

    client := &http.Client{Transport: tr}

    resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.TLS == nil {
		fmt.Println("No TLS connection established")
		return 
	}

	cert := resp.TLS.PeerCertificates[0]

    notBefore := cert.NotBefore.Local().Format(time.RFC1123)
    notAfter := cert.NotAfter.Local().Format(time.RFC1123)

    certInfo = CertInfo{ServerName: resp.TLS.ServerName, 
        TlsVersion: tls.VersionName(resp.TLS.Version),
        Subject: cert.Subject.String(),
        CipherSuite: tls.CipherSuiteName(resp.TLS.CipherSuite),
        NotBefore: notBefore, 
        NotAfter: notAfter}

    return nil
}

func (ci *CertInfo) generateTlsInfoList() {

    cdt := time.Now()
    clientTime := cdt.Format(time.RFC1123)

    na, err := time.Parse(time.RFC1123, ci.NotBefore)
    if err != nil {
        fmt.Println(err)
    }

    l := list.NewWriter()
    l.SetStyle(list.StyleConnectedRounded)

    l.AppendItem(fmt.Sprintf("Server: %s", ci.ServerName))
    l.Indent()
    l.AppendItem(fmt.Sprintf("TLS Version: %s", ci.TlsVersion))
    l.AppendItem(fmt.Sprintf("Cipher Suite: %s", ci.CipherSuite))
    l.AppendItem(fmt.Sprintf("Subject: %s", ci.Subject))
    l.UnIndent()
    l.AppendItem("Certificate Dates:")
    l.Indent()
    l.AppendItem(fmt.Sprintf("Not Before: %s", ci.NotBefore))
    l.AppendItem(fmt.Sprintf("Not After: %s", ci.NotAfter))
    if na.Before(cdt) {
        l.Indent()
        l.AppendItem("Date Check: Pass")
        l.UnIndent()
    } else {
        l.Indent()
        l.AppendItem("Date Check: Failed")
        l.AppendItem("Reason: Exceeds after date")
        l.UnIndent()
    }
    l.UnIndent()
    l.AppendItem("Client Information:")
    l.Indent()
    l.AppendItem(fmt.Sprintf("Local Date/Time: %v", clientTime))

    fmt.Printf("|---------------------------------------|\n")
    fmt.Printf("|-- Certificate Checker (certchecker) --|\n")
    fmt.Printf("|---------------------------------------|\n")
    fmt.Println(l.Render())
}
