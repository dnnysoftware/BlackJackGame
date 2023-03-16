<a name="readme-top"></a>

[![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![made-with-golang](https://img.shields.io/badge/Go-v1.20-blue)](https://go.dev/)


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><<a href="#installation">Installation</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This Project is a blackjack game that is designed to have a player set options for setting up a deck of cards for a game of blackjack, however it is recommended to do a 8 deck shoe and have that shoe be randomly shuffled as the arguments. The goal of this game is to practice the game of blackjack but to prevent card ncounting the game ends after theres only a half deck (31 cards) in the shoe. The design of this game takes into consideration of software engineering Object-Oriented, SOLID and GRASP principles to allow expanison of the game.


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

Languages:
* Golang
* Shell Script

Libraries:
* math/rand
* time
* sort
* strings

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

Clone the repository, via HTTP or SSH above, download golang v1.20 if you havent already, finally start the program by running the shell script file `./run.sh` or run `main.go` manually. Once the game starts the user will have the option to hit by pressing 'h', stand by pressing 's', or quit the game by pressing 'q'. The game will continue until you reach a half deck of cards left in the shoe in a 8 deck starting shoe as per rules of the game.

### Installation

1. Clone the repo
  ```sh
  git clone https://github.com/dnnysoftware/BlackJackGame.git
  ```
2. Give file execute permissions to run.sh
  ```sh
  chmod +x run.sh
  ```
3. Run Program
  * In root directory run by typing in CLI
  ```sh
  ./run.sh
  ```
  * Or you can manually run the program and providing values to optional arguments by doing:
  ```sh
  go build .
  go run main.go
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>