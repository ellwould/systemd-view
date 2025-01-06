/*
MIT License

Copyright (c) 2024 Elliot Michael Keavney

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"systemdviewresource"
	"runtime"
	"strings"
)

func systemd() {
	var startHTML string
	startHTML = systemdviewresource.StartHTML()

	var endHTML string
	endHTML = systemdviewresource.EndHTML()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, startHTML)
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><a href=\"https://ell.today\" class=\"tableButton externalButton\">Written By Elliot Keavney (Website)</a></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><a href=\"https://github.com/Ellwould/systemd-view\" class=\"tableButton externalButton\">Systemd View Source Code (GitHub)</a></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th>&nbsp<button onclick=\"toggleKeyTable() \"class=\"tableButton\">Hide/Show<br>Key</button>&nbsp</th>")
		fmt.Fprintf(w, "    <th><h3>&nbsp &nbsp &nbsp Background Process (Deamon) Infomation &nbsp &nbsp &nbsp</h3></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<div id=\"keyTable\">")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th>Status/Symbol</th>")
		fmt.Fprintf(w, "    <th>Description</th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>&#128994</td>")
		fmt.Fprintf(w, "    <td>Service is active</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>&#128308</td>")
		fmt.Fprintf(w, "    <td>Service is not active</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Enabled &#9989</td>")
		fmt.Fprintf(w, "    <td>Service automatically starts on boot</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Disabled &#10060</td>")
		fmt.Fprintf(w, "    <td>Service does not automatically start on boot</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "     <td>Masked &#127917</td>")
		fmt.Fprintf(w, "    <td>Completely disabled, any start operation on it fails</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Static &#9940</td>")
		fmt.Fprintf(w, "    <td>The unit file is not enabled, and has no provisions <br>for enabling in the [Install] unit file section</td>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Alias &#128195&#8594&#128196</td>")
		fmt.Fprintf(w, "    <td>The name is an alias (symlink to another unit file)</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Indirect &#8669</td>")
		fmt.Fprintf(w, "    <td>The unit file itself is not enabled, but it has non-empty Also=<br>setting in the [Install] unit file section</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Enabled Runtime &#127939&#9201</td>")
		fmt.Fprintf(w, "    <td>Service automatically starts on boot</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>Unknown&#10067</td>")
		fmt.Fprintf(w, "    <td>Unknown Service</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>N/A</td>")
		fmt.Fprintf(w, "    <td>Not Applicable</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "</div>")
		fmt.Fprintf(w, "<script>")
		fmt.Fprintf(w, "function toggleKeyTable() {")
		fmt.Fprintf(w, "  var x = document.getElementById(\"keyTable\");")
		fmt.Fprintf(w, "  if (x.style.display === \"none\") {")
		fmt.Fprintf(w, "    x.style.display = \"table\";")
		fmt.Fprintf(w, "  } else {")
		fmt.Fprintf(w, "    x.style.display = \"none\";")
		fmt.Fprintf(w, "  }")
		fmt.Fprintf(w, "}")
		fmt.Fprintf(w, "</script>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th>&nbsp &nbsp Search &nbsp &nbsp</th>")
		fmt.Fprintf(w, "    <th><input type=\"text\" id=\"tableInput\" onkeyup=\"tableFunction()\" placeholder=\"Type to look for a service...\" title=\"search\"></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table id=\"table\">")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><b>Service</b></th>")
		fmt.Fprintf(w, "    <th><b>Status</b></th>")
		fmt.Fprintf(w, "    <th><b>Enabled on boot</b></th>")
		fmt.Fprintf(w, "  </tr>")

		systemdFiles, _ := exec.Command("find", "/lib/systemd/system/", "-maxdepth", "1", "-mindepth", "1", "-not", "-path", "*@*", "-and", "-not", "-path", "*wants*", "-and", "-not", "-path", "*.d*", "-execdir", "basename", "{}", ";").CombinedOutput()
		service := strings.Split(string(systemdFiles), "\n")
		service = service[:len(service)-1]
		for _, serviceName := range service {
			fmt.Fprintf(w, "  <tr>")
			serviceStatus, _ := exec.Command("systemctl", "status", serviceName).CombinedOutput()
			isEnabled, _ := exec.Command("systemctl", "is-enabled", serviceName).CombinedOutput()
			if string(serviceStatus) == ("Unit " + serviceName + ".service could not be found.\n") {
				fmt.Fprintf(w, "<td>Unit "+serviceName+" could not be found</td>")
				fmt.Fprintf(w, "<td>N/A</td>")
				fmt.Fprintf(w, "<td>N/A</td>")
			} else {
				serviceActive, _ := exec.Command("systemctl", "is-active", serviceName).CombinedOutput()
				if string(serviceActive) == "active\n" {
					fmt.Fprintf(w, "<td>"+serviceName+"</td>")
					fmt.Fprintf(w, "<td>&#128994</td>")
				} else if string(serviceActive) == "inactive\n" {
					fmt.Fprintf(w, "<td>"+serviceName+"</td>")
					fmt.Fprintf(w, "<td>&#128308</td>")
				} else {
					fmt.Println("Error")
					fmt.Println(string(serviceName))
				}

				if string(isEnabled) == "enabled\n" {
					fmt.Fprintf(w, "<td>Enabled &#9989</td>")
				} else if string(isEnabled) == "disabled\n" {
					fmt.Fprintf(w, "<td>Disabled &#10060</td>")
				} else if string(isEnabled) == "masked\n" {
					fmt.Fprintf(w, "<td>Masked &#127917</td>")
				} else if string(isEnabled) == "static\n" {
					fmt.Fprintf(w, "<td>Static &#9940</td>")
				} else if string(isEnabled) == "alias\n" {
					fmt.Fprintf(w, "<td>Alias &#128195&#8594&#128196</td>")
				} else if string(isEnabled) == "indirect\n" {
					fmt.Fprintf(w, "<td>Indirect &#8669</td>")
				} else if string(isEnabled) == "enabled-runtime\n" {
					fmt.Fprintf(w, "<td>Enabled Runtime &#127939&#9201</td>")
				} else {
					fmt.Fprintf(w, "<td>Unknown&#10067</td>")
				}
			}
		}
		fmt.Fprintf(w, "</tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<script>")
		fmt.Fprintf(w, "function tableFunction() {")
		fmt.Fprintf(w, "  var input, filter, table, tr, td, i, txtValue;")
		fmt.Fprintf(w, "  input = document.getElementById(\"tableInput\");")
		fmt.Fprintf(w, "  filter = input.value.toUpperCase();")
		fmt.Fprintf(w, "  table = document.getElementById(\"table\");")
		fmt.Fprintf(w, "  tr = table.getElementsByTagName(\"tr\");")
		fmt.Fprintf(w, "  for (i = 0; i < tr.length; i++) {")
		fmt.Fprintf(w, "    td = tr[i].getElementsByTagName(\"td\")[0];")
		fmt.Fprintf(w, "    if (td) {")
		fmt.Fprintf(w, "      txtValue = td.textContent || td.innerText;")
		fmt.Fprintf(w, "      if (txtValue.toUpperCase().indexOf(filter) > -1) {")
		fmt.Fprintf(w, "        tr[i].style.display = \"\";")
		fmt.Fprintf(w, "      } else {")
		fmt.Fprintf(w, "        tr[i].style.display = \"none\";")
		fmt.Fprintf(w, "      }")
		fmt.Fprintf(w, "    }")
		fmt.Fprintf(w, "  }")
		fmt.Fprintf(w, "}")
		fmt.Fprintf(w, "</script>")
		fmt.Fprintf(w, endHTML)
	})

	// Different port numbers will be assigned to each Go program later
	port := "localhost:8000"
	fmt.Println("Systemd View is running on port " + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {

	if runtime.GOOS != "linux" {
		fmt.Println("Operating system must be GNU/Linux to work")
	} else {
		systemd()
	}
}
