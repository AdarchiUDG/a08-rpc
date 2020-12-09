package main

import (
	"bufio"
	"os"
	"fmt"
	"net/rpc"
	"log"
	"strings"
)

type Grade struct {
	Student, Class string
	Value float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	defer client.Close()

	option := 0
	for option != 5{
		fmt.Println("1. Capturar Calificacion")
		fmt.Println("2. Ver Promedio de Alumno")
		fmt.Println("3. Ver Promedio de Materia")
		fmt.Println("4. Ver Promedio General")
		fmt.Println("5. Salir")
		fmt.Println("Teclea el numero de la opcion deseada")
		fmt.Print("> ")
	
		fmt.Scanf("%d", &option)
		fmt.Scanln()

		switch option {
			case 1:
				var value float64
				var reply float64

				fmt.Print("Calificacion: ")
				fmt.Scanf("%f", &value) 
				fmt.Scanln()
				grade := Grade{
					Student: readLine(in, "Nombre: "),
					Class: readLine(in, "Clase: "),
					Value: value }

				client.Call("School.AddGrade", grade, &reply)

				fmt.Println("Se agrego la calificacion")
			case 2:
				var reply float64
				err = client.Call("School.GetStudentAverage", readLine(in, "Estudiante: "), &reply)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("El promedio es: %.4f\n", reply)
				}
			case 3:
				var reply float64
				err = client.Call("School.GetClassAverage", readLine(in, "Clase: "), &reply)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("El promedio es: %.4f\n", reply)
				}
			case 4:
				var reply float64
				err = client.Call("School.GetGeneralAverage", true, &reply)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("El promedio es: %.4f\n", reply)
				}
			default:
		}

		if option != 5 {
			fmt.Scanln()
		}
	}
	fmt.Println("Saliendo . . .")
}

func readLine(in *bufio.Reader, str string) string {
	fmt.Print(str)
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}