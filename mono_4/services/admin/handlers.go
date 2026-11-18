package main

import (
	"net/http"
	"os"
	"os/exec"
)

// CWE-78 chain: handler -> RunCommand -> sh -c.
func ExecHandler(w http.ResponseWriter, r *http.Request) {
	command := r.FormValue("cmd")
	output, err := RunCommand(command)
	if err != nil {
		http.Error(w, string(output)+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

// CWE-78 direct: unquoted host reaches a shell in the same function.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	host := r.FormValue("host")
	command := exec.Command("sh", "-c", "ping -c 1 "+host)
	output, _ := command.CombinedOutput()
	w.Write(output)
}

// CWE-95: request body evaluated as a shell expression.
func EvalHandler(w http.ResponseWriter, r *http.Request) {
	expression := r.FormValue("expression")
	output := RunShell("echo $(( " + expression + " ))")
	w.Write(output)
}

// CWE-73: recursive delete from a request value.
func WipeHandler(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("path")
	DeleteEverything(target)
	w.Write([]byte("wiped"))
}

func ChmodHandler(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("path")
	os.Chmod(target, 0777)
	w.Write([]byte("world writable"))
}

func KillHandler(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("pid")
	KillProcess(target)
	w.Write([]byte("killed"))
}
