package utils

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/dropbox/godropbox/errors"
	"github.com/pritunl/pritunl-link/errortypes"
)

func Exec(dir, name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if dir != "" {
		cmd.Dir = dir
	}

	err = cmd.Run()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to exec '%s'", name),
		}
		return
	}

	return
}

func ExecInput(dir, input, name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err,
				"utils: Failed to get stdin in exec '%s'", name),
		}
		return
	}

	if dir != "" {
		cmd.Dir = dir
	}

	err = cmd.Start()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to exec '%s'", name),
		}
		return
	}

	_, err = io.WriteString(stdin, input)
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to write stdin in exec '%s'",
				name),
		}
		return
	}

	err = stdin.Close()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to close stdin in exec '%s'",
				name),
		}
		return
	}

	err = cmd.Wait()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to exec '%s'", name),
		}
		return
	}

	return
}

func ExecOutput(dir, name string, arg ...string) (output string, err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr

	if dir != "" {
		cmd.Dir = dir
	}

	outputByt, err := cmd.Output()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to exec '%s'", name),
		}
		return
	}
	output = string(outputByt)

	return
}

func ExecCombinedOutput(dir, name string, arg ...string) (
	output string, err error) {

	cmd := exec.Command(name, arg...)

	if dir != "" {
		cmd.Dir = dir
	}

	outputByt, err := cmd.CombinedOutput()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to exec '%s'", name),
		}
		return
	}
	output = string(outputByt)

	return
}

func ExecSilent(dir, name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = ioutil.Discard

	if dir != "" {
		cmd.Dir = dir
	}

	err = cmd.Run()
	if err != nil {
		err = &errortypes.ExecError{
			errors.Wrapf(err, "utils: Failed to exec '%s'", name),
		}
		return
	}

	return
}
