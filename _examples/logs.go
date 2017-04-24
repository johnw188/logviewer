package main

import (
	"github.com/johnw188/logviewer"
	"time"
)

func main() {
	v := logviewer.NewViewer("john-ci")
	f1 := v.AddLogFeed("dns-1936050449-0hv0j", 100)
	f2 := v.AddLogFeed("node-exporter-95pc4", 100)
	f3 := v.AddLogFeed("john-ci-jenkins-694845882-lsc9h", 100)

	dns, exporter, jenkins := fakeData()

	t0 := time.NewTicker(time.Second * 1)
	t1 := time.NewTicker(time.Second * 1)
	t2 := time.NewTicker(time.Second * 1)

	i := 0
	j := 0
	k := 0

	go func() {
		for {
			select {
			case t := <-t0.C:
				if i < len(dns) {
					f1.AddLogLine(&logviewer.Line{
						Log:       dns[i],
						Timestamp: t,
					})
					i++
				}
			case t := <-t1.C:
				if j < len(exporter) {
					f2.AddLogLine(&logviewer.Line{
						Log:       exporter[j],
						Timestamp: t,
					})
					j++
				}
			case t := <-t2.C:
				if k < len(jenkins) {
					f3.AddLogLine(&logviewer.Line{
						Log:       jenkins[k],
						Timestamp: t,
					})
					k++
				}
			}
		}
	}()

	v.Display()
}

