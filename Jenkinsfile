// pipeline {
//     agent any
    
//     environment {
//         // Define environment variables for Docker Hub credentials
//         DOCKERHUB_CREDENTIALS = credentials('dockerhub_cred')
//         DOCKER_IMAGE_NAME = 'prady0t/pipeline'
//     }

//     stages {
//         stage('Checkout') {
//             steps {
//                 // Checkout code from a version control system (e.g., Git)
//                 echo 'Checkout'
//                 script {
//                     git branch: 'main',
//                         credentialsId: 'Github',
//                         url: 'https://github.com/prady0t/CI-CD-pipeline-app'
//                 }
//             }
//         }

//         // stage('Build and push Docker Image') {
//         //     steps {

//         //         echo 'Build'

//         //         script {
//         //             // Build the Docker image using the Dockerfile in the repository
//         //             dockerImage = docker.build env.DOCKER_IMAGE_NAME

//         //             // Authenticate with Docker Hub
//         //             withCredentials(DOCKERHUB_CREDENTIALS) {
//         //                 // Push the Docker image to Docker Hub
//         //                 docker.withRegistry('', "prady0t/pipeline") {
//         //                     dockerImage.push()
//         //                 }
//         //             }
//         //         }
//         //     }
//         // }
// //test
//         stage('Build and Push Docker Image') {
//       environment {
//         DOCKER_IMAGE = "prady0t/pipeline"
//         REGISTRY_CREDENTIALS = credentials('dockerhub-cred')
//       }
//       steps {
//         script {
  
//             def dockerImage = docker.image("${DOCKER_IMAGE}")
//             docker.withRegistry('https://index.docker.io/v1/', "dockerhub-cred") {
//                 dockerImage.push()
//             }
//         }
//       }    
//     }

//         // stage('Test') {
//         //     steps {
//         //         // Run tests (adjust as needed based on your testing framework)
//         //         echo 'Test'
//         //     }
//         // }

//         stage('Deploy') {
//             steps {
//                 // Perform deployment steps (e.g., deploy to a server)
//                 // Note: This stage is just a placeholder; adjust based on your deployment process
//                 echo 'Deploy'
//             }
//         }
//     }

//     post {
//         success {
//             // Actions to take if the pipeline is successful
//             echo 'Pipeline succeeded! Notify or perform additional tasks here.'
//         }

//         failure {
//             // Actions to take if the pipeline fails
//             echo 'Pipeline failed! Notify or perform cleanup tasks here.'
//         }
//     }
// }

pipeline {
    agent any
    
    environment {
        DOCKER_IMAGE_NAME = 'prady0t/pipeline'
        DOCKERFILE_LOCATION = ''
        GITHUB_REPO_URL = 'https://github.com/prady0t/CI-CD-pipeline-app'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    // Checkout code from GitHub
                    checkout([$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[url: GITHUB_REPO_URL]]])
                }
            }
        }

        stage('Build and Push Docker Image') {
            steps {
                script {
                    // Build the Docker image
                    def dockerImage = docker.build("${DOCKER_IMAGE_NAME}", "-f ${DOCKERFILE_LOCATION} .")

                    // Authenticate with Docker Hub
                    withCredentials([usernamePassword(credentialsId: 'dockerhub_cred', passwordVariable: 'DOCKERHUB_PASSWORD', usernameVariable: 'DOCKERHUB_USERNAME')]) {
                        // Log in to Docker Hub
                        docker.withRegistry("https://index.docker.io/v1/", "dockerhub_cred") {
                            // Push the Docker image to Docker Hub
                            dockerImage.push()
                        }
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline succeeded! Notify or perform additional tasks here.'
        }

        failure {
            echo 'Pipeline failed! Notify or perform cleanup tasks here.'
        }
    }
}
