package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Docker subcommand
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Explains something about Docker",
	Long: `This command provides explanations and examples related to Docker.
For example:

- Explaining basic Docker commands
- Giving extensive information about advanced Docker features`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(`Docker command has no subcommands. Please provide one of the following flags:
--command
--advanced`)
			return
		}

		command, _ := cmd.Flags().GetString("command")
		if command != "" {
			explainDockerCommand(command)
			return
		}

		advanced, _ := cmd.Flags().GetString("advanced")
		if advanced != "" {
			explainAdvancedDockerConcepts(advanced)
			return
		}

		fmt.Println("Docker is a platform for developing, shipping, and running applications in containers.")
		fmt.Println("- docker run: Run a command in a new container.")
		fmt.Println("- docker build: Build an image from a Dockerfile.")
		fmt.Println("- docker push: Push an image or a repository to a registry.")
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)

	dockerCmd.Flags().StringVarP(&commandFlag, "command", "c", "", "Specify a Docker command to explain")
	dockerCmd.Flags().StringVarP(&advancedFlag, "advanced", "a", "", "Explain advanced Docker concepts")
}

func explainDockerCommand(command string) {
	switch command {
	case "run":
		fmt.Println("The 'docker run' command runs a command in a new container.")
		fmt.Println("Example: docker run -it ubuntu bash")
		fmt.Println("This command runs an interactive shell in a new Ubuntu container.")
	case "build":
		fmt.Println("The 'docker build' command builds an image from a Dockerfile.")
		fmt.Println("Example: docker build -t my-image .")
		fmt.Println("This command builds a Docker image named 'my-image' from the current directory.")
	case "push":
		fmt.Println("The 'docker push' command pushes an image or a repository to a registry.")
		fmt.Println("Example: docker push my-registry/my-image:latest")
		fmt.Println("This command pushes the 'my-image' image to the 'my-registry' registry with the 'latest' tag.")
	// Add more cases for other Docker commands
	default:
		fmt.Printf("Explanation for '%s' Docker command is not available. Try another Docker command.\n", command)
	}
}

func explainAdvancedDockerConcepts(command string) {
	switch command {
	case "compose":
		fmt.Println("- Docker Compose is a tool for defining and running multi-container Docker applications.")
		fmt.Println("It allows you to define the services, networks, and volumes in a YAML file, and then spin up the entire application stack with a single command.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Docker Compose:
docker-compose up               Build and start the entire application stack
docker-compose down             Stop and remove the entire application stack
docker-compose ps               List the status of containers defined in the Docker Compose file
docker-compose logs             View output from containers
docker-compose exec <service>   Run a command in a running service container`)
		fmt.Print("\n")
		fmt.Println("Example:")
		fmt.Println("Create a docker-compose.yml file defining services and then run:")
		fmt.Println("$ docker-compose up")
	case "swarm":
		fmt.Println("- Docker Swarm is a native clustering and orchestration solution for Docker.")
		fmt.Println("It turns a pool of Docker hosts into a single, virtual Docker host.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Docker Swarm:
docker swarm init               Initialize a new Docker Swarm
docker swarm join               Join a Docker host to a Swarm as a worker or manager
docker node ls                  List nodes in the Swarm
docker service ls               List services in the Swarm
docker stack deploy             Deploy a new stack or update an existing stack`)
		fmt.Print("\n")
		fmt.Println("Example:")
		fmt.Println("Initialize a new Docker Swarm:")
		fmt.Println("$ docker swarm init")
	case "network":
		fmt.Println("- Docker networking allows containers to communicate with each other and the outside world.")
		fmt.Println("Docker provides various network drivers for different use cases.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Docker networking:
docker network create            Create a new Docker network
docker network ls                List Docker networks
docker network inspect           Display detailed information about a Docker network`)
		fmt.Print("\n")
		fmt.Println("Example:")
		fmt.Println("Create a new bridge network:")
		fmt.Println("$ docker network create my-network")
	case "volume":
		fmt.Println("- Docker volumes are used to persist data generated by and used by Docker containers.")
		fmt.Println("They are a way to share data between containers or persist data across container restarts.")
		fmt.Print("\n")
		fmt.Println(`Here’s a summary of the different commands associated with Docker volumes:
docker volume create             Create a new Docker volume
docker volume ls                 List Docker volumes
docker volume inspect            Display detailed information about a Docker volume`)
		fmt.Print("\n")
		fmt.Println("Example:")
		fmt.Println("Create a new named volume:")
		fmt.Println("$ docker volume create my-data-volume")
	default:
		fmt.Printf("Explanation for '%s' Docker concept is not available. Try another advanced Docker concept.\n", command)
	}
}
