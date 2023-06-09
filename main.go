package main

import (
	"fmt"
	"time"
)

// Karyawan Struct untuk merepresentasikan data karyawan
type Karyawan struct {
	ID        int
	Name      string
	Position  string
	CreatedAt time.Time
}

// EmployeeRepository Struct untuk mengelola data karyawan
type EmployeeRepository struct {
	employees []Karyawan
}

// AddEmployee Metode untuk menambahkan karyawan baru
func (er *EmployeeRepository) AddEmployee(employee Karyawan) {
	er.employees = append(er.employees, employee)
}

// RemoveEmployee Metode untuk menghapus karyawan berdasarkan ID
func (er *EmployeeRepository) RemoveEmployee(employeeID int) {
	for i, employee := range er.employees {
		if employee.ID == employeeID {
			er.employees = append(er.employees[:i], er.employees[i+1:]...)
			break
		}
	}
}

// FindEmployeeByID Metode untuk mencari karyawan berdasarkan ID
func (er *EmployeeRepository) FindEmployeeByID(employeeID int) *Karyawan {
	for _, employee := range er.employees {
		if employee.ID == employeeID {
			return &employee
		}
	}
	return nil
}

// AttendanceService Struct untuk mengelola absensi karyawan
type AttendanceService struct {
	employeeRepository *EmployeeRepository
}

// ClockIn Metode untuk melakukan absensi masuk
func (as *AttendanceService) ClockIn(employeeID int) {
	employee := as.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi masuk berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

// ClockOut Metode untuk melakukan absensi keluar
func (as *AttendanceService) ClockOut(employeeID int) {
	employee := as.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi keluar berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

func main() {
	employeeRepo := &EmployeeRepository{}
	attendanceService := &AttendanceService{employeeRepository: employeeRepo}

	employee1 := Karyawan{ID: 1, Name: "John Doe", Position: "Manager", CreatedAt: time.Now()}
	employee2 := Karyawan{ID: 2, Name: "Jane Smith", Position: "Staff", CreatedAt: time.Now()}

	employeeRepo.AddEmployee(employee1)
	employeeRepo.AddEmployee(employee2)

	attendanceService.ClockIn(1)
	attendanceService.ClockIn(2)

	attendanceService.ClockOut(1)
	attendanceService.ClockOut(3) // Karyawan dengan ID 3 tidak ada
}
