package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

// Struktur untuk menyimpan data calon legislatif dan partai
type Candidate struct {
	Name  string
	Party string
	Votes int
}

// Struktur untuk menyimpan data pemilih
type Voter struct {
	ID       string
	HasVoted bool
}

// Daftar kandidat dan pemilih
var candidates []Candidate
var voters []Voter
var electionStart time.Time
var electionEnd time.Time
var voteThreshold int

// Fungsi untuk menambahkan kandidat
func addCandidate(name string, party string) {
	candidates = append(candidates, Candidate{Name: name, Party: party, Votes: 0})
}

// Fungsi untuk mengedit kandidat
func editCandidate(oldName string, newName string, newParty string) {
	for i, candidate := range candidates {
		if candidate.Name == oldName {
			candidates[i].Name = newName
			candidates[i].Party = newParty
			fmt.Println("Candidate updated successfully.")
			return
		}
	}
	fmt.Println("Candidate not found.")
}

// Fungsi untuk menghapus kandidat
func deleteCandidate(name string) {
	for i, candidate := range candidates {
		if candidate.Name == name {
			candidates = append(candidates[:i], candidates[i+1:]...)
			fmt.Println("Candidate deleted successfully.")
			return
		}
	}
	fmt.Println("Candidate not found.")
}

// Fungsi untuk mendaftarkan pemilih
func registerVoter(id string) {
	voters = append(voters, Voter{ID: id, HasVoted: false})
}

// Fungsi untuk menghapus pemilih
func deleteVoter(id string) {
	for i, voter := range voters {
		if voter.ID == id {
			voters = append(voters[:i], voters[i+1:]...)
			fmt.Println("Voter deleted successfully.")
			return
		}
	}
	fmt.Println("Voter not found.")
}

// Fungsi untuk memilih kandidat
func vote(voterID string, candidateName string) {
	if time.Now().Before(electionStart) || time.Now().After(electionEnd) {
		fmt.Println("Voting is not allowed at this time.")
		return
	}
	for i, voter := range voters {
		if voter.ID == voterID {
			if voter.HasVoted {
				fmt.Println("You have already voted!")
				return
			}
			for j, candidate := range candidates {
				if candidate.Name == candidateName {
					candidates[j].Votes++
					voters[i].HasVoted = true
					fmt.Println("Vote successfully cast!")
					return
				}
			}
			fmt.Println("Candidate not found!")
			return
		}
	}
	fmt.Println("Voter not registered!")
}

// Fungsi untuk menampilkan hasil pemilu
func displayResults() {
	fmt.Println("Election Results:")
	for _, candidate := range candidates {
		fmt.Printf("%s (%s): %d votes\n", candidate.Name, candidate.Party, candidate.Votes)
	}
}

// Fungsi untuk menampilkan data terurut
func displaySortedResults(by string) {
	switch by {
	case "votes":
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].Votes > candidates[j].Votes
		})
	case "party":
		sort.Slice(candidates, func(i, j int) bool {
			if candidates[i].Party == candidates[j].Party {
				return candidates[i].Name < candidates[j].Name
			}
			return candidates[i].Party < candidates[j].Party
		})
	case "name":
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].Name < candidates[j].Name
		})
	default:
		fmt.Println("Invalid sort option.")
		return
	}
	fmt.Println("Sorted Results:")
	displayResults()
}

// Fungsi untuk mencari data
func searchCandidates(by string, value string) {
	fmt.Println("Search Results:")
	for _, candidate := range candidates {
		switch by {
		case "party":
			if candidate.Party == value {
				fmt.Printf("%s (%s): %d votes\n", candidate.Name, candidate.Party, candidate.Votes)
			}
		case "name":
			if candidate.Name == value {
				fmt.Printf("%s (%s): %d votes\n", candidate.Name, candidate.Party, candidate.Votes)
			}
		default:
			fmt.Println("Invalid search option.")
			return
		}
	}
}

func main() {
	// Menentukan durasi pemilu
	electionStart = time.Now()
	electionEnd = electionStart.Add(24 * time.Hour) // Durasi 1 hari
	voteThreshold = 10 // Ambang batas

	// Menu utama aplikasi
	for {
		fmt.Println("\nElection Application")
		fmt.Println("1. Add Candidate")
		fmt.Println("2. Edit Candidate")
		fmt.Println("3. Delete Candidate")
		fmt.Println("4. Register Voter")
		fmt.Println("5. Delete Voter")
		fmt.Println("6. Vote")
		fmt.Println("7. Display Results")
		fmt.Println("8. Display Sorted Results")
		fmt.Println("9. Search Candidates")
		fmt.Println("10. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, party string
			fmt.Print("Enter candidate name: ")
			fmt.Scan(&name)
			fmt.Print("Enter party name: ")
			fmt.Scan(&party)
			addCandidate(name, party)
		case 2:
			var oldName, newName, newParty string
			fmt.Print("Enter current candidate name: ")
			fmt.Scan(&oldName)
			fmt.Print("Enter new candidate name: ")
			fmt.Scan(&newName)
			fmt.Print("Enter new party name: ")
			fmt.Scan(&newParty)
			editCandidate(oldName, newName, newParty)
		case 3:
			var name string
			fmt.Print("Enter candidate name to delete: ")
			fmt.Scan(&name)
			deleteCandidate(name)
		case 4:
			var id string
			fmt.Print("Enter voter ID: ")
			fmt.Scan(&id)
			registerVoter(id)
		case 5:
			var id string
			fmt.Print("Enter voter ID to delete: ")
			fmt.Scan(&id)
			deleteVoter(id)
		case 6:
			var voterID, candidateName string
			fmt.Print("Enter voter ID: ")
			fmt.Scan(&voterID)
			fmt.Print("Enter candidate name: ")
			fmt.Scan(&candidateName)
			vote(voterID, candidateName)
		case 7:
			displayResults()
		case 8:
			var sortBy string
			fmt.Print("Sort by (votes/party/name): ")
			fmt.Scan(&sortBy)
			displaySortedResults(sortBy)
		case 9:
			var searchBy, value string
			fmt.Print("Search by (party/name): ")
			fmt.Scan(&searchBy)
			fmt.Print("Enter value: ")
			fmt.Scan(&value)
			searchCandidates(searchBy, value)
		case 10:
			fmt.Println("Exiting application...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}