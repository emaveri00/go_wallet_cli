package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Portafoglio struct {
	balance float64
}

func (p *Portafoglio) Deposit (earnings float64){
	p.balance += earnings
}

func (p *Portafoglio) Spend (spending float64) bool{
	if spending > p.balance {
		return false
	}
	p.balance -= spending
	return true
}

func (p *Portafoglio) getBalance() float64{
	return p.balance
}

func main() {
	portafoglio := Portafoglio{}
	reader:= bufio.NewReader(os.Stdin)
	Loop:
	for{
		fmt.Println("Operazioni sul portafoglio")
		fmt.Println("1. Depositare")
		fmt.Println("2. Spendere")
		fmt.Println("3. Visualizzare il saldo")
		fmt.Println("4. Uscire")
		operation, _ := reader.ReadString('\n')
		operation = strings.TrimSpace(operation)

		switch operation {
		case "1":
			fmt.Println("Inserisci l'importo da depositare:")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			var amount float64
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Errore: per favore inserisci un numero valido!")
				continue // Il ciclo for riparte e non si blocca
			}
			portafoglio.Deposit(amount)
			fmt.Printf("Hai depositato: %.2f\n", amount)
		case "2":
			fmt.Println("Inserisci l'importo da spendere:")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			var amount float64
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Errore: per favore inserisci un numero valido!")
				continue // Il ciclo for riparte e non si blocca
			}
			success := portafoglio.Spend(amount)
			if success {
				fmt.Printf("Hai speso: %.2f\n", amount)
			}else{
				fmt.Println("Saldo insufficiente per questa spesa.")
			}
		case "3":
			fmt.Printf("Il tuo saldo è: %.2f\n", portafoglio.getBalance())
		case "4":
			fmt.Println("Uscita dal programma.")
			break Loop
		}
	}
	
}
