pipeline {
    agent any
    
    environment {
        // Define environment variables for Docker Hub credentials
        DOCKERHUB_CREDENTIALS = credentials('dockerhub_cred')
        DOCKER_IMAGE_NAME = 'prady0t/pipeline'
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout code from a version control system (e.g., Git)
                echo 'Checkout'
                script {
                    git branch: 'main',
                        credentialsId: 'Github',
                        url: 'https://github.com/prady0t/CI-CD-pipeline-app'
                }
            }
        }

        // stage('Build and push Docker Image') {
        //     steps {

        //         echo 'Build'

        //         script {
        //             // Build the Docker image using the Dockerfile in the repository
        //             dockerImage = docker.build env.DOCKER_IMAGE_NAME

        //             // Authenticate with Docker Hub
        //             withCredentials(DOCKERHUB_CREDENTIALS) {
        //                 // Push the Docker image to Docker Hub
        //                 docker.withRegistry('', "prady0t/pipeline") {
        //                     dockerImage.push()
        //                 }
        //             }
        //         }
        //     }
        // }
//test
        stage('Build and Push Docker Image') {
      environment {
        DOCKER_IMAGE = "prady0t/pipeline"
        // DOCKERFILE_LOCATION = "java-maven-sonar-argocd-helm-k8s/spring-boot-app/Dockerfile"
        REGISTRY_CREDENTIALS = credentials('docker-cred')
      }
      steps {
        script {
  
            def dockerImage = docker.image("${DOCKER_IMAGE}")
            docker.withRegistry('https://index.docker.io/v1/', "docker-cred") {
                dockerImage.push()
            }
        }
      }

        // stage('Test') {
        //     steps {
        //         // Run tests (adjust as needed based on your testing framework)
        //         echo 'Test'
        //     }
        // }

        stage('Deploy') {
            steps {
                // Perform deployment steps (e.g., deploy to a server)
                // Note: This stage is just a placeholder; adjust based on your deployment process
                echo 'Deploy'
            }
        }
    }

    post {
        success {
            // Actions to take if the pipeline is successful
            echo 'Pipeline succeeded! Notify or perform additional tasks here.'
        }

        failure {
            // Actions to take if the pipeline fails
            echo 'Pipeline failed! Notify or perform cleanup tasks here.'
        }
    }
}
