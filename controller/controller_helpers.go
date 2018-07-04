package controller

import "github.com/kenigbolo/go-web-app/viewmodel"

//LoginAlertHelper function
func LoginAlertHelper(alert, message string, vm viewmodel.Login) viewmodel.Login {
	vm.AlertMessage = message
	if alert == "success" {
		vm.Alert = "visible"
		vm.AlertSuccess = "visible"
		vm.AlertDanger = "invisble"
		return vm
	}
	vm.Alert = "visible"
	vm.AlertSuccess = "invisible"
	vm.AlertDanger = "visible"
	return vm
}

//SignupAlertHelper function
func SignupAlertHelper(alert, message string, vm viewmodel.Signup) viewmodel.Signup {
	vm.AlertMessage = message
	if alert == "success" {
		vm.Alert = "visible"
		vm.AlertSuccess = "visible"
		vm.AlertDanger = "invisble"
		return vm
	}
	vm.Alert = "visible"
	vm.AlertSuccess = "invisible"
	vm.AlertDanger = "visible"
	return vm
}
