package main

import "fmt"
import "os"
import "path/filepath"
import "strings"
import "syscall"
import "golang.org/x/sys/windows"

var defective_updates []string = []string{
	"C-00000291*.sys",
}

func isAdministrator() bool {

	var result bool = false

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")

	if err == nil {
		result = true
	}

	return result

}

func runAsElevatedProcess() {

	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verb_pointer, _ := syscall.UTF16PtrFromString(verb)
	exe_pointer, _ := syscall.UTF16PtrFromString(exe)
	cwd_pointer, _ := syscall.UTF16PtrFromString(cwd)
	arg_pointer, _ := syscall.UTF16PtrFromString(args)

	err := windows.ShellExecute(0, verb_pointer, exe_pointer, arg_pointer, cwd_pointer, int32(1))

	if err != nil {
		fmt.Println(err)
	}

}

func isDefectiveUpdate(name string) bool {

	var result bool = false

	for d := 0; d < len(defective_updates); d++ {

		pattern := defective_updates[d]

		if strings.Contains(pattern, "*") {

			prefix := pattern[0:strings.Index(pattern, "*")]
			suffix := pattern[strings.Index(pattern, "*")+1:]

			if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
				result = true
				break
			}

		}

	}

	return result

}

func showWarning() {
	fmt.Println("--------")
	fmt.Println("THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.")
	fmt.Println("--------")
}

func main() {

	showWarning()

	if isAdministrator() {

		win_path := os.Getenv("WINDIR")

		if win_path == "" {
			win_path = filepath.Join("C:", "WINDOWS")
		}

		dir_path := filepath.Join(win_path, "System32", "drivers", "CrowdStrike")

		files, err1 := os.ReadDir(dir_path)

		if err1 == nil {

			for f := 0; f < len(files); f++ {

				name := files[f].Name()

				if isDefectiveUpdate(name) {

					file_path := filepath.Join(dir_path, name)
					stat, err2 := os.Stat(file_path)

					if err2 == nil && stat.IsDir() == false {

						err3 := os.Remove(file_path)

						if err3 == nil {

							fmt.Println("Deleted defective update \"" + name + "\"")

						} else {
							fmt.Println("Cannot delete defective update \"" + name + "\"")
							fmt.Println("Please execute this program with admin rights!")
						}

					} else {
						fmt.Println("Cannot access defective update \"" + name + "\"")
						fmt.Println("Please execute this program with admin rights!")
					}

				}

			}

		} else {

			fmt.Println("Cannot access directory \"" + dir_path + "\".")
			fmt.Println("Please execute this program with admin rights!")

			fmt.Println("")
			fmt.Println("Detailed error:")
			fmt.Println(err1.Error())

		}

	} else {

		runAsElevatedProcess()

	}

	// 	showUsage()
	// 	os.Exit(1)

}
