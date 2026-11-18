package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// CWE-306: none of these routes has authentication or authorization.
	mux.HandleFunc("/exec", ExecHandler)
	mux.HandleFunc("/ping", PingHandler)
	mux.HandleFunc("/eval", EvalHandler)
	mux.HandleFunc("/audit", AuditHandler)
	mux.HandleFunc("/logs", LogHandler)
	mux.HandleFunc("/wipe", WipeHandler)
	mux.HandleFunc("/env", EnvHandler)
	mux.HandleFunc("/pprof", PprofHandler)
	mux.HandleFunc("/unsafe", UnsafeHandler)
	mux.HandleFunc("/chmod", ChmodHandler)
	mux.HandleFunc("/kill", KillHandler)

	http.ListenAndServe("0.0.0.0:8080", mux)
}