func fakeData() (dns []string, exporter []string, jenkins []string) {
	dns = []string{
		`I0414 23:05:48.311683       1 server.go:94] Using https://10.3.0.1:443 for kubernetes master, kubernetes API: <nil>`,
		`I0414 23:05:48.314159       1 server.go:99] v1.4.0-alpha.2.1652+c69e3d32a29cfa-dirty`,
		`I0414 23:05:48.314213       1 server.go:101] FLAG: --alsologtostderr="false"`,
		`I0414 23:05:48.314233       1 server.go:101] FLAG: --dns-port="10053"`,
		`I0414 23:05:48.314246       1 server.go:101] FLAG: --domain="cluster.local."`,
		`I0414 23:05:48.314261       1 server.go:101] FLAG: --federations=""`,
		`I0414 23:05:48.314271       1 server.go:101] FLAG: --healthz-port="8081"`,
		`I0414 23:05:48.314281       1 server.go:101] FLAG: --kube-master-url=""`,
		`I0414 23:05:48.314290       1 server.go:101] FLAG: --kubecfg-file=""`,
		`I0414 23:05:48.314307       1 server.go:101] FLAG: --log-backtrace-at=":0"`,
		`I0414 23:05:48.314318       1 server.go:101] FLAG: --log-dir=""`,
		`I0414 23:05:48.314327       1 server.go:101] FLAG: --log-flush-frequency="5s"`,
		`I0414 23:05:48.314336       1 server.go:101] FLAG: --logtostderr="true"`,
		`I0414 23:05:48.314347       1 server.go:101] FLAG: --stderrthreshold="2"`,
		`I0414 23:05:48.314359       1 server.go:101] FLAG: --v="0"`,
		`I0414 23:05:48.314366       1 server.go:101] FLAG: --version="false"`,
		`I0414 23:05:48.314377       1 server.go:101] FLAG: --vmodule=""`,
		`I0414 23:05:48.314646       1 server.go:138] Starting SkyDNS server. Listening on port:10053`,
		`I0414 23:05:48.316345       1 server.go:145] skydns: metrics enabled on : /metrics:tcp://10.3.0.207:9090`,
		`I0414 23:05:48.316394       1 dns.go:167] Waiting for service: default/kubernetes`,
		`I0414 23:05:48.409325       1 logs.go:41] skydns: ready for queries on cluster.local. for tcp://0.0.0.0:10053 [rcache 0]`,
		`I0414 23:05:48.409652       1 logs.go:41] skydns: ready for queries on cluster.local. for udp://0.0.0.0:10053 [rcache 0]`,
		`I0414 23:05:49.309207       1 server.go:107] Setting up Healthz Handler(/readiness, /cache) on port :8081`,
	}

	exporter = []string{
		`level=info msg="Starting node_exporter (version=0.12.0, branch=master, revision=61f36ac1ab87c40f7ce42de00908d7101c223217)" source="node_exporter.go:135"`,
		`level=info msg="Build context (go=go1.6.3, user=root@185b7034f1ed, date=20160911-16:10:38)" source="node_exporter.go:136"`,
		`level=info msg="No directory specified, see --collector.textfile.directory" source="textfile.go:57"`,
		`level=info msg="Enabled collectors:" source="node_exporter.go:155"`,
		`level=info msg=" - diskstats" source="node_exporter.go:157"`,
		`level=info msg=" - sockstat" source="node_exporter.go:157"`,
		`level=info msg=" - vmstat" source="node_exporter.go:157"`,
		`level=info msg=" - uname" source="node_exporter.go:157"`,
		`level=info msg=" - conntrack" source="node_exporter.go:157"`,
		`level=info msg=" - meminfo" source="node_exporter.go:157"`,
		`level=info msg=" - netstat" source="node_exporter.go:157"`,
		`level=info msg=" - stat" source="node_exporter.go:157"`,
		`level=info msg=" - textfile" source="node_exporter.go:157"`,
		`level=info msg=" - entropy" source="node_exporter.go:157"`,
		`level=info msg=" - loadavg" source="node_exporter.go:157"`,
		`level=info msg=" - mdadm" source="node_exporter.go:157"`,
		`level=info msg=" - netdev" source="node_exporter.go:157"`,
		`level=info msg=" - time" source="node_exporter.go:157"`,
		`level=info msg=" - filefd" source="node_exporter.go:157"`,
		`level=info msg=" - filesystem" source="node_exporter.go:157"`,
		`level=info msg="Listening on :9100" source="node_exporter.go:176"`,
	}

	jenkins = []string{
		`javax.naming.AuthenticationException: [LDAP: error code 49 - 80090308: LdapErr: DSID-0C0903A9, comment: AcceptSecurityContext error, data 775, v1db1]`,
		`	at com.sun.jndi.ldap.LdapCtx.mapErrorCode(LdapCtx.java:3136)`,
		`	at com.sun.jndi.ldap.LdapCtx.processReturnCode(LdapCtx.java:3082)`,
		`	at com.sun.jndi.ldap.LdapCtx.processReturnCode(LdapCtx.java:2883)`,
		`	at com.sun.jndi.ldap.LdapCtx.connect(LdapCtx.java:2797)`,
		`	at com.sun.jndi.ldap.LdapCtx.<init>(LdapCtx.java:319)`,
		`	at com.sun.jndi.ldap.LdapCtxFactory.getUsingURL(LdapCtxFactory.java:192)`,
		`	at com.sun.jndi.ldap.LdapCtxFactory.getUsingURLs(LdapCtxFactory.java:210)`,
		`	at com.sun.jndi.ldap.LdapCtxFactory.getLdapCtxInstance(LdapCtxFactory.java:153)`,
		`	at com.sun.jndi.ldap.LdapCtxFactory.getInitialContext(LdapCtxFactory.java:83)`,
		`	at javax.naming.spi.NamingManager.getInitialContext(NamingManager.java:684)`,
		`	at javax.naming.InitialContext.getDefaultInitCtx(InitialContext.java:313)`,
		`	at javax.naming.InitialContext.init(InitialContext.java:244)`,
		`	at javax.naming.InitialContext.<init>(InitialContext.java:216)`,
		`	at javax.naming.directory.InitialDirContext.<init>(InitialDirContext.java:101)`,
		`	at org.acegisecurity.ldap.DefaultInitialDirContextFactory.connect(DefaultInitialDirContextFactory.java:180)`,
		`	at org.acegisecurity.ldap.DefaultInitialDirContextFactory.newInitialDirContext(DefaultInitialDirContextFactory.java:261)`,
		`	at org.acegisecurity.ldap.LdapTemplate.execute(LdapTemplate.java:123)`,
		`	at org.acegisecurity.ldap.LdapTemplate.retrieveEntry(LdapTemplate.java:165)`,
		`	at org.acegisecurity.providers.ldap.authenticator.BindAuthenticator.bindWithDn(BindAuthenticator.java:87)`,
		`	at org.acegisecurity.providers.ldap.authenticator.BindAuthenticator.authenticate(BindAuthenticator.java:72)`,
		`	at org.acegisecurity.providers.ldap.authenticator.BindAuthenticator2.authenticate(BindAuthenticator2.java:49)`,
		`	at org.acegisecurity.providers.ldap.LdapAuthenticationProvider.retrieveUser(LdapAuthenticationProvider.java:233)`,
		`	at org.acegisecurity.providers.dao.AbstractUserDetailsAuthenticationProvider.authenticate(AbstractUserDetailsAuthenticationProvider.java:122)`,
		`	at org.acegisecurity.providers.ProviderManager.doAuthentication(ProviderManager.java:200)`,
		`	at org.acegisecurity.AbstractAuthenticationManager.authenticate(AbstractAuthenticationManager.java:47)`,
		`	at hudson.security.LDAPSecurityRealm$LDAPAuthenticationManager.authenticate(LDAPSecurityRealm.java:846)`,
		`	at jenkins.security.BasicHeaderRealPasswordAuthenticator.authenticate(BasicHeaderRealPasswordAuthenticator.java:56)`,
		`	at jenkins.security.BasicHeaderProcessor.doFilter(BasicHeaderProcessor.java:79)`,
		`	at hudson.security.ChainedServletFilter$1.doFilter(ChainedServletFilter.java:87)`,
		`	at org.acegisecurity.context.HttpSessionContextIntegrationFilter.doFilter(HttpSessionContextIntegrationFilter.java:249)`,
		`	at hudson.security.HttpSessionContextIntegrationFilter2.doFilter(HttpSessionContextIntegrationFilter2.java:67)`,
		`	at hudson.security.ChainedServletFilter$1.doFilter(ChainedServletFilter.java:87)`,
		`	at hudson.security.ChainedServletFilter.doFilter(ChainedServletFilter.java:76)`,
		`	at hudson.security.HudsonFilter.doFilter(HudsonFilter.java:171)`,
		`	at org.eclipse.jetty.servlet.ServletHandler$CachedChain.doFilter(ServletHandler.java:1652)`,
		`	at org.kohsuke.stapler.compression.CompressionFilter.doFilter(CompressionFilter.java:49)`,
		`	at org.eclipse.jetty.servlet.ServletHandler$CachedChain.doFilter(ServletHandler.java:1652)`,
		`	at hudson.util.CharacterEncodingFilter.doFilter(CharacterEncodingFilter.java:82)`,
		`	at org.eclipse.jetty.servlet.ServletHandler$CachedChain.doFilter(ServletHandler.java:1652)`,
		`	at org.kohsuke.stapler.DiagnosticThreadNameFilter.doFilter(DiagnosticThreadNameFilter.java:30)`,
		`	at org.eclipse.jetty.servlet.ServletHandler$CachedChain.doFilter(ServletHandler.java:1652)`,
		`	at org.eclipse.jetty.servlet.ServletHandler.doHandle(ServletHandler.java:585)`,
		`	at org.eclipse.jetty.server.handler.ScopedHandler.handle(ScopedHandler.java:143)`,
		`	at org.eclipse.jetty.security.SecurityHandler.handle(SecurityHandler.java:553)`,
		`	at org.eclipse.jetty.server.session.SessionHandler.doHandle(SessionHandler.java:223)`,
		`	at org.eclipse.jetty.server.handler.ContextHandler.doHandle(ContextHandler.java:1127)`,
		`	at org.eclipse.jetty.servlet.ServletHandler.doScope(ServletHandler.java:515)`,
		`	at org.eclipse.jetty.server.session.SessionHandler.doScope(SessionHandler.java:185)`,
		`	at org.eclipse.jetty.server.handler.ContextHandler.doScope(ContextHandler.java:1061)`,
		`	at org.eclipse.jetty.server.handler.ScopedHandler.handle(ScopedHandler.java:141)`,
		`	at org.eclipse.jetty.server.handler.HandlerWrapper.handle(HandlerWrapper.java:97)`,
		`	at org.eclipse.jetty.server.Server.handle(Server.java:499)`,
		`	at org.eclipse.jetty.server.HttpChannel.handle(HttpChannel.java:311)`,
		`	at org.eclipse.jetty.server.HttpConnection.onFillable(HttpConnection.java:257)`,
		`	at org.eclipse.jetty.io.AbstractConnection$2.run(AbstractConnection.java:544)`,
		`	at winstone.BoundedExecutorService$1.run(BoundedExecutorService.java:77)`,
		`	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)`,
		`	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)`,
		`	at java.lang.Thread.run(Thread.java:745)`,
		``,
		`Apr 19, 2017 10:33:40 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Download metadata. 20,035 ms`,
		`Apr 20, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Workspace clean-up`,
		`Apr 20, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Workspace clean-up. 0 ms`,
		`Apr 20, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Fingerprint cleanup`,
		`Apr 20, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Fingerprint cleanup. 0 ms`,
		`Apr 20, 2017 6:24:34 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Connection Activity monitoring to agents`,
		`Apr 20, 2017 6:24:34 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Connection Activity monitoring to agents. 2 ms`,
		`Apr 20, 2017 10:33:20 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Download metadata`,
		`Apr 20, 2017 10:33:40 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Download metadata. 20,037 ms`,
		`Apr 21, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Workspace clean-up`,
		`Apr 21, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Workspace clean-up. 2 ms`,
		`Apr 21, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Fingerprint cleanup`,
		`Apr 21, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Fingerprint cleanup. 2 ms`,
		`Apr 21, 2017 10:33:20 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Download metadata`,
		`Apr 21, 2017 10:33:40 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Download metadata. 20,107 ms`,
		`Apr 22, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Workspace clean-up`,
		`Apr 22, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Workspace clean-up. 0 ms`,
		`Apr 22, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Fingerprint cleanup`,
		`Apr 22, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Fingerprint cleanup. 0 ms`,
		`Apr 22, 2017 10:33:20 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Download metadata`,
		`Apr 22, 2017 10:33:40 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Download metadata. 20,080 ms`,
		`Apr 23, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Workspace clean-up`,
		`Apr 23, 2017 7:23:02 AM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Workspace clean-up. 4 ms`,
		`Apr 23, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Fingerprint cleanup`,
		`Apr 23, 2017 2:38:49 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Fingerprint cleanup. 5 ms`,
		`Apr 23, 2017 10:33:20 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Started Download metadata`,
		`Apr 23, 2017 10:33:40 PM hudson.model.AsyncPeriodicWork$1 run`,
		`INFO: Finished Download metadata. 20,098 ms`,
	}

	return
}
